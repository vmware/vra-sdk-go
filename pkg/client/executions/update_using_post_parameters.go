// Code generated by go-swagger; DO NOT EDIT.

package executions

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

// NewUpdateUsingPOSTParams creates a new UpdateUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUsingPOSTParams() *UpdateUsingPOSTParams {
	return &UpdateUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUsingPOSTParamsWithTimeout creates a new UpdateUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewUpdateUsingPOSTParamsWithTimeout(timeout time.Duration) *UpdateUsingPOSTParams {
	return &UpdateUsingPOSTParams{
		timeout: timeout,
	}
}

// NewUpdateUsingPOSTParamsWithContext creates a new UpdateUsingPOSTParams object
// with the ability to set a context for a request.
func NewUpdateUsingPOSTParamsWithContext(ctx context.Context) *UpdateUsingPOSTParams {
	return &UpdateUsingPOSTParams{
		Context: ctx,
	}
}

// NewUpdateUsingPOSTParamsWithHTTPClient creates a new UpdateUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUsingPOSTParamsWithHTTPClient(client *http.Client) *UpdateUsingPOSTParams {
	return &UpdateUsingPOSTParams{
		HTTPClient: client,
	}
}

/* UpdateUsingPOSTParams contains all the parameters to send to the API endpoint
   for the update using p o s t operation.

   Typically these are written to a http.Request.
*/
type UpdateUsingPOSTParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Request.

	   request
	*/
	Request models.BatchUserOperationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUsingPOSTParams) WithDefaults() *UpdateUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update using p o s t params
func (o *UpdateUsingPOSTParams) WithTimeout(timeout time.Duration) *UpdateUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update using p o s t params
func (o *UpdateUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update using p o s t params
func (o *UpdateUsingPOSTParams) WithContext(ctx context.Context) *UpdateUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update using p o s t params
func (o *UpdateUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update using p o s t params
func (o *UpdateUsingPOSTParams) WithHTTPClient(client *http.Client) *UpdateUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update using p o s t params
func (o *UpdateUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the update using p o s t params
func (o *UpdateUsingPOSTParams) WithAuthorization(authorization string) *UpdateUsingPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the update using p o s t params
func (o *UpdateUsingPOSTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the update using p o s t params
func (o *UpdateUsingPOSTParams) WithAPIVersion(aPIVersion string) *UpdateUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update using p o s t params
func (o *UpdateUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithRequest adds the request to the update using p o s t params
func (o *UpdateUsingPOSTParams) WithRequest(request models.BatchUserOperationRequest) *UpdateUsingPOSTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update using p o s t params
func (o *UpdateUsingPOSTParams) SetRequest(request models.BatchUserOperationRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
