// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// MetroclusterDiagState metrocluster diag state
//
// swagger:model metrocluster_diag_state
type MetroclusterDiagState string

func NewMetroclusterDiagState(value MetroclusterDiagState) *MetroclusterDiagState {
	v := value
	return &v
}

const (

	// MetroclusterDiagStateOk captures enum value "ok"
	MetroclusterDiagStateOk MetroclusterDiagState = "ok"

	// MetroclusterDiagStateWarning captures enum value "warning"
	MetroclusterDiagStateWarning MetroclusterDiagState = "warning"

	// MetroclusterDiagStateNotRun captures enum value "not_run"
	MetroclusterDiagStateNotRun MetroclusterDiagState = "not_run"

	// MetroclusterDiagStateNotApplicable captures enum value "not_applicable"
	MetroclusterDiagStateNotApplicable MetroclusterDiagState = "not_applicable"
)

// for schema
var metroclusterDiagStateEnum []interface{}

func init() {
	var res []MetroclusterDiagState
	if err := json.Unmarshal([]byte(`["ok","warning","not_run","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterDiagStateEnum = append(metroclusterDiagStateEnum, v)
	}
}

func (m MetroclusterDiagState) validateMetroclusterDiagStateEnum(path, location string, value MetroclusterDiagState) error {
	if err := validate.EnumCase(path, location, value, metroclusterDiagStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this metrocluster diag state
func (m MetroclusterDiagState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMetroclusterDiagStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this metrocluster diag state based on context it is used
func (m MetroclusterDiagState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
