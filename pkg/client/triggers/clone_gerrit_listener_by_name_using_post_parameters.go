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
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCloneGerritListenerByNameUsingPOSTParams creates a new CloneGerritListenerByNameUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloneGerritListenerByNameUsingPOSTParams() *CloneGerritListenerByNameUsingPOSTParams {
	return &CloneGerritListenerByNameUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloneGerritListenerByNameUsingPOSTParamsWithTimeout creates a new CloneGerritListenerByNameUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCloneGerritListenerByNameUsingPOSTParamsWithTimeout(timeout time.Duration) *CloneGerritListenerByNameUsingPOSTParams {
	return &CloneGerritListenerByNameUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCloneGerritListenerByNameUsingPOSTParamsWithContext creates a new CloneGerritListenerByNameUsingPOSTParams object
// with the ability to set a context for a request.
func NewCloneGerritListenerByNameUsingPOSTParamsWithContext(ctx context.Context) *CloneGerritListenerByNameUsingPOSTParams {
	return &CloneGerritListenerByNameUsingPOSTParams{
		Context: ctx,
	}
}

// NewCloneGerritListenerByNameUsingPOSTParamsWithHTTPClient creates a new CloneGerritListenerByNameUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloneGerritListenerByNameUsingPOSTParamsWithHTTPClient(client *http.Client) *CloneGerritListenerByNameUsingPOSTParams {
	return &CloneGerritListenerByNameUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
CloneGerritListenerByNameUsingPOSTParams contains all the parameters to send to the API endpoint

	for the clone gerrit listener by name using p o s t operation.

	Typically these are written to a http.Request.
*/
type CloneGerritListenerByNameUsingPOSTParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Body.

	   Request object for actions such as cloning
	*/
	Body models.ServiceRequest

	/* Name.

	   The name of the Gerrit Listener
	*/
	Name string

	/* Project.

	   The project the Gerrit Listener belongs to
	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the clone gerrit listener by name using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneGerritListenerByNameUsingPOSTParams) WithDefaults() *CloneGerritListenerByNameUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the clone gerrit listener by name using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneGerritListenerByNameUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) WithTimeout(timeout time.Duration) *CloneGerritListenerByNameUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) WithContext(ctx context.Context) *CloneGerritListenerByNameUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) WithHTTPClient(client *http.Client) *CloneGerritListenerByNameUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) WithAuthorization(authorization string) *CloneGerritListenerByNameUsingPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) WithAPIVersion(aPIVersion string) *CloneGerritListenerByNameUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) WithBody(body models.ServiceRequest) *CloneGerritListenerByNameUsingPOSTParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) SetBody(body models.ServiceRequest) {
	o.Body = body
}

// WithName adds the name to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) WithName(name string) *CloneGerritListenerByNameUsingPOSTParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) WithProject(project string) *CloneGerritListenerByNameUsingPOSTParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the clone gerrit listener by name using p o s t params
func (o *CloneGerritListenerByNameUsingPOSTParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *CloneGerritListenerByNameUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion

	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
