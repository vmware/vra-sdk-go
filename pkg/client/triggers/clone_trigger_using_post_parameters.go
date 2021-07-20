// Code generated by go-swagger; DO NOT EDIT.

package triggers

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

// NewCloneTriggerUsingPOSTParams creates a new CloneTriggerUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloneTriggerUsingPOSTParams() *CloneTriggerUsingPOSTParams {
	return &CloneTriggerUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloneTriggerUsingPOSTParamsWithTimeout creates a new CloneTriggerUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCloneTriggerUsingPOSTParamsWithTimeout(timeout time.Duration) *CloneTriggerUsingPOSTParams {
	return &CloneTriggerUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCloneTriggerUsingPOSTParamsWithContext creates a new CloneTriggerUsingPOSTParams object
// with the ability to set a context for a request.
func NewCloneTriggerUsingPOSTParamsWithContext(ctx context.Context) *CloneTriggerUsingPOSTParams {
	return &CloneTriggerUsingPOSTParams{
		Context: ctx,
	}
}

// NewCloneTriggerUsingPOSTParamsWithHTTPClient creates a new CloneTriggerUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloneTriggerUsingPOSTParamsWithHTTPClient(client *http.Client) *CloneTriggerUsingPOSTParams {
	return &CloneTriggerUsingPOSTParams{
		HTTPClient: client,
	}
}

/* CloneTriggerUsingPOSTParams contains all the parameters to send to the API endpoint
   for the clone trigger using p o s t operation.

   Typically these are written to a http.Request.
*/
type CloneTriggerUsingPOSTParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   id
	*/
	ID string

	/* ServiceRequest.

	   serviceRequest
	*/
	ServiceRequest models.ServiceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the clone trigger using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneTriggerUsingPOSTParams) WithDefaults() *CloneTriggerUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the clone trigger using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneTriggerUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) WithTimeout(timeout time.Duration) *CloneTriggerUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) WithContext(ctx context.Context) *CloneTriggerUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) WithHTTPClient(client *http.Client) *CloneTriggerUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) WithAuthorization(authorization string) *CloneTriggerUsingPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) WithAPIVersion(aPIVersion string) *CloneTriggerUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) WithID(id string) *CloneTriggerUsingPOSTParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) SetID(id string) {
	o.ID = id
}

// WithServiceRequest adds the serviceRequest to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) WithServiceRequest(serviceRequest models.ServiceRequest) *CloneTriggerUsingPOSTParams {
	o.SetServiceRequest(serviceRequest)
	return o
}

// SetServiceRequest adds the serviceRequest to the clone trigger using p o s t params
func (o *CloneTriggerUsingPOSTParams) SetServiceRequest(serviceRequest models.ServiceRequest) {
	o.ServiceRequest = serviceRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CloneTriggerUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.ServiceRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
