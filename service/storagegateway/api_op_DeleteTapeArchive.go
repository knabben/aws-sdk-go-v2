// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// DeleteTapeArchiveInput
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DeleteTapeArchiveInput
type DeleteTapeArchiveInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the virtual tape to delete from the virtual
	// tape shelf (VTS).
	//
	// TapeARN is a required field
	TapeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTapeArchiveInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTapeArchiveInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTapeArchiveInput"}

	if s.TapeARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("TapeARN"))
	}
	if s.TapeARN != nil && len(*s.TapeARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("TapeARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// DeleteTapeArchiveOutput
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DeleteTapeArchiveOutput
type DeleteTapeArchiveOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the virtual tape that was deleted from
	// the virtual tape shelf (VTS).
	TapeARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DeleteTapeArchiveOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTapeArchive = "DeleteTapeArchive"

// DeleteTapeArchiveRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Deletes the specified virtual tape from the virtual tape shelf (VTS). This
// operation is only supported in the tape gateway type.
//
//    // Example sending a request using DeleteTapeArchiveRequest.
//    req := client.DeleteTapeArchiveRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DeleteTapeArchive
func (c *Client) DeleteTapeArchiveRequest(input *DeleteTapeArchiveInput) DeleteTapeArchiveRequest {
	op := &aws.Operation{
		Name:       opDeleteTapeArchive,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTapeArchiveInput{}
	}

	req := c.newRequest(op, input, &DeleteTapeArchiveOutput{})
	return DeleteTapeArchiveRequest{Request: req, Input: input, Copy: c.DeleteTapeArchiveRequest}
}

// DeleteTapeArchiveRequest is the request type for the
// DeleteTapeArchive API operation.
type DeleteTapeArchiveRequest struct {
	*aws.Request
	Input *DeleteTapeArchiveInput
	Copy  func(*DeleteTapeArchiveInput) DeleteTapeArchiveRequest
}

// Send marshals and sends the DeleteTapeArchive API request.
func (r DeleteTapeArchiveRequest) Send(ctx context.Context) (*DeleteTapeArchiveResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTapeArchiveResponse{
		DeleteTapeArchiveOutput: r.Request.Data.(*DeleteTapeArchiveOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTapeArchiveResponse is the response type for the
// DeleteTapeArchive API operation.
type DeleteTapeArchiveResponse struct {
	*DeleteTapeArchiveOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTapeArchive request.
func (r *DeleteTapeArchiveResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}