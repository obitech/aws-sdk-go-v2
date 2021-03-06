// Code generated by smithy-go-codegen DO NOT EDIT.

package ioteventsdata

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchPutMessage struct {
}

func (*validateOpBatchPutMessage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchPutMessage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchPutMessageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchPutMessageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchUpdateDetector struct {
}

func (*validateOpBatchUpdateDetector) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchUpdateDetector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchUpdateDetectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchUpdateDetectorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeDetector struct {
}

func (*validateOpDescribeDetector) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeDetector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeDetectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeDetectorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListDetectors struct {
}

func (*validateOpListDetectors) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListDetectors) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListDetectorsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListDetectorsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchPutMessageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchPutMessage{}, middleware.After)
}

func addOpBatchUpdateDetectorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchUpdateDetector{}, middleware.After)
}

func addOpDescribeDetectorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeDetector{}, middleware.After)
}

func addOpListDetectorsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListDetectors{}, middleware.After)
}

func validateDetectorStateDefinition(v *types.DetectorStateDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DetectorStateDefinition"}
	if v.StateName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateName"))
	}
	if v.Variables == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Variables"))
	} else if v.Variables != nil {
		if err := validateVariableDefinitions(v.Variables); err != nil {
			invalidParams.AddNested("Variables", err.(smithy.InvalidParamsError))
		}
	}
	if v.Timers == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timers"))
	} else if v.Timers != nil {
		if err := validateTimerDefinitions(v.Timers); err != nil {
			invalidParams.AddNested("Timers", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMessage(v *types.Message) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Message"}
	if v.MessageId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MessageId"))
	}
	if v.InputName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputName"))
	}
	if v.Payload == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Payload"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMessages(v []types.Message) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Messages"}
	for i := range v {
		if err := validateMessage(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTimerDefinition(v *types.TimerDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TimerDefinition"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Seconds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Seconds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTimerDefinitions(v []types.TimerDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TimerDefinitions"}
	for i := range v {
		if err := validateTimerDefinition(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateDetectorRequest(v *types.UpdateDetectorRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDetectorRequest"}
	if v.MessageId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MessageId"))
	}
	if v.DetectorModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DetectorModelName"))
	}
	if v.State == nil {
		invalidParams.Add(smithy.NewErrParamRequired("State"))
	} else if v.State != nil {
		if err := validateDetectorStateDefinition(v.State); err != nil {
			invalidParams.AddNested("State", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateDetectorRequests(v []types.UpdateDetectorRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDetectorRequests"}
	for i := range v {
		if err := validateUpdateDetectorRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVariableDefinition(v *types.VariableDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VariableDefinition"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVariableDefinitions(v []types.VariableDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VariableDefinitions"}
	for i := range v {
		if err := validateVariableDefinition(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchPutMessageInput(v *BatchPutMessageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchPutMessageInput"}
	if v.Messages == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Messages"))
	} else if v.Messages != nil {
		if err := validateMessages(v.Messages); err != nil {
			invalidParams.AddNested("Messages", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchUpdateDetectorInput(v *BatchUpdateDetectorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchUpdateDetectorInput"}
	if v.Detectors == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Detectors"))
	} else if v.Detectors != nil {
		if err := validateUpdateDetectorRequests(v.Detectors); err != nil {
			invalidParams.AddNested("Detectors", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeDetectorInput(v *DescribeDetectorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeDetectorInput"}
	if v.DetectorModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DetectorModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListDetectorsInput(v *ListDetectorsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListDetectorsInput"}
	if v.DetectorModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DetectorModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
