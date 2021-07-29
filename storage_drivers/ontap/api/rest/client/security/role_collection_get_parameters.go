// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewRoleCollectionGetParams creates a new RoleCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRoleCollectionGetParams() *RoleCollectionGetParams {
	return &RoleCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRoleCollectionGetParamsWithTimeout creates a new RoleCollectionGetParams object
// with the ability to set a timeout on a request.
func NewRoleCollectionGetParamsWithTimeout(timeout time.Duration) *RoleCollectionGetParams {
	return &RoleCollectionGetParams{
		timeout: timeout,
	}
}

// NewRoleCollectionGetParamsWithContext creates a new RoleCollectionGetParams object
// with the ability to set a context for a request.
func NewRoleCollectionGetParamsWithContext(ctx context.Context) *RoleCollectionGetParams {
	return &RoleCollectionGetParams{
		Context: ctx,
	}
}

// NewRoleCollectionGetParamsWithHTTPClient creates a new RoleCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewRoleCollectionGetParamsWithHTTPClient(client *http.Client) *RoleCollectionGetParams {
	return &RoleCollectionGetParams{
		HTTPClient: client,
	}
}

/* RoleCollectionGetParams contains all the parameters to send to the API endpoint
   for the role collection get operation.

   Typically these are written to a http.Request.
*/
type RoleCollectionGetParams struct {

	/* Builtin.

	   Filter by builtin
	*/
	BuiltinQueryParameter *bool

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* OwnerName.

	   Filter by owner.name
	*/
	OwnerNameQueryParameter *string

	/* OwnerUUID.

	   Filter by owner.uuid
	*/
	OwnerUUIDQueryParameter *string

	/* PrivilegesAccess.

	   Filter by privileges.access
	*/
	PrivilegesAccessQueryParameter *string

	/* PrivilegesPath.

	   Filter by privileges.path
	*/
	PrivilegesPathQueryParameter *string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* Scope.

	   Filter by scope
	*/
	ScopeQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the role collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RoleCollectionGetParams) WithDefaults() *RoleCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the role collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RoleCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := RoleCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the role collection get params
