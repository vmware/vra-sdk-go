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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAllUsingGET2Params creates a new GetAllUsingGET2Params object
// with the default values initialized.
func NewGetAllUsingGET2Params() *GetAllUsingGET2Params {
	var (
		dollarOrderbyDefault = string("name desc")
		dollarSkipDefault    = string("0")
		dollarTopDefault     = string("25")
		pageDefault          = string("0")
	)
	return &GetAllUsingGET2Params{
		DollarOrderby: &dollarOrderbyDefault,
		DollarSkip:    &dollarSkipDefault,
		DollarTop:     &dollarTopDefault,
		Page:          &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllUsingGET2ParamsWithTimeout creates a new GetAllUsingGET2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllUsingGET2ParamsWithTimeout(timeout time.Duration) *GetAllUsingGET2Params {
	var (
		dollarOrderbyDefault = string("name desc")
		dollarSkipDefault    = string("0")
		dollarTopDefault     = string("25")
		pageDefault          = string("0")
	)
	return &GetAllUsingGET2Params{
		DollarOrderby: &dollarOrderbyDefault,
		DollarSkip:    &dollarSkipDefault,
		DollarTop:     &dollarTopDefault,
		Page:          &pageDefault,

		timeout: timeout,
	}
}

// NewGetAllUsingGET2ParamsWithContext creates a new GetAllUsingGET2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllUsingGET2ParamsWithContext(ctx context.Context) *GetAllUsingGET2Params {
	var (
		dollarOrderbyDefault = string("name desc")
		dollarSkipDefault    = string("0")
		dollarTopDefault     = string("25")
		pageDefault          = string("0")
	)
	return &GetAllUsingGET2Params{
		DollarOrderby: &dollarOrderbyDefault,
		DollarSkip:    &dollarSkipDefault,
		DollarTop:     &dollarTopDefault,
		Page:          &pageDefault,

		Context: ctx,
	}
}

// NewGetAllUsingGET2ParamsWithHTTPClient creates a new GetAllUsingGET2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllUsingGET2ParamsWithHTTPClient(client *http.Client) *GetAllUsingGET2Params {
	var (
		dollarOrderbyDefault = string("name desc")
		dollarSkipDefault    = string("0")
		dollarTopDefault     = string("25")
		pageDefault          = string("0")
	)
	return &GetAllUsingGET2Params{
		DollarOrderby: &dollarOrderbyDefault,
		DollarSkip:    &dollarSkipDefault,
		DollarTop:     &dollarTopDefault,
		Page:          &pageDefault,
		HTTPClient:    client,
	}
}

/*GetAllUsingGET2Params contains all the parameters to send to the API endpoint
for the get all using g e t 2 operation typically these are written to a http.Request
*/
type GetAllUsingGET2Params struct {

	/*DollarFilter
	  To list gerrit triggers with OData like filter

	*/
	DollarFilter *string
	/*DollarOrderby
	  Order by attribute

	*/
	DollarOrderby *string
	/*DollarSkip
	  To skip 'n' gerrit triggers for listing

	*/
	DollarSkip *string
	/*DollarTop
	  To list top 'n' gerrit triggers for listing

	*/
	DollarTop *string
	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*Page
	  To select 'n'th page for listing

	*/
	Page *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) WithTimeout(timeout time.Duration) *GetAllUsingGET2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) WithContext(ctx context.Context) *GetAllUsingGET2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) WithHTTPClient(client *http.Client) *GetAllUsingGET2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) WithDollarFilter(dollarFilter *string) *GetAllUsingGET2Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarOrderby adds the dollarOrderby to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) WithDollarOrderby(dollarOrderby *string) *GetAllUsingGET2Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) WithDollarSkip(dollarSkip *string) *GetAllUsingGET2Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) WithDollarTop(dollarTop *string) *GetAllUsingGET2Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) WithAPIVersion(aPIVersion string) *GetAllUsingGET2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithPage adds the page to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) WithPage(page *string) *GetAllUsingGET2Params {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get all using g e t 2 params
func (o *GetAllUsingGET2Params) SetPage(page *string) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllUsingGET2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
