package cmdb_meta

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
)

// NewCreateProjectParams creates a new CreateProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTableItemParams() *GetTableItemsParams {
	return &GetTableItemsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectParamsWithTimeout creates a new CreateProjectParams object
// with the ability to set a timeout on a request.
func NewGetTableItemParamsWithTimeout(timeout time.Duration) *GetTableItemsParams {
	return &GetTableItemsParams{
		timeout: timeout,
	}
}

// NewCreateProjectParamsWithContext creates a new CreateProjectParams object
// with the ability to set a context for a request.
func NewGetTableItemParamsWithContext(ctx context.Context) *GetTableItemsParams {
	return &GetTableItemsParams{
		Context: ctx,
	}
}

// NewCreateProjectParamsWithHTTPClient creates a new CreateProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTableItemParamsWithHTTPClient(client *http.Client) *GetTableItemsParams {
	return &GetTableItemsParams{
		HTTPClient: client,
	}
}

/* CreateProjectParams contains all the parameters to send to the API endpoint
   for the create project operation.

   Typically these are written to a http.Request.
*/
type GetTableItemsParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Project Specification instance
	*/
	ClassName  string
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTableItemsParams) WithDefaults() *GetTableItemsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTableItemsParams) SetDefaults() {
	// no default values defined for this parameter
}

func (o *GetTableItemsParams) WithClassName(className string) *GetTableItemsParams {
	o.SetClassName(className)
	return o
}

func (o *GetTableItemsParams) SetClassName(className string) {
	o.ClassName = className
}

// WithTimeout adds the timeout to the create project params
func (o *GetTableItemsParams) WithTimeout(timeout time.Duration) *GetTableItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project params
func (o *GetTableItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project params
func (o *GetTableItemsParams) WithContext(ctx context.Context) *GetTableItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project params
func (o *GetTableItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project params
func (o *GetTableItemsParams) WithHTTPClient(client *http.Client) *GetTableItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project params
func (o *GetTableItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create project params
func (o *GetTableItemsParams) WithAPIVersion(aPIVersion *string) *GetTableItemsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create project params
func (o *GetTableItemsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetTableItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