func (o *RoleCollectionGetParams) WithTimeout(timeout time.Duration) *RoleCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the role collection get params
func (o *RoleCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the role collection get params
func (o *RoleCollectionGetParams) WithContext(ctx context.Context) *RoleCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the role collection get params
func (o *RoleCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the role collection get params
func (o *RoleCollectionGetParams) WithHTTPClient(client *http.Client) *RoleCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the role collection get params
func (o *RoleCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuiltinQueryParameter adds the builtin to the role collection get params
func (o *RoleCollectionGetParams) WithBuiltinQueryParameter(builtin *bool) *RoleCollectionGetParams {
	o.SetBuiltinQueryParameter(builtin)
	return o
}

// SetBuiltinQueryParameter adds the builtin to the role collection get params
func (o *RoleCollectionGetParams) SetBuiltinQueryParameter(builtin *bool) {
	o.BuiltinQueryParameter = builtin
}

// WithFields adds the fields to the role collection get params
func (o *RoleCollectionGetParams) WithFields(fields []string) *RoleCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the role collection get params
func (o *RoleCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the role collection get params
func (o *RoleCollectionGetParams) WithMaxRecords(maxRecords *int64) *RoleCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the role collection get params
func (o *RoleCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNameQueryParameter adds the name to the role collection get params
func (o *RoleCollectionGetParams) WithNameQueryParameter(name *string) *RoleCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the role collection get params
func (o *RoleCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderBy adds the orderBy to the role collection get params
func (o *RoleCollectionGetParams) WithOrderBy(orderBy []string) *RoleCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the role collection get params
func (o *RoleCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithOwnerNameQueryParameter adds the ownerName to the role collection get params
func (o *RoleCollectionGetParams) WithOwnerNameQueryParameter(ownerName *string) *RoleCollectionGetParams {
	o.SetOwnerNameQueryParameter(ownerName)
	return o
}

// SetOwnerNameQueryParameter adds the ownerName to the role collection get params
func (o *RoleCollectionGetParams) SetOwnerNameQueryParameter(ownerName *string) {
	o.OwnerNameQueryParameter = ownerName
}

// WithOwnerUUIDQueryParameter adds the ownerUUID to the role collection get params
func (o *RoleCollectionGetParams) WithOwnerUUIDQueryParameter(ownerUUID *string) *RoleCollectionGetParams {
	o.SetOwnerUUIDQueryParameter(ownerUUID)
	return o
}

// SetOwnerUUIDQueryParameter adds the ownerUuid to the role collection get params
func (o *RoleCollectionGetParams) SetOwnerUUIDQueryParameter(ownerUUID *string) {
	o.OwnerUUIDQueryParameter = ownerUUID
}

// WithPrivilegesAccessQueryParameter adds the privilegesAccess to the role collection get params
func (o *RoleCollectionGetParams) WithPrivilegesAccessQueryParameter(privilegesAccess *string) *RoleCollectionGetParams {
	o.SetPrivilegesAccessQueryParameter(privilegesAccess)
	return o
}

// SetPrivilegesAccessQueryParameter adds the privilegesAccess to the role collection get params
func (o *RoleCollectionGetParams) SetPrivilegesAccessQueryParameter(privilegesAccess *string) {
	o.PrivilegesAccessQueryParameter = privilegesAccess
}

// WithPrivilegesPathQueryParameter adds the privilegesPath to the role collection get params
func (o *RoleCollectionGetParams) WithPrivilegesPathQueryParameter(privilegesPath *string) *RoleCollectionGetParams {
	o.SetPrivilegesPathQueryParameter(privilegesPath)
	return o
}

// SetPrivilegesPathQueryParameter adds the privilegesPath to the role collection get params
func (o *RoleCollectionGetParams) SetPrivilegesPathQueryParameter(privilegesPath *string) {
	o.PrivilegesPathQueryParameter = privilegesPath
}

// WithReturnRecords adds the returnRecords to the role collection get params
func (o *RoleCollectionGetParams) WithReturnRecords(returnRecords *bool) *RoleCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the role collection get params
func (o *RoleCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the role collection get params
func (o *RoleCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *RoleCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the role collection get params
func (o *RoleCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScopeQueryParameter adds the scope to the role collection get params
func (o *RoleCollectionGetParams) WithScopeQueryParameter(scope *string) *RoleCollectionGetParams {
	o.SetScopeQueryParameter(scope)
	return o
}

// SetScopeQueryParameter adds the scope to the role collection get params
func (o *RoleCollectionGetParams) SetScopeQueryParameter(scope *string) {
	o.ScopeQueryParameter = scope
}

// WriteToRequest writes these params to a swagger request
func (o *RoleCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BuiltinQueryParameter != nil {

		// query param builtin
		var qrBuiltin bool

		if o.BuiltinQueryParameter != nil {
			qrBuiltin = *o.BuiltinQueryParameter
		}
		qBuiltin := swag.FormatBool(qrBuiltin)
		if qBuiltin != "" {

			if err := r.SetQueryParam("builtin", qBuiltin); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.OwnerNameQueryParameter != nil {

		// query param owner.name
		var qrOwnerName string

		if o.OwnerNameQueryParameter != nil {
			qrOwnerName = *o.OwnerNameQueryParameter
		}
		qOwnerName := qrOwnerName
		if qOwnerName != "" {

			if err := r.SetQueryParam("owner.name", qOwnerName); err != nil {
				return err
			}
		}
	}

	if o.OwnerUUIDQueryParameter != nil {

		// query param owner.uuid
		var qrOwnerUUID string

		if o.OwnerUUIDQueryParameter != nil {
			qrOwnerUUID = *o.OwnerUUIDQueryParameter
		}
		qOwnerUUID := qrOwnerUUID
		if qOwnerUUID != "" {

			if err := r.SetQueryParam("owner.uuid", qOwnerUUID); err != nil {
				return err
			}
		}
	}

	if o.PrivilegesAccessQueryParameter != nil {

		// query param privileges.access
		var qrPrivilegesAccess string

		if o.PrivilegesAccessQueryParameter != nil {
			qrPrivilegesAccess = *o.PrivilegesAccessQueryParameter
		}
		qPrivilegesAccess := qrPrivilegesAccess
		if qPrivilegesAccess != "" {

			if err := r.SetQueryParam("privileges.access", qPrivilegesAccess); err != nil {
				return err
			}
		}
	}

	if o.PrivilegesPathQueryParameter != nil {

		// query param privileges.path
		var qrPrivilegesPath string

		if o.PrivilegesPathQueryParameter != nil {
			qrPrivilegesPath = *o.PrivilegesPathQueryParameter
		}
		qPrivilegesPath := qrPrivilegesPath
		if qPrivilegesPath != "" {

			if err := r.SetQueryParam("privileges.path", qPrivilegesPath); err != nil {
				return err
			}
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.ScopeQueryParameter != nil {

		// query param scope
		var qrScope string

		if o.ScopeQueryParameter != nil {
			qrScope = *o.ScopeQueryParameter
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamRoleCollectionGet binds the parameter fields
func (o *RoleCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamRoleCollectionGet binds the parameter order_by
func (o *RoleCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}