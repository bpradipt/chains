// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of aliases in the caller's Amazon Web Services account and region.
// For more information about aliases, see CreateAlias.
//
// By default, the ListAliases operation returns all aliases in the account and
// region. To get only the aliases associated with a particular KMS key, use the
// KeyId parameter.
//
// The ListAliases response can include aliases that you created and associated
// with your customer managed keys, and aliases that Amazon Web Services created
// and associated with Amazon Web Services managed keys in your account. You can
// recognize Amazon Web Services aliases because their names have the format aws/ ,
// such as aws/dynamodb .
//
// The response might also include aliases that have no TargetKeyId field. These
// are predefined aliases that Amazon Web Services has created but has not yet
// associated with a KMS key. Aliases that Amazon Web Services creates in your
// account, including predefined aliases, do not count against your [KMS aliases quota].
//
// Cross-account use: No. ListAliases does not return aliases in other Amazon Web
// Services accounts.
//
// Required permissions: [kms:ListAliases] (IAM policy)
//
// For details, see [Controlling access to aliases] in the Key Management Service Developer Guide.
//
// Related operations:
//
// # CreateAlias
//
// # DeleteAlias
//
// # UpdateAlias
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [KMS aliases quota]: https://docs.aws.amazon.com/kms/latest/developerguide/limits.html#aliases-limit
// [kms:ListAliases]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-eventual-consistency.html
// [Controlling access to aliases]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html#alias-access
func (c *Client) ListAliases(ctx context.Context, params *ListAliasesInput, optFns ...func(*Options)) (*ListAliasesOutput, error) {
	if params == nil {
		params = &ListAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAliases", params, optFns, c.addOperationListAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAliasesInput struct {

	// Lists only aliases that are associated with the specified KMS key. Enter a KMS
	// key in your Amazon Web Services account.
	//
	// This parameter is optional. If you omit it, ListAliases returns all aliases in
	// the account and Region.
	//
	// Specify the key ID or key ARN of the KMS key.
	//
	// For example:
	//
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.
	KeyId *string

	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, KMS does not return more than the specified number of items,
	// but it might return fewer.
	//
	// This value is optional. If you include a value, it must be between 1 and 100,
	// inclusive. If you do not include a value, it defaults to 50.
	Limit *int32

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated
	// response you just received.
	Marker *string

	noSmithyDocumentSerde
}

type ListAliasesOutput struct {

	// A list of aliases.
	Aliases []types.AliasListEntry

	// When Truncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent request.
	NextMarker *string

	// A flag that indicates whether there are more items in the list. When this value
	// is true, the list in this response is truncated. To get more items, pass the
	// value of the NextMarker element in this response to the Marker parameter in a
	// subsequent request.
	Truncated bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAliases{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAliases"); err != nil {
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
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAliases(options.Region), middleware.Before); err != nil {
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
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// ListAliasesPaginatorOptions is the paginator options for ListAliases
type ListAliasesPaginatorOptions struct {
	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, KMS does not return more than the specified number of items,
	// but it might return fewer.
	//
	// This value is optional. If you include a value, it must be between 1 and 100,
	// inclusive. If you do not include a value, it defaults to 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAliasesPaginator is a paginator for ListAliases
type ListAliasesPaginator struct {
	options   ListAliasesPaginatorOptions
	client    ListAliasesAPIClient
	params    *ListAliasesInput
	nextToken *string
	firstPage bool
}

// NewListAliasesPaginator returns a new ListAliasesPaginator
func NewListAliasesPaginator(client ListAliasesAPIClient, params *ListAliasesInput, optFns ...func(*ListAliasesPaginatorOptions)) *ListAliasesPaginator {
	if params == nil {
		params = &ListAliasesInput{}
	}

	options := ListAliasesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAliasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAliasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAliases page.
func (p *ListAliasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAliasesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListAliases(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListAliasesAPIClient is a client that implements the ListAliases operation.
type ListAliasesAPIClient interface {
	ListAliases(context.Context, *ListAliasesInput, ...func(*Options)) (*ListAliasesOutput, error)
}

var _ ListAliasesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAliases",
	}
}
