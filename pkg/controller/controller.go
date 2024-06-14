package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"time"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/alustan/terraform-controller/pkg/container"
	"github.com/alustan/terraform-controller/pkg/kubernetes"
    "github.com/alustan/terraform-controller/pkg/util"
	"github.com/alustan/terraform-controller/plugin"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	dynclient "k8s.io/client-go/dynamic"
	k8sclient "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	maxRetries = 10
)

type Controller struct {
	clientset *k8sclient.Clientset
	dynClient dynclient.Interface
}

type TerraformConfigSpec struct {
	Variables        map[string]string `json:"variables"`
	Backend          map[string]string `json:"backend"`
	Scripts          Scripts           `json:"scripts"`
	GitRepo          GitRepo           `json:"gitRepo"`
	ContainerRegistry ContainerRegistry `json:"containerRegistry"`
	
}

type Scripts struct {
	Deploy   string `json:"deploy"`
	Destroy string `json:"destroy"`
}

type GitRepo struct {
	URL          string       `json:"url"`
	Branch       string       `json:"branch"`
}

type ContainerRegistry struct {
	ImageName string    `json:"imageName"`
}

type ParentResource struct {
	ApiVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Metadata   metav1.ObjectMeta      `json:"metadata"`
	Spec       TerraformConfigSpec    `json:"spec"`
	Status     map[string]interface{} `json:"status"`
}

type SyncRequest struct {
	Parent     ParentResource `json:"parent"`
	Finalizing bool           `json:"finalizing"`
}

func NewController(clientset *k8sclient.Clientset, dynClient dynclient.Interface) *Controller {
	return &Controller{
		clientset: clientset,
		dynClient: dynClient,
	}
}

func NewInClusterController() *Controller {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error creating in-cluster config: %v", err)
	}

	clientset, err := k8sclient.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes clientset: %v", err)
	}

	dynClient, err := dynclient.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating dynamic Kubernetes client: %v", err)
	}

	return NewController(clientset, dynClient)
}


func (c *Controller) ServeHTTP(r *gin.Context) {
	var observed SyncRequest
	err := json.NewDecoder(r.Request.Body).Decode(&observed)
	if err != nil {
		r.String(http.StatusBadRequest, err.Error())
		return
	}
	defer func() {
		if err := r.Request.Body.Close(); err != nil {
			log.Printf("Error closing request body: %v", err)
		}
	}()

	response := c.handleSyncRequest(observed)

	r.Writer.Header().Set("Content-Type", "application/json")
	r.JSON(http.StatusOK, gin.H{"body": response})
}



func (c *Controller) handleSyncRequest(observed SyncRequest) map[string]interface{} {
	envVars := c.extractEnvVars(observed.Parent.Spec.Variables)
	log.Printf("Observed Parent Spec: %+v", observed.Parent.Spec)

	scriptContent := observed.Parent.Spec.Scripts.Deploy
	if scriptContent == "" {
		status := c.errorResponse("executing deploy script", fmt.Errorf("deploy script is missing"))
		c.updateStatus(observed, status)
		return status
	}

	if observed.Finalizing {
		scriptContent = observed.Parent.Spec.Scripts.Destroy
		if scriptContent == "" {
			status := c.errorResponse("executing destroy script", fmt.Errorf("destroy script is missing"))
			c.updateStatus(observed, status)
			return status
		}
	}

	repoDir := filepath.Join("/tmp", observed.Parent.Metadata.Name)
	sshKey := os.Getenv("GIT_SSH_SECRET")

	dockerfileAdditions, providerExists, err := c.setupBackend(observed.Parent.Spec.Backend)
	if err != nil {
		status := c.errorResponse("setting up backend", err)
		c.updateStatus(observed, status)
		return status
	}

	configMapName, err := container.CreateDockerfileConfigMap(c.clientset, observed.Parent.Metadata.Name, observed.Parent.Metadata.Namespace, dockerfileAdditions, providerExists)
	if err != nil {
		status := c.errorResponse("creating Dockerfile ConfigMap", err)
		c.updateStatus(observed, status)
		return status
	}

	encodedDockerConfigJSON := os.Getenv("CONTAINER_REGISTRY_SECRET")
	secretName := fmt.Sprintf("%s-container-secret", observed.Parent.Metadata.Name)
	err = container.CreateDockerConfigSecret(c.clientset, secretName, observed.Parent.Metadata.Namespace, encodedDockerConfigJSON)
	if err != nil {
		status := c.errorResponse("creating Docker config secret", err)
		c.updateStatus(observed, status)
		return status
	}

	taggedImageName, _, err := c.buildAndTagImage(observed, configMapName, repoDir, sshKey, secretName)
	if err != nil {
		status := c.errorResponse("creating build job", err)
		c.updateStatus(observed, status)
		return status
	}

	status := c.runTerraform(observed, scriptContent, taggedImageName, secretName, envVars)
	
    c.updateStatus(observed, status)

	return status
}


