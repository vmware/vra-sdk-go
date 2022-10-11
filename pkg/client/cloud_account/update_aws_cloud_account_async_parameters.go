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

// NewUpdateAWSCloudAccountAsyncParams creates a new UpdateAWSCloudAccountAsyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAWSCloudAccountAsyncParams() *UpdateAWSCloudAccountAsyncParams {
	return &UpdateAWSCloudAccountAsyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAWSCloudAccountAsyncParamsWithTimeout creates a new UpdateAWSCloudAccountAsyncParams object
// with the ability to set a timeout on a request.
func NewUpdateAWSCloudAccountAsyncParamsWithTimeout(timeout time.Duration) *UpdateAWSCloudAccountAsyncParams {
	return &UpdateAWSCloudAccountAsyncParams{
		timeout: timeout,
	}
}

// NewUpdateAWSCloudAccountAsyncParamsWithContext creates a new UpdateAWSCloudAccountAsyncParams object
// with the ability to set a context for a request.
func NewUpdateAWSCloudAccountAsyncParamsWithContext(ctx context.Context) *UpdateAWSCloudAccountAsyncParams {
	return &UpdateAWSCloudAccountAsyncParams{
		Context: ctx,
	}
}

// NewUpdateAWSCloudAccountAsyncParamsWithHTTPClient creates a new UpdateAWSCloudAccountAsyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAWSCloudAccountAsyncParamsWithHTTPClient(client *http.Client) *UpdateAWSCloudAccountAsyncParams {
	return &UpdateAWSCloudAccountAsyncParams{
		HTTPClient: client,
	}
}

/*
UpdateAWSCloudAccountAsyncParams contains all the parameters to send to the API endpoint

	for the update a w s cloud account async operation.

	Typically these are written to a http.Request.
*/
type UpdateAWSCloudAccountAsyncParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   AWS cloud account details to be updated
	*/
	Body *models.UpdateCloudAccountAwsSpecification

	/* ID.

	   Cloud account id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update a w s cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAWSCloudAccountAsyncParams) WithDefaults() *UpdateAWSCloudAccountAsyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update a w s cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAWSCloudAccountAsyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update a w s cloud account async params
func (o *UpdateAWSCloudAccountAsyncParams) WithTimeout(timeout time.Duration) *UpdateAWSCloudAccountAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update a w s cloud account async params
func (o *UpdateAWSCloudAccountAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update a w s cloud account async params
func (o *UpdateAWSCloudAccountAsyncParams) WithContext(ctx context.Context) *UpdateAWSCloudAccountAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update a w s cloud account async params
func (o *UpdateAWSCloudAccountAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update a w s cloud account async params
func (o *UpdateAWSCloudAccountAsyncParams) WithHTTPClient(client *http.Client) *UpdateAWSCloudAccountAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update a w s cloud account async params
func (o *UpdateAWSCloudAccountAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update a w s cloud account async params
func (o *UpdateAWSCloudAccountAsyncParams) WithAPIVersion(aPIVersion *string) *UpdateAWSCloudAccountAsyncParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update a w s cloud account async params
func (o *UpdateAWSCloudAccountAsyncParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update a w s cloud account async params
func (o *UpdateAWSCloudAccountAsyncParams) WithBody(body *models.UpdateCloudAccountAwsSpecification) *UpdateAWSCloudAccountAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update a w s cloud account async params
func (o *UpdateAWSCloudAccountAsyncParams) SetBody(body *models.UpdateCloudAccountAwsSpecification) {
	o.Body = body
}

// WithID adds the id to the update a w s cloud account async params
func (o *UpdateAWSCloudAccountAsyncParams) WithID(id string) *UpdateAWSCloudAccountAsyncParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update a w s cloud account async params
func (o *UpdateAWSCloudAccountAsyncParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAWSCloudAccountAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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