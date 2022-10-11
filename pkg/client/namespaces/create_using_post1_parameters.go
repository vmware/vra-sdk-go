// Code generated by go-swagger; DO NOT EDIT.

package namespaces

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

// NewCreateUsingPOST1Params creates a new CreateUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUsingPOST1Params() *CreateUsingPOST1Params {
	return &CreateUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUsingPOST1ParamsWithTimeout creates a new CreateUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewCreateUsingPOST1ParamsWithTimeout(timeout time.Duration) *CreateUsingPOST1Params {
	return &CreateUsingPOST1Params{
		timeout: timeout,
	}
}

// NewCreateUsingPOST1ParamsWithContext creates a new CreateUsingPOST1Params object
// with the ability to set a context for a request.
func NewCreateUsingPOST1ParamsWithContext(ctx context.Context) *CreateUsingPOST1Params {
	return &CreateUsingPOST1Params{
		Context: ctx,
	}
}

// NewCreateUsingPOST1ParamsWithHTTPClient creates a new CreateUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUsingPOST1ParamsWithHTTPClient(client *http.Client) *CreateUsingPOST1Params {
	return &CreateUsingPOST1Params{
		HTTPClient: client,
	}
}

/*
CreateUsingPOST1Params contains all the parameters to send to the API endpoint

	for the create using p o s t 1 operation.

	Typically these are written to a http.Request.
*/
type CreateUsingPOST1Params struct {

	/* Namespace.

	   namespace
	*/
	Namespace *models.K8SNamespace

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUsingPOST1Params) WithDefaults() *CreateUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) WithTimeout(timeout time.Duration) *CreateUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) WithContext(ctx context.Context) *CreateUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) WithHTTPClient(client *http.Client) *CreateUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) WithNamespace(namespace *models.K8SNamespace) *CreateUsingPOST1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) SetNamespace(namespace *models.K8SNamespace) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Namespace != nil {
		if err := r.SetBodyParam(o.Namespace); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}