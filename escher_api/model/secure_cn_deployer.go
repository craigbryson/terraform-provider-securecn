// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SecureCnDeployer secure cn deployer
// swagger:model SecureCnDeployer
type SecureCnDeployer struct {
	deployerField string

	deployerIdField *strfmt.UUID

	idField strfmt.UUID

	SecureCnDeployerAllOf1
}

// Deployer gets the deployer of this subtype
func (m *SecureCnDeployer) Deployer() string {
	return m.deployerField
}

// SetDeployer sets the deployer of this subtype
func (m *SecureCnDeployer) SetDeployer(val string) {
	m.deployerField = val
}

// DeployerID gets the deployer Id of this subtype
func (m *SecureCnDeployer) DeployerID() *strfmt.UUID {
	return m.deployerIdField
}

// SetDeployerID sets the deployer Id of this subtype
func (m *SecureCnDeployer) SetDeployerID(val *strfmt.UUID) {
	m.deployerIdField = val
}

// DeployerType gets the deployer type of this subtype
func (m *SecureCnDeployer) DeployerType() string {
	return "SecureCnDeployer"
}

// SetDeployerType sets the deployer type of this subtype
func (m *SecureCnDeployer) SetDeployerType(val string) {

}

// ID gets the id of this subtype
func (m *SecureCnDeployer) ID() strfmt.UUID {
	return m.idField
}

// SetID sets the id of this subtype
func (m *SecureCnDeployer) SetID(val strfmt.UUID) {
	m.idField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SecureCnDeployer) UnmarshalJSON(raw []byte) error {
	var data struct {
		SecureCnDeployerAllOf1
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Deployer string `json:"deployer,omitempty"`

		DeployerID *strfmt.UUID `json:"deployerId"`

		DeployerType string `json:"deployerType"`

		ID strfmt.UUID `json:"id,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result SecureCnDeployer

	result.deployerField = base.Deployer

	result.deployerIdField = base.DeployerID

	if base.DeployerType != result.DeployerType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid deployerType value: %q", base.DeployerType)
	}

	result.idField = base.ID

	result.SecureCnDeployerAllOf1 = data.SecureCnDeployerAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SecureCnDeployer) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		SecureCnDeployerAllOf1
	}{

		SecureCnDeployerAllOf1: m.SecureCnDeployerAllOf1,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Deployer string `json:"deployer,omitempty"`

		DeployerID *strfmt.UUID `json:"deployerId"`

		DeployerType string `json:"deployerType"`

		ID strfmt.UUID `json:"id,omitempty"`
	}{

		Deployer: m.Deployer(),

		DeployerID: m.DeployerID(),

		DeployerType: m.DeployerType(),

		ID: m.ID(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this secure cn deployer
func (m *SecureCnDeployer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with SecureCnDeployerAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecureCnDeployer) validateDeployerID(formats strfmt.Registry) error {

	if err := validate.Required("deployerId", "body", m.DeployerID()); err != nil {
		return err
	}

	if err := validate.FormatOf("deployerId", "body", "uuid", m.DeployerID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SecureCnDeployer) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID().String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecureCnDeployer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecureCnDeployer) UnmarshalBinary(b []byte) error {
	var res SecureCnDeployer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecureCnDeployerAllOf1 secure cn deployer all of1
// swagger:model SecureCnDeployerAllOf1
type SecureCnDeployerAllOf1 interface{}