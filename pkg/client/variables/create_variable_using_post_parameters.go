// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// NewCreateVariableUsingPOSTParams creates a new CreateVariableUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVariableUsingPOSTParams() *CreateVariableUsingPOSTParams {
	return &CreateVariableUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVariableUsingPOSTParamsWithTimeout creates a new CreateVariableUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreateVariableUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateVariableUsingPOSTParams {
	return &CreateVariableUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreateVariableUsingPOSTParamsWithContext creates a new CreateVariableUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreateVariableUsingPOSTParamsWithContext(ctx context.Context) *CreateVariableUsingPOSTParams {
	return &CreateVariableUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreateVariableUsingPOSTParamsWithHTTPClient creates a new CreateVariableUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVariableUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateVariableUsingPOSTParams {
	return &CreateVariableUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
CreateVariableUsingPOSTParams contains all the parameters to send to the API endpoint

	for the create variable using p o s t operation.

	Typically these are written to a http.Request.
*/
type CreateVariableUsingPOSTParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Body.

	   Variable specification
	*/
	Body models.VariableSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create variable using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVariableUsingPOSTParams) WithDefaults() *CreateVariableUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create variable using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVariableUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create variable using p o s t params
func (o *CreateVariableUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateVariableUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create variable using p o s t params
func (o *CreateVariableUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create variable using p o s t params
func (o *CreateVariableUsingPOSTParams) WithContext(ctx context.Context) *CreateVariableUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create variable using p o s t params
func (o *CreateVariableUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create variable using p o s t params
func (o *CreateVariableUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateVariableUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create variable using p o s t params
func (o *CreateVariableUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create variable using p o s t params
func (o *CreateVariableUsingPOSTParams) WithAuthorization(authorization string) *CreateVariableUsingPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create variable using p o s t params
func (o *CreateVariableUsingPOSTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the create variable using p o s t params
func (o *CreateVariableUsingPOSTParams) WithAPIVersion(aPIVersion string) *CreateVariableUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create variable using p o s t params
func (o *CreateVariableUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create variable using p o s t params
func (o *CreateVariableUsingPOSTParams) WithBody(body models.VariableSpec) *CreateVariableUsingPOSTParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create variable using p o s t params
func (o *CreateVariableUsingPOSTParams) SetBody(body models.VariableSpec) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVariableUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion

	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
