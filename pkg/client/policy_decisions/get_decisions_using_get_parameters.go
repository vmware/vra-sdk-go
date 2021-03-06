// Code generated by go-swagger; DO NOT EDIT.

package policy_decisions

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

// NewGetDecisionsUsingGETParams creates a new GetDecisionsUsingGETParams object
// with the default values initialized.
func NewGetDecisionsUsingGETParams() *GetDecisionsUsingGETParams {
	var ()
	return &GetDecisionsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDecisionsUsingGETParamsWithTimeout creates a new GetDecisionsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDecisionsUsingGETParamsWithTimeout(timeout time.Duration) *GetDecisionsUsingGETParams {
	var ()
	return &GetDecisionsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetDecisionsUsingGETParamsWithContext creates a new GetDecisionsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDecisionsUsingGETParamsWithContext(ctx context.Context) *GetDecisionsUsingGETParams {
	var ()
	return &GetDecisionsUsingGETParams{

		Context: ctx,
	}
}

// NewGetDecisionsUsingGETParamsWithHTTPClient creates a new GetDecisionsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDecisionsUsingGETParamsWithHTTPClient(client *http.Client) *GetDecisionsUsingGETParams {
	var ()
	return &GetDecisionsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetDecisionsUsingGETParams contains all the parameters to send to the API endpoint
for the get decisions using g e t operation typically these are written to a http.Request
*/
type GetDecisionsUsingGETParams struct {

	/*DollarOrderby
	  Sorting criteria in the format: property (asc|desc). Default sort order is ascending. Multiple sort criteria are supported.

	*/
	DollarOrderby []string
	/*DollarSkip
	  Number of records you want to skip

	*/
	DollarSkip *int32
	/*DollarTop
	  Number of records you want

	*/
	DollarTop *int32
	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.

	*/
	APIVersion *string
	/*DryRunID
	  dryRunId

	*/
	DryRunID *strfmt.UUID
	/*PolicyTypeID
	  Matches will only include policies of this type

	*/
	PolicyTypeID *string
	/*ProjectID
	  Matches will only include decisions with this project ID

	*/
	ProjectID *string
	/*Search
	  Matches will start with this string in their policy name or target name or have this string somewhere in their description.

	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get decisions using get params
func (o *GetDecisionsUsingGETParams) WithTimeout(timeout time.Duration) *GetDecisionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get decisions using get params
func (o *GetDecisionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get decisions using get params
func (o *GetDecisionsUsingGETParams) WithContext(ctx context.Context) *GetDecisionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get decisions using get params
func (o *GetDecisionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get decisions using get params
func (o *GetDecisionsUsingGETParams) WithHTTPClient(client *http.Client) *GetDecisionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get decisions using get params
func (o *GetDecisionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the get decisions using get params
func (o *GetDecisionsUsingGETParams) WithDollarOrderby(dollarOrderby []string) *GetDecisionsUsingGETParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get decisions using get params
func (o *GetDecisionsUsingGETParams) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get decisions using get params
func (o *GetDecisionsUsingGETParams) WithDollarSkip(dollarSkip *int32) *GetDecisionsUsingGETParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get decisions using get params
func (o *GetDecisionsUsingGETParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get decisions using get params
func (o *GetDecisionsUsingGETParams) WithDollarTop(dollarTop *int32) *GetDecisionsUsingGETParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get decisions using get params
func (o *GetDecisionsUsingGETParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get decisions using get params
func (o *GetDecisionsUsingGETParams) WithAPIVersion(aPIVersion *string) *GetDecisionsUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get decisions using get params
func (o *GetDecisionsUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDryRunID adds the dryRunID to the get decisions using get params
func (o *GetDecisionsUsingGETParams) WithDryRunID(dryRunID *strfmt.UUID) *GetDecisionsUsingGETParams {
	o.SetDryRunID(dryRunID)
	return o
}

// SetDryRunID adds the dryRunId to the get decisions using get params
func (o *GetDecisionsUsingGETParams) SetDryRunID(dryRunID *strfmt.UUID) {
	o.DryRunID = dryRunID
}

// WithPolicyTypeID adds the policyTypeID to the get decisions using get params
func (o *GetDecisionsUsingGETParams) WithPolicyTypeID(policyTypeID *string) *GetDecisionsUsingGETParams {
	o.SetPolicyTypeID(policyTypeID)
	return o
}

// SetPolicyTypeID adds the policyTypeId to the get decisions using get params
func (o *GetDecisionsUsingGETParams) SetPolicyTypeID(policyTypeID *string) {
	o.PolicyTypeID = policyTypeID
}

// WithProjectID adds the projectID to the get decisions using get params
func (o *GetDecisionsUsingGETParams) WithProjectID(projectID *string) *GetDecisionsUsingGETParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get decisions using get params
func (o *GetDecisionsUsingGETParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithSearch adds the search to the get decisions using get params
func (o *GetDecisionsUsingGETParams) WithSearch(search *string) *GetDecisionsUsingGETParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get decisions using get params
func (o *GetDecisionsUsingGETParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetDecisionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDollarOrderby := o.DollarOrderby

	joinedDollarOrderby := swag.JoinByFormat(valuesDollarOrderby, "multi")
	// query array param $orderby
	if err := r.SetQueryParam("$orderby", joinedDollarOrderby...); err != nil {
		return err
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

	if o.DryRunID != nil {

		// query param dryRunId
		var qrDryRunID strfmt.UUID
		if o.DryRunID != nil {
			qrDryRunID = *o.DryRunID
		}
		qDryRunID := qrDryRunID.String()
		if qDryRunID != "" {
			if err := r.SetQueryParam("dryRunId", qDryRunID); err != nil {
				return err
			}
		}

	}

	if o.PolicyTypeID != nil {

		// query param policyTypeId
		var qrPolicyTypeID string
		if o.PolicyTypeID != nil {
			qrPolicyTypeID = *o.PolicyTypeID
		}
		qPolicyTypeID := qrPolicyTypeID
		if qPolicyTypeID != "" {
			if err := r.SetQueryParam("policyTypeId", qPolicyTypeID); err != nil {
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
