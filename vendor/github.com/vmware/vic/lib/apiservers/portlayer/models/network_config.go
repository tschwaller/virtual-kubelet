package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkConfig network config
// swagger:model NetworkConfig
type NetworkConfig struct {

	// address
	Address string `json:"address,omitempty"`

	// aliases
	Aliases []string `json:"aliases"`

	// network name
	// Required: true
	NetworkName string `json:"networkName"`

	// ports
	Ports []string `json:"ports"`
}

// Validate validates this network config
func (m *NetworkConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAliases(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetworkName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkConfig) validateAliases(formats strfmt.Registry) error {

	if swag.IsZero(m.Aliases) { // not required
		return nil
	}

	return nil
}

func (m *NetworkConfig) validateNetworkName(formats strfmt.Registry) error {

	if err := validate.RequiredString("networkName", "body", string(m.NetworkName)); err != nil {
		return err
	}

	return nil
}

func (m *NetworkConfig) validatePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	return nil
}
