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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateBlueprintMappingUsingPOSTParams creates a new CreateBlueprintMappingUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBlueprintMappingUsingPOSTParams() *CreateBlueprintMappingUsingPOSTParams {
	return &CreateBlueprintMappingUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlueprintMappingUsingPOSTParamsWithTimeout creates a new CreateBlueprintMappingUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreateBlueprintMappingUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateBlueprintMappingUsingPOSTParams {
	return &CreateBlueprintMappingUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreateBlueprintMappingUsingPOSTParamsWithContext creates a new CreateBlueprintMappingUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreateBlueprintMappingUsingPOSTParamsWithContext(ctx context.Context) *CreateBlueprintMappingUsingPOSTParams {
	return &CreateBlueprintMappingUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreateBlueprintMappingUsingPOSTParamsWithHTTPClient creates a new CreateBlueprintMappingUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBlueprintMappingUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateBlueprintMappingUsingPOSTParams {
	return &CreateBlueprintMappingUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
CreateBlueprintMappingUsingPOSTParams contains all the parameters to send to the API endpoint

	for the create blueprint mapping using p o s t operation.

	Typically these are written to a http.Request.
*/
type CreateBlueprintMappingUsingPOSTParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* ConfigurationSourceReference.

	   configurationSourceReference
	*/
	ConfigurationSourceReference *models.TerraformConfigurationSourceReference

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create blueprint mapping using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlueprintMappingUsingPOSTParams) WithDefaults() *CreateBlueprintMappingUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create blueprint mapping using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlueprintMappingUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create blueprint mapping using p o s t params
func (o *CreateBlueprintMappingUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateBlueprintMappingUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create blueprint mapping using p o s t params
func (o *CreateBlueprintMappingUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create blueprint mapping using p o s t params
func (o *CreateBlueprintMappingUsingPOSTParams) WithContext(ctx context.Context) *CreateBlueprintMappingUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create blueprint mapping using p o s t params
func (o *CreateBlueprintMappingUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create blueprint mapping using p o s t params
func (o *CreateBlueprintMappingUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateBlueprintMappingUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create blueprint mapping using p o s t params
func (o *CreateBlueprintMappingUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create blueprint mapping using p o s t params
func (o *CreateBlueprintMappingUsingPOSTParams) WithAPIVersion(aPIVersion *string) *CreateBlueprintMappingUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create blueprint mapping using p o s t params
func (o *CreateBlueprintMappingUsingPOSTParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithConfigurationSourceReference adds the configurationSourceReference to the create blueprint mapping using p o s t params
func (o *CreateBlueprintMappingUsingPOSTParams) WithConfigurationSourceReference(configurationSourceReference *models.TerraformConfigurationSourceReference) *CreateBlueprintMappingUsingPOSTParams {
	o.SetConfigurationSourceReference(configurationSourceReference)
	return o
}

// SetConfigurationSourceReference adds the configurationSourceReference to the create blueprint mapping using p o s t params
func (o *CreateBlueprintMappingUsingPOSTParams) SetConfigurationSourceReference(configurationSourceReference *models.TerraformConfigurationSourceReference) {
	o.ConfigurationSourceReference = configurationSourceReference
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlueprintMappingUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.ConfigurationSourceReference != nil {
		if err := r.SetBodyParam(o.ConfigurationSourceReference); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
