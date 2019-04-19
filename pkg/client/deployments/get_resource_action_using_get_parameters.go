// Code generated by go-swagger; DO NOT EDIT.

package deployments

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

// NewGetResourceActionUsingGETParams creates a new GetResourceActionUsingGETParams object
// with the default values initialized.
func NewGetResourceActionUsingGETParams() *GetResourceActionUsingGETParams {
	var ()
	return &GetResourceActionUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceActionUsingGETParamsWithTimeout creates a new GetResourceActionUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResourceActionUsingGETParamsWithTimeout(timeout time.Duration) *GetResourceActionUsingGETParams {
	var ()
	return &GetResourceActionUsingGETParams{

		timeout: timeout,
	}
}

// NewGetResourceActionUsingGETParamsWithContext creates a new GetResourceActionUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetResourceActionUsingGETParamsWithContext(ctx context.Context) *GetResourceActionUsingGETParams {
	var ()
	return &GetResourceActionUsingGETParams{

		Context: ctx,
	}
}

// NewGetResourceActionUsingGETParamsWithHTTPClient creates a new GetResourceActionUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetResourceActionUsingGETParamsWithHTTPClient(client *http.Client) *GetResourceActionUsingGETParams {
	var ()
	return &GetResourceActionUsingGETParams{
		HTTPClient: client,
	}
}

/*GetResourceActionUsingGETParams contains all the parameters to send to the API endpoint
for the get resource action using g e t operation typically these are written to a http.Request
*/
type GetResourceActionUsingGETParams struct {

	/*ActionID
	  Action ID

	*/
	ActionID strfmt.UUID
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

// WithTimeout adds the timeout to the get resource action using get params
func (o *GetResourceActionUsingGETParams) WithTimeout(timeout time.Duration) *GetResourceActionUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource action using get params
func (o *GetResourceActionUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource action using get params
func (o *GetResourceActionUsingGETParams) WithContext(ctx context.Context) *GetResourceActionUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource action using get params
func (o *GetResourceActionUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource action using get params
func (o *GetResourceActionUsingGETParams) WithHTTPClient(client *http.Client) *GetResourceActionUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource action using get params
func (o *GetResourceActionUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the get resource action using get params
func (o *GetResourceActionUsingGETParams) WithActionID(actionID strfmt.UUID) *GetResourceActionUsingGETParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the get resource action using get params
func (o *GetResourceActionUsingGETParams) SetActionID(actionID strfmt.UUID) {
	o.ActionID = actionID
}

// WithDepID adds the depID to the get resource action using get params
func (o *GetResourceActionUsingGETParams) WithDepID(depID strfmt.UUID) *GetResourceActionUsingGETParams {
	o.SetDepID(depID)
	return o
}

// SetDepID adds the depId to the get resource action using get params
func (o *GetResourceActionUsingGETParams) SetDepID(depID strfmt.UUID) {
	o.DepID = depID
}

// WithResourceID adds the resourceID to the get resource action using get params
func (o *GetResourceActionUsingGETParams) WithResourceID(resourceID strfmt.UUID) *GetResourceActionUsingGETParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get resource action using get params
func (o *GetResourceActionUsingGETParams) SetResourceID(resourceID strfmt.UUID) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceActionUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID.String()); err != nil {
		return err
	}

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
