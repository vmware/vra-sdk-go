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

// NewCreateNsxTCloudAccountAsyncParams creates a new CreateNsxTCloudAccountAsyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNsxTCloudAccountAsyncParams() *CreateNsxTCloudAccountAsyncParams {
	return &CreateNsxTCloudAccountAsyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNsxTCloudAccountAsyncParamsWithTimeout creates a new CreateNsxTCloudAccountAsyncParams object
// with the ability to set a timeout on a request.
func NewCreateNsxTCloudAccountAsyncParamsWithTimeout(timeout time.Duration) *CreateNsxTCloudAccountAsyncParams {
	return &CreateNsxTCloudAccountAsyncParams{
		timeout: timeout,
	}
}

// NewCreateNsxTCloudAccountAsyncParamsWithContext creates a new CreateNsxTCloudAccountAsyncParams object
// with the ability to set a context for a request.
func NewCreateNsxTCloudAccountAsyncParamsWithContext(ctx context.Context) *CreateNsxTCloudAccountAsyncParams {
	return &CreateNsxTCloudAccountAsyncParams{
		Context: ctx,
	}
}

// NewCreateNsxTCloudAccountAsyncParamsWithHTTPClient creates a new CreateNsxTCloudAccountAsyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNsxTCloudAccountAsyncParamsWithHTTPClient(client *http.Client) *CreateNsxTCloudAccountAsyncParams {
	return &CreateNsxTCloudAccountAsyncParams{
		HTTPClient: client,
	}
}

/*
CreateNsxTCloudAccountAsyncParams contains all the parameters to send to the API endpoint

	for the create nsx t cloud account async operation.

	Typically these are written to a http.Request.
*/
type CreateNsxTCloudAccountAsyncParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion string

	/* Body.

	   Cloud account NsxT specification
	*/
	Body *models.CloudAccountNsxTSpecification

	/* ValidateOnly.

	   If provided, it only validates the credentials in the Cloud Account Specification, and cloud account will not be created.
	*/
	ValidateOnly *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create nsx t cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNsxTCloudAccountAsyncParams) WithDefaults() *CreateNsxTCloudAccountAsyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create nsx t cloud account async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNsxTCloudAccountAsyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create nsx t cloud account async params
func (o *CreateNsxTCloudAccountAsyncParams) WithTimeout(timeout time.Duration) *CreateNsxTCloudAccountAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create nsx t cloud account async params
func (o *CreateNsxTCloudAccountAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create nsx t cloud account async params
func (o *CreateNsxTCloudAccountAsyncParams) WithContext(ctx context.Context) *CreateNsxTCloudAccountAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create nsx t cloud account async params
func (o *CreateNsxTCloudAccountAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create nsx t cloud account async params
func (o *CreateNsxTCloudAccountAsyncParams) WithHTTPClient(client *http.Client) *CreateNsxTCloudAccountAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create nsx t cloud account async params
func (o *CreateNsxTCloudAccountAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create nsx t cloud account async params
func (o *CreateNsxTCloudAccountAsyncParams) WithAPIVersion(aPIVersion string) *CreateNsxTCloudAccountAsyncParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create nsx t cloud account async params
func (o *CreateNsxTCloudAccountAsyncParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create nsx t cloud account async params
func (o *CreateNsxTCloudAccountAsyncParams) WithBody(body *models.CloudAccountNsxTSpecification) *CreateNsxTCloudAccountAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create nsx t cloud account async params
func (o *CreateNsxTCloudAccountAsyncParams) SetBody(body *models.CloudAccountNsxTSpecification) {
	o.Body = body
}

// WithValidateOnly adds the validateOnly to the create nsx t cloud account async params
func (o *CreateNsxTCloudAccountAsyncParams) WithValidateOnly(validateOnly *string) *CreateNsxTCloudAccountAsyncParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the create nsx t cloud account async params
func (o *CreateNsxTCloudAccountAsyncParams) SetValidateOnly(validateOnly *string) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNsxTCloudAccountAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ValidateOnly != nil {

		// query param validateOnly
		var qrValidateOnly string

		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := qrValidateOnly
		if qValidateOnly != "" {

			if err := r.SetQueryParam("validateOnly", qValidateOnly); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
