// Code generated by go-swagger; DO NOT EDIT.

package catalog_items

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

// NewSubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params creates a new SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params() *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params {
	return &SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitUpfrontPriceRequestForCatalogItemUsingPOST2ParamsWithTimeout creates a new SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params object
// with the ability to set a timeout on a request.
func NewSubmitUpfrontPriceRequestForCatalogItemUsingPOST2ParamsWithTimeout(timeout time.Duration) *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params {
	return &SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params{
		timeout: timeout,
	}
}

// NewSubmitUpfrontPriceRequestForCatalogItemUsingPOST2ParamsWithContext creates a new SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params object
// with the ability to set a context for a request.
func NewSubmitUpfrontPriceRequestForCatalogItemUsingPOST2ParamsWithContext(ctx context.Context) *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params {
	return &SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params{
		Context: ctx,
	}
}

// NewSubmitUpfrontPriceRequestForCatalogItemUsingPOST2ParamsWithHTTPClient creates a new SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params object
// with the ability to set a custom HTTPClient for a request.
func NewSubmitUpfrontPriceRequestForCatalogItemUsingPOST2ParamsWithHTTPClient(client *http.Client) *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params {
	return &SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params{
		HTTPClient: client,
	}
}

/*
SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params contains all the parameters to send to the API endpoint

	for the submit upfront price request for catalog item using p o s t 2 operation.

	Typically these are written to a http.Request.
*/
type SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* ID.

	   Catalog Item ID

	   Format: uuid
	*/
	ID strfmt.UUID

	/* Request.

	   request
	*/
	Request *models.CatalogItemRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the submit upfront price request for catalog item using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) WithDefaults() *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the submit upfront price request for catalog item using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the submit upfront price request for catalog item using p o s t 2 params
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) WithTimeout(timeout time.Duration) *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit upfront price request for catalog item using p o s t 2 params
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit upfront price request for catalog item using p o s t 2 params
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) WithContext(ctx context.Context) *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit upfront price request for catalog item using p o s t 2 params
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit upfront price request for catalog item using p o s t 2 params
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) WithHTTPClient(client *http.Client) *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit upfront price request for catalog item using p o s t 2 params
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the submit upfront price request for catalog item using p o s t 2 params
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) WithAPIVersion(aPIVersion *string) *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the submit upfront price request for catalog item using p o s t 2 params
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the submit upfront price request for catalog item using p o s t 2 params
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) WithID(id strfmt.UUID) *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the submit upfront price request for catalog item using p o s t 2 params
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithRequest adds the request to the submit upfront price request for catalog item using p o s t 2 params
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) WithRequest(request *models.CatalogItemRequest) *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the submit upfront price request for catalog item using p o s t 2 params
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) SetRequest(request *models.CatalogItemRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitUpfrontPriceRequestForCatalogItemUsingPOST2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
