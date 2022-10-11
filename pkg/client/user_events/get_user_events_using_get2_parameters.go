// Code generated by go-swagger; DO NOT EDIT.

package user_events

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

// NewGetUserEventsUsingGET2Params creates a new GetUserEventsUsingGET2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserEventsUsingGET2Params() *GetUserEventsUsingGET2Params {
	return &GetUserEventsUsingGET2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserEventsUsingGET2ParamsWithTimeout creates a new GetUserEventsUsingGET2Params object
// with the ability to set a timeout on a request.
func NewGetUserEventsUsingGET2ParamsWithTimeout(timeout time.Duration) *GetUserEventsUsingGET2Params {
	return &GetUserEventsUsingGET2Params{
		timeout: timeout,
	}
}

// NewGetUserEventsUsingGET2ParamsWithContext creates a new GetUserEventsUsingGET2Params object
// with the ability to set a context for a request.
func NewGetUserEventsUsingGET2ParamsWithContext(ctx context.Context) *GetUserEventsUsingGET2Params {
	return &GetUserEventsUsingGET2Params{
		Context: ctx,
	}
}

// NewGetUserEventsUsingGET2ParamsWithHTTPClient creates a new GetUserEventsUsingGET2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserEventsUsingGET2ParamsWithHTTPClient(client *http.Client) *GetUserEventsUsingGET2Params {
	return &GetUserEventsUsingGET2Params{
		HTTPClient: client,
	}
}

/*
GetUserEventsUsingGET2Params contains all the parameters to send to the API endpoint

	for the get user events using g e t 2 operation.

	Typically these are written to a http.Request.
*/
type GetUserEventsUsingGET2Params struct {

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

	/* DeploymentID.

	   Deployment ID

	   Format: uuid
	*/
	DeploymentID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user events using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserEventsUsingGET2Params) WithDefaults() *GetUserEventsUsingGET2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user events using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserEventsUsingGET2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) WithTimeout(timeout time.Duration) *GetUserEventsUsingGET2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) WithContext(ctx context.Context) *GetUserEventsUsingGET2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) WithHTTPClient(client *http.Client) *GetUserEventsUsingGET2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) WithDollarOrderby(dollarOrderby []string) *GetUserEventsUsingGET2Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) WithDollarSkip(dollarSkip *int32) *GetUserEventsUsingGET2Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) WithDollarTop(dollarTop *int32) *GetUserEventsUsingGET2Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) WithAPIVersion(aPIVersion *string) *GetUserEventsUsingGET2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDeploymentID adds the deploymentID to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) WithDeploymentID(deploymentID strfmt.UUID) *GetUserEventsUsingGET2Params {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the get user events using g e t 2 params
func (o *GetUserEventsUsingGET2Params) SetDeploymentID(deploymentID strfmt.UUID) {
	o.DeploymentID = deploymentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserEventsUsingGET2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deploymentId
	if err := r.SetPathParam("deploymentId", o.DeploymentID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetUserEventsUsingGET2 binds the parameter $orderby
func (o *GetUserEventsUsingGET2Params) bindParamDollarOrderby(formats strfmt.Registry) []string {
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