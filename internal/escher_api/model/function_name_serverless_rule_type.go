// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FunctionNameServerlessRuleType function name serverless rule type
//
// swagger:model FunctionNameServerlessRuleType
type FunctionNameServerlessRuleType struct {
	serverlessFunctionValidationField *ServerlessFunctionValidation

	// functions that match one of the given ARNs
	Names []string `json:"names"`
}

// ServerlessFunctionValidation gets the serverless function validation of this subtype
func (m *FunctionNameServerlessRuleType) ServerlessFunctionValidation() *ServerlessFunctionValidation {
	return m.serverlessFunctionValidationField
}

// SetServerlessFunctionValidation sets the serverless function validation of this subtype
func (m *FunctionNameServerlessRuleType) SetServerlessFunctionValidation(val *ServerlessFunctionValidation) {
	m.serverlessFunctionValidationField = val
}

// ServerlessRuleType gets the serverless rule type of this subtype
func (m *FunctionNameServerlessRuleType) ServerlessRuleType() string {
	return "FunctionNameServerlessRuleType"
}

// SetServerlessRuleType sets the serverless rule type of this subtype
func (m *FunctionNameServerlessRuleType) SetServerlessRuleType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *FunctionNameServerlessRuleType) UnmarshalJSON(raw []byte) error {
	var data struct {

		// functions that match one of the given ARNs
		Names []string `json:"names"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ServerlessFunctionValidation *ServerlessFunctionValidation `json:"serverlessFunctionValidation,omitempty"`

		ServerlessRuleType string `json:"serverlessRuleType"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result FunctionNameServerlessRuleType

	result.serverlessFunctionValidationField = base.ServerlessFunctionValidation

	if base.ServerlessRuleType != result.ServerlessRuleType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid serverlessRuleType value: %q", base.ServerlessRuleType)
	}

	result.Names = data.Names

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m FunctionNameServerlessRuleType) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// functions that match one of the given ARNs
		Names []string `json:"names"`
	}{

		Names: m.Names,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ServerlessFunctionValidation *ServerlessFunctionValidation `json:"serverlessFunctionValidation,omitempty"`

		ServerlessRuleType string `json:"serverlessRuleType"`
	}{

		ServerlessFunctionValidation: m.ServerlessFunctionValidation(),

		ServerlessRuleType: m.ServerlessRuleType(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this function name serverless rule type
func (m *FunctionNameServerlessRuleType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerlessFunctionValidation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FunctionNameServerlessRuleType) validateServerlessFunctionValidation(formats strfmt.Registry) error {

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

// ContextValidate validate this function name serverless rule type based on the context it is used
func (m *FunctionNameServerlessRuleType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServerlessFunctionValidation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FunctionNameServerlessRuleType) contextValidateServerlessFunctionValidation(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *FunctionNameServerlessRuleType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionNameServerlessRuleType) UnmarshalBinary(b []byte) error {
	var res FunctionNameServerlessRuleType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}