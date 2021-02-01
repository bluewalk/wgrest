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

// WireguardDevice wireguard device
//
// swagger:model WireguardDevice
type WireguardDevice struct {

	// listen port
	// Required: true
	ListenPort *int64 `json:"listen_port"`

	// name
	// Required: true
	Name *string `json:"name"`

	// network
	// Required: true
	Network *string `json:"network"`

	// private key
	// Required: true
	// Min Length: 32
	PrivateKey *string `json:"private_key"`

	// public key
	// Read Only: true
	// Min Length: 32
	PublicKey string `json:"public_key,omitempty"`
}

// Validate validates this wireguard device
func (m *WireguardDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateListenPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WireguardDevice) validateListenPort(formats strfmt.Registry) error {

	if err := validate.Required("listen_port", "body", m.ListenPort); err != nil {
		return err
	}

	return nil
}

func (m *WireguardDevice) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *WireguardDevice) validateNetwork(formats strfmt.Registry) error {

	if err := validate.Required("network", "body", m.Network); err != nil {
		return err
	}

	return nil
}

func (m *WireguardDevice) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("private_key", "body", m.PrivateKey); err != nil {
		return err
	}

	if err := validate.MinLength("private_key", "body", *m.PrivateKey, 32); err != nil {
		return err
	}

	return nil
}

func (m *WireguardDevice) validatePublicKey(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicKey) { // not required
		return nil
	}

	if err := validate.MinLength("public_key", "body", m.PublicKey, 32); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this wireguard device based on the context it is used
func (m *WireguardDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePublicKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WireguardDevice) contextValidatePublicKey(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "public_key", "body", string(m.PublicKey)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WireguardDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WireguardDevice) UnmarshalBinary(b []byte) error {
	var res WireguardDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
