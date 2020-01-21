// Code generated by go-swagger; DO NOT EDIT.

package source_control_sync

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

// NewScheduleSyncAllUsingPOSTParams creates a new ScheduleSyncAllUsingPOSTParams object
// with the default values initialized.
func NewScheduleSyncAllUsingPOSTParams() *ScheduleSyncAllUsingPOSTParams {
	var ()
	return &ScheduleSyncAllUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewScheduleSyncAllUsingPOSTParamsWithTimeout creates a new ScheduleSyncAllUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewScheduleSyncAllUsingPOSTParamsWithTimeout(timeout time.Duration) *ScheduleSyncAllUsingPOSTParams {
	var ()
	return &ScheduleSyncAllUsingPOSTParams{

		timeout: timeout,
	}
}

// NewScheduleSyncAllUsingPOSTParamsWithContext creates a new ScheduleSyncAllUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewScheduleSyncAllUsingPOSTParamsWithContext(ctx context.Context) *ScheduleSyncAllUsingPOSTParams {
	var ()
	return &ScheduleSyncAllUsingPOSTParams{

		Context: ctx,
	}
}

// NewScheduleSyncAllUsingPOSTParamsWithHTTPClient creates a new ScheduleSyncAllUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewScheduleSyncAllUsingPOSTParamsWithHTTPClient(client *http.Client) *ScheduleSyncAllUsingPOSTParams {
	var ()
	return &ScheduleSyncAllUsingPOSTParams{
		HTTPClient: client,
	}
}

/*ScheduleSyncAllUsingPOSTParams contains all the parameters to send to the API endpoint
for the schedule sync all using p o s t operation typically these are written to a http.Request
*/
type ScheduleSyncAllUsingPOSTParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information, please refer to /content/api/about

	*/
	APIVersion *string
	/*Request
	  request

	*/
	Request *models.SourceControlSyncAllRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schedule sync all using p o s t params
func (o *ScheduleSyncAllUsingPOSTParams) WithTimeout(timeout time.Duration) *ScheduleSyncAllUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedule sync all using p o s t params
func (o *ScheduleSyncAllUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedule sync all using p o s t params
func (o *ScheduleSyncAllUsingPOSTParams) WithContext(ctx context.Context) *ScheduleSyncAllUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedule sync all using p o s t params
func (o *ScheduleSyncAllUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedule sync all using p o s t params
func (o *ScheduleSyncAllUsingPOSTParams) WithHTTPClient(client *http.Client) *ScheduleSyncAllUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedule sync all using p o s t params
func (o *ScheduleSyncAllUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the schedule sync all using p o s t params
func (o *ScheduleSyncAllUsingPOSTParams) WithAPIVersion(aPIVersion *string) *ScheduleSyncAllUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the schedule sync all using p o s t params
func (o *ScheduleSyncAllUsingPOSTParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithRequest adds the request to the schedule sync all using p o s t params
func (o *ScheduleSyncAllUsingPOSTParams) WithRequest(request *models.SourceControlSyncAllRequest) *ScheduleSyncAllUsingPOSTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the schedule sync all using p o s t params
func (o *ScheduleSyncAllUsingPOSTParams) SetRequest(request *models.SourceControlSyncAllRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *ScheduleSyncAllUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
