// Code generated by go-swagger; DO NOT EDIT.

package project

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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateUsingPOSTMixin5Params creates a new CreateUsingPOSTMixin5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUsingPOSTMixin5Params() *CreateUsingPOSTMixin5Params {
	return &CreateUsingPOSTMixin5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUsingPOSTMixin5ParamsWithTimeout creates a new CreateUsingPOSTMixin5Params object
// with the ability to set a timeout on a request.
func NewCreateUsingPOSTMixin5ParamsWithTimeout(timeout time.Duration) *CreateUsingPOSTMixin5Params {
	return &CreateUsingPOSTMixin5Params{
		timeout: timeout,
	}
}

// NewCreateUsingPOSTMixin5ParamsWithContext creates a new CreateUsingPOSTMixin5Params object
// with the ability to set a context for a request.
func NewCreateUsingPOSTMixin5ParamsWithContext(ctx context.Context) *CreateUsingPOSTMixin5Params {
	return &CreateUsingPOSTMixin5Params{
		Context: ctx,
	}
}

// NewCreateUsingPOSTMixin5ParamsWithHTTPClient creates a new CreateUsingPOSTMixin5Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUsingPOSTMixin5ParamsWithHTTPClient(client *http.Client) *CreateUsingPOSTMixin5Params {
	return &CreateUsingPOSTMixin5Params{
		HTTPClient: client,
	}
}

/*
CreateUsingPOSTMixin5Params contains all the parameters to send to the API endpoint

	for the create using p o s t mixin5 operation.

	Typically these are written to a http.Request.
*/
type CreateUsingPOSTMixin5Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format. For versioning information refer to /project-service/api/about.
	*/
	APIVersion *string

	/* Request.

	   The project to create.
	*/
	Request models.ProjectSpecification

	/* ValidatePrincipals.

	   If true, a limit of 20 principals is enforced. Additionally each principal is validated in the Identity provider and important rules for group email formats are enforced.
	*/
	ValidatePrincipals *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create using p o s t mixin5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUsingPOSTMixin5Params) WithDefaults() *CreateUsingPOSTMixin5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create using p o s t mixin5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUsingPOSTMixin5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create using p o s t mixin5 params
func (o *CreateUsingPOSTMixin5Params) WithTimeout(timeout time.Duration) *CreateUsingPOSTMixin5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create using p o s t mixin5 params
func (o *CreateUsingPOSTMixin5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create using p o s t mixin5 params
func (o *CreateUsingPOSTMixin5Params) WithContext(ctx context.Context) *CreateUsingPOSTMixin5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create using p o s t mixin5 params
func (o *CreateUsingPOSTMixin5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create using p o s t mixin5 params
func (o *CreateUsingPOSTMixin5Params) WithHTTPClient(client *http.Client) *CreateUsingPOSTMixin5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create using p o s t mixin5 params
func (o *CreateUsingPOSTMixin5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create using p o s t mixin5 params
func (o *CreateUsingPOSTMixin5Params) WithAPIVersion(aPIVersion *string) *CreateUsingPOSTMixin5Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create using p o s t mixin5 params
func (o *CreateUsingPOSTMixin5Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithRequest adds the request to the create using p o s t mixin5 params
func (o *CreateUsingPOSTMixin5Params) WithRequest(request models.ProjectSpecification) *CreateUsingPOSTMixin5Params {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create using p o s t mixin5 params
func (o *CreateUsingPOSTMixin5Params) SetRequest(request models.ProjectSpecification) {
	o.Request = request
}

// WithValidatePrincipals adds the validatePrincipals to the create using p o s t mixin5 params
func (o *CreateUsingPOSTMixin5Params) WithValidatePrincipals(validatePrincipals *bool) *CreateUsingPOSTMixin5Params {
	o.SetValidatePrincipals(validatePrincipals)
	return o
}

// SetValidatePrincipals adds the validatePrincipals to the create using p o s t mixin5 params
func (o *CreateUsingPOSTMixin5Params) SetValidatePrincipals(validatePrincipals *bool) {
	o.ValidatePrincipals = validatePrincipals
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUsingPOSTMixin5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if o.ValidatePrincipals != nil {

		// query param validatePrincipals
		var qrValidatePrincipals bool

		if o.ValidatePrincipals != nil {
			qrValidatePrincipals = *o.ValidatePrincipals
		}
		qValidatePrincipals := swag.FormatBool(qrValidatePrincipals)
		if qValidatePrincipals != "" {

			if err := r.SetQueryParam("validatePrincipals", qValidatePrincipals); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}