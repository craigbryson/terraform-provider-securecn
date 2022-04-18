// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServerlessRuleType identify the serverless functions matching type. Only one of the below should be not null, and  used.
//
// swagger:discriminator ServerlessRuleType serverlessRuleType
type ServerlessRuleType interface {
	runtime.Validatable
	runtime.ContextValidatable

	// serverless function validation
	ServerlessFunctionValidation() *ServerlessFunctionValidation
	SetServerlessFunctionValidation(*ServerlessFunctionValidation)

	// serverless rule type
	// Required: true
	// Enum: [FunctionAnyServerlessRuleType FunctionNameServerlessRuleType FunctionArnServerlessRuleType]
	ServerlessRuleType() string
	SetServerlessRuleType(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type serverlessRuleType struct {
	serverlessFunctionValidationField *ServerlessFunctionValidation

	serverlessRuleTypeField string
}

// ServerlessFunctionValidation gets the serverless function validation of this polymorphic type
func (m *serverlessRuleType) ServerlessFunctionValidation() *ServerlessFunctionValidation {
	return m.serverlessFunctionValidationField
}

// SetServerlessFunctionValidation sets the serverless function validation of this polymorphic type
func (m *serverlessRuleType) SetServerlessFunctionValidation(val *ServerlessFunctionValidation) {
	m.serverlessFunctionValidationField = val
}

// ServerlessRuleType gets the serverless rule type of this polymorphic type
func (m *serverlessRuleType) ServerlessRuleType() string {
	return "ServerlessRuleType"
}

// SetServerlessRuleType sets the serverless rule type of this polymorphic type
func (m *serverlessRuleType) SetServerlessRuleType(val string) {
}

// UnmarshalServerlessRuleTypeSlice unmarshals polymorphic slices of ServerlessRuleType
func UnmarshalServerlessRuleTypeSlice(reader io.Reader, consumer runtime.Consumer) ([]ServerlessRuleType, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []ServerlessRuleType
	for _, element := range elements {
		obj, err := unmarshalServerlessRuleType(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalServerlessRuleType unmarshals polymorphic ServerlessRuleType
func UnmarshalServerlessRuleType(reader io.Reader, consumer runtime.Consumer) (ServerlessRuleType, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalServerlessRuleType(data, consumer)
}

func unmarshalServerlessRuleType(data []byte, consumer runtime.Consumer) (ServerlessRuleType, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the serverlessRuleType property.
	var getType struct {
		ServerlessRuleType string `json:"serverlessRuleType"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("serverlessRuleType", "body", getType.ServerlessRuleType); err != nil {
		return nil, err
	}

	// The value of serverlessRuleType is used to determine which type to create and unmarshal the data into
	switch getType.ServerlessRuleType {
	case "FunctionAnyServerlessRuleType":
		var result FunctionAnyServerlessRuleType
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "FunctionArnServerlessRuleType":
		var result FunctionArnServerlessRuleType
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "FunctionNameServerlessRuleType":
		var result FunctionNameServerlessRuleType
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "ServerlessRuleType":
		var result serverlessRuleType
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid serverlessRuleType value: %q", getType.ServerlessRuleType)
}

// Validate validates this serverless rule type
func (m *serverlessRuleType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerlessFunctionValidation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *serverlessRuleType) validateServerlessFunctionValidation(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerlessFunctionValidation()) { // not required
		return nil
	}

	if m.ServerlessFunctionValidation() != nil {
		if err := m.ServerlessFunctionValidation().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serverlessFunctionValidation")
			}
			return err
		}
	}

	return nil
}

var serverlessRuleTypeTypeServerlessRuleTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FunctionAnyServerlessRuleType","FunctionNameServerlessRuleType","FunctionArnServerlessRuleType"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverlessRuleTypeTypeServerlessRuleTypePropEnum = append(serverlessRuleTypeTypeServerlessRuleTypePropEnum, v)
	}
}

const (

	// ServerlessRuleTypeServerlessRuleTypeFunctionAnyServerlessRuleType captures enum value "FunctionAnyServerlessRuleType"
	ServerlessRuleTypeServerlessRuleTypeFunctionAnyServerlessRuleType string = "FunctionAnyServerlessRuleType"

	// ServerlessRuleTypeServerlessRuleTypeFunctionNameServerlessRuleType captures enum value "FunctionNameServerlessRuleType"
	ServerlessRuleTypeServerlessRuleTypeFunctionNameServerlessRuleType string = "FunctionNameServerlessRuleType"

	// ServerlessRuleTypeServerlessRuleTypeFunctionArnServerlessRuleType captures enum value "FunctionArnServerlessRuleType"
	ServerlessRuleTypeServerlessRuleTypeFunctionArnServerlessRuleType string = "FunctionArnServerlessRuleType"
)

// prop value enum
func (m *serverlessRuleType) validateServerlessRuleTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverlessRuleTypeTypeServerlessRuleTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

// ContextValidate validate this serverless rule type based on the context it is used
func (m *serverlessRuleType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServerlessFunctionValidation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *serverlessRuleType) contextValidateServerlessFunctionValidation(ctx context.Context, formats strfmt.Registry) error {

	if m.ServerlessFunctionValidation() != nil {
		if err := m.ServerlessFunctionValidation().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serverlessFunctionValidation")
			}
			return err
		}
	}

	return nil
}
