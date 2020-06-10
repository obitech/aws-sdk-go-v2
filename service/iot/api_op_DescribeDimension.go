// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeDimensionInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the dimension.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDimensionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDimensionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDimensionInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDimensionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeDimensionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN (Amazon resource name) for the dimension.
	Arn *string `locationName:"arn" type:"string"`

	// The date the dimension was created.
	CreationDate *time.Time `locationName:"creationDate" type:"timestamp"`

	// The date the dimension was last modified.
	LastModifiedDate *time.Time `locationName:"lastModifiedDate" type:"timestamp"`

	// The unique identifier for the dimension.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The value or list of values used to scope the dimension. For example, for
	// topic filters, this is the pattern used to match the MQTT topic name.
	StringValues []string `locationName:"stringValues" min:"1" type:"list"`

	// The type of the dimension.
	Type DimensionType `locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeDimensionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDimensionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.LastModifiedDate != nil {
		v := *s.LastModifiedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastModifiedDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StringValues != nil {
		v := s.StringValues

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "stringValues", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opDescribeDimension = "DescribeDimension"

// DescribeDimensionRequest returns a request value for making API operation for
// AWS IoT.
//
// Provides details about a dimension that is defined in your AWS account.
//
//    // Example sending a request using DescribeDimensionRequest.
//    req := client.DescribeDimensionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeDimensionRequest(input *DescribeDimensionInput) DescribeDimensionRequest {
	op := &aws.Operation{
		Name:       opDescribeDimension,
		HTTPMethod: "GET",
		HTTPPath:   "/dimensions/{name}",
	}

	if input == nil {
		input = &DescribeDimensionInput{}
	}

	req := c.newRequest(op, input, &DescribeDimensionOutput{})

	return DescribeDimensionRequest{Request: req, Input: input, Copy: c.DescribeDimensionRequest}
}

// DescribeDimensionRequest is the request type for the
// DescribeDimension API operation.
type DescribeDimensionRequest struct {
	*aws.Request
	Input *DescribeDimensionInput
	Copy  func(*DescribeDimensionInput) DescribeDimensionRequest
}

// Send marshals and sends the DescribeDimension API request.
func (r DescribeDimensionRequest) Send(ctx context.Context) (*DescribeDimensionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDimensionResponse{
		DescribeDimensionOutput: r.Request.Data.(*DescribeDimensionOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDimensionResponse is the response type for the
// DescribeDimension API operation.
type DescribeDimensionResponse struct {
	*DescribeDimensionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDimension request.
func (r *DescribeDimensionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}