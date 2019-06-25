// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetDevEndpointsRequest
type GetDevEndpointsInput struct {
	_ struct{} `type:"structure"`

	// The maximum size of information to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetDevEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDevEndpointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDevEndpointsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetDevEndpointsResponse
type GetDevEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// A list of DevEndpoint definitions.
	DevEndpoints []DevEndpoint `type:"list"`

	// A continuation token, if not all DevEndpoint definitions have yet been returned.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetDevEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDevEndpoints = "GetDevEndpoints"

// GetDevEndpointsRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves all the DevEndpoints in this AWS account.
//
// When you create a development endpoint in a virtual private cloud (VPC),
// AWS Glue returns only a private IP address and the public IP address field
// is not populated. When you create a non-VPC development endpoint, AWS Glue
// returns only a public IP address.
//
//    // Example sending a request using GetDevEndpointsRequest.
//    req := client.GetDevEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetDevEndpoints
func (c *Client) GetDevEndpointsRequest(input *GetDevEndpointsInput) GetDevEndpointsRequest {
	op := &aws.Operation{
		Name:       opGetDevEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetDevEndpointsInput{}
	}

	req := c.newRequest(op, input, &GetDevEndpointsOutput{})
	return GetDevEndpointsRequest{Request: req, Input: input, Copy: c.GetDevEndpointsRequest}
}

// GetDevEndpointsRequest is the request type for the
// GetDevEndpoints API operation.
type GetDevEndpointsRequest struct {
	*aws.Request
	Input *GetDevEndpointsInput
	Copy  func(*GetDevEndpointsInput) GetDevEndpointsRequest
}

// Send marshals and sends the GetDevEndpoints API request.
func (r GetDevEndpointsRequest) Send(ctx context.Context) (*GetDevEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDevEndpointsResponse{
		GetDevEndpointsOutput: r.Request.Data.(*GetDevEndpointsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetDevEndpointsRequestPaginator returns a paginator for GetDevEndpoints.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetDevEndpointsRequest(input)
//   p := glue.NewGetDevEndpointsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetDevEndpointsPaginator(req GetDevEndpointsRequest) GetDevEndpointsPaginator {
	return GetDevEndpointsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetDevEndpointsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetDevEndpointsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetDevEndpointsPaginator struct {
	aws.Pager
}

func (p *GetDevEndpointsPaginator) CurrentPage() *GetDevEndpointsOutput {
	return p.Pager.CurrentPage().(*GetDevEndpointsOutput)
}

// GetDevEndpointsResponse is the response type for the
// GetDevEndpoints API operation.
type GetDevEndpointsResponse struct {
	*GetDevEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDevEndpoints request.
func (r *GetDevEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}