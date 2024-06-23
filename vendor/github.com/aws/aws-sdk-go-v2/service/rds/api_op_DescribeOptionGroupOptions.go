// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes all available options for the specified engine.
func (c *Client) DescribeOptionGroupOptions(ctx context.Context, params *DescribeOptionGroupOptionsInput, optFns ...func(*Options)) (*DescribeOptionGroupOptionsOutput, error) {
	if params == nil {
		params = &DescribeOptionGroupOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOptionGroupOptions", params, optFns, c.addOperationDescribeOptionGroupOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOptionGroupOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOptionGroupOptionsInput struct {

	// The name of the engine to describe options for.
	//
	// Valid Values:
	//
	//   - db2-ae
	//
	//   - db2-se
	//
	//   - mariadb
	//
	//   - mysql
	//
	//   - oracle-ee
	//
	//   - oracle-ee-cdb
	//
	//   - oracle-se2
	//
	//   - oracle-se2-cdb
	//
	//   - postgres
	//
	//   - sqlserver-ee
	//
	//   - sqlserver-se
	//
	//   - sqlserver-ex
	//
	//   - sqlserver-web
	//
	// This member is required.
	EngineName *string

	// This parameter isn't currently supported.
	Filters []types.Filter

	// If specified, filters the results to include only options for the specified
	// major engine version.
	MajorEngineVersion *string

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeOptionGroupOptionsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// List of available option group options.
	OptionGroupOptions []types.OptionGroupOption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOptionGroupOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeOptionGroupOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeOptionGroupOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeOptionGroupOptions"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeOptionGroupOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOptionGroupOptions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeOptionGroupOptionsPaginatorOptions is the paginator options for
// DescribeOptionGroupOptions
type DescribeOptionGroupOptionsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOptionGroupOptionsPaginator is a paginator for
// DescribeOptionGroupOptions
type DescribeOptionGroupOptionsPaginator struct {
	options   DescribeOptionGroupOptionsPaginatorOptions
	client    DescribeOptionGroupOptionsAPIClient
	params    *DescribeOptionGroupOptionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeOptionGroupOptionsPaginator returns a new
// DescribeOptionGroupOptionsPaginator
func NewDescribeOptionGroupOptionsPaginator(client DescribeOptionGroupOptionsAPIClient, params *DescribeOptionGroupOptionsInput, optFns ...func(*DescribeOptionGroupOptionsPaginatorOptions)) *DescribeOptionGroupOptionsPaginator {
	if params == nil {
		params = &DescribeOptionGroupOptionsInput{}
	}

	options := DescribeOptionGroupOptionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOptionGroupOptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOptionGroupOptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOptionGroupOptions page.
func (p *DescribeOptionGroupOptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOptionGroupOptionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeOptionGroupOptions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeOptionGroupOptionsAPIClient is a client that implements the
// DescribeOptionGroupOptions operation.
type DescribeOptionGroupOptionsAPIClient interface {
	DescribeOptionGroupOptions(context.Context, *DescribeOptionGroupOptionsInput, ...func(*Options)) (*DescribeOptionGroupOptionsOutput, error)
}

var _ DescribeOptionGroupOptionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeOptionGroupOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeOptionGroupOptions",
	}
}