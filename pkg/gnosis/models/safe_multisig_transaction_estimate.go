// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SafeMultisigTransactionEstimate safe multisig transaction estimate
//
// swagger:model SafeMultisigTransactionEstimate
type SafeMultisigTransactionEstimate struct {

	// Data
	Data *string `json:"data,omitempty"`

	// Operation
	// Required: true
	// Minimum: 0
	Operation *int64 `json:"operation"`

	// To
	// Required: true
	To *string `json:"to"`

	// Value
	// Required: true
	// Minimum: 0
	Value *int64 `json:"value"`
}

// Validate validates this safe multisig transaction estimate
func (m *SafeMultisigTransactionEstimate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SafeMultisigTransactionEstimate) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	if err := validate.MinimumInt("operation", "body", *m.Operation, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SafeMultisigTransactionEstimate) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

func (m *SafeMultisigTransactionEstimate) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.MinimumInt("value", "body", *m.Value, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this safe multisig transaction estimate based on context it is used
func (m *SafeMultisigTransactionEstimate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SafeMultisigTransactionEstimate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SafeMultisigTransactionEstimate) UnmarshalBinary(b []byte) error {
	var res SafeMultisigTransactionEstimate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
