// Code generated by smithy-go-codegen DO NOT EDIT.

package pricing

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/encoding/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/encoding/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsAwsjson11_serializeOpDescribeServices struct {
}

func (*awsAwsjson11_serializeOpDescribeServices) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeServices) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeServicesInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSPriceListService.DescribeServices")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeServicesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpGetAttributeValues struct {
}

func (*awsAwsjson11_serializeOpGetAttributeValues) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGetAttributeValues) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetAttributeValuesInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSPriceListService.GetAttributeValues")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGetAttributeValuesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpGetProducts struct {
}

func (*awsAwsjson11_serializeOpGetProducts) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGetProducts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetProductsInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSPriceListService.GetProducts")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGetProductsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentFilter(v *types.Filter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Field != nil {
		ok := object.Key("Field")
		ok.String(*v.Field)
	}

	if len(v.Type) > 0 {
		ok := object.Key("Type")
		ok.String(string(v.Type))
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson11_serializeDocumentFilters(v []types.Filter, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentFilter(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeServicesInput(v *DescribeServicesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.FormatVersion != nil {
		ok := object.Key("FormatVersion")
		ok.String(*v.FormatVersion)
	}

	if v.MaxResults != 0 {
		ok := object.Key("MaxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.ServiceCode != nil {
		ok := object.Key("ServiceCode")
		ok.String(*v.ServiceCode)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentGetAttributeValuesInput(v *GetAttributeValuesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AttributeName != nil {
		ok := object.Key("AttributeName")
		ok.String(*v.AttributeName)
	}

	if v.MaxResults != 0 {
		ok := object.Key("MaxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.ServiceCode != nil {
		ok := object.Key("ServiceCode")
		ok.String(*v.ServiceCode)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentGetProductsInput(v *GetProductsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Filters != nil {
		ok := object.Key("Filters")
		if err := awsAwsjson11_serializeDocumentFilters(v.Filters, ok); err != nil {
			return err
		}
	}

	if v.FormatVersion != nil {
		ok := object.Key("FormatVersion")
		ok.String(*v.FormatVersion)
	}

	if v.MaxResults != 0 {
		ok := object.Key("MaxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.ServiceCode != nil {
		ok := object.Key("ServiceCode")
		ok.String(*v.ServiceCode)
	}

	return nil
}