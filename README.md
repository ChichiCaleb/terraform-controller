# Give it a Little Time

## Introduction

This document provides a detailed design for a Kubernetes controller integrated with Terraform. 

The controller manages infrastructure as code (IaC) by applying Terraform configurations stored in Git repositories. This integration automates the reconciliation of infrastructure states defined by custom resources in Kubernetes.

## Objectives

> - Automate Infrastructure Management: Automatically apply, reconcile and manage infrastructure using Terraform based on custom Kubernetes resources.

> - Scalability: Handle multiple infrastructure configurations efficiently.

> - Extensibility: Allow easy addition of new features and backends.

> - Reliability: Ensure robust error handling and retry mechanisms.

> - Observability: Provide comprehensive logging and monitoring for debugging and performance analysis.

## Architecture

The project architecture consists of the following main components:

- Controller: The central component responsible for managing custom resources and executing Terraform commands.

- API Server: Exposes endpoints to handle incoming sync requests from metacontroller

- Terraform Integration: Executes Terraform scripts based on the provided configurations.

- Git Integration: Clones and manages Terraform configurations from Git repositories.

- Kubernetes Integration: Interacts with the Kubernetes API to manage resources and update statuses.

## Components

#### Controller

- Controller Struct: Manages Kubernetes clients used in  interacting with resources.

- ServeHTTP Method: Handles incoming HTTP requests for synchronization from metacontroller.

- Reconcile Method: Periodically reconciles the state of custom resources.

- handleSyncRequest Method: Processes the synchronization requests and executes Terraform commands.

#### API Server

- Gin Framework: Used for setting up the HTTP server and routing.

- Sync Endpoint: Exposes a POST endpoint /sync to receive synchronization requests.

#### Terraform Integration

- TerraformConfigSpec: Defines the structure of the custom resource spec for Terraform configurations.

- Scripts: Handles execution of Terraform  `apply`, and `destroy` scripts.

- Backend Setup: Configures backends like AWS S3 and Vault for storing Terraform states.

#### Git Integration

- CloneGitRepo Method: Clones the specified Git repository containing Terraform configurations.

#### Kubernetes Integration

- Kubernetes Clients: Handles interactions with the Kubernetes API.

- UpdateStatus Method: Updates the status of custom resources based on the outcome of Terraform commands.

#### in-cluster container build

- Using kaniko

#### Workflow

##### - Initialization:

> The controller initializes Kubernetes clients and dynamic clients.
> Sets up the API server using the Gin framework.

##### - Sync Request Handling:

> The API server listens for incoming sync requests at /sync.
> Upon receiving a request, the controller decodes the request body to a SyncRequest struct.

##### - Terraform Execution:

> The controller determines the appropriate Terraform script to run (apply or destroy).
> Clones the Terraform configuration from the specified Git repository.
> Sets up the backend (e.g., AWS S3, Vault) for Terraform state management.
> Executes the Terraform script with environment variables.

##### - Status Update:

> After executing Terraform, the controller updates the status of the custom resource with the outcome.
> If errors occur, the controller retries the operation with a maximum retry limit.

##### - Reconciliation Loop:

> The controller periodically reconciles the state of all custom resources to ensure consistency.

#### Environment Setup

 - Setup Script:
A Bash script `setup.sh` and `makefile` is provided to initialize the environment.


- Dockerfile:
Defines the container image for the controller.

- Kustomize:
packaging of the controller.

#### Logging and Monitoring

- Standard Logging: Uses the log package to write logs to standard error, ensuring compatibility with Kubernetes logging mechanisms.

- Log Levels: Info, error, and debug logs to provide detailed insights into the controller's operations.

- Monitoring: Integrate with tools like Prometheus and Grafana for metrics and dashboarding.

#### Testing and Coverage

- Unit Tests: Located in the test directory, covering individual components and methods.

- Integration Tests: Validate the end-to-end workflow from sync request handling to status updates.


#### Security Considerations

- Authentication and Authorization: Ensure that the controller has the necessary permissions to interact with Kubernetes resources.

- Secret Management: Securely handle sensitive information like Git SSH keys and backend credentials.


#### Future Enhancements

- Support for Additional Backends: Add support for more Terraform backends.

- Enhanced Error Handling: Improve error handling and retry mechanisms.

- Custom Metrics: Expose custom metrics for better observability.

- Webhooks: Implement webhooks for real-time notifications and updates.

```yaml
apiVersion: alustan.io/v1alpha1
kind: Terraform
metadata:
  name: terraformconfig
  namespace: default
spec:
  variables:
    TF_VAR_provision_cluster: true
    TF_VAR_provision_db: false
    TF_VAR_vpc_cidr: "10.1.0.0/16"
  backend:
    provider: aws
    s3: s3-store
    dynamoDB: db-table
    region: us-east-1
  scripts:
    deploy: deploy.sh
    destroy: destroy.sh
  gitRepo:
    url: https://github.com/example/terraform-repo
    branch: main
  containerRegistry:
    imageName: docker.io/alustan/terraform-controller
     
#  status:
#   state: "Pending"
#   message: "Awaiting processing"
```

**This design document outlines the architecture, components, and workflows for the Kubernetes controller integrated with Terraform. It serves as a reference for development, deployment, and future enhancements.**