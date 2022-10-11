// Code generated by go-swagger; DO NOT EDIT.

package requests

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

// NewGetDeploymentRequestsUsingGET2Params creates a new GetDeploymentRequestsUsingGET2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentRequestsUsingGET2Params() *GetDeploymentRequestsUsingGET2Params {
	return &GetDeploymentRequestsUsingGET2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentRequestsUsingGET2ParamsWithTimeout creates a new GetDeploymentRequestsUsingGET2Params object
// with the ability to set a timeout on a request.
func NewGetDeploymentRequestsUsingGET2ParamsWithTimeout(timeout time.Duration) *GetDeploymentRequestsUsingGET2Params {
	return &GetDeploymentRequestsUsingGET2Params{
		timeout: timeout,
	}
}

// NewGetDeploymentRequestsUsingGET2ParamsWithContext creates a new GetDeploymentRequestsUsingGET2Params object
// with the ability to set a context for a request.
func NewGetDeploymentRequestsUsingGET2ParamsWithContext(ctx context.Context) *GetDeploymentRequestsUsingGET2Params {
	return &GetDeploymentRequestsUsingGET2Params{
		Context: ctx,
	}
}

// NewGetDeploymentRequestsUsingGET2ParamsWithHTTPClient creates a new GetDeploymentRequestsUsingGET2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentRequestsUsingGET2ParamsWithHTTPClient(client *http.Client) *GetDeploymentRequestsUsingGET2Params {
	return &GetDeploymentRequestsUsingGET2Params{
		HTTPClient: client,
	}
}

/*
GetDeploymentRequestsUsingGET2Params contains all the parameters to send to the API endpoint

	for the get deployment requests using g e t 2 operation.

	Typically these are written to a http.Request.
*/
type GetDeploymentRequestsUsingGET2Params struct {

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

	/* Deleted.

	   Retrieves the soft-deleted requests that have not yet been completely deleted.
	*/
	Deleted *bool

	/* DeploymentID.

	   Deployment ID

	   Format: uuid
	*/
	DeploymentID strfmt.UUID

	/* InprogressRequests.

	   Retrieves the requests that are currently in-progress for a deployment. Incase of a false value the param is ignored.
	*/
	InprogressRequests *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deployment requests using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentRequestsUsingGET2Params) WithDefaults() *GetDeploymentRequestsUsingGET2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployment requests using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentRequestsUsingGET2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) WithTimeout(timeout time.Duration) *GetDeploymentRequestsUsingGET2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) WithContext(ctx context.Context) *GetDeploymentRequestsUsingGET2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) WithHTTPClient(client *http.Client) *GetDeploymentRequestsUsingGET2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) WithDollarOrderby(dollarOrderby []string) *GetDeploymentRequestsUsingGET2Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) WithDollarSkip(dollarSkip *int32) *GetDeploymentRequestsUsingGET2Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) WithDollarTop(dollarTop *int32) *GetDeploymentRequestsUsingGET2Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) WithAPIVersion(aPIVersion *string) *GetDeploymentRequestsUsingGET2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDeleted adds the deleted to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) WithDeleted(deleted *bool) *GetDeploymentRequestsUsingGET2Params {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithDeploymentID adds the deploymentID to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) WithDeploymentID(deploymentID strfmt.UUID) *GetDeploymentRequestsUsingGET2Params {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) SetDeploymentID(deploymentID strfmt.UUID) {
	o.DeploymentID = deploymentID
}

// WithInprogressRequests adds the inprogressRequests to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) WithInprogressRequests(inprogressRequests *bool) *GetDeploymentRequestsUsingGET2Params {
	o.SetInprogressRequests(inprogressRequests)
	return o
}

// SetInprogressRequests adds the inprogressRequests to the get deployment requests using g e t 2 params
func (o *GetDeploymentRequestsUsingGET2Params) SetInprogressRequests(inprogressRequests *bool) {
	o.InprogressRequests = inprogressRequests
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentRequestsUsingGET2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool

		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {

			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}
	}

	// path param deploymentId
	if err := r.SetPathParam("deploymentId", o.DeploymentID.String()); err != nil {
		return err
	}

	if o.InprogressRequests != nil {

		// query param inprogressRequests
		var qrInprogressRequests bool

		if o.InprogressRequests != nil {
			qrInprogressRequests = *o.InprogressRequests
		}
		qInprogressRequests := swag.FormatBool(qrInprogressRequests)
		if qInprogressRequests != "" {

			if err := r.SetQueryParam("inprogressRequests", qInprogressRequests); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetDeploymentRequestsUsingGET2 binds the parameter $orderby
func (o *GetDeploymentRequestsUsingGET2Params) bindParamDollarOrderby(formats strfmt.Registry) []string {
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
