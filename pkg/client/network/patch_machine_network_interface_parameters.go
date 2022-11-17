// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewPatchMachineNetworkInterfaceParams creates a new PatchMachineNetworkInterfaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchMachineNetworkInterfaceParams() *PatchMachineNetworkInterfaceParams {
	return &PatchMachineNetworkInterfaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchMachineNetworkInterfaceParamsWithTimeout creates a new PatchMachineNetworkInterfaceParams object
// with the ability to set a timeout on a request.
func NewPatchMachineNetworkInterfaceParamsWithTimeout(timeout time.Duration) *PatchMachineNetworkInterfaceParams {
	return &PatchMachineNetworkInterfaceParams{
		timeout: timeout,
	}
}

// NewPatchMachineNetworkInterfaceParamsWithContext creates a new PatchMachineNetworkInterfaceParams object
// with the ability to set a context for a request.
func NewPatchMachineNetworkInterfaceParamsWithContext(ctx context.Context) *PatchMachineNetworkInterfaceParams {
	return &PatchMachineNetworkInterfaceParams{
		Context: ctx,
	}
}

// NewPatchMachineNetworkInterfaceParamsWithHTTPClient creates a new PatchMachineNetworkInterfaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchMachineNetworkInterfaceParamsWithHTTPClient(client *http.Client) *PatchMachineNetworkInterfaceParams {
	return &PatchMachineNetworkInterfaceParams{
		HTTPClient: client,
	}
}

/*
PatchMachineNetworkInterfaceParams contains all the parameters to send to the API endpoint

	for the patch machine network interface operation.

	Typically these are written to a http.Request.
*/
type PatchMachineNetworkInterfaceParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion string

	/* Body.

	   NetworkInterface Specification
	*/
	Body *models.UpdateNetworkInterfaceSpecification

	/* ID.

	   The ID of the machine.
	*/
	ID string

	/* NetworkID.

	   The ID of the network interface.
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch machine network interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchMachineNetworkInterfaceParams) WithDefaults() *PatchMachineNetworkInterfaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch machine network interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchMachineNetworkInterfaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) WithTimeout(timeout time.Duration) *PatchMachineNetworkInterfaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) WithContext(ctx context.Context) *PatchMachineNetworkInterfaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) WithHTTPClient(client *http.Client) *PatchMachineNetworkInterfaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) WithAPIVersion(aPIVersion string) *PatchMachineNetworkInterfaceParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) WithBody(body *models.UpdateNetworkInterfaceSpecification) *PatchMachineNetworkInterfaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) SetBody(body *models.UpdateNetworkInterfaceSpecification) {
	o.Body = body
}

// WithID adds the id to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) WithID(id string) *PatchMachineNetworkInterfaceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) SetID(id string) {
	o.ID = id
}

// WithNetworkID adds the networkID to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) WithNetworkID(networkID string) *PatchMachineNetworkInterfaceParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the patch machine network interface params
func (o *PatchMachineNetworkInterfaceParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchMachineNetworkInterfaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {

		if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
