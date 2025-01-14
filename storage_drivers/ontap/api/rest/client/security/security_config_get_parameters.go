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

// NewSecurityConfigGetParams creates a new SecurityConfigGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityConfigGetParams() *SecurityConfigGetParams {
	return &SecurityConfigGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityConfigGetParamsWithTimeout creates a new SecurityConfigGetParams object
// with the ability to set a timeout on a request.
func NewSecurityConfigGetParamsWithTimeout(timeout time.Duration) *SecurityConfigGetParams {
	return &SecurityConfigGetParams{
		timeout: timeout,
	}
}

// NewSecurityConfigGetParamsWithContext creates a new SecurityConfigGetParams object
// with the ability to set a context for a request.
func NewSecurityConfigGetParamsWithContext(ctx context.Context) *SecurityConfigGetParams {
	return &SecurityConfigGetParams{
		Context: ctx,
	}
}

// NewSecurityConfigGetParamsWithHTTPClient creates a new SecurityConfigGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityConfigGetParamsWithHTTPClient(client *http.Client) *SecurityConfigGetParams {
	return &SecurityConfigGetParams{
		HTTPClient: client,
	}
}

/* SecurityConfigGetParams contains all the parameters to send to the API endpoint
   for the security config get operation.

   Typically these are written to a http.Request.
*/
type SecurityConfigGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security config get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityConfigGetParams) WithDefaults() *SecurityConfigGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security config get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityConfigGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := SecurityConfigGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security config get params
func (o *SecurityConfigGetParams) WithTimeout(timeout time.Duration) *SecurityConfigGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security config get params
func (o *SecurityConfigGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security config get params
func (o *SecurityConfigGetParams) WithContext(ctx context.Context) *SecurityConfigGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security config get params
func (o *SecurityConfigGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security config get params
func (o *SecurityConfigGetParams) WithHTTPClient(client *http.Client) *SecurityConfigGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security config get params
func (o *SecurityConfigGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the security config get params
func (o *SecurityConfigGetParams) WithFieldsQueryParameter(fields []string) *SecurityConfigGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the security config get params
func (o *SecurityConfigGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the security config get params
func (o *SecurityConfigGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *SecurityConfigGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the security config get params
func (o *SecurityConfigGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the security config get params
func (o *SecurityConfigGetParams) WithOrderByQueryParameter(orderBy []string) *SecurityConfigGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the security config get params
func (o *SecurityConfigGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the security config get params
func (o *SecurityConfigGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *SecurityConfigGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the security config get params
func (o *SecurityConfigGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the security config get params
func (o *SecurityConfigGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SecurityConfigGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the security config get params
func (o *SecurityConfigGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityConfigGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSecurityConfigGet binds the parameter fields
func (o *SecurityConfigGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamSecurityConfigGet binds the parameter order_by
func (o *SecurityConfigGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
