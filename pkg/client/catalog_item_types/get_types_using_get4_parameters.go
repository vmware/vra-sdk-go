// Code generated by go-swagger; DO NOT EDIT.

package catalog_item_types

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
	"github.com/go-openapi/swag"
)

// NewGetTypesUsingGET4Params creates a new GetTypesUsingGET4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTypesUsingGET4Params() *GetTypesUsingGET4Params {
	return &GetTypesUsingGET4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTypesUsingGET4ParamsWithTimeout creates a new GetTypesUsingGET4Params object
// with the ability to set a timeout on a request.
func NewGetTypesUsingGET4ParamsWithTimeout(timeout time.Duration) *GetTypesUsingGET4Params {
	return &GetTypesUsingGET4Params{
		timeout: timeout,
	}
}

// NewGetTypesUsingGET4ParamsWithContext creates a new GetTypesUsingGET4Params object
// with the ability to set a context for a request.
func NewGetTypesUsingGET4ParamsWithContext(ctx context.Context) *GetTypesUsingGET4Params {
	return &GetTypesUsingGET4Params{
		Context: ctx,
	}
}

// NewGetTypesUsingGET4ParamsWithHTTPClient creates a new GetTypesUsingGET4Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetTypesUsingGET4ParamsWithHTTPClient(client *http.Client) *GetTypesUsingGET4Params {
	return &GetTypesUsingGET4Params{
		HTTPClient: client,
	}
}

/*
GetTypesUsingGET4Params contains all the parameters to send to the API endpoint

	for the get types using g e t 4 operation.

	Typically these are written to a http.Request.
*/
type GetTypesUsingGET4Params struct {

	/* DollarOrderby.

	   Sorting criteria in the format: property (asc|desc). Default sort order is ascending. Multiple sort criteria are supported.
	*/
	DollarOrderby []string

	/* DollarSkip.

	   Number of records you want to skip

	   Format: int32
	*/
	DollarSkip *int32

	/* DollarTop.

	   Number of records you want

	   Format: int32
	*/
	DollarTop *int32

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get types using g e t 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTypesUsingGET4Params) WithDefaults() *GetTypesUsingGET4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get types using g e t 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTypesUsingGET4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) WithTimeout(timeout time.Duration) *GetTypesUsingGET4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) WithContext(ctx context.Context) *GetTypesUsingGET4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) WithHTTPClient(client *http.Client) *GetTypesUsingGET4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) WithDollarOrderby(dollarOrderby []string) *GetTypesUsingGET4Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) WithDollarSkip(dollarSkip *int32) *GetTypesUsingGET4Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) WithDollarTop(dollarTop *int32) *GetTypesUsingGET4Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) WithAPIVersion(aPIVersion *string) *GetTypesUsingGET4Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get types using g e t 4 params
func (o *GetTypesUsingGET4Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetTypesUsingGET4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarOrderby != nil {

		// binding items for $orderby
		joinedDollarOrderby := o.bindParamDollarOrderby(reg)

		// query array param $orderby
		if err := r.SetQueryParam("$orderby", joinedDollarOrderby...); err != nil {
			return err
		}
	}

	if o.DollarSkip != nil {

		// query param $skip
		var qrDollarSkip int32

		if o.DollarSkip != nil {
			qrDollarSkip = *o.DollarSkip
		}
		qDollarSkip := swag.FormatInt32(qrDollarSkip)
		if qDollarSkip != "" {

			if err := r.SetQueryParam("$skip", qDollarSkip); err != nil {
				return err
			}
		}
	}

	if o.DollarTop != nil {

		// query param $top
		var qrDollarTop int32

		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := swag.FormatInt32(qrDollarTop)
		if qDollarTop != "" {

			if err := r.SetQueryParam("$top", qDollarTop); err != nil {
				return err
			}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetTypesUsingGET4 binds the parameter $orderby
func (o *GetTypesUsingGET4Params) bindParamDollarOrderby(formats strfmt.Registry) []string {
	dollarOrderbyIR := o.DollarOrderby

	var dollarOrderbyIC []string
	for _, dollarOrderbyIIR := range dollarOrderbyIR { // explode []string

		dollarOrderbyIIV := dollarOrderbyIIR // string as string
		dollarOrderbyIC = append(dollarOrderbyIC, dollarOrderbyIIV)
	}

	// items.CollectionFormat: "multi"
	dollarOrderbyIS := swag.JoinByFormat(dollarOrderbyIC, "multi")

	return dollarOrderbyIS
}
