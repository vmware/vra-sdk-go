// Code generated by go-swagger; DO NOT EDIT.

package property_groups

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

// NewListPropertyGroupsUsingGETParams creates a new ListPropertyGroupsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPropertyGroupsUsingGETParams() *ListPropertyGroupsUsingGETParams {
	return &ListPropertyGroupsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPropertyGroupsUsingGETParamsWithTimeout creates a new ListPropertyGroupsUsingGETParams object
// with the ability to set a timeout on a request.
func NewListPropertyGroupsUsingGETParamsWithTimeout(timeout time.Duration) *ListPropertyGroupsUsingGETParams {
	return &ListPropertyGroupsUsingGETParams{
		timeout: timeout,
	}
}

// NewListPropertyGroupsUsingGETParamsWithContext creates a new ListPropertyGroupsUsingGETParams object
// with the ability to set a context for a request.
func NewListPropertyGroupsUsingGETParamsWithContext(ctx context.Context) *ListPropertyGroupsUsingGETParams {
	return &ListPropertyGroupsUsingGETParams{
		Context: ctx,
	}
}

// NewListPropertyGroupsUsingGETParamsWithHTTPClient creates a new ListPropertyGroupsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPropertyGroupsUsingGETParamsWithHTTPClient(client *http.Client) *ListPropertyGroupsUsingGETParams {
	return &ListPropertyGroupsUsingGETParams{
		HTTPClient: client,
	}
}

/* ListPropertyGroupsUsingGETParams contains all the parameters to send to the API endpoint
   for the list property groups using g e t operation.

   Typically these are written to a http.Request.
*/
type ListPropertyGroupsUsingGETParams struct {

	/* DollarOrderby.

	   Sorting criteria in the format: property (asc|desc). Default sort order is descending on updatedAt. Sorting is supported on fields createdAt, updatedAt, createdBy, updatedBy, name, type.
	*/
	DollarOrderby []string

	/* DollarSelect.

	   Fields to include.
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

	/* Name.

	   Filter by name
	*/
	Name *string

	/* Projects.

	   A comma-separated list. Results must be associated with one of these project IDs.
	*/
	Projects []string

	/* Search.

	   Search by name and description
	*/
	Search *string

	/* Type.

	   Filter by type
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list property groups using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPropertyGroupsUsingGETParams) WithDefaults() *ListPropertyGroupsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list property groups using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPropertyGroupsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) WithTimeout(timeout time.Duration) *ListPropertyGroupsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) WithContext(ctx context.Context) *ListPropertyGroupsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) WithHTTPClient(client *http.Client) *ListPropertyGroupsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) WithDollarOrderby(dollarOrderby []string) *ListPropertyGroupsUsingGETParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) WithDollarSelect(dollarSelect []string) *ListPropertyGroupsUsingGETParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) SetDollarSelect(dollarSelect []string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) WithDollarSkip(dollarSkip *int32) *ListPropertyGroupsUsingGETParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) WithDollarTop(dollarTop *int32) *ListPropertyGroupsUsingGETParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) WithAPIVersion(aPIVersion *string) *ListPropertyGroupsUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithName adds the name to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) WithName(name *string) *ListPropertyGroupsUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) SetName(name *string) {
	o.Name = name
}

// WithProjects adds the projects to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) WithProjects(projects []string) *ListPropertyGroupsUsingGETParams {
	o.SetProjects(projects)
	return o
}

// SetProjects adds the projects to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) SetProjects(projects []string) {
	o.Projects = projects
}

// WithSearch adds the search to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) WithSearch(search *string) *ListPropertyGroupsUsingGETParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) SetSearch(search *string) {
	o.Search = search
}

// WithType adds the typeVar to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) WithType(typeVar *string) *ListPropertyGroupsUsingGETParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the list property groups using get params
func (o *ListPropertyGroupsUsingGETParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ListPropertyGroupsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Projects != nil {

		// binding items for projects
		joinedProjects := o.bindParamProjects(reg)

		// query array param projects
		if err := r.SetQueryParam("projects", joinedProjects...); err != nil {
			return err
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

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListPropertyGroupsUsingGET binds the parameter $orderby
func (o *ListPropertyGroupsUsingGETParams) bindParamDollarOrderby(formats strfmt.Registry) []string {
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

// bindParamListPropertyGroupsUsingGET binds the parameter $select
func (o *ListPropertyGroupsUsingGETParams) bindParamDollarSelect(formats strfmt.Registry) []string {
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

// bindParamListPropertyGroupsUsingGET binds the parameter projects
func (o *ListPropertyGroupsUsingGETParams) bindParamProjects(formats strfmt.Registry) []string {
	projectsIR := o.Projects

	var projectsIC []string
	for _, projectsIIR := range projectsIR { // explode []string

		projectsIIV := projectsIIR // string as string
		projectsIC = append(projectsIC, projectsIIV)
	}

	// items.CollectionFormat: "multi"
	projectsIS := swag.JoinByFormat(projectsIC, "multi")

	return projectsIS
}
