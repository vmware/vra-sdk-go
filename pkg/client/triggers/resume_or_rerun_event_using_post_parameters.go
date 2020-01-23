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

	strfmt "github.com/go-openapi/strfmt"
)

// NewResumeOrRerunEventUsingPOSTParams creates a new ResumeOrRerunEventUsingPOSTParams object
// with the default values initialized.
func NewResumeOrRerunEventUsingPOSTParams() *ResumeOrRerunEventUsingPOSTParams {
	var ()
	return &ResumeOrRerunEventUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResumeOrRerunEventUsingPOSTParamsWithTimeout creates a new ResumeOrRerunEventUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResumeOrRerunEventUsingPOSTParamsWithTimeout(timeout time.Duration) *ResumeOrRerunEventUsingPOSTParams {
	var ()
	return &ResumeOrRerunEventUsingPOSTParams{

		timeout: timeout,
	}
}

// NewResumeOrRerunEventUsingPOSTParamsWithContext creates a new ResumeOrRerunEventUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewResumeOrRerunEventUsingPOSTParamsWithContext(ctx context.Context) *ResumeOrRerunEventUsingPOSTParams {
	var ()
	return &ResumeOrRerunEventUsingPOSTParams{

		Context: ctx,
	}
}

// NewResumeOrRerunEventUsingPOSTParamsWithHTTPClient creates a new ResumeOrRerunEventUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResumeOrRerunEventUsingPOSTParamsWithHTTPClient(client *http.Client) *ResumeOrRerunEventUsingPOSTParams {
	var ()
	return &ResumeOrRerunEventUsingPOSTParams{
		HTTPClient: client,
	}
}

/*ResumeOrRerunEventUsingPOSTParams contains all the parameters to send to the API endpoint
for the resume or rerun event using p o s t operation typically these are written to a http.Request
*/
type ResumeOrRerunEventUsingPOSTParams struct {

	/*Action
	  Resume/Rerun

	*/
	Action string
	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*ID
	  id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resume or rerun event using p o s t params
func (o *ResumeOrRerunEventUsingPOSTParams) WithTimeout(timeout time.Duration) *ResumeOrRerunEventUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resume or rerun event using p o s t params
func (o *ResumeOrRerunEventUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resume or rerun event using p o s t params
func (o *ResumeOrRerunEventUsingPOSTParams) WithContext(ctx context.Context) *ResumeOrRerunEventUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resume or rerun event using p o s t params
func (o *ResumeOrRerunEventUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resume or rerun event using p o s t params
func (o *ResumeOrRerunEventUsingPOSTParams) WithHTTPClient(client *http.Client) *ResumeOrRerunEventUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resume or rerun event using p o s t params
func (o *ResumeOrRerunEventUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the resume or rerun event using p o s t params
func (o *ResumeOrRerunEventUsingPOSTParams) WithAction(action string) *ResumeOrRerunEventUsingPOSTParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the resume or rerun event using p o s t params
func (o *ResumeOrRerunEventUsingPOSTParams) SetAction(action string) {
	o.Action = action
}

// WithAPIVersion adds the aPIVersion to the resume or rerun event using p o s t params
func (o *ResumeOrRerunEventUsingPOSTParams) WithAPIVersion(aPIVersion string) *ResumeOrRerunEventUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the resume or rerun event using p o s t params
func (o *ResumeOrRerunEventUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the resume or rerun event using p o s t params
func (o *ResumeOrRerunEventUsingPOSTParams) WithID(id string) *ResumeOrRerunEventUsingPOSTParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the resume or rerun event using p o s t params
func (o *ResumeOrRerunEventUsingPOSTParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResumeOrRerunEventUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
