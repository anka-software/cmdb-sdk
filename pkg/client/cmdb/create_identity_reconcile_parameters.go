package cmdb

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/anka-software/cmdb-sdk/pkg/models"
)

// NewCreateProjectParams creates a new CreateProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateIdentifyReconcileParams() *CreateIdentifyReconcileParams {
	return &CreateIdentifyReconcileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectParamsWithTimeout creates a new CreateProjectParams object
// with the ability to set a timeout on a request.
func NewCreateIdentifyReconcileParamsWithTimeout(timeout time.Duration) *CreateIdentifyReconcileParams {
	return &CreateIdentifyReconcileParams{
		timeout: timeout,
	}
}

// NewCreateProjectParamsWithContext creates a new CreateProjectParams object
// with the ability to set a context for a request.
func NewCreateIdentifyReconcileParamsWithContext(ctx context.Context) *CreateIdentifyReconcileParams {
	return &CreateIdentifyReconcileParams{
		Context: ctx,
	}
}

// NewCreateProjectParamsWithHTTPClient creates a new CreateProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateIdentifyReconcilearamsWithHTTPClient(client *http.Client) *CreateIdentifyReconcileParams {
	return &CreateIdentifyReconcileParams{
		HTTPClient: client,
	}
}

/* CreateProjectParams contains all the parameters to send to the API endpoint
   for the create project operation.

   Typically these are written to a http.Request.
*/
type CreateIdentifyReconcileParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Project Specification instance
	*/
	Body *models.IdentifyReconcileItemList

	SysparmDataSource *string
	timeout           time.Duration
	Context           context.Context
	HTTPClient        *http.Client
}

// WithDefaults hydrates default values in the create project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIdentifyReconcileParams) WithDefaults() *CreateIdentifyReconcileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIdentifyReconcileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create project params
func (o *CreateIdentifyReconcileParams) WithTimeout(timeout time.Duration) *CreateIdentifyReconcileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project params
func (o *CreateIdentifyReconcileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project params
func (o *CreateIdentifyReconcileParams) WithContext(ctx context.Context) *CreateIdentifyReconcileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project params
func (o *CreateIdentifyReconcileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project params
func (o *CreateIdentifyReconcileParams) WithHTTPClient(client *http.Client) *CreateIdentifyReconcileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project params
func (o *CreateIdentifyReconcileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create project params
func (o *CreateIdentifyReconcileParams) WithAPIVersion(aPIVersion *string) *CreateIdentifyReconcileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create project params
func (o *CreateIdentifyReconcileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

func (o *CreateIdentifyReconcileParams) WithSysParamDataSource(dataSource *string) *CreateIdentifyReconcileParams {
	o.SetSysParamDataSource(dataSource)
	return o
}

// SetAPIVersion adds the apiVersion to the create project params
func (o *CreateIdentifyReconcileParams) SetSysParamDataSource(dataSource *string) {
	o.SysparmDataSource = dataSource
}

// WithBody adds the body to the create project params
func (o *CreateIdentifyReconcileParams) WithBody(body *models.IdentifyReconcileItemList) *CreateIdentifyReconcileParams {
	o.SetBody(body)
	return o
}

func (o *CreateIdentifyReconcileParams) SetBody(body *models.IdentifyReconcileItemList) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIdentifyReconcileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.SysparmDataSource != nil {

		if err := r.SetQueryParam("sysparm_data_source", *o.SysparmDataSource); err != nil {
			return err
		}

	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