func (c *Controller) updateStatus(observed SyncRequest, status map[string]interface{}) {
	err := kubernetes.UpdateStatus(c.dynClient, observed.Parent.Metadata.Namespace, observed.Parent.Metadata.Name, status)
	if err != nil {
		log.Printf("Error updating status for %s: %v", observed.Parent.Metadata.Name, err)
	}
}

func (c *Controller) extractEnvVars(variables map[string]string) map[string]string {
	if variables == nil {
		return nil
	}
	return util.ExtractEnvVars(variables)
}


func (c *Controller) setupBackend(backend map[string]string) (string, bool, error) {
	if backend == nil || len(backend) == 0 {
		log.Println("No backend provided, continuing without backend setup")
		return "", false, nil
	}

	providerType, providerExists := backend["provider"]
	if !providerExists || providerType == "" {
		log.Println("Backend provided without specifying provider, continuing without backend setup")
		return "", false, nil
	}

	provider, err := plugin.GetProvider(providerType)
	if err != nil {
		return "", false, fmt.Errorf("error getting provider: %v", err)
	}

	if err := provider.SetupBackend(backend); err != nil {
		return "", false, fmt.Errorf("error setting up %s backend: %v", providerType, err)
	}

	return provider.GetDockerfileAdditions(), true, nil
}

func (c *Controller) buildAndTagImage(observed SyncRequest, configMapName, repoDir, sshKey,secretName string) (string,string, error) {
	imageName := observed.Parent.Spec.ContainerRegistry.ImageName
	

	return container.CreateBuildPod(c.clientset, 
		  observed.Parent.Metadata.Name,
		  observed.Parent.Metadata.Namespace,
		  configMapName, 
		  imageName, 
		  secretName,
		  repoDir,
		  observed.Parent.Spec.GitRepo.URL,
		  observed.Parent.Spec.GitRepo.Branch,
		  sshKey)
}




func (c *Controller) runTerraform(observed SyncRequest, scriptContent, taggedImageName, secretName string, envVars map[string]string) map[string]interface{} {


	var terraformErr error
	for i := 0; i < maxRetries; i++ {
		terraformErr = container.CreateRunPod(c.clientset, observed.Parent.Metadata.Name, observed.Parent.Metadata.Namespace, envVars, scriptContent, taggedImageName, secretName)
		if terraformErr == nil {
			break
		}
		log.Printf("Retrying Terraform command due to error: %v", terraformErr)
		time.Sleep(1 * time.Minute)
	}

	status := map[string]interface{}{
		"state":   "Success",
		"message": "Terraform applied successfully",
	}
	if terraformErr != nil {
		status["state"] = "Failed"
		status["message"] = terraformErr.Error()
	}

	return status
}

func (c *Controller) errorResponse(action string, err error) map[string]interface{} {
	log.Printf("Error %s: %v", action, err)
	return map[string]interface{}{
		"state":  "error",
		"message": fmt.Sprintf("Error %s: %v", action, err),
	}
}


func (c *Controller) Reconcile(syncInterval time.Duration) {
	for {
		c.reconcileLoop()
		time.Sleep(syncInterval)
	}
}

func (c *Controller) reconcileLoop() {
	log.Println("Starting reconciliation loop")
	resourceList, err := c.dynClient.Resource(schema.GroupVersionResource{
		Group:    "alustan.io",
		Version:  "v1alpha1",
		Resource: "terraforms",
	}).Namespace("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error fetching Terraform resources: %v", err)
		return
	}

	log.Printf("Fetched %d Terraform resources", len(resourceList.Items))

	for _, item := range resourceList.Items {
		go func(item unstructured.Unstructured) {
			var observed SyncRequest
			raw, err := item.MarshalJSON()
			if err != nil {
				log.Printf("Error marshalling item: %v", err)
				return
			}
			err = json.Unmarshal(raw, &observed)
			if err != nil {
				log.Printf("Error unmarshalling item: %v", err)
				return
			}

			log.Printf("Handling resource: %s", observed.Parent.Metadata.Name)
			c.handleSyncRequest(observed)
		}(item)
	}
}
