// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetAllProjectsUsingGETParams creates a new GetAllProjectsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllProjectsUsingGETParams() *GetAllProjectsUsingGETParams {
	return &GetAllProjectsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllProjectsUsingGETParamsWithTimeout creates a new GetAllProjectsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAllProjectsUsingGETParamsWithTimeout(timeout time.Duration) *GetAllProjectsUsingGETParams {
	return &GetAllProjectsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAllProjectsUsingGETParamsWithContext creates a new GetAllProjectsUsingGETParams object
// with the ability to set a context for a request.
func NewGetAllProjectsUsingGETParamsWithContext(ctx context.Context) *GetAllProjectsUsingGETParams {
	return &GetAllProjectsUsingGETParams{
		Context: ctx,
	}
}

// NewGetAllProjectsUsingGETParamsWithHTTPClient creates a new GetAllProjectsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllProjectsUsingGETParamsWithHTTPClient(client *http.Client) *GetAllProjectsUsingGETParams {
	return &GetAllProjectsUsingGETParams{
		HTTPClient: client,
	}
}

/* GetAllProjectsUsingGETParams contains all the parameters to send to the API endpoint
   for the get all projects using g e t operation.

   Typically these are written to a http.Request.
*/
type GetAllProjectsUsingGETParams struct {

	/* DollarSelect.

	   Select a subset of properties to include in the response. Possible values for this parameter are id, name, operationTimeout, constraints.
	*/
	DollarSelect *string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format. For versioning information refer to /project-service/api/about.
	*/
	APIVersion *string

	/* ExcludeNotSharedProjectsForMember.

	   Filters projects based on the member role and the access to the resources of the project. When the value is true it will not return the projects in which the current user is only member and the project is not with shared resources.
	*/
	ExcludeNotSharedProjectsForMember *bool

	/* ExcludeViewer.

	   Filters projects based on the viewer role. When the value is true it will not return the projects in which the current user is only viewer and will ignore privileged roles: CodeStream:Developer and CodeStream:Executor, if the user has them. Else it will return all projects that the user can read. The default value is false.
	*/
	ExcludeViewer *bool

	/* WithAnyPermission.

	   Optional permissions that, if granted to the user, allow him access to the proper set of projects. If the user actually has any of those permissions, the 'excludeViewer' parameter has no effect.
	*/
	WithAnyPermission []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all projects using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllProjectsUsingGETParams) WithDefaults() *GetAllProjectsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all projects using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllProjectsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) WithTimeout(timeout time.Duration) *GetAllProjectsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) WithContext(ctx context.Context) *GetAllProjectsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) WithHTTPClient(client *http.Client) *GetAllProjectsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarSelect adds the dollarSelect to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) WithDollarSelect(dollarSelect *string) *GetAllProjectsUsingGETParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithAPIVersion adds the aPIVersion to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) WithAPIVersion(aPIVersion *string) *GetAllProjectsUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithExcludeNotSharedProjectsForMember adds the excludeNotSharedProjectsForMember to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) WithExcludeNotSharedProjectsForMember(excludeNotSharedProjectsForMember *bool) *GetAllProjectsUsingGETParams {
	o.SetExcludeNotSharedProjectsForMember(excludeNotSharedProjectsForMember)
	return o
}

// SetExcludeNotSharedProjectsForMember adds the excludeNotSharedProjectsForMember to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) SetExcludeNotSharedProjectsForMember(excludeNotSharedProjectsForMember *bool) {
	o.ExcludeNotSharedProjectsForMember = excludeNotSharedProjectsForMember
}

// WithExcludeViewer adds the excludeViewer to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) WithExcludeViewer(excludeViewer *bool) *GetAllProjectsUsingGETParams {
	o.SetExcludeViewer(excludeViewer)
	return o
}

// SetExcludeViewer adds the excludeViewer to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) SetExcludeViewer(excludeViewer *bool) {
	o.ExcludeViewer = excludeViewer
}

// WithWithAnyPermission adds the withAnyPermission to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) WithWithAnyPermission(withAnyPermission []string) *GetAllProjectsUsingGETParams {
	o.SetWithAnyPermission(withAnyPermission)
	return o
}

// SetWithAnyPermission adds the withAnyPermission to the get all projects using get params
func (o *GetAllProjectsUsingGETParams) SetWithAnyPermission(withAnyPermission []string) {
	o.WithAnyPermission = withAnyPermission
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllProjectsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarSelect != nil {

		// query param $select
		var qrDollarSelect string

		if o.DollarSelect != nil {
			qrDollarSelect = *o.DollarSelect
		}
		qDollarSelect := qrDollarSelect
		if qDollarSelect != "" {

			if err := r.SetQueryParam("$select", qDollarSelect); err != nil {
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

	if o.ExcludeNotSharedProjectsForMember != nil {

		// query param excludeNotSharedProjectsForMember
		var qrExcludeNotSharedProjectsForMember bool

		if o.ExcludeNotSharedProjectsForMember != nil {
			qrExcludeNotSharedProjectsForMember = *o.ExcludeNotSharedProjectsForMember
		}
		qExcludeNotSharedProjectsForMember := swag.FormatBool(qrExcludeNotSharedProjectsForMember)
		if qExcludeNotSharedProjectsForMember != "" {

			if err := r.SetQueryParam("excludeNotSharedProjectsForMember", qExcludeNotSharedProjectsForMember); err != nil {
				return err
			}
		}
	}

	if o.ExcludeViewer != nil {

		// query param excludeViewer
		var qrExcludeViewer bool

		if o.ExcludeViewer != nil {
			qrExcludeViewer = *o.ExcludeViewer
		}
		qExcludeViewer := swag.FormatBool(qrExcludeViewer)
		if qExcludeViewer != "" {

			if err := r.SetQueryParam("excludeViewer", qExcludeViewer); err != nil {
				return err
			}
		}
	}

	if o.WithAnyPermission != nil {

		// binding items for withAnyPermission
		joinedWithAnyPermission := o.bindParamWithAnyPermission(reg)

		// query array param withAnyPermission
		if err := r.SetQueryParam("withAnyPermission", joinedWithAnyPermission...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAllProjectsUsingGET binds the parameter withAnyPermission
func (o *GetAllProjectsUsingGETParams) bindParamWithAnyPermission(formats strfmt.Registry) []string {
	withAnyPermissionIR := o.WithAnyPermission

	var withAnyPermissionIC []string
	for _, withAnyPermissionIIR := range withAnyPermissionIR { // explode []string

		withAnyPermissionIIV := withAnyPermissionIIR // string as string
		withAnyPermissionIC = append(withAnyPermissionIC, withAnyPermissionIIV)
	}

	// items.CollectionFormat: "multi"
	withAnyPermissionIS := swag.JoinByFormat(withAnyPermissionIC, "multi")

	return withAnyPermissionIS
}
