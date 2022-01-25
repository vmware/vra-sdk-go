// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

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
)

// NewGetTerraformConfigurationSourceCommitListUsingGETParams creates a new GetTerraformConfigurationSourceCommitListUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTerraformConfigurationSourceCommitListUsingGETParams() *GetTerraformConfigurationSourceCommitListUsingGETParams {
	return &GetTerraformConfigurationSourceCommitListUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTerraformConfigurationSourceCommitListUsingGETParamsWithTimeout creates a new GetTerraformConfigurationSourceCommitListUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetTerraformConfigurationSourceCommitListUsingGETParamsWithTimeout(timeout time.Duration) *GetTerraformConfigurationSourceCommitListUsingGETParams {
	return &GetTerraformConfigurationSourceCommitListUsingGETParams{
		timeout: timeout,
	}
}

// NewGetTerraformConfigurationSourceCommitListUsingGETParamsWithContext creates a new GetTerraformConfigurationSourceCommitListUsingGETParams object
// with the ability to set a context for a request.
func NewGetTerraformConfigurationSourceCommitListUsingGETParamsWithContext(ctx context.Context) *GetTerraformConfigurationSourceCommitListUsingGETParams {
	return &GetTerraformConfigurationSourceCommitListUsingGETParams{
		Context: ctx,
	}
}

// NewGetTerraformConfigurationSourceCommitListUsingGETParamsWithHTTPClient creates a new GetTerraformConfigurationSourceCommitListUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTerraformConfigurationSourceCommitListUsingGETParamsWithHTTPClient(client *http.Client) *GetTerraformConfigurationSourceCommitListUsingGETParams {
	return &GetTerraformConfigurationSourceCommitListUsingGETParams{
		HTTPClient: client,
	}
}

/* GetTerraformConfigurationSourceCommitListUsingGETParams contains all the parameters to send to the API endpoint
   for the get terraform configuration source commit list using g e t operation.

   Typically these are written to a http.Request.
*/
type GetTerraformConfigurationSourceCommitListUsingGETParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* ConfigurationSourceID.

	   A Configuration Source ID.
	*/
	ConfigurationSourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get terraform configuration source commit list using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) WithDefaults() *GetTerraformConfigurationSourceCommitListUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get terraform configuration source commit list using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get terraform configuration source commit list using get params
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) WithTimeout(timeout time.Duration) *GetTerraformConfigurationSourceCommitListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get terraform configuration source commit list using get params
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get terraform configuration source commit list using get params
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) WithContext(ctx context.Context) *GetTerraformConfigurationSourceCommitListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get terraform configuration source commit list using get params
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get terraform configuration source commit list using get params
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) WithHTTPClient(client *http.Client) *GetTerraformConfigurationSourceCommitListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get terraform configuration source commit list using get params
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get terraform configuration source commit list using get params
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) WithAPIVersion(aPIVersion *string) *GetTerraformConfigurationSourceCommitListUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get terraform configuration source commit list using get params
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithConfigurationSourceID adds the configurationSourceID to the get terraform configuration source commit list using get params
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) WithConfigurationSourceID(configurationSourceID string) *GetTerraformConfigurationSourceCommitListUsingGETParams {
	o.SetConfigurationSourceID(configurationSourceID)
	return o
}

// SetConfigurationSourceID adds the configurationSourceId to the get terraform configuration source commit list using get params
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) SetConfigurationSourceID(configurationSourceID string) {
	o.ConfigurationSourceID = configurationSourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTerraformConfigurationSourceCommitListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param configurationSourceId
	qrConfigurationSourceID := o.ConfigurationSourceID
	qConfigurationSourceID := qrConfigurationSourceID
	if qConfigurationSourceID != "" {

		if err := r.SetQueryParam("configurationSourceId", qConfigurationSourceID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
