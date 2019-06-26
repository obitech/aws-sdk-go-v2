// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastoredata

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-data-2017-09-01/DescribeObjectRequest
type DescribeObjectInput struct {
	_ struct{} `type:"structure"`

	// The path (including the file name) where the object is stored in the container.
	// Format: <folder name>/<folder name>/<file name>
	//
	// Path is a required field
	Path *string `location:"uri" locationName:"Path" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeObjectInput"}

	if s.Path == nil {
		invalidParams.Add(aws.NewErrParamRequired("Path"))
	}
	if s.Path != nil && len(*s.Path) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Path", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeObjectInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Path != nil {
		v := *s.Path

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Path", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-data-2017-09-01/DescribeObjectResponse
type DescribeObjectOutput struct {
	_ struct{} `type:"structure"`

	// An optional CacheControl header that allows the caller to control the object's
	// cache behavior. Headers can be passed in as specified in the HTTP at https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9
	// (https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9).
	//
	// Headers with a custom user-defined value are also accepted.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// The length of the object in bytes.
	ContentLength *int64 `location:"header" locationName:"Content-Length" type:"long"`

	// The content type of the object.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// The ETag that represents a unique instance of the object.
	ETag *string `location:"header" locationName:"ETag" min:"1" type:"string"`

	// The date and time that the object was last modified.
	LastModified *time.Time `location:"header" locationName:"Last-Modified" type:"timestamp" timestampFormat:"rfc822"`
}

// String returns the string representation
func (s DescribeObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeObjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CacheControl != nil {
		v := *s.CacheControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Cache-Control", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentLength != nil {
		v := *s.ContentLength

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Length", protocol.Int64Value(v), metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Last-Modified", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
	return nil
}

const opDescribeObject = "DescribeObject"

// DescribeObjectRequest returns a request value for making API operation for
// AWS Elemental MediaStore Data Plane.
//
// Gets the headers for an object at the specified path.
//
//    // Example sending a request using DescribeObjectRequest.
//    req := client.DescribeObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-data-2017-09-01/DescribeObject
func (c *Client) DescribeObjectRequest(input *DescribeObjectInput) DescribeObjectRequest {
	op := &aws.Operation{
		Name:       opDescribeObject,
		HTTPMethod: "HEAD",
		HTTPPath:   "/{Path+}",
	}

	if input == nil {
		input = &DescribeObjectInput{}
	}

	req := c.newRequest(op, input, &DescribeObjectOutput{})
	return DescribeObjectRequest{Request: req, Input: input, Copy: c.DescribeObjectRequest}
}

// DescribeObjectRequest is the request type for the
// DescribeObject API operation.
type DescribeObjectRequest struct {
	*aws.Request
	Input *DescribeObjectInput
	Copy  func(*DescribeObjectInput) DescribeObjectRequest
}

// Send marshals and sends the DescribeObject API request.
func (r DescribeObjectRequest) Send(ctx context.Context) (*DescribeObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeObjectResponse{
		DescribeObjectOutput: r.Request.Data.(*DescribeObjectOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeObjectResponse is the response type for the
// DescribeObject API operation.
type DescribeObjectResponse struct {
	*DescribeObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeObject request.
func (r *DescribeObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}