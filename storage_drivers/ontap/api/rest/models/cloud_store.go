// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudStore cloud store
//
// swagger:model cloud_store
type CloudStore struct {

	// links
	Links *CloudStoreLinks `json:"_links,omitempty"`

	// aggregate
	Aggregate *CloudStoreAggregate `json:"aggregate,omitempty"`

	// Availability of the object store.
	// Read Only: true
	// Enum: [available unavailable]
	Availability string `json:"availability,omitempty"`

	// This field identifies if the mirror cloud store is in sync with the primary cloud store of a FabricPool.
	// Read Only: true
	MirrorDegraded *bool `json:"mirror_degraded,omitempty"`

	// This field indicates whether the cloud store is the primary cloud store of a mirrored FabricPool.
	Primary *bool `json:"primary,omitempty"`

	// target
	Target *CloudStoreTarget `json:"target,omitempty"`

	// unavailable reason
	UnavailableReason *CloudStoreUnavailableReason `json:"unavailable_reason,omitempty"`

	// Usage threshold for reclaiming unused space in the cloud store. Valid values are 0 to 99. The default value depends on the provider type. This can be specified in PATCH but not POST.
	// Example: 20
	UnreclaimedSpaceThreshold int64 `json:"unreclaimed_space_threshold,omitempty"`

	// The amount of object space used. Calculated every 5 minutes and cached.
	// Read Only: true
	Used int64 `json:"used,omitempty"`
}

// Validate validates this cloud store
func (m *CloudStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAggregate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnavailableReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStore) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *CloudStore) validateAggregate(formats strfmt.Registry) error {
	if swag.IsZero(m.Aggregate) { // not required
		return nil
	}

	if m.Aggregate != nil {
		if err := m.Aggregate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate")
			}
			return err
		}
	}

	return nil
}

var cloudStoreTypeAvailabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["available","unavailable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudStoreTypeAvailabilityPropEnum = append(cloudStoreTypeAvailabilityPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// cloud_store
	// CloudStore
	// availability
	// Availability
	// available
	// END DEBUGGING
	// CloudStoreAvailabilityAvailable captures enum value "available"
	CloudStoreAvailabilityAvailable string = "available"

	// BEGIN DEBUGGING
	// cloud_store
	// CloudStore
	// availability
	// Availability
	// unavailable
	// END DEBUGGING
	// CloudStoreAvailabilityUnavailable captures enum value "unavailable"
	CloudStoreAvailabilityUnavailable string = "unavailable"
)

// prop value enum
func (m *CloudStore) validateAvailabilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cloudStoreTypeAvailabilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CloudStore) validateAvailability(formats strfmt.Registry) error {
	if swag.IsZero(m.Availability) { // not required
		return nil
	}

	// value enum
	if err := m.validateAvailabilityEnum("availability", "body", m.Availability); err != nil {
		return err
	}

	return nil
}

func (m *CloudStore) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *CloudStore) validateUnavailableReason(formats strfmt.Registry) error {
	if swag.IsZero(m.UnavailableReason) { // not required
		return nil
	}

	if m.UnavailableReason != nil {
		if err := m.UnavailableReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unavailable_reason")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud store based on the context it is used
func (m *CloudStore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAggregate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAvailability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMirrorDegraded(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnavailableReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStore) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *CloudStore) contextValidateAggregate(ctx context.Context, formats strfmt.Registry) error {

	if m.Aggregate != nil {
		if err := m.Aggregate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate")
			}
			return err
		}
	}

	return nil
}

func (m *CloudStore) contextValidateAvailability(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "availability", "body", string(m.Availability)); err != nil {
		return err
	}

	return nil
}

func (m *CloudStore) contextValidateMirrorDegraded(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mirror_degraded", "body", m.MirrorDegraded); err != nil {
		return err
	}

	return nil
}

func (m *CloudStore) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {
		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *CloudStore) contextValidateUnavailableReason(ctx context.Context, formats strfmt.Registry) error {

	if m.UnavailableReason != nil {
		if err := m.UnavailableReason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unavailable_reason")
			}
			return err
		}
	}

	return nil
}

func (m *CloudStore) contextValidateUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "used", "body", int64(m.Used)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStore) UnmarshalBinary(b []byte) error {
	var res CloudStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudStoreAggregate Aggregate
//
// swagger:model CloudStoreAggregate
type CloudStoreAggregate struct {

	// name
	// Example: aggr1
	Name string `json:"name,omitempty"`
}

// Validate validates this cloud store aggregate
func (m *CloudStoreAggregate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloud store aggregate based on context it is used
func (m *CloudStoreAggregate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudStoreAggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStoreAggregate) UnmarshalBinary(b []byte) error {
	var res CloudStoreAggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudStoreLinks cloud store links
//
// swagger:model CloudStoreLinks
type CloudStoreLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cloud store links
func (m *CloudStoreLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud store links based on the context it is used
func (m *CloudStoreLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudStoreLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStoreLinks) UnmarshalBinary(b []byte) error {
	var res CloudStoreLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudStoreTarget Cloud target
//
// swagger:model CloudStoreTarget
type CloudStoreTarget struct {

	// links
	Links *CloudStoreTargetLinks `json:"_links,omitempty"`

	// name
	// Example: target1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this cloud store target
func (m *CloudStoreTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreTarget) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud store target based on the context it is used
func (m *CloudStoreTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreTarget) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudStoreTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStoreTarget) UnmarshalBinary(b []byte) error {
	var res CloudStoreTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudStoreTargetLinks cloud store target links
//
// swagger:model CloudStoreTargetLinks
type CloudStoreTargetLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cloud store target links
func (m *CloudStoreTargetLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreTargetLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud store target links based on the context it is used
func (m *CloudStoreTargetLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreTargetLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudStoreTargetLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStoreTargetLinks) UnmarshalBinary(b []byte) error {
	var res CloudStoreTargetLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudStoreUnavailableReason cloud store unavailable reason
//
// swagger:model CloudStoreUnavailableReason
type CloudStoreUnavailableReason struct {

	// Indicates why the object store is unavailable.
	// Read Only: true
	Message string `json:"message,omitempty"`
}

// Validate validates this cloud store unavailable reason
func (m *CloudStoreUnavailableReason) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cloud store unavailable reason based on the context it is used
func (m *CloudStoreUnavailableReason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreUnavailableReason) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "unavailable_reason"+"."+"message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudStoreUnavailableReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStoreUnavailableReason) UnmarshalBinary(b []byte) error {
	var res CloudStoreUnavailableReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
