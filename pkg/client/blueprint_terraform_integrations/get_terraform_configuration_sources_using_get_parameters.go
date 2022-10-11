// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

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

// NewGetTerraformConfigurationSourcesUsingGETParams creates a new GetTerraformConfigurationSourcesUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTerraformConfigurationSourcesUsingGETParams() *GetTerraformConfigurationSourcesUsingGETParams {
	return &GetTerraformConfigurationSourcesUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTerraformConfigurationSourcesUsingGETParamsWithTimeout creates a new GetTerraformConfigurationSourcesUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetTerraformConfigurationSourcesUsingGETParamsWithTimeout(timeout time.Duration) *GetTerraformConfigurationSourcesUsingGETParams {
	return &GetTerraformConfigurationSourcesUsingGETParams{
		timeout: timeout,
	}
}

// NewGetTerraformConfigurationSourcesUsingGETParamsWithContext creates a new GetTerraformConfigurationSourcesUsingGETParams object
// with the ability to set a context for a request.
func NewGetTerraformConfigurationSourcesUsingGETParamsWithContext(ctx context.Context) *GetTerraformConfigurationSourcesUsingGETParams {
	return &GetTerraformConfigurationSourcesUsingGETParams{
		Context: ctx,
	}
}

// NewGetTerraformConfigurationSourcesUsingGETParamsWithHTTPClient creates a new GetTerraformConfigurationSourcesUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTerraformConfigurationSourcesUsingGETParamsWithHTTPClient(client *http.Client) *GetTerraformConfigurationSourcesUsingGETParams {
	return &GetTerraformConfigurationSourcesUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetTerraformConfigurationSourcesUsingGETParams contains all the parameters to send to the API endpoint

	for the get terraform configuration sources using g e t operation.

	Typically these are written to a http.Request.
*/
type GetTerraformConfigurationSourcesUsingGETParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* Projects.

	   A comma-separated list of project IDs. Results will be filtered so only configuration sources accessible from one or more of these projects will be returned.
	*/
	Projects []string

	/* Search.

	   A search parameter to search based on configuration source name or repository.
	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get terraform configuration sources using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformConfigurationSourcesUsingGETParams) WithDefaults() *GetTerraformConfigurationSourcesUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get terraform configuration sources using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformConfigurationSourcesUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get terraform configuration sources using get params
func (o *GetTerraformConfigurationSourcesUsingGETParams) WithTimeout(timeout time.Duration) *GetTerraformConfigurationSourcesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get terraform configuration sources using get params
func (o *GetTerraformConfigurationSourcesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get terraform configuration sources using get params
func (o *GetTerraformConfigurationSourcesUsingGETParams) WithContext(ctx context.Context) *GetTerraformConfigurationSourcesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get terraform configuration sources using get params
func (o *GetTerraformConfigurationSourcesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get terraform configuration sources using get params
func (o *GetTerraformConfigurationSourcesUsingGETParams) WithHTTPClient(client *http.Client) *GetTerraformConfigurationSourcesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get terraform configuration sources using get params
func (o *GetTerraformConfigurationSourcesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get terraform configuration sources using get params
func (o *GetTerraformConfigurationSourcesUsingGETParams) WithAPIVersion(aPIVersion *string) *GetTerraformConfigurationSourcesUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get terraform configuration sources using get params
func (o *GetTerraformConfigurationSourcesUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithProjects adds the projects to the get terraform configuration sources using get params
func (o *GetTerraformConfigurationSourcesUsingGETParams) WithProjects(projects []string) *GetTerraformConfigurationSourcesUsingGETParams {
	o.SetProjects(projects)
	return o
}

// SetProjects adds the projects to the get terraform configuration sources using get params
func (o *GetTerraformConfigurationSourcesUsingGETParams) SetProjects(projects []string) {
	o.Projects = projects
}

// WithSearch adds the search to the get terraform configuration sources using get params
func (o *GetTerraformConfigurationSourcesUsingGETParams) WithSearch(search *string) *GetTerraformConfigurationSourcesUsingGETParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get terraform configuration sources using get params
func (o *GetTerraformConfigurationSourcesUsingGETParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetTerraformConfigurationSourcesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetTerraformConfigurationSourcesUsingGET binds the parameter projects
func (o *GetTerraformConfigurationSourcesUsingGETParams) bindParamProjects(formats strfmt.Registry) []string {
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