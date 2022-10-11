// Code generated by go-swagger; DO NOT EDIT.

package pricing_cards

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

// NewGetPoliciesUsingGET4Params creates a new GetPoliciesUsingGET4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPoliciesUsingGET4Params() *GetPoliciesUsingGET4Params {
	return &GetPoliciesUsingGET4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPoliciesUsingGET4ParamsWithTimeout creates a new GetPoliciesUsingGET4Params object
// with the ability to set a timeout on a request.
func NewGetPoliciesUsingGET4ParamsWithTimeout(timeout time.Duration) *GetPoliciesUsingGET4Params {
	return &GetPoliciesUsingGET4Params{
		timeout: timeout,
	}
}

// NewGetPoliciesUsingGET4ParamsWithContext creates a new GetPoliciesUsingGET4Params object
// with the ability to set a context for a request.
func NewGetPoliciesUsingGET4ParamsWithContext(ctx context.Context) *GetPoliciesUsingGET4Params {
	return &GetPoliciesUsingGET4Params{
		Context: ctx,
	}
}

// NewGetPoliciesUsingGET4ParamsWithHTTPClient creates a new GetPoliciesUsingGET4Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetPoliciesUsingGET4ParamsWithHTTPClient(client *http.Client) *GetPoliciesUsingGET4Params {
	return &GetPoliciesUsingGET4Params{
		HTTPClient: client,
	}
}

/*
GetPoliciesUsingGET4Params contains all the parameters to send to the API endpoint

	for the get policies using g e t 4 operation.

	Typically these are written to a http.Request.
*/
type GetPoliciesUsingGET4Params struct {

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

	/* ExpandAssignmentInfo.

	   Whether or not returns count of assignments.
	*/
	ExpandAssignmentInfo *bool

	/* ExpandPricingCard.

	   Whether or not returns detailed pricing card for each result.
	*/
	ExpandPricingCard *bool

	/* Search.

	   Search by name and description
	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get policies using g e t 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPoliciesUsingGET4Params) WithDefaults() *GetPoliciesUsingGET4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get policies using g e t 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPoliciesUsingGET4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) WithTimeout(timeout time.Duration) *GetPoliciesUsingGET4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) WithContext(ctx context.Context) *GetPoliciesUsingGET4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) WithHTTPClient(client *http.Client) *GetPoliciesUsingGET4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) WithDollarOrderby(dollarOrderby []string) *GetPoliciesUsingGET4Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) WithDollarSkip(dollarSkip *int32) *GetPoliciesUsingGET4Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) WithDollarTop(dollarTop *int32) *GetPoliciesUsingGET4Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) WithAPIVersion(aPIVersion *string) *GetPoliciesUsingGET4Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithExpandAssignmentInfo adds the expandAssignmentInfo to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) WithExpandAssignmentInfo(expandAssignmentInfo *bool) *GetPoliciesUsingGET4Params {
	o.SetExpandAssignmentInfo(expandAssignmentInfo)
	return o
}

// SetExpandAssignmentInfo adds the expandAssignmentInfo to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) SetExpandAssignmentInfo(expandAssignmentInfo *bool) {
	o.ExpandAssignmentInfo = expandAssignmentInfo
}

// WithExpandPricingCard adds the expandPricingCard to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) WithExpandPricingCard(expandPricingCard *bool) *GetPoliciesUsingGET4Params {
	o.SetExpandPricingCard(expandPricingCard)
	return o
}

// SetExpandPricingCard adds the expandPricingCard to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) SetExpandPricingCard(expandPricingCard *bool) {
	o.ExpandPricingCard = expandPricingCard
}

// WithSearch adds the search to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) WithSearch(search *string) *GetPoliciesUsingGET4Params {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get policies using g e t 4 params
func (o *GetPoliciesUsingGET4Params) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetPoliciesUsingGET4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ExpandAssignmentInfo != nil {

		// query param expandAssignmentInfo
		var qrExpandAssignmentInfo bool

		if o.ExpandAssignmentInfo != nil {
			qrExpandAssignmentInfo = *o.ExpandAssignmentInfo
		}
		qExpandAssignmentInfo := swag.FormatBool(qrExpandAssignmentInfo)
		if qExpandAssignmentInfo != "" {

			if err := r.SetQueryParam("expandAssignmentInfo", qExpandAssignmentInfo); err != nil {
				return err
			}
		}
	}

	if o.ExpandPricingCard != nil {

		// query param expandPricingCard
		var qrExpandPricingCard bool

		if o.ExpandPricingCard != nil {
			qrExpandPricingCard = *o.ExpandPricingCard
		}
		qExpandPricingCard := swag.FormatBool(qrExpandPricingCard)
		if qExpandPricingCard != "" {

			if err := r.SetQueryParam("expandPricingCard", qExpandPricingCard); err != nil {
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

// bindParamGetPoliciesUsingGET4 binds the parameter $orderby
func (o *GetPoliciesUsingGET4Params) bindParamDollarOrderby(formats strfmt.Registry) []string {
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
