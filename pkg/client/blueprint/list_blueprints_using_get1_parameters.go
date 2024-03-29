// Code generated by go-swagger; DO NOT EDIT.

package blueprint

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

// NewListBlueprintsUsingGET1Params creates a new ListBlueprintsUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListBlueprintsUsingGET1Params() *ListBlueprintsUsingGET1Params {
	return &ListBlueprintsUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewListBlueprintsUsingGET1ParamsWithTimeout creates a new ListBlueprintsUsingGET1Params object
// with the ability to set a timeout on a request.
func NewListBlueprintsUsingGET1ParamsWithTimeout(timeout time.Duration) *ListBlueprintsUsingGET1Params {
	return &ListBlueprintsUsingGET1Params{
		timeout: timeout,
	}
}

// NewListBlueprintsUsingGET1ParamsWithContext creates a new ListBlueprintsUsingGET1Params object
// with the ability to set a context for a request.
func NewListBlueprintsUsingGET1ParamsWithContext(ctx context.Context) *ListBlueprintsUsingGET1Params {
	return &ListBlueprintsUsingGET1Params{
		Context: ctx,
	}
}

// NewListBlueprintsUsingGET1ParamsWithHTTPClient creates a new ListBlueprintsUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewListBlueprintsUsingGET1ParamsWithHTTPClient(client *http.Client) *ListBlueprintsUsingGET1Params {
	return &ListBlueprintsUsingGET1Params{
		HTTPClient: client,
	}
}

/*
ListBlueprintsUsingGET1Params contains all the parameters to send to the API endpoint

	for the list blueprints using get1 operation.

	Typically these are written to a http.Request.
*/
type ListBlueprintsUsingGET1Params struct {

	/* DollarOrderby.

	   Sorting criteria in the format: property (asc|desc). Default sort order is descending on updatedAt. Sorting is supported on fields createdAt, updatedAt, createdBy, updatedBy, name.
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

	/* Name.

	   Filter by name
	*/
	Name *string

	/* Projects.

	   A comma-separated list. Results must be associated with one of these project IDs.
	*/
	Projects []string

	/* PropertyGroups.

	   Filter blueprints with any of the specified property groups
	*/
	PropertyGroups []string

	/* Released.

	   Filter blueprints with at least one released version
	*/
	Released *bool

	/* Search.

	   Search by name and description
	*/
	Search *string

	/* Versioned.

	   Filter blueprints with at least one version
	*/
	Versioned *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list blueprints using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBlueprintsUsingGET1Params) WithDefaults() *ListBlueprintsUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list blueprints using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBlueprintsUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithTimeout(timeout time.Duration) *ListBlueprintsUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithContext(ctx context.Context) *ListBlueprintsUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithHTTPClient(client *http.Client) *ListBlueprintsUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithDollarOrderby(dollarOrderby []string) *ListBlueprintsUsingGET1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithDollarSelect(dollarSelect []string) *ListBlueprintsUsingGET1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetDollarSelect(dollarSelect []string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithDollarSkip(dollarSkip *int32) *ListBlueprintsUsingGET1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithDollarTop(dollarTop *int32) *ListBlueprintsUsingGET1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithAPIVersion(aPIVersion *string) *ListBlueprintsUsingGET1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithName adds the name to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithName(name *string) *ListBlueprintsUsingGET1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetName(name *string) {
	o.Name = name
}

// WithProjects adds the projects to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithProjects(projects []string) *ListBlueprintsUsingGET1Params {
	o.SetProjects(projects)
	return o
}

// SetProjects adds the projects to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetProjects(projects []string) {
	o.Projects = projects
}

// WithPropertyGroups adds the propertyGroups to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithPropertyGroups(propertyGroups []string) *ListBlueprintsUsingGET1Params {
	o.SetPropertyGroups(propertyGroups)
	return o
}

// SetPropertyGroups adds the propertyGroups to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetPropertyGroups(propertyGroups []string) {
	o.PropertyGroups = propertyGroups
}

// WithReleased adds the released to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithReleased(released *bool) *ListBlueprintsUsingGET1Params {
	o.SetReleased(released)
	return o
}

// SetReleased adds the released to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetReleased(released *bool) {
	o.Released = released
}

// WithSearch adds the search to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithSearch(search *string) *ListBlueprintsUsingGET1Params {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetSearch(search *string) {
	o.Search = search
}

// WithVersioned adds the versioned to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) WithVersioned(versioned *bool) *ListBlueprintsUsingGET1Params {
	o.SetVersioned(versioned)
	return o
}

// SetVersioned adds the versioned to the list blueprints using get1 params
func (o *ListBlueprintsUsingGET1Params) SetVersioned(versioned *bool) {
	o.Versioned = versioned
}

// WriteToRequest writes these params to a swagger request
func (o *ListBlueprintsUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PropertyGroups != nil {

		// binding items for propertyGroups
		joinedPropertyGroups := o.bindParamPropertyGroups(reg)

		// query array param propertyGroups
		if err := r.SetQueryParam("propertyGroups", joinedPropertyGroups...); err != nil {
			return err
		}
	}

	if o.Released != nil {

		// query param released
		var qrReleased bool

		if o.Released != nil {
			qrReleased = *o.Released
		}
		qReleased := swag.FormatBool(qrReleased)
		if qReleased != "" {

			if err := r.SetQueryParam("released", qReleased); err != nil {
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

	if o.Versioned != nil {

		// query param versioned
		var qrVersioned bool

		if o.Versioned != nil {
			qrVersioned = *o.Versioned
		}
		qVersioned := swag.FormatBool(qrVersioned)
		if qVersioned != "" {

			if err := r.SetQueryParam("versioned", qVersioned); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListBlueprintsUsingGET1 binds the parameter $orderby
func (o *ListBlueprintsUsingGET1Params) bindParamDollarOrderby(formats strfmt.Registry) []string {
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

// bindParamListBlueprintsUsingGET1 binds the parameter $select
func (o *ListBlueprintsUsingGET1Params) bindParamDollarSelect(formats strfmt.Registry) []string {
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

// bindParamListBlueprintsUsingGET1 binds the parameter projects
func (o *ListBlueprintsUsingGET1Params) bindParamProjects(formats strfmt.Registry) []string {
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

// bindParamListBlueprintsUsingGET1 binds the parameter propertyGroups
func (o *ListBlueprintsUsingGET1Params) bindParamPropertyGroups(formats strfmt.Registry) []string {
	propertyGroupsIR := o.PropertyGroups

	var propertyGroupsIC []string
	for _, propertyGroupsIIR := range propertyGroupsIR { // explode []string

		propertyGroupsIIV := propertyGroupsIIR // string as string
		propertyGroupsIC = append(propertyGroupsIC, propertyGroupsIIV)
	}

	// items.CollectionFormat: "multi"
	propertyGroupsIS := swag.JoinByFormat(propertyGroupsIC, "multi")

	return propertyGroupsIS
}
