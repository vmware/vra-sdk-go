// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdateCloudAccountAsyncParams creates a new UpdateCloudAccountAsyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCloudAccountAsyncParams() *UpdateCloudAccountAsyncParams {
	return &UpdateCloudAccountAsyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCloudAccountAsyncParamsWithTimeout creates a new UpdateCloudAccountAsyncParams object
// with the ability to set a timeout on a request.
func NewUpdateCloudAccountAsyncParamsWithTimeout(timeout time.Duration) *UpdateCloudAccountAsyncParams {
	return &UpdateCloudAccountAsyncParams{
		timeout: timeout,
	}
}

// NewUpdateCloudAccountAsyncParamsWithContext creates a new UpdateCloudAccountAsyncParams object
// with the ability to set a context for a request.
func NewUpdateCloudAccountAsyncParamsWithContext(ctx context.Context) *UpdateCloudAccountAsyncParams {
	return &UpdateCloudAccountAsyncParams{
		Context: ctx,
	}
}

// NewUpdateCloudAccountAsyncParamsWithHTTPClient creates a new UpdateCloudAccountAsyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCloudAccountAsyncParamsWithHTTPClient(client *http.Client) *UpdateCloudAccountAsyncParams {
	return &UpdateCloudAccountAsyncParams{
		HTTPClient: client,
	}
}

/* UpdateCloudAccountAsyncParams contains all the parameters to send to the API endpoint
   for the update cloud account async operation.

   Typically these are written to a http.Request.
*/
type UpdateCloudAccountAsyncParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Cloud account details to be updated
	*/
	Body *models.UpdateCloudAccountSpecification

	/* ID.

	   The ID of the cloudAccount
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCloudAccountAsyncParams) WithDefaults() *UpdateCloudAccountAsyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCloudAccountAsyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update cloud account async params
func (o *UpdateCloudAccountAsyncParams) WithTimeout(timeout time.Duration) *UpdateCloudAccountAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cloud account async params
func (o *UpdateCloudAccountAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cloud account async params
func (o *UpdateCloudAccountAsyncParams) WithContext(ctx context.Context) *UpdateCloudAccountAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cloud account async params
func (o *UpdateCloudAccountAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cloud account async params
func (o *UpdateCloudAccountAsyncParams) WithHTTPClient(client *http.Client) *UpdateCloudAccountAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cloud account async params
func (o *UpdateCloudAccountAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update cloud account async params
func (o *UpdateCloudAccountAsyncParams) WithAPIVersion(aPIVersion *string) *UpdateCloudAccountAsyncParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update cloud account async params
func (o *UpdateCloudAccountAsyncParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update cloud account async params
func (o *UpdateCloudAccountAsyncParams) WithBody(body *models.UpdateCloudAccountSpecification) *UpdateCloudAccountAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update cloud account async params
func (o *UpdateCloudAccountAsyncParams) SetBody(body *models.UpdateCloudAccountSpecification) {
	o.Body = body
}

// WithID adds the id to the update cloud account async params
func (o *UpdateCloudAccountAsyncParams) WithID(id string) *UpdateCloudAccountAsyncParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update cloud account async params
func (o *UpdateCloudAccountAsyncParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCloudAccountAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
