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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewModifyExecutionUsingPOSTParams creates a new ModifyExecutionUsingPOSTParams object
// with the default values initialized.
func NewModifyExecutionUsingPOSTParams() *ModifyExecutionUsingPOSTParams {
	var ()
	return &ModifyExecutionUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyExecutionUsingPOSTParamsWithTimeout creates a new ModifyExecutionUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyExecutionUsingPOSTParamsWithTimeout(timeout time.Duration) *ModifyExecutionUsingPOSTParams {
	var ()
	return &ModifyExecutionUsingPOSTParams{

		timeout: timeout,
	}
}

// NewModifyExecutionUsingPOSTParamsWithContext creates a new ModifyExecutionUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyExecutionUsingPOSTParamsWithContext(ctx context.Context) *ModifyExecutionUsingPOSTParams {
	var ()
	return &ModifyExecutionUsingPOSTParams{

		Context: ctx,
	}
}

// NewModifyExecutionUsingPOSTParamsWithHTTPClient creates a new ModifyExecutionUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyExecutionUsingPOSTParamsWithHTTPClient(client *http.Client) *ModifyExecutionUsingPOSTParams {
	var ()
	return &ModifyExecutionUsingPOSTParams{
		HTTPClient: client,
	}
}

/*ModifyExecutionUsingPOSTParams contains all the parameters to send to the API endpoint
for the modify execution using p o s t operation typically these are written to a http.Request
*/
type ModifyExecutionUsingPOSTParams struct {

	/*Action
	  Action to perform on the Execution. Can be any of pause, resume, cancel and tag

	*/
	Action string
	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*Body
	  Execution Request specification

	*/
	Body models.ExecutionActionRequest
	/*ID
	  The ID of the Execution

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) WithTimeout(timeout time.Duration) *ModifyExecutionUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) WithContext(ctx context.Context) *ModifyExecutionUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) WithHTTPClient(client *http.Client) *ModifyExecutionUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) WithAction(action string) *ModifyExecutionUsingPOSTParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) SetAction(action string) {
	o.Action = action
}

// WithAPIVersion adds the aPIVersion to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) WithAPIVersion(aPIVersion string) *ModifyExecutionUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) WithBody(body models.ExecutionActionRequest) *ModifyExecutionUsingPOSTParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) SetBody(body models.ExecutionActionRequest) {
	o.Body = body
}

// WithID adds the id to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) WithID(id string) *ModifyExecutionUsingPOSTParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the modify execution using p o s t params
func (o *ModifyExecutionUsingPOSTParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyExecutionUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param action
	qrAction := o.Action
	qAction := qrAction
	if qAction != "" {
		if err := r.SetQueryParam("action", qAction); err != nil {
			return err
		}
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
