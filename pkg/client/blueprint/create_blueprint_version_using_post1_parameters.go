// Code generated by go-swagger; DO NOT EDIT.

package blueprint

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

// NewCreateBlueprintVersionUsingPOST1Params creates a new CreateBlueprintVersionUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBlueprintVersionUsingPOST1Params() *CreateBlueprintVersionUsingPOST1Params {
	return &CreateBlueprintVersionUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlueprintVersionUsingPOST1ParamsWithTimeout creates a new CreateBlueprintVersionUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewCreateBlueprintVersionUsingPOST1ParamsWithTimeout(timeout time.Duration) *CreateBlueprintVersionUsingPOST1Params {
	return &CreateBlueprintVersionUsingPOST1Params{
		timeout: timeout,
	}
}

// NewCreateBlueprintVersionUsingPOST1ParamsWithContext creates a new CreateBlueprintVersionUsingPOST1Params object
// with the ability to set a context for a request.
func NewCreateBlueprintVersionUsingPOST1ParamsWithContext(ctx context.Context) *CreateBlueprintVersionUsingPOST1Params {
	return &CreateBlueprintVersionUsingPOST1Params{
		Context: ctx,
	}
}

// NewCreateBlueprintVersionUsingPOST1ParamsWithHTTPClient creates a new CreateBlueprintVersionUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBlueprintVersionUsingPOST1ParamsWithHTTPClient(client *http.Client) *CreateBlueprintVersionUsingPOST1Params {
	return &CreateBlueprintVersionUsingPOST1Params{
		HTTPClient: client,
	}
}

/*
CreateBlueprintVersionUsingPOST1Params contains all the parameters to send to the API endpoint

	for the create blueprint version using p o s t 1 operation.

	Typically these are written to a http.Request.
*/
type CreateBlueprintVersionUsingPOST1Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* BlueprintID.

	   blueprintId

	   Format: uuid
	*/
	BlueprintID strfmt.UUID

	/* VersionRequest.

	   Blueprint version object
	*/
	VersionRequest *models.BlueprintVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create blueprint version using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlueprintVersionUsingPOST1Params) WithDefaults() *CreateBlueprintVersionUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create blueprint version using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlueprintVersionUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create blueprint version using p o s t 1 params
func (o *CreateBlueprintVersionUsingPOST1Params) WithTimeout(timeout time.Duration) *CreateBlueprintVersionUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create blueprint version using p o s t 1 params
func (o *CreateBlueprintVersionUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create blueprint version using p o s t 1 params
func (o *CreateBlueprintVersionUsingPOST1Params) WithContext(ctx context.Context) *CreateBlueprintVersionUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create blueprint version using p o s t 1 params
func (o *CreateBlueprintVersionUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create blueprint version using p o s t 1 params
func (o *CreateBlueprintVersionUsingPOST1Params) WithHTTPClient(client *http.Client) *CreateBlueprintVersionUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create blueprint version using p o s t 1 params
func (o *CreateBlueprintVersionUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create blueprint version using p o s t 1 params
func (o *CreateBlueprintVersionUsingPOST1Params) WithAPIVersion(aPIVersion *string) *CreateBlueprintVersionUsingPOST1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create blueprint version using p o s t 1 params
func (o *CreateBlueprintVersionUsingPOST1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBlueprintID adds the blueprintID to the create blueprint version using p o s t 1 params
func (o *CreateBlueprintVersionUsingPOST1Params) WithBlueprintID(blueprintID strfmt.UUID) *CreateBlueprintVersionUsingPOST1Params {
	o.SetBlueprintID(blueprintID)
	return o
}

// SetBlueprintID adds the blueprintId to the create blueprint version using p o s t 1 params
func (o *CreateBlueprintVersionUsingPOST1Params) SetBlueprintID(blueprintID strfmt.UUID) {
	o.BlueprintID = blueprintID
}

// WithVersionRequest adds the versionRequest to the create blueprint version using p o s t 1 params
func (o *CreateBlueprintVersionUsingPOST1Params) WithVersionRequest(versionRequest *models.BlueprintVersionRequest) *CreateBlueprintVersionUsingPOST1Params {
	o.SetVersionRequest(versionRequest)
	return o
}

// SetVersionRequest adds the versionRequest to the create blueprint version using p o s t 1 params
func (o *CreateBlueprintVersionUsingPOST1Params) SetVersionRequest(versionRequest *models.BlueprintVersionRequest) {
	o.VersionRequest = versionRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlueprintVersionUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param blueprintId
	if err := r.SetPathParam("blueprintId", o.BlueprintID.String()); err != nil {
		return err
	}
	if o.VersionRequest != nil {
		if err := r.SetBodyParam(o.VersionRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
