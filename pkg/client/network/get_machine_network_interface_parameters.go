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
)

// NewGetMachineNetworkInterfaceParams creates a new GetMachineNetworkInterfaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMachineNetworkInterfaceParams() *GetMachineNetworkInterfaceParams {
	return &GetMachineNetworkInterfaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMachineNetworkInterfaceParamsWithTimeout creates a new GetMachineNetworkInterfaceParams object
// with the ability to set a timeout on a request.
func NewGetMachineNetworkInterfaceParamsWithTimeout(timeout time.Duration) *GetMachineNetworkInterfaceParams {
	return &GetMachineNetworkInterfaceParams{
		timeout: timeout,
	}
}

// NewGetMachineNetworkInterfaceParamsWithContext creates a new GetMachineNetworkInterfaceParams object
// with the ability to set a context for a request.
func NewGetMachineNetworkInterfaceParamsWithContext(ctx context.Context) *GetMachineNetworkInterfaceParams {
	return &GetMachineNetworkInterfaceParams{
		Context: ctx,
	}
}

// NewGetMachineNetworkInterfaceParamsWithHTTPClient creates a new GetMachineNetworkInterfaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMachineNetworkInterfaceParamsWithHTTPClient(client *http.Client) *GetMachineNetworkInterfaceParams {
	return &GetMachineNetworkInterfaceParams{
		HTTPClient: client,
	}
}

/* GetMachineNetworkInterfaceParams contains all the parameters to send to the API endpoint
   for the get machine network interface operation.

   Typically these are written to a http.Request.
*/
type GetMachineNetworkInterfaceParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

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

// WithDefaults hydrates default values in the get machine network interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMachineNetworkInterfaceParams) WithDefaults() *GetMachineNetworkInterfaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get machine network interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMachineNetworkInterfaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get machine network interface params
func (o *GetMachineNetworkInterfaceParams) WithTimeout(timeout time.Duration) *GetMachineNetworkInterfaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get machine network interface params
func (o *GetMachineNetworkInterfaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get machine network interface params
func (o *GetMachineNetworkInterfaceParams) WithContext(ctx context.Context) *GetMachineNetworkInterfaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get machine network interface params
func (o *GetMachineNetworkInterfaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get machine network interface params
func (o *GetMachineNetworkInterfaceParams) WithHTTPClient(client *http.Client) *GetMachineNetworkInterfaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get machine network interface params
func (o *GetMachineNetworkInterfaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get machine network interface params
func (o *GetMachineNetworkInterfaceParams) WithAPIVersion(aPIVersion *string) *GetMachineNetworkInterfaceParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get machine network interface params
func (o *GetMachineNetworkInterfaceParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get machine network interface params
func (o *GetMachineNetworkInterfaceParams) WithID(id string) *GetMachineNetworkInterfaceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get machine network interface params
func (o *GetMachineNetworkInterfaceParams) SetID(id string) {
	o.ID = id
}

// WithNetworkID adds the networkID to the get machine network interface params
func (o *GetMachineNetworkInterfaceParams) WithNetworkID(networkID string) *GetMachineNetworkInterfaceParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get machine network interface params
func (o *GetMachineNetworkInterfaceParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMachineNetworkInterfaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
