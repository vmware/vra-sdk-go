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

// NewCreateBlueprintFromMappingUsingPOSTParams creates a new CreateBlueprintFromMappingUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBlueprintFromMappingUsingPOSTParams() *CreateBlueprintFromMappingUsingPOSTParams {
	return &CreateBlueprintFromMappingUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlueprintFromMappingUsingPOSTParamsWithTimeout creates a new CreateBlueprintFromMappingUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreateBlueprintFromMappingUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateBlueprintFromMappingUsingPOSTParams {
	return &CreateBlueprintFromMappingUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreateBlueprintFromMappingUsingPOSTParamsWithContext creates a new CreateBlueprintFromMappingUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreateBlueprintFromMappingUsingPOSTParamsWithContext(ctx context.Context) *CreateBlueprintFromMappingUsingPOSTParams {
	return &CreateBlueprintFromMappingUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreateBlueprintFromMappingUsingPOSTParamsWithHTTPClient creates a new CreateBlueprintFromMappingUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBlueprintFromMappingUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateBlueprintFromMappingUsingPOSTParams {
	return &CreateBlueprintFromMappingUsingPOSTParams{
		HTTPClient: client,
	}
}

/* CreateBlueprintFromMappingUsingPOSTParams contains all the parameters to send to the API endpoint
   for the create blueprint from mapping using p o s t operation.

   Typically these are written to a http.Request.
*/
type CreateBlueprintFromMappingUsingPOSTParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* TerraformBlueprintConfig.

	   A Terraform blueprint configuration.
	*/
	TerraformBlueprintConfig *models.TerraformBlueprintConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create blueprint from mapping using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlueprintFromMappingUsingPOSTParams) WithDefaults() *CreateBlueprintFromMappingUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create blueprint from mapping using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlueprintFromMappingUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create blueprint from mapping using p o s t params
func (o *CreateBlueprintFromMappingUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateBlueprintFromMappingUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create blueprint from mapping using p o s t params
func (o *CreateBlueprintFromMappingUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create blueprint from mapping using p o s t params
func (o *CreateBlueprintFromMappingUsingPOSTParams) WithContext(ctx context.Context) *CreateBlueprintFromMappingUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create blueprint from mapping using p o s t params
func (o *CreateBlueprintFromMappingUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create blueprint from mapping using p o s t params
func (o *CreateBlueprintFromMappingUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateBlueprintFromMappingUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create blueprint from mapping using p o s t params
func (o *CreateBlueprintFromMappingUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create blueprint from mapping using p o s t params
func (o *CreateBlueprintFromMappingUsingPOSTParams) WithAPIVersion(aPIVersion *string) *CreateBlueprintFromMappingUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create blueprint from mapping using p o s t params
func (o *CreateBlueprintFromMappingUsingPOSTParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithTerraformBlueprintConfig adds the terraformBlueprintConfig to the create blueprint from mapping using p o s t params
func (o *CreateBlueprintFromMappingUsingPOSTParams) WithTerraformBlueprintConfig(terraformBlueprintConfig *models.TerraformBlueprintConfig) *CreateBlueprintFromMappingUsingPOSTParams {
	o.SetTerraformBlueprintConfig(terraformBlueprintConfig)
	return o
}

// SetTerraformBlueprintConfig adds the terraformBlueprintConfig to the create blueprint from mapping using p o s t params
func (o *CreateBlueprintFromMappingUsingPOSTParams) SetTerraformBlueprintConfig(terraformBlueprintConfig *models.TerraformBlueprintConfig) {
	o.TerraformBlueprintConfig = terraformBlueprintConfig
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlueprintFromMappingUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.TerraformBlueprintConfig != nil {
		if err := r.SetBodyParam(o.TerraformBlueprintConfig); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
