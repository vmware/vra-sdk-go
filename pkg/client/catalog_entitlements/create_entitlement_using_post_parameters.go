// Code generated by go-swagger; DO NOT EDIT.

package catalog_entitlements

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

// NewCreateEntitlementUsingPOSTParams creates a new CreateEntitlementUsingPOSTParams object
// with the default values initialized.
func NewCreateEntitlementUsingPOSTParams() *CreateEntitlementUsingPOSTParams {
	var ()
	return &CreateEntitlementUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEntitlementUsingPOSTParamsWithTimeout creates a new CreateEntitlementUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateEntitlementUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateEntitlementUsingPOSTParams {
	var ()
	return &CreateEntitlementUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateEntitlementUsingPOSTParamsWithContext creates a new CreateEntitlementUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateEntitlementUsingPOSTParamsWithContext(ctx context.Context) *CreateEntitlementUsingPOSTParams {
	var ()
	return &CreateEntitlementUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateEntitlementUsingPOSTParamsWithHTTPClient creates a new CreateEntitlementUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateEntitlementUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateEntitlementUsingPOSTParams {
	var ()
	return &CreateEntitlementUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateEntitlementUsingPOSTParams contains all the parameters to send to the API endpoint
for the create entitlement using p o s t operation typically these are written to a http.Request
*/
type CreateEntitlementUsingPOSTParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.

	*/
	APIVersion *string
	/*Entitlement
	  The entitlement to be created

	*/
	Entitlement *models.Entitlement

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create entitlement using p o s t params
func (o *CreateEntitlementUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateEntitlementUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create entitlement using p o s t params
func (o *CreateEntitlementUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create entitlement using p o s t params
func (o *CreateEntitlementUsingPOSTParams) WithContext(ctx context.Context) *CreateEntitlementUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create entitlement using p o s t params
func (o *CreateEntitlementUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create entitlement using p o s t params
func (o *CreateEntitlementUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateEntitlementUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create entitlement using p o s t params
func (o *CreateEntitlementUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create entitlement using p o s t params
func (o *CreateEntitlementUsingPOSTParams) WithAPIVersion(aPIVersion *string) *CreateEntitlementUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create entitlement using p o s t params
func (o *CreateEntitlementUsingPOSTParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithEntitlement adds the entitlement to the create entitlement using p o s t params
func (o *CreateEntitlementUsingPOSTParams) WithEntitlement(entitlement *models.Entitlement) *CreateEntitlementUsingPOSTParams {
	o.SetEntitlement(entitlement)
	return o
}

// SetEntitlement adds the entitlement to the create entitlement using p o s t params
func (o *CreateEntitlementUsingPOSTParams) SetEntitlement(entitlement *models.Entitlement) {
	o.Entitlement = entitlement
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEntitlementUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Entitlement != nil {
		if err := r.SetBodyParam(o.Entitlement); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
