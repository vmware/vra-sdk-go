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

// NewUpdateNsxTCloudAccountAsyncParams creates a new UpdateNsxTCloudAccountAsyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNsxTCloudAccountAsyncParams() *UpdateNsxTCloudAccountAsyncParams {
	return &UpdateNsxTCloudAccountAsyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNsxTCloudAccountAsyncParamsWithTimeout creates a new UpdateNsxTCloudAccountAsyncParams object
// with the ability to set a timeout on a request.
func NewUpdateNsxTCloudAccountAsyncParamsWithTimeout(timeout time.Duration) *UpdateNsxTCloudAccountAsyncParams {
	return &UpdateNsxTCloudAccountAsyncParams{
		timeout: timeout,
	}
}

// NewUpdateNsxTCloudAccountAsyncParamsWithContext creates a new UpdateNsxTCloudAccountAsyncParams object
// with the ability to set a context for a request.
func NewUpdateNsxTCloudAccountAsyncParamsWithContext(ctx context.Context) *UpdateNsxTCloudAccountAsyncParams {
	return &UpdateNsxTCloudAccountAsyncParams{
		Context: ctx,
	}
}

// NewUpdateNsxTCloudAccountAsyncParamsWithHTTPClient creates a new UpdateNsxTCloudAccountAsyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNsxTCloudAccountAsyncParamsWithHTTPClient(client *http.Client) *UpdateNsxTCloudAccountAsyncParams {
	return &UpdateNsxTCloudAccountAsyncParams{
		HTTPClient: client,
	}
}

/*
UpdateNsxTCloudAccountAsyncParams contains all the parameters to send to the API endpoint

	for the update nsx t cloud account async operation.

	Typically these are written to a http.Request.
*/
type UpdateNsxTCloudAccountAsyncParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion string

	/* Body.

	   NSX-T cloud account details to be updated
	*/
	Body *models.UpdateCloudAccountNsxTSpecification

	/* ID.

	   Cloud account id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update nsx t cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNsxTCloudAccountAsyncParams) WithDefaults() *UpdateNsxTCloudAccountAsyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update nsx t cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNsxTCloudAccountAsyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update nsx t cloud account async params
func (o *UpdateNsxTCloudAccountAsyncParams) WithTimeout(timeout time.Duration) *UpdateNsxTCloudAccountAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update nsx t cloud account async params
func (o *UpdateNsxTCloudAccountAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update nsx t cloud account async params
func (o *UpdateNsxTCloudAccountAsyncParams) WithContext(ctx context.Context) *UpdateNsxTCloudAccountAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update nsx t cloud account async params
func (o *UpdateNsxTCloudAccountAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update nsx t cloud account async params
func (o *UpdateNsxTCloudAccountAsyncParams) WithHTTPClient(client *http.Client) *UpdateNsxTCloudAccountAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update nsx t cloud account async params
func (o *UpdateNsxTCloudAccountAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update nsx t cloud account async params
func (o *UpdateNsxTCloudAccountAsyncParams) WithAPIVersion(aPIVersion string) *UpdateNsxTCloudAccountAsyncParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update nsx t cloud account async params
func (o *UpdateNsxTCloudAccountAsyncParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update nsx t cloud account async params
func (o *UpdateNsxTCloudAccountAsyncParams) WithBody(body *models.UpdateCloudAccountNsxTSpecification) *UpdateNsxTCloudAccountAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update nsx t cloud account async params
func (o *UpdateNsxTCloudAccountAsyncParams) SetBody(body *models.UpdateCloudAccountNsxTSpecification) {
	o.Body = body
}

// WithID adds the id to the update nsx t cloud account async params
func (o *UpdateNsxTCloudAccountAsyncParams) WithID(id string) *UpdateNsxTCloudAccountAsyncParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update nsx t cloud account async params
func (o *UpdateNsxTCloudAccountAsyncParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNsxTCloudAccountAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {

		if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
			return err
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
