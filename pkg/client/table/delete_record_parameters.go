package table

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteRecordParams Deletes a new DeleteRecordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRecordParams() *DeleteRecordParams {
	return &DeleteRecordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRecordParamsWithTimeout Deletes a new DeleteRecordParams object
// with the ability to set a timeout on a request.
func NewDeleteRecordParamsWithTimeout(timeout time.Duration) *DeleteRecordParams {
	return &DeleteRecordParams{
		timeout: timeout,
	}
}

// NewDeleteRecordParamsWithContext Deletes a new DeleteRecordParams object
// with the ability to set a context for a request.
func NewDeleteRecordParamsWithContext(ctx context.Context) *DeleteRecordParams {
	return &DeleteRecordParams{
		Context: ctx,
	}
}

// NewDeleteRecordParamsWithHTTPClient Deletes a new DeleteRecordParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRecordParamsWithHTTPClient(client *http.Client) *DeleteRecordParams {
	return &DeleteRecordParams{
		HTTPClient: client,
	}
}

/* DeleteRecordParams contains all the parameters to send to the API endpoint
   for the Delete Record operation.

   Typically these are written to a http.Request.
*/
type DeleteRecordParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Record Specification instance
	*/
	TableName  string
	SysId      string
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Delete Record params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRecordParams) WithDefaults() *DeleteRecordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Delete Record params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRecordParams) SetDefaults() {
	// no default values defined for this parameter
}

func (o *DeleteRecordParams) WithTableName(tableName string) *DeleteRecordParams {
	o.SetTableName(tableName)
	return o
}

func (o *DeleteRecordParams) SetTableName(tableName string) {
	o.TableName = tableName
}

func (o *DeleteRecordParams) WithSysID(id string) *DeleteRecordParams {
	o.SetSysID(id)
	return o
}

func (o *DeleteRecordParams) SetSysID(id string) {
	o.SysId = id
}

// WithTimeout adds the timeout to the Delete Record params
func (o *DeleteRecordParams) WithTimeout(timeout time.Duration) *DeleteRecordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Delete Record params
func (o *DeleteRecordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Delete Record params
func (o *DeleteRecordParams) WithContext(ctx context.Context) *DeleteRecordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Delete Record params
func (o *DeleteRecordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Delete Record params
func (o *DeleteRecordParams) WithHTTPClient(client *http.Client) *DeleteRecordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Delete Record params
func (o *DeleteRecordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the Delete Record params
func (o *DeleteRecordParams) WithAPIVersion(aPIVersion *string) *DeleteRecordParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the Delete Record params
func (o *DeleteRecordParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRecordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	r.SetQueryParam("sysparm_query_no_domain", "true")

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string

		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {

			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
		}
	}

	if err := r.SetPathParam("tableName", o.TableName); err != nil {
		return err
	}

	if err := r.SetPathParam("sys_id", o.SysId); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
