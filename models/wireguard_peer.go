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

// WireguardPeer wireguard peer
//
// swagger:model WireguardPeer
type WireguardPeer struct {

	// allowed ips
	AllowedIps []string `json:"allowed_ips"`

	// peer id
	// Read Only: true
	PeerID string `json:"peer_id,omitempty"`

	// preshared key
	// Min Length: 32
	PresharedKey string `json:"preshared_key,omitempty"`

	// private key
	// Min Length: 32
	PrivateKey string `json:"private_key,omitempty"`

	// public key
	// Required: true
	// Min Length: 32
	PublicKey *string `json:"public_key"`
}

// Validate validates this wireguard peer
func (m *WireguardPeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePresharedKey(formats); err != nil {
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

func (m *WireguardPeer) validatePresharedKey(formats strfmt.Registry) error {
	if swag.IsZero(m.PresharedKey) { // not required
		return nil
	}

	if err := validate.MinLength("preshared_key", "body", m.PresharedKey, 32); err != nil {
		return err
	}

	return nil
}

func (m *WireguardPeer) validatePrivateKey(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivateKey) { // not required
		return nil
	}

	if err := validate.MinLength("private_key", "body", m.PrivateKey, 32); err != nil {
		return err
	}

	return nil
}

func (m *WireguardPeer) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("public_key", "body", m.PublicKey); err != nil {
		return err
	}

	if err := validate.MinLength("public_key", "body", *m.PublicKey, 32); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this wireguard peer based on the context it is used
func (m *WireguardPeer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePeerID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WireguardPeer) contextValidatePeerID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "peer_id", "body", string(m.PeerID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WireguardPeer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WireguardPeer) UnmarshalBinary(b []byte) error {
	var res WireguardPeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
