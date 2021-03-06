// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamquery

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelQuery struct {
}

func (*validateOpCancelQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpQuery struct {
}

func (*validateOpQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*QueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelQuery{}, middleware.After)
}

func addOpQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpQuery{}, middleware.After)
}

func validateOpCancelQueryInput(v *CancelQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelQueryInput"}
	if v.QueryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpQueryInput(v *QueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "QueryInput"}
	if v.QueryString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryString"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
