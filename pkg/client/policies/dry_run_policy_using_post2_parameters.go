// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewDryRunPolicyUsingPOST2Params creates a new DryRunPolicyUsingPOST2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDryRunPolicyUsingPOST2Params() *DryRunPolicyUsingPOST2Params {
	return &DryRunPolicyUsingPOST2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDryRunPolicyUsingPOST2ParamsWithTimeout creates a new DryRunPolicyUsingPOST2Params object
// with the ability to set a timeout on a request.
func NewDryRunPolicyUsingPOST2ParamsWithTimeout(timeout time.Duration) *DryRunPolicyUsingPOST2Params {
	return &DryRunPolicyUsingPOST2Params{
		timeout: timeout,
	}
}

// NewDryRunPolicyUsingPOST2ParamsWithContext creates a new DryRunPolicyUsingPOST2Params object
// with the ability to set a context for a request.
func NewDryRunPolicyUsingPOST2ParamsWithContext(ctx context.Context) *DryRunPolicyUsingPOST2Params {
	return &DryRunPolicyUsingPOST2Params{
		Context: ctx,
	}
}

// NewDryRunPolicyUsingPOST2ParamsWithHTTPClient creates a new DryRunPolicyUsingPOST2Params object
// with the ability to set a custom HTTPClient for a request.
func NewDryRunPolicyUsingPOST2ParamsWithHTTPClient(client *http.Client) *DryRunPolicyUsingPOST2Params {
	return &DryRunPolicyUsingPOST2Params{
		HTTPClient: client,
	}
}

/*
DryRunPolicyUsingPOST2Params contains all the parameters to send to the API endpoint

	for the dry run policy using p o s t 2 operation.

	Typically these are written to a http.Request.
*/
type DryRunPolicyUsingPOST2Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	// DryRun.
	DryRun string

	/* Policy.

	   The policy to be created
	*/
	Policy *models.Policy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dry run policy using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DryRunPolicyUsingPOST2Params) WithDefaults() *DryRunPolicyUsingPOST2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dry run policy using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DryRunPolicyUsingPOST2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dry run policy using p o s t 2 params
func (o *DryRunPolicyUsingPOST2Params) WithTimeout(timeout time.Duration) *DryRunPolicyUsingPOST2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dry run policy using p o s t 2 params
func (o *DryRunPolicyUsingPOST2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dry run policy using p o s t 2 params
func (o *DryRunPolicyUsingPOST2Params) WithContext(ctx context.Context) *DryRunPolicyUsingPOST2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dry run policy using p o s t 2 params
func (o *DryRunPolicyUsingPOST2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dry run policy using p o s t 2 params
func (o *DryRunPolicyUsingPOST2Params) WithHTTPClient(client *http.Client) *DryRunPolicyUsingPOST2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dry run policy using p o s t 2 params
func (o *DryRunPolicyUsingPOST2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the dry run policy using p o s t 2 params
func (o *DryRunPolicyUsingPOST2Params) WithAPIVersion(aPIVersion *string) *DryRunPolicyUsingPOST2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the dry run policy using p o s t 2 params
func (o *DryRunPolicyUsingPOST2Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDryRun adds the dryRun to the dry run policy using p o s t 2 params
func (o *DryRunPolicyUsingPOST2Params) WithDryRun(dryRun string) *DryRunPolicyUsingPOST2Params {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the dry run policy using p o s t 2 params
func (o *DryRunPolicyUsingPOST2Params) SetDryRun(dryRun string) {
	o.DryRun = dryRun
}

// WithPolicy adds the policy to the dry run policy using p o s t 2 params
func (o *DryRunPolicyUsingPOST2Params) WithPolicy(policy *models.Policy) *DryRunPolicyUsingPOST2Params {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the dry run policy using p o s t 2 params
func (o *DryRunPolicyUsingPOST2Params) SetPolicy(policy *models.Policy) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *DryRunPolicyUsingPOST2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param dryRun
	qrDryRun := o.DryRun
	qDryRun := qrDryRun
	if qDryRun != "" {

		if err := r.SetQueryParam("dryRun", qDryRun); err != nil {
			return err
		}
	}
	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}