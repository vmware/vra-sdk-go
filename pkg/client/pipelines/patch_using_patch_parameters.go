// Code generated by go-swagger; DO NOT EDIT.

package pipelines

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

// NewPatchUsingPATCHParams creates a new PatchUsingPATCHParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchUsingPATCHParams() *PatchUsingPATCHParams {
	return &PatchUsingPATCHParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchUsingPATCHParamsWithTimeout creates a new PatchUsingPATCHParams object
// with the ability to set a timeout on a request.
func NewPatchUsingPATCHParamsWithTimeout(timeout time.Duration) *PatchUsingPATCHParams {
	return &PatchUsingPATCHParams{
		timeout: timeout,
	}
}

// NewPatchUsingPATCHParamsWithContext creates a new PatchUsingPATCHParams object
// with the ability to set a context for a request.
func NewPatchUsingPATCHParamsWithContext(ctx context.Context) *PatchUsingPATCHParams {
	return &PatchUsingPATCHParams{
		Context: ctx,
	}
}

// NewPatchUsingPATCHParamsWithHTTPClient creates a new PatchUsingPATCHParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchUsingPATCHParamsWithHTTPClient(client *http.Client) *PatchUsingPATCHParams {
	return &PatchUsingPATCHParams{
		HTTPClient: client,
	}
}

/* PatchUsingPATCHParams contains all the parameters to send to the API endpoint
   for the patch using p a t c h operation.

   Typically these are written to a http.Request.
*/
type PatchUsingPATCHParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Body.

	   Patch Request for a pipeline
	*/
	Body models.PipelinePatchRequest

	/* ID.

	   The ID of the Pipeline
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchUsingPATCHParams) WithDefaults() *PatchUsingPATCHParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchUsingPATCHParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch using p a t c h params
func (o *PatchUsingPATCHParams) WithTimeout(timeout time.Duration) *PatchUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch using p a t c h params
func (o *PatchUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch using p a t c h params
func (o *PatchUsingPATCHParams) WithContext(ctx context.Context) *PatchUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch using p a t c h params
func (o *PatchUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch using p a t c h params
func (o *PatchUsingPATCHParams) WithHTTPClient(client *http.Client) *PatchUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch using p a t c h params
func (o *PatchUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch using p a t c h params
func (o *PatchUsingPATCHParams) WithAuthorization(authorization string) *PatchUsingPATCHParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch using p a t c h params
func (o *PatchUsingPATCHParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the patch using p a t c h params
func (o *PatchUsingPATCHParams) WithAPIVersion(aPIVersion string) *PatchUsingPATCHParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the patch using p a t c h params
func (o *PatchUsingPATCHParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the patch using p a t c h params
func (o *PatchUsingPATCHParams) WithBody(body models.PipelinePatchRequest) *PatchUsingPATCHParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch using p a t c h params
func (o *PatchUsingPATCHParams) SetBody(body models.PipelinePatchRequest) {
	o.Body = body
}

// WithID adds the id to the patch using p a t c h params
func (o *PatchUsingPATCHParams) WithID(id string) *PatchUsingPATCHParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch using p a t c h params
func (o *PatchUsingPATCHParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
