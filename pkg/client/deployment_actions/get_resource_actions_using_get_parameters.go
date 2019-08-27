// Code generated by go-swagger; DO NOT EDIT.

package deployment_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetResourceActionsUsingGETParams creates a new GetResourceActionsUsingGETParams object
// with the default values initialized.
func NewGetResourceActionsUsingGETParams() *GetResourceActionsUsingGETParams {
	var ()
	return &GetResourceActionsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceActionsUsingGETParamsWithTimeout creates a new GetResourceActionsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResourceActionsUsingGETParamsWithTimeout(timeout time.Duration) *GetResourceActionsUsingGETParams {
	var ()
	return &GetResourceActionsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetResourceActionsUsingGETParamsWithContext creates a new GetResourceActionsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetResourceActionsUsingGETParamsWithContext(ctx context.Context) *GetResourceActionsUsingGETParams {
	var ()
	return &GetResourceActionsUsingGETParams{

		Context: ctx,
	}
}

// NewGetResourceActionsUsingGETParamsWithHTTPClient creates a new GetResourceActionsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetResourceActionsUsingGETParamsWithHTTPClient(client *http.Client) *GetResourceActionsUsingGETParams {
	var ()
	return &GetResourceActionsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetResourceActionsUsingGETParams contains all the parameters to send to the API endpoint
for the get resource actions using g e t operation typically these are written to a http.Request
*/
type GetResourceActionsUsingGETParams struct {

	/*DepID
	  Deployment ID

	*/
	DepID strfmt.UUID
	/*ResourceID
	  Resource ID

	*/
	ResourceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) WithTimeout(timeout time.Duration) *GetResourceActionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) WithContext(ctx context.Context) *GetResourceActionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) WithHTTPClient(client *http.Client) *GetResourceActionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDepID adds the depID to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) WithDepID(depID strfmt.UUID) *GetResourceActionsUsingGETParams {
	o.SetDepID(depID)
	return o
}

// SetDepID adds the depId to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) SetDepID(depID strfmt.UUID) {
	o.DepID = depID
}

// WithResourceID adds the resourceID to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) WithResourceID(resourceID strfmt.UUID) *GetResourceActionsUsingGETParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) SetResourceID(resourceID strfmt.UUID) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceActionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param depId
	if err := r.SetPathParam("depId", o.DepID.String()); err != nil {
		return err
	}

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}