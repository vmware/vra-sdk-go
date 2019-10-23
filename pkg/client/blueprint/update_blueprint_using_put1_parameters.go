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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdateBlueprintUsingPUT1Params creates a new UpdateBlueprintUsingPUT1Params object
// with the default values initialized.
func NewUpdateBlueprintUsingPUT1Params() *UpdateBlueprintUsingPUT1Params {
	var ()
	return &UpdateBlueprintUsingPUT1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBlueprintUsingPUT1ParamsWithTimeout creates a new UpdateBlueprintUsingPUT1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateBlueprintUsingPUT1ParamsWithTimeout(timeout time.Duration) *UpdateBlueprintUsingPUT1Params {
	var ()
	return &UpdateBlueprintUsingPUT1Params{

		timeout: timeout,
	}
}

// NewUpdateBlueprintUsingPUT1ParamsWithContext creates a new UpdateBlueprintUsingPUT1Params object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateBlueprintUsingPUT1ParamsWithContext(ctx context.Context) *UpdateBlueprintUsingPUT1Params {
	var ()
	return &UpdateBlueprintUsingPUT1Params{

		Context: ctx,
	}
}

// NewUpdateBlueprintUsingPUT1ParamsWithHTTPClient creates a new UpdateBlueprintUsingPUT1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateBlueprintUsingPUT1ParamsWithHTTPClient(client *http.Client) *UpdateBlueprintUsingPUT1Params {
	var ()
	return &UpdateBlueprintUsingPUT1Params{
		HTTPClient: client,
	}
}

/*UpdateBlueprintUsingPUT1Params contains all the parameters to send to the API endpoint
for the update blueprint using p u t 1 operation typically these are written to a http.Request
*/
type UpdateBlueprintUsingPUT1Params struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about

	*/
	APIVersion *string
	/*Blueprint
	  Blueprint object

	*/
	Blueprint *models.Blueprint
	/*BlueprintID
	  blueprintId

	*/
	BlueprintID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update blueprint using p u t 1 params
func (o *UpdateBlueprintUsingPUT1Params) WithTimeout(timeout time.Duration) *UpdateBlueprintUsingPUT1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update blueprint using p u t 1 params
func (o *UpdateBlueprintUsingPUT1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update blueprint using p u t 1 params
func (o *UpdateBlueprintUsingPUT1Params) WithContext(ctx context.Context) *UpdateBlueprintUsingPUT1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update blueprint using p u t 1 params
func (o *UpdateBlueprintUsingPUT1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update blueprint using p u t 1 params
func (o *UpdateBlueprintUsingPUT1Params) WithHTTPClient(client *http.Client) *UpdateBlueprintUsingPUT1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update blueprint using p u t 1 params
func (o *UpdateBlueprintUsingPUT1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update blueprint using p u t 1 params
func (o *UpdateBlueprintUsingPUT1Params) WithAPIVersion(aPIVersion *string) *UpdateBlueprintUsingPUT1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update blueprint using p u t 1 params
func (o *UpdateBlueprintUsingPUT1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBlueprint adds the blueprint to the update blueprint using p u t 1 params
func (o *UpdateBlueprintUsingPUT1Params) WithBlueprint(blueprint *models.Blueprint) *UpdateBlueprintUsingPUT1Params {
	o.SetBlueprint(blueprint)
	return o
}

// SetBlueprint adds the blueprint to the update blueprint using p u t 1 params
func (o *UpdateBlueprintUsingPUT1Params) SetBlueprint(blueprint *models.Blueprint) {
	o.Blueprint = blueprint
}

// WithBlueprintID adds the blueprintID to the update blueprint using p u t 1 params
func (o *UpdateBlueprintUsingPUT1Params) WithBlueprintID(blueprintID strfmt.UUID) *UpdateBlueprintUsingPUT1Params {
	o.SetBlueprintID(blueprintID)
	return o
}

// SetBlueprintID adds the blueprintId to the update blueprint using p u t 1 params
func (o *UpdateBlueprintUsingPUT1Params) SetBlueprintID(blueprintID strfmt.UUID) {
	o.BlueprintID = blueprintID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBlueprintUsingPUT1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Blueprint != nil {
		if err := r.SetBodyParam(o.Blueprint); err != nil {
			return err
		}
	}

	// path param blueprintId
	if err := r.SetPathParam("blueprintId", o.BlueprintID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
