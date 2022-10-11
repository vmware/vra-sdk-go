// Code generated by go-swagger; DO NOT EDIT.

package triggers

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

// NewGetAllGerritTriggersUsingGETParams creates a new GetAllGerritTriggersUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllGerritTriggersUsingGETParams() *GetAllGerritTriggersUsingGETParams {
	return &GetAllGerritTriggersUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllGerritTriggersUsingGETParamsWithTimeout creates a new GetAllGerritTriggersUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAllGerritTriggersUsingGETParamsWithTimeout(timeout time.Duration) *GetAllGerritTriggersUsingGETParams {
	return &GetAllGerritTriggersUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAllGerritTriggersUsingGETParamsWithContext creates a new GetAllGerritTriggersUsingGETParams object
// with the ability to set a context for a request.
func NewGetAllGerritTriggersUsingGETParamsWithContext(ctx context.Context) *GetAllGerritTriggersUsingGETParams {
	return &GetAllGerritTriggersUsingGETParams{
		Context: ctx,
	}
}

// NewGetAllGerritTriggersUsingGETParamsWithHTTPClient creates a new GetAllGerritTriggersUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllGerritTriggersUsingGETParamsWithHTTPClient(client *http.Client) *GetAllGerritTriggersUsingGETParams {
	return &GetAllGerritTriggersUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetAllGerritTriggersUsingGETParams contains all the parameters to send to the API endpoint

	for the get all gerrit triggers using g e t operation.

	Typically these are written to a http.Request.
*/
type GetAllGerritTriggersUsingGETParams struct {

	/* DollarFilter.

	   To list gerrit triggers with OData like filter
	*/
	DollarFilter *string

	/* DollarOrderby.

	   Order by attribute

	   Default: "name desc"
	*/
	DollarOrderby *string

	/* DollarSkip.

	   To skip 'n' gerrit triggers for listing

	   Default: "0"
	*/
	DollarSkip *string

	/* DollarTop.

	   To list top 'n' gerrit triggers for listing

	   Default: "25"
	*/
	DollarTop *string

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Page.

	   To select 'n'th page for listing

	   Default: "0"
	*/
	Page *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all gerrit triggers using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllGerritTriggersUsingGETParams) WithDefaults() *GetAllGerritTriggersUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all gerrit triggers using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllGerritTriggersUsingGETParams) SetDefaults() {
	var (
		dollarOrderbyDefault = string("name desc")

		dollarSkipDefault = string("0")

		dollarTopDefault = string("25")

		pageDefault = string("0")
	)

	val := GetAllGerritTriggersUsingGETParams{
		DollarOrderby: &dollarOrderbyDefault,
		DollarSkip:    &dollarSkipDefault,
		DollarTop:     &dollarTopDefault,
		Page:          &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) WithTimeout(timeout time.Duration) *GetAllGerritTriggersUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) WithContext(ctx context.Context) *GetAllGerritTriggersUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) WithHTTPClient(client *http.Client) *GetAllGerritTriggersUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) WithDollarFilter(dollarFilter *string) *GetAllGerritTriggersUsingGETParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarOrderby adds the dollarOrderby to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) WithDollarOrderby(dollarOrderby *string) *GetAllGerritTriggersUsingGETParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) WithDollarSkip(dollarSkip *string) *GetAllGerritTriggersUsingGETParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) WithDollarTop(dollarTop *string) *GetAllGerritTriggersUsingGETParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithAuthorization adds the authorization to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) WithAuthorization(authorization string) *GetAllGerritTriggersUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) WithAPIVersion(aPIVersion string) *GetAllGerritTriggersUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithPage adds the page to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) WithPage(page *string) *GetAllGerritTriggersUsingGETParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get all gerrit triggers using get params
func (o *GetAllGerritTriggersUsingGETParams) SetPage(page *string) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllGerritTriggersUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarFilter != nil {

		// query param $filter
		var qrDollarFilter string

		if o.DollarFilter != nil {
			qrDollarFilter = *o.DollarFilter
		}
		qDollarFilter := qrDollarFilter
		if qDollarFilter != "" {

			if err := r.SetQueryParam("$filter", qDollarFilter); err != nil {
				return err
			}
		}
	}

	if o.DollarOrderby != nil {

		// query param $orderby
		var qrDollarOrderby string

		if o.DollarOrderby != nil {
			qrDollarOrderby = *o.DollarOrderby
		}
		qDollarOrderby := qrDollarOrderby
		if qDollarOrderby != "" {

			if err := r.SetQueryParam("$orderby", qDollarOrderby); err != nil {
				return err
			}
		}
	}

	if o.DollarSkip != nil {

		// query param $skip
		var qrDollarSkip string

		if o.DollarSkip != nil {
			qrDollarSkip = *o.DollarSkip
		}
		qDollarSkip := qrDollarSkip
		if qDollarSkip != "" {

			if err := r.SetQueryParam("$skip", qDollarSkip); err != nil {
				return err
			}
		}
	}

	if o.DollarTop != nil {

		// query param $top
		var qrDollarTop string

		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := qrDollarTop
		if qDollarTop != "" {

			if err := r.SetQueryParam("$top", qDollarTop); err != nil {
				return err
			}
		}
	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion

	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage string

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
