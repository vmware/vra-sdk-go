// Code generated by go-swagger; DO NOT EDIT.

package image_profile

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

// NewUpdateImageProfileParams creates a new UpdateImageProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateImageProfileParams() *UpdateImageProfileParams {
	return &UpdateImageProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateImageProfileParamsWithTimeout creates a new UpdateImageProfileParams object
// with the ability to set a timeout on a request.
func NewUpdateImageProfileParamsWithTimeout(timeout time.Duration) *UpdateImageProfileParams {
	return &UpdateImageProfileParams{
		timeout: timeout,
	}
}

// NewUpdateImageProfileParamsWithContext creates a new UpdateImageProfileParams object
// with the ability to set a context for a request.
func NewUpdateImageProfileParamsWithContext(ctx context.Context) *UpdateImageProfileParams {
	return &UpdateImageProfileParams{
		Context: ctx,
	}
}

// NewUpdateImageProfileParamsWithHTTPClient creates a new UpdateImageProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateImageProfileParamsWithHTTPClient(client *http.Client) *UpdateImageProfileParams {
	return &UpdateImageProfileParams{
		HTTPClient: client,
	}
}

/*
UpdateImageProfileParams contains all the parameters to send to the API endpoint

	for the update image profile operation.

	Typically these are written to a http.Request.
*/
type UpdateImageProfileParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   ImageProfile instance
	*/
	Body *models.UpdateImageProfileSpecification

	/* ID.

	   The ID of the image.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update image profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateImageProfileParams) WithDefaults() *UpdateImageProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update image profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateImageProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update image profile params
func (o *UpdateImageProfileParams) WithTimeout(timeout time.Duration) *UpdateImageProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update image profile params
func (o *UpdateImageProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update image profile params
func (o *UpdateImageProfileParams) WithContext(ctx context.Context) *UpdateImageProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update image profile params
func (o *UpdateImageProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update image profile params
func (o *UpdateImageProfileParams) WithHTTPClient(client *http.Client) *UpdateImageProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update image profile params
func (o *UpdateImageProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update image profile params
func (o *UpdateImageProfileParams) WithAPIVersion(aPIVersion *string) *UpdateImageProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update image profile params
func (o *UpdateImageProfileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update image profile params
func (o *UpdateImageProfileParams) WithBody(body *models.UpdateImageProfileSpecification) *UpdateImageProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update image profile params
func (o *UpdateImageProfileParams) SetBody(body *models.UpdateImageProfileSpecification) {
	o.Body = body
}

// WithID adds the id to the update image profile params
func (o *UpdateImageProfileParams) WithID(id string) *UpdateImageProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update image profile params
func (o *UpdateImageProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateImageProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
