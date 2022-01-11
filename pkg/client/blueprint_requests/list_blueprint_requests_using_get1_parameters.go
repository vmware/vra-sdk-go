// Code generated by go-swagger; DO NOT EDIT.

package blueprint_requests

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

// NewListBlueprintRequestsUsingGET1Params creates a new ListBlueprintRequestsUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListBlueprintRequestsUsingGET1Params() *ListBlueprintRequestsUsingGET1Params {
	return &ListBlueprintRequestsUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewListBlueprintRequestsUsingGET1ParamsWithTimeout creates a new ListBlueprintRequestsUsingGET1Params object
// with the ability to set a timeout on a request.
func NewListBlueprintRequestsUsingGET1ParamsWithTimeout(timeout time.Duration) *ListBlueprintRequestsUsingGET1Params {
	return &ListBlueprintRequestsUsingGET1Params{
		timeout: timeout,
	}
}

// NewListBlueprintRequestsUsingGET1ParamsWithContext creates a new ListBlueprintRequestsUsingGET1Params object
// with the ability to set a context for a request.
func NewListBlueprintRequestsUsingGET1ParamsWithContext(ctx context.Context) *ListBlueprintRequestsUsingGET1Params {
	return &ListBlueprintRequestsUsingGET1Params{
		Context: ctx,
	}
}

// NewListBlueprintRequestsUsingGET1ParamsWithHTTPClient creates a new ListBlueprintRequestsUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewListBlueprintRequestsUsingGET1ParamsWithHTTPClient(client *http.Client) *ListBlueprintRequestsUsingGET1Params {
	return &ListBlueprintRequestsUsingGET1Params{
		HTTPClient: client,
	}
}

/* ListBlueprintRequestsUsingGET1Params contains all the parameters to send to the API endpoint
   for the list blueprint requests using get1 operation.

   Typically these are written to a http.Request.
*/
type ListBlueprintRequestsUsingGET1Params struct {

	/* DollarOrderby.

	   Sorting criteria in the format: property (asc|desc). Default sort order is descending on updatedAt. Sorting is supported on fields createdAt, updatedAt, createdBy, updatedBy.
	*/
	DollarOrderby []string

	/* DollarSelect.

	   Fields to include in content.
	*/
	DollarSelect []string

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

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* DeploymentID.

	   Deployment Id filter
	*/
	DeploymentID *string

	/* IncludePlan.

	   Plan Requests filter
	*/
	IncludePlan *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list blueprint requests using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBlueprintRequestsUsingGET1Params) WithDefaults() *ListBlueprintRequestsUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list blueprint requests using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBlueprintRequestsUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) WithTimeout(timeout time.Duration) *ListBlueprintRequestsUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) WithContext(ctx context.Context) *ListBlueprintRequestsUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) WithHTTPClient(client *http.Client) *ListBlueprintRequestsUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) WithDollarOrderby(dollarOrderby []string) *ListBlueprintRequestsUsingGET1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) WithDollarSelect(dollarSelect []string) *ListBlueprintRequestsUsingGET1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) SetDollarSelect(dollarSelect []string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) WithDollarSkip(dollarSkip *int32) *ListBlueprintRequestsUsingGET1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) WithDollarTop(dollarTop *int32) *ListBlueprintRequestsUsingGET1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) WithAPIVersion(aPIVersion *string) *ListBlueprintRequestsUsingGET1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDeploymentID adds the deploymentID to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) WithDeploymentID(deploymentID *string) *ListBlueprintRequestsUsingGET1Params {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) SetDeploymentID(deploymentID *string) {
	o.DeploymentID = deploymentID
}

// WithIncludePlan adds the includePlan to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) WithIncludePlan(includePlan *bool) *ListBlueprintRequestsUsingGET1Params {
	o.SetIncludePlan(includePlan)
	return o
}

// SetIncludePlan adds the includePlan to the list blueprint requests using get1 params
func (o *ListBlueprintRequestsUsingGET1Params) SetIncludePlan(includePlan *bool) {
	o.IncludePlan = includePlan
}

// WriteToRequest writes these params to a swagger request
func (o *ListBlueprintRequestsUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DollarSelect != nil {

		// binding items for $select
		joinedDollarSelect := o.bindParamDollarSelect(reg)

		// query array param $select
		if err := r.SetQueryParam("$select", joinedDollarSelect...); err != nil {
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

	if o.DeploymentID != nil {

		// query param deploymentId
		var qrDeploymentID string

		if o.DeploymentID != nil {
			qrDeploymentID = *o.DeploymentID
		}
		qDeploymentID := qrDeploymentID
		if qDeploymentID != "" {

			if err := r.SetQueryParam("deploymentId", qDeploymentID); err != nil {
				return err
			}
		}
	}

	if o.IncludePlan != nil {

		// query param includePlan
		var qrIncludePlan bool

		if o.IncludePlan != nil {
			qrIncludePlan = *o.IncludePlan
		}
		qIncludePlan := swag.FormatBool(qrIncludePlan)
		if qIncludePlan != "" {

			if err := r.SetQueryParam("includePlan", qIncludePlan); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListBlueprintRequestsUsingGET1 binds the parameter $orderby
func (o *ListBlueprintRequestsUsingGET1Params) bindParamDollarOrderby(formats strfmt.Registry) []string {
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

// bindParamListBlueprintRequestsUsingGET1 binds the parameter $select
func (o *ListBlueprintRequestsUsingGET1Params) bindParamDollarSelect(formats strfmt.Registry) []string {
	dollarSelectIR := o.DollarSelect

	var dollarSelectIC []string
	for _, dollarSelectIIR := range dollarSelectIR { // explode []string

		dollarSelectIIV := dollarSelectIIR // string as string
		dollarSelectIC = append(dollarSelectIC, dollarSelectIIV)
	}

	// items.CollectionFormat: "multi"
	dollarSelectIS := swag.JoinByFormat(dollarSelectIC, "multi")

	return dollarSelectIS
}
