// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"strconv"
	"time"
)

// Returns information about DB snapshots. This API action supports pagination.
func (c *Client) DescribeDBSnapshots(ctx context.Context, params *DescribeDBSnapshotsInput, optFns ...func(*Options)) (*DescribeDBSnapshotsOutput, error) {
	if params == nil {
		params = &DescribeDBSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBSnapshots", params, optFns, c.addOperationDescribeDBSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBSnapshotsInput struct {

	// The ID of the DB instance to retrieve the list of DB snapshots for. This
	// parameter isn't case-sensitive.
	//
	// Constraints:
	//
	//   - If supplied, must match the identifier of an existing DBInstance.
	DBInstanceIdentifier *string

	// A specific DB snapshot identifier to describe. This value is stored as a
	// lowercase string.
	//
	// Constraints:
	//
	//   - If supplied, must match the identifier of an existing DBSnapshot.
	//
	//   - If this identifier is for an automated snapshot, the SnapshotType parameter
	//   must also be specified.
	DBSnapshotIdentifier *string

	// A specific DB resource ID to describe.
	DbiResourceId *string

	// A filter that specifies one or more DB snapshots to describe.
	//
	// Supported filters:
	//
	//   - db-instance-id - Accepts DB instance identifiers and DB instance Amazon
	//   Resource Names (ARNs).
	//
	//   - db-snapshot-id - Accepts DB snapshot identifiers.
	//
	//   - dbi-resource-id - Accepts identifiers of source DB instances.
	//
	//   - snapshot-type - Accepts types of DB snapshots.
	//
	//   - engine - Accepts names of database engines.
	Filters []types.Filter

	// Specifies whether to include manual DB cluster snapshots that are public and
	// can be copied or restored by any Amazon Web Services account. By default, the
	// public snapshots are not included.
	//
	// You can share a manual DB snapshot as public by using the ModifyDBSnapshotAttribute API.
	//
	// This setting doesn't apply to RDS Custom.
	IncludePublic *bool

	// Specifies whether to include shared manual DB cluster snapshots from other
	// Amazon Web Services accounts that this Amazon Web Services account has been
	// given permission to copy or restore. By default, these snapshots are not
	// included.
	//
	// You can give an Amazon Web Services account permission to restore a manual DB
	// snapshot from another Amazon Web Services account by using the
	// ModifyDBSnapshotAttribute API action.
	//
	// This setting doesn't apply to RDS Custom.
	IncludeShared *bool

	// An optional pagination token provided by a previous DescribeDBSnapshots
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The type of snapshots to be returned. You can specify one of the following
	// values:
	//
	//   - automated - Return all DB snapshots that have been automatically taken by
	//   Amazon RDS for my Amazon Web Services account.
	//
	//   - manual - Return all DB snapshots that have been taken by my Amazon Web
	//   Services account.
	//
	//   - shared - Return all manual DB snapshots that have been shared to my Amazon
	//   Web Services account.
	//
	//   - public - Return all DB snapshots that have been marked as public.
	//
	//   - awsbackup - Return the DB snapshots managed by the Amazon Web Services
	//   Backup service.
	//
	// For information about Amazon Web Services Backup, see the [Amazon Web Services Backup Developer Guide.]
	//
	// The awsbackup type does not apply to Aurora.
	//
	// If you don't specify a SnapshotType value, then both automated and manual
	// snapshots are returned. Shared and public DB snapshots are not included in the
	// returned results by default. You can include shared snapshots with these results
	// by enabling the IncludeShared parameter. You can include public snapshots with
	// these results by enabling the IncludePublic parameter.
	//
	// The IncludeShared and IncludePublic parameters don't apply for SnapshotType
	// values of manual or automated . The IncludePublic parameter doesn't apply when
	// SnapshotType is set to shared . The IncludeShared parameter doesn't apply when
	// SnapshotType is set to public .
	//
	// [Amazon Web Services Backup Developer Guide.]: https://docs.aws.amazon.com/aws-backup/latest/devguide/whatisbackup.html
	SnapshotType *string

	noSmithyDocumentSerde
}

// Contains the result of a successful invocation of the DescribeDBSnapshots
// action.
type DescribeDBSnapshotsOutput struct {

	// A list of DBSnapshot instances.
	DBSnapshots []types.DBSnapshot

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDBSnapshots"); err != nil {
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
	if err = addOpDescribeDBSnapshotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBSnapshots(options.Region), middleware.Before); err != nil {
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

// DBSnapshotAvailableWaiterOptions are waiter options for
// DBSnapshotAvailableWaiter
type DBSnapshotAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// DBSnapshotAvailableWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, DBSnapshotAvailableWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeDBSnapshotsInput, *DescribeDBSnapshotsOutput, error) (bool, error)
}

// DBSnapshotAvailableWaiter defines the waiters for DBSnapshotAvailable
type DBSnapshotAvailableWaiter struct {
	client DescribeDBSnapshotsAPIClient

	options DBSnapshotAvailableWaiterOptions
}

// NewDBSnapshotAvailableWaiter constructs a DBSnapshotAvailableWaiter.
func NewDBSnapshotAvailableWaiter(client DescribeDBSnapshotsAPIClient, optFns ...func(*DBSnapshotAvailableWaiterOptions)) *DBSnapshotAvailableWaiter {
	options := DBSnapshotAvailableWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = dBSnapshotAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &DBSnapshotAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for DBSnapshotAvailable waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *DBSnapshotAvailableWaiter) Wait(ctx context.Context, params *DescribeDBSnapshotsInput, maxWaitDur time.Duration, optFns ...func(*DBSnapshotAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for DBSnapshotAvailable waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *DBSnapshotAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeDBSnapshotsInput, maxWaitDur time.Duration, optFns ...func(*DBSnapshotAvailableWaiterOptions)) (*DescribeDBSnapshotsOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeDBSnapshots(ctx, params, func(o *Options) {
			baseOpts := []func(*Options){
				addIsWaiterUserAgent,
			}
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range baseOpts {
				opt(o)
			}
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for DBSnapshotAvailable waiter")
}

func dBSnapshotAvailableStateRetryable(ctx context.Context, input *DescribeDBSnapshotsInput, output *DescribeDBSnapshotsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("DBSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleted"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleting"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "failed"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-restore"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-parameters"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

// DBSnapshotDeletedWaiterOptions are waiter options for DBSnapshotDeletedWaiter
type DBSnapshotDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// DBSnapshotDeletedWaiter will use default minimum delay of 30 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, DBSnapshotDeletedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeDBSnapshotsInput, *DescribeDBSnapshotsOutput, error) (bool, error)
}

// DBSnapshotDeletedWaiter defines the waiters for DBSnapshotDeleted
type DBSnapshotDeletedWaiter struct {
	client DescribeDBSnapshotsAPIClient

	options DBSnapshotDeletedWaiterOptions
}

// NewDBSnapshotDeletedWaiter constructs a DBSnapshotDeletedWaiter.
func NewDBSnapshotDeletedWaiter(client DescribeDBSnapshotsAPIClient, optFns ...func(*DBSnapshotDeletedWaiterOptions)) *DBSnapshotDeletedWaiter {
	options := DBSnapshotDeletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = dBSnapshotDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &DBSnapshotDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for DBSnapshotDeleted waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *DBSnapshotDeletedWaiter) Wait(ctx context.Context, params *DescribeDBSnapshotsInput, maxWaitDur time.Duration, optFns ...func(*DBSnapshotDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for DBSnapshotDeleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *DBSnapshotDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeDBSnapshotsInput, maxWaitDur time.Duration, optFns ...func(*DBSnapshotDeletedWaiterOptions)) (*DescribeDBSnapshotsOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeDBSnapshots(ctx, params, func(o *Options) {
			baseOpts := []func(*Options){
				addIsWaiterUserAgent,
			}
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range baseOpts {
				opt(o)
			}
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for DBSnapshotDeleted waiter")
}

func dBSnapshotDeletedStateRetryable(ctx context.Context, input *DescribeDBSnapshotsInput, output *DescribeDBSnapshotsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("length(DBSnapshots) == `0`", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "true"
		bv, err := strconv.ParseBool(expectedValue)
		if err != nil {
			return false, fmt.Errorf("error parsing boolean from string %w", err)
		}
		value, ok := pathValue.(bool)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected bool value got %T", pathValue)
		}

		if value == bv {
			return false, nil
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "DBSnapshotNotFound" == apiErr.ErrorCode() {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "creating"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "modifying"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "rebooting"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "resetting-master-credentials"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

// DescribeDBSnapshotsPaginatorOptions is the paginator options for
// DescribeDBSnapshots
type DescribeDBSnapshotsPaginatorOptions struct {
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

// DescribeDBSnapshotsPaginator is a paginator for DescribeDBSnapshots
type DescribeDBSnapshotsPaginator struct {
	options   DescribeDBSnapshotsPaginatorOptions
	client    DescribeDBSnapshotsAPIClient
	params    *DescribeDBSnapshotsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBSnapshotsPaginator returns a new DescribeDBSnapshotsPaginator
func NewDescribeDBSnapshotsPaginator(client DescribeDBSnapshotsAPIClient, params *DescribeDBSnapshotsInput, optFns ...func(*DescribeDBSnapshotsPaginatorOptions)) *DescribeDBSnapshotsPaginator {
	if params == nil {
		params = &DescribeDBSnapshotsInput{}
	}

	options := DescribeDBSnapshotsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBSnapshotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBSnapshotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBSnapshots page.
func (p *DescribeDBSnapshotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBSnapshotsOutput, error) {
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
	result, err := p.client.DescribeDBSnapshots(ctx, &params, optFns...)
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

// DescribeDBSnapshotsAPIClient is a client that implements the
// DescribeDBSnapshots operation.
type DescribeDBSnapshotsAPIClient interface {
	DescribeDBSnapshots(context.Context, *DescribeDBSnapshotsInput, ...func(*Options)) (*DescribeDBSnapshotsOutput, error)
}

var _ DescribeDBSnapshotsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeDBSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDBSnapshots",
	}
}