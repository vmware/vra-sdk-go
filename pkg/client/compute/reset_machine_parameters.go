// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewResetMachineParams creates a new ResetMachineParams object
// with the default values initialized.
func NewResetMachineParams() *ResetMachineParams {
	var ()
	return &ResetMachineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResetMachineParamsWithTimeout creates a new ResetMachineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResetMachineParamsWithTimeout(timeout time.Duration) *ResetMachineParams {
	var ()
	return &ResetMachineParams{

		timeout: timeout,
	}
}

// NewResetMachineParamsWithContext creates a new ResetMachineParams object
// with the default values initialized, and the ability to set a context for a request
func NewResetMachineParamsWithContext(ctx context.Context) *ResetMachineParams {
	var ()
	return &ResetMachineParams{

		Context: ctx,
	}
}

// NewResetMachineParamsWithHTTPClient creates a new ResetMachineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResetMachineParamsWithHTTPClient(client *http.Client) *ResetMachineParams {
	var ()
	return &ResetMachineParams{
		HTTPClient: client,
	}
}

/*ResetMachineParams contains all the parameters to send to the API endpoint
for the reset machine operation typically these are written to a http.Request
*/
type ResetMachineParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The id of the Machine.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reset machine params
func (o *ResetMachineParams) WithTimeout(timeout time.Duration) *ResetMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset machine params
func (o *ResetMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset machine params
func (o *ResetMachineParams) WithContext(ctx context.Context) *ResetMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset machine params
func (o *ResetMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset machine params
func (o *ResetMachineParams) WithHTTPClient(client *http.Client) *ResetMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset machine params
func (o *ResetMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the reset machine params
func (o *ResetMachineParams) WithAPIVersion(aPIVersion *string) *ResetMachineParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the reset machine params
func (o *ResetMachineParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the reset machine params
func (o *ResetMachineParams) WithID(id string) *ResetMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the reset machine params
func (o *ResetMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResetMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
