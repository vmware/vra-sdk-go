// Code generated by go-swagger; DO NOT EDIT.

package notification_scenario_configuration

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

// NewCreateOrUpdateUsingPOST2Params creates a new CreateOrUpdateUsingPOST2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrUpdateUsingPOST2Params() *CreateOrUpdateUsingPOST2Params {
	return &CreateOrUpdateUsingPOST2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrUpdateUsingPOST2ParamsWithTimeout creates a new CreateOrUpdateUsingPOST2Params object
// with the ability to set a timeout on a request.
func NewCreateOrUpdateUsingPOST2ParamsWithTimeout(timeout time.Duration) *CreateOrUpdateUsingPOST2Params {
	return &CreateOrUpdateUsingPOST2Params{
		timeout: timeout,
	}
}

// NewCreateOrUpdateUsingPOST2ParamsWithContext creates a new CreateOrUpdateUsingPOST2Params object
// with the ability to set a context for a request.
func NewCreateOrUpdateUsingPOST2ParamsWithContext(ctx context.Context) *CreateOrUpdateUsingPOST2Params {
	return &CreateOrUpdateUsingPOST2Params{
		Context: ctx,
	}
}

// NewCreateOrUpdateUsingPOST2ParamsWithHTTPClient creates a new CreateOrUpdateUsingPOST2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrUpdateUsingPOST2ParamsWithHTTPClient(client *http.Client) *CreateOrUpdateUsingPOST2Params {
	return &CreateOrUpdateUsingPOST2Params{
		HTTPClient: client,
	}
}

/*
CreateOrUpdateUsingPOST2Params contains all the parameters to send to the API endpoint

	for the create or update using p o s t 2 operation.

	Typically these are written to a http.Request.
*/
type CreateOrUpdateUsingPOST2Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* Config.

	   The notification scenario configuration to be created or updated
	*/
	Config *models.NotificationScenarioConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create or update using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrUpdateUsingPOST2Params) WithDefaults() *CreateOrUpdateUsingPOST2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create or update using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrUpdateUsingPOST2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create or update using p o s t 2 params
func (o *CreateOrUpdateUsingPOST2Params) WithTimeout(timeout time.Duration) *CreateOrUpdateUsingPOST2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create or update using p o s t 2 params
func (o *CreateOrUpdateUsingPOST2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create or update using p o s t 2 params
func (o *CreateOrUpdateUsingPOST2Params) WithContext(ctx context.Context) *CreateOrUpdateUsingPOST2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create or update using p o s t 2 params
func (o *CreateOrUpdateUsingPOST2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create or update using p o s t 2 params
func (o *CreateOrUpdateUsingPOST2Params) WithHTTPClient(client *http.Client) *CreateOrUpdateUsingPOST2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create or update using p o s t 2 params
func (o *CreateOrUpdateUsingPOST2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create or update using p o s t 2 params
func (o *CreateOrUpdateUsingPOST2Params) WithAPIVersion(aPIVersion *string) *CreateOrUpdateUsingPOST2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create or update using p o s t 2 params
func (o *CreateOrUpdateUsingPOST2Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithConfig adds the config to the create or update using p o s t 2 params
func (o *CreateOrUpdateUsingPOST2Params) WithConfig(config *models.NotificationScenarioConfig) *CreateOrUpdateUsingPOST2Params {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the create or update using p o s t 2 params
func (o *CreateOrUpdateUsingPOST2Params) SetConfig(config *models.NotificationScenarioConfig) {
	o.Config = config
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrUpdateUsingPOST2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}