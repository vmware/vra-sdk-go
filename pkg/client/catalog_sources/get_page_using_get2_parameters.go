// Code generated by go-swagger; DO NOT EDIT.

package catalog_sources

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

// NewGetPageUsingGET2Params creates a new GetPageUsingGET2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPageUsingGET2Params() *GetPageUsingGET2Params {
	return &GetPageUsingGET2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPageUsingGET2ParamsWithTimeout creates a new GetPageUsingGET2Params object
// with the ability to set a timeout on a request.
func NewGetPageUsingGET2ParamsWithTimeout(timeout time.Duration) *GetPageUsingGET2Params {
	return &GetPageUsingGET2Params{
		timeout: timeout,
	}
}

// NewGetPageUsingGET2ParamsWithContext creates a new GetPageUsingGET2Params object
// with the ability to set a context for a request.
func NewGetPageUsingGET2ParamsWithContext(ctx context.Context) *GetPageUsingGET2Params {
	return &GetPageUsingGET2Params{
		Context: ctx,
	}
}

// NewGetPageUsingGET2ParamsWithHTTPClient creates a new GetPageUsingGET2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetPageUsingGET2ParamsWithHTTPClient(client *http.Client) *GetPageUsingGET2Params {
	return &GetPageUsingGET2Params{
		HTTPClient: client,
	}
}

/*
GetPageUsingGET2Params contains all the parameters to send to the API endpoint

	for the get page using g e t 2 operation.

	Typically these are written to a http.Request.
*/
type GetPageUsingGET2Params struct {

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

	/* ProjectID.

	   Find sources which contains items that can be requested in the given projectId
	*/
	ProjectID *string

	/* Search.

	   Matches will have this string in their name or description.
	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get page using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPageUsingGET2Params) WithDefaults() *GetPageUsingGET2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get page using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPageUsingGET2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) WithTimeout(timeout time.Duration) *GetPageUsingGET2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) WithContext(ctx context.Context) *GetPageUsingGET2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) WithHTTPClient(client *http.Client) *GetPageUsingGET2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) WithDollarOrderby(dollarOrderby []string) *GetPageUsingGET2Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) WithDollarSkip(dollarSkip *int32) *GetPageUsingGET2Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) WithDollarTop(dollarTop *int32) *GetPageUsingGET2Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) WithAPIVersion(aPIVersion *string) *GetPageUsingGET2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithProjectID adds the projectID to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) WithProjectID(projectID *string) *GetPageUsingGET2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithSearch adds the search to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) WithSearch(search *string) *GetPageUsingGET2Params {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get page using g e t 2 params
func (o *GetPageUsingGET2Params) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetPageUsingGET2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID string

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {

			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
				return err
			}
		}
	}

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetPageUsingGET2 binds the parameter $orderby
func (o *GetPageUsingGET2Params) bindParamDollarOrderby(formats strfmt.Registry) []string {
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