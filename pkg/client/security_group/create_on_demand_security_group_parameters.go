// Code generated by go-swagger; DO NOT EDIT.

package security_group

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

// NewCreateOnDemandSecurityGroupParams creates a new CreateOnDemandSecurityGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOnDemandSecurityGroupParams() *CreateOnDemandSecurityGroupParams {
	return &CreateOnDemandSecurityGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOnDemandSecurityGroupParamsWithTimeout creates a new CreateOnDemandSecurityGroupParams object
// with the ability to set a timeout on a request.
func NewCreateOnDemandSecurityGroupParamsWithTimeout(timeout time.Duration) *CreateOnDemandSecurityGroupParams {
	return &CreateOnDemandSecurityGroupParams{
		timeout: timeout,
	}
}

// NewCreateOnDemandSecurityGroupParamsWithContext creates a new CreateOnDemandSecurityGroupParams object
// with the ability to set a context for a request.
func NewCreateOnDemandSecurityGroupParamsWithContext(ctx context.Context) *CreateOnDemandSecurityGroupParams {
	return &CreateOnDemandSecurityGroupParams{
		Context: ctx,
	}
}

// NewCreateOnDemandSecurityGroupParamsWithHTTPClient creates a new CreateOnDemandSecurityGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOnDemandSecurityGroupParamsWithHTTPClient(client *http.Client) *CreateOnDemandSecurityGroupParams {
	return &CreateOnDemandSecurityGroupParams{
		HTTPClient: client,
	}
}

/*
CreateOnDemandSecurityGroupParams contains all the parameters to send to the API endpoint

	for the create on demand security group operation.

	Typically these are written to a http.Request.
*/
type CreateOnDemandSecurityGroupParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Security group Specification instance
	*/
	Body *models.SecurityGroupSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create on demand security group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOnDemandSecurityGroupParams) WithDefaults() *CreateOnDemandSecurityGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create on demand security group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOnDemandSecurityGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create on demand security group params
func (o *CreateOnDemandSecurityGroupParams) WithTimeout(timeout time.Duration) *CreateOnDemandSecurityGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create on demand security group params
func (o *CreateOnDemandSecurityGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create on demand security group params
func (o *CreateOnDemandSecurityGroupParams) WithContext(ctx context.Context) *CreateOnDemandSecurityGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create on demand security group params
func (o *CreateOnDemandSecurityGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create on demand security group params
func (o *CreateOnDemandSecurityGroupParams) WithHTTPClient(client *http.Client) *CreateOnDemandSecurityGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create on demand security group params
func (o *CreateOnDemandSecurityGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create on demand security group params
func (o *CreateOnDemandSecurityGroupParams) WithAPIVersion(aPIVersion *string) *CreateOnDemandSecurityGroupParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create on demand security group params
func (o *CreateOnDemandSecurityGroupParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create on demand security group params
func (o *CreateOnDemandSecurityGroupParams) WithBody(body *models.SecurityGroupSpecification) *CreateOnDemandSecurityGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create on demand security group params
func (o *CreateOnDemandSecurityGroupParams) SetBody(body *models.SecurityGroupSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOnDemandSecurityGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
