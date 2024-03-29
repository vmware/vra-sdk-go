// Code generated by go-swagger; DO NOT EDIT.

package deployment_actions

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

// NewSubmitResourceActionRequestUsingPOST4Params creates a new SubmitResourceActionRequestUsingPOST4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubmitResourceActionRequestUsingPOST4Params() *SubmitResourceActionRequestUsingPOST4Params {
	return &SubmitResourceActionRequestUsingPOST4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitResourceActionRequestUsingPOST4ParamsWithTimeout creates a new SubmitResourceActionRequestUsingPOST4Params object
// with the ability to set a timeout on a request.
func NewSubmitResourceActionRequestUsingPOST4ParamsWithTimeout(timeout time.Duration) *SubmitResourceActionRequestUsingPOST4Params {
	return &SubmitResourceActionRequestUsingPOST4Params{
		timeout: timeout,
	}
}

// NewSubmitResourceActionRequestUsingPOST4ParamsWithContext creates a new SubmitResourceActionRequestUsingPOST4Params object
// with the ability to set a context for a request.
func NewSubmitResourceActionRequestUsingPOST4ParamsWithContext(ctx context.Context) *SubmitResourceActionRequestUsingPOST4Params {
	return &SubmitResourceActionRequestUsingPOST4Params{
		Context: ctx,
	}
}

// NewSubmitResourceActionRequestUsingPOST4ParamsWithHTTPClient creates a new SubmitResourceActionRequestUsingPOST4Params object
// with the ability to set a custom HTTPClient for a request.
func NewSubmitResourceActionRequestUsingPOST4ParamsWithHTTPClient(client *http.Client) *SubmitResourceActionRequestUsingPOST4Params {
	return &SubmitResourceActionRequestUsingPOST4Params{
		HTTPClient: client,
	}
}

/*
SubmitResourceActionRequestUsingPOST4Params contains all the parameters to send to the API endpoint

	for the submit resource action request using p o s t 4 operation.

	Typically these are written to a http.Request.
*/
type SubmitResourceActionRequestUsingPOST4Params struct {

	/* ActionRequest.

	   actionRequest
	*/
	ActionRequest *models.ResourceActionRequest

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* DeploymentID.

	   Deployment ID

	   Format: uuid
	*/
	DeploymentID strfmt.UUID

	/* ResourceID.

	   Resource ID

	   Format: uuid
	*/
	ResourceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the submit resource action request using p o s t 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitResourceActionRequestUsingPOST4Params) WithDefaults() *SubmitResourceActionRequestUsingPOST4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the submit resource action request using p o s t 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitResourceActionRequestUsingPOST4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) WithTimeout(timeout time.Duration) *SubmitResourceActionRequestUsingPOST4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) WithContext(ctx context.Context) *SubmitResourceActionRequestUsingPOST4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) WithHTTPClient(client *http.Client) *SubmitResourceActionRequestUsingPOST4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionRequest adds the actionRequest to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) WithActionRequest(actionRequest *models.ResourceActionRequest) *SubmitResourceActionRequestUsingPOST4Params {
	o.SetActionRequest(actionRequest)
	return o
}

// SetActionRequest adds the actionRequest to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) SetActionRequest(actionRequest *models.ResourceActionRequest) {
	o.ActionRequest = actionRequest
}

// WithAPIVersion adds the aPIVersion to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) WithAPIVersion(aPIVersion *string) *SubmitResourceActionRequestUsingPOST4Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDeploymentID adds the deploymentID to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) WithDeploymentID(deploymentID strfmt.UUID) *SubmitResourceActionRequestUsingPOST4Params {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) SetDeploymentID(deploymentID strfmt.UUID) {
	o.DeploymentID = deploymentID
}

// WithResourceID adds the resourceID to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) WithResourceID(resourceID strfmt.UUID) *SubmitResourceActionRequestUsingPOST4Params {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the submit resource action request using p o s t 4 params
func (o *SubmitResourceActionRequestUsingPOST4Params) SetResourceID(resourceID strfmt.UUID) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitResourceActionRequestUsingPOST4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ActionRequest != nil {
		if err := r.SetBodyParam(o.ActionRequest); err != nil {
			return err
		}
	}

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

	// path param deploymentId
	if err := r.SetPathParam("deploymentId", o.DeploymentID.String()); err != nil {
		return err
	}

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
