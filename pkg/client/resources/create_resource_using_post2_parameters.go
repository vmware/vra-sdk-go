// Code generated by go-swagger; DO NOT EDIT.

package resources

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

// NewCreateResourceUsingPOST2Params creates a new CreateResourceUsingPOST2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateResourceUsingPOST2Params() *CreateResourceUsingPOST2Params {
	return &CreateResourceUsingPOST2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateResourceUsingPOST2ParamsWithTimeout creates a new CreateResourceUsingPOST2Params object
// with the ability to set a timeout on a request.
func NewCreateResourceUsingPOST2ParamsWithTimeout(timeout time.Duration) *CreateResourceUsingPOST2Params {
	return &CreateResourceUsingPOST2Params{
		timeout: timeout,
	}
}

// NewCreateResourceUsingPOST2ParamsWithContext creates a new CreateResourceUsingPOST2Params object
// with the ability to set a context for a request.
func NewCreateResourceUsingPOST2ParamsWithContext(ctx context.Context) *CreateResourceUsingPOST2Params {
	return &CreateResourceUsingPOST2Params{
		Context: ctx,
	}
}

// NewCreateResourceUsingPOST2ParamsWithHTTPClient creates a new CreateResourceUsingPOST2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateResourceUsingPOST2ParamsWithHTTPClient(client *http.Client) *CreateResourceUsingPOST2Params {
	return &CreateResourceUsingPOST2Params{
		HTTPClient: client,
	}
}

/*
CreateResourceUsingPOST2Params contains all the parameters to send to the API endpoint

	for the create resource using p o s t 2 operation.

	Typically these are written to a http.Request.
*/
type CreateResourceUsingPOST2Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* Resource.

	   The resource spec to create a resource.
	*/
	Resource *models.ResourceSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create resource using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateResourceUsingPOST2Params) WithDefaults() *CreateResourceUsingPOST2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create resource using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateResourceUsingPOST2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create resource using p o s t 2 params
func (o *CreateResourceUsingPOST2Params) WithTimeout(timeout time.Duration) *CreateResourceUsingPOST2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create resource using p o s t 2 params
func (o *CreateResourceUsingPOST2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create resource using p o s t 2 params
func (o *CreateResourceUsingPOST2Params) WithContext(ctx context.Context) *CreateResourceUsingPOST2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create resource using p o s t 2 params
func (o *CreateResourceUsingPOST2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create resource using p o s t 2 params
func (o *CreateResourceUsingPOST2Params) WithHTTPClient(client *http.Client) *CreateResourceUsingPOST2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create resource using p o s t 2 params
func (o *CreateResourceUsingPOST2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create resource using p o s t 2 params
func (o *CreateResourceUsingPOST2Params) WithAPIVersion(aPIVersion *string) *CreateResourceUsingPOST2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create resource using p o s t 2 params
func (o *CreateResourceUsingPOST2Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithResource adds the resource to the create resource using p o s t 2 params
func (o *CreateResourceUsingPOST2Params) WithResource(resource *models.ResourceSpecification) *CreateResourceUsingPOST2Params {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the create resource using p o s t 2 params
func (o *CreateResourceUsingPOST2Params) SetResource(resource *models.ResourceSpecification) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *CreateResourceUsingPOST2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Resource != nil {
		if err := r.SetBodyParam(o.Resource); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
