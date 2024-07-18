// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1Vtxo v1 vtxo
// swagger:model v1Vtxo
type V1Vtxo struct {

	// expire at
	ExpireAt string `json:"expireAt,omitempty"`

	// outpoint
	Outpoint *V1Input `json:"outpoint,omitempty"`

	// pool txid
	PoolTxid string `json:"poolTxid,omitempty"`

	// receiver
	Receiver *V1Output `json:"receiver,omitempty"`

	// spent
	Spent bool `json:"spent,omitempty"`

	// spent by
	SpentBy string `json:"spentBy,omitempty"`

	// swept
	Swept bool `json:"swept,omitempty"`
}

// Validate validates this v1 vtxo
func (m *V1Vtxo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceiver(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Vtxo) validateOutpoint(formats strfmt.Registry) error {

	if swag.IsZero(m.Outpoint) { // not required
		return nil
	}

	if m.Outpoint != nil {
		if err := m.Outpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outpoint")
			}
			return err
		}
	}

	return nil
}

func (m *V1Vtxo) validateReceiver(formats strfmt.Registry) error {

	if swag.IsZero(m.Receiver) { // not required
		return nil
	}

	if m.Receiver != nil {
		if err := m.Receiver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("receiver")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Vtxo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Vtxo) UnmarshalBinary(b []byte) error {
	var res V1Vtxo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
