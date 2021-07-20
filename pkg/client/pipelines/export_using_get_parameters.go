// Code generated by go-swagger; DO NOT EDIT.

package pipelines

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
)

// NewExportUsingGETParams creates a new ExportUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportUsingGETParams() *ExportUsingGETParams {
	return &ExportUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportUsingGETParamsWithTimeout creates a new ExportUsingGETParams object
// with the ability to set a timeout on a request.
func NewExportUsingGETParamsWithTimeout(timeout time.Duration) *ExportUsingGETParams {
	return &ExportUsingGETParams{
		timeout: timeout,
	}
}

// NewExportUsingGETParamsWithContext creates a new ExportUsingGETParams object
// with the ability to set a context for a request.
func NewExportUsingGETParamsWithContext(ctx context.Context) *ExportUsingGETParams {
	return &ExportUsingGETParams{
		Context: ctx,
	}
}

// NewExportUsingGETParamsWithHTTPClient creates a new ExportUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportUsingGETParamsWithHTTPClient(client *http.Client) *ExportUsingGETParams {
	return &ExportUsingGETParams{
		HTTPClient: client,
	}
}

/* ExportUsingGETParams contains all the parameters to send to the API endpoint
   for the export using g e t operation.

   Typically these are written to a http.Request.
*/
type ExportUsingGETParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Endpoints.

	   Comma separated list of endpoints to be exported in a given project

	   Default: "Jenkins, Jira"
	*/
	Endpoints *string

	/* Pipeline.

	   Name of the Pipeline to be exported. Here, all endpoints referred in the pipeline also get exported

	   Default: "Deploy Production"
	*/
	Pipeline *string

	/* Pipelines.

	   Comma separated list of pipelines to be exported in a given project

	   Default: "Deploy Production, Dev"
	*/
	Pipelines *string

	/* Project.

	   Name of the Project to which Endpoint(s)/Pipeline(s) belong to

	   Default: "Project-1"
	*/
	Project *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportUsingGETParams) WithDefaults() *ExportUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportUsingGETParams) SetDefaults() {
	var (
		endpointsDefault = string("Jenkins, Jira")

		pipelineDefault = string("Deploy Production")

		pipelinesDefault = string("Deploy Production, Dev")

		projectDefault = string("Project-1")
	)

	val := ExportUsingGETParams{
		Endpoints: &endpointsDefault,
		Pipeline:  &pipelineDefault,
		Pipelines: &pipelinesDefault,
		Project:   &projectDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the export using get params
func (o *ExportUsingGETParams) WithTimeout(timeout time.Duration) *ExportUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export using get params
func (o *ExportUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export using get params
func (o *ExportUsingGETParams) WithContext(ctx context.Context) *ExportUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export using get params
func (o *ExportUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export using get params
func (o *ExportUsingGETParams) WithHTTPClient(client *http.Client) *ExportUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export using get params
func (o *ExportUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the export using get params
func (o *ExportUsingGETParams) WithAuthorization(authorization string) *ExportUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the export using get params
func (o *ExportUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the export using get params
func (o *ExportUsingGETParams) WithAPIVersion(aPIVersion string) *ExportUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the export using get params
func (o *ExportUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithEndpoints adds the endpoints to the export using get params
func (o *ExportUsingGETParams) WithEndpoints(endpoints *string) *ExportUsingGETParams {
	o.SetEndpoints(endpoints)
	return o
}

// SetEndpoints adds the endpoints to the export using get params
func (o *ExportUsingGETParams) SetEndpoints(endpoints *string) {
	o.Endpoints = endpoints
}

// WithPipeline adds the pipeline to the export using get params
func (o *ExportUsingGETParams) WithPipeline(pipeline *string) *ExportUsingGETParams {
	o.SetPipeline(pipeline)
	return o
}

// SetPipeline adds the pipeline to the export using get params
func (o *ExportUsingGETParams) SetPipeline(pipeline *string) {
	o.Pipeline = pipeline
}

// WithPipelines adds the pipelines to the export using get params
func (o *ExportUsingGETParams) WithPipelines(pipelines *string) *ExportUsingGETParams {
	o.SetPipelines(pipelines)
	return o
}

// SetPipelines adds the pipelines to the export using get params
func (o *ExportUsingGETParams) SetPipelines(pipelines *string) {
	o.Pipelines = pipelines
}

// WithProject adds the project to the export using get params
func (o *ExportUsingGETParams) WithProject(project *string) *ExportUsingGETParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the export using get params
func (o *ExportUsingGETParams) SetProject(project *string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *ExportUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Endpoints != nil {

		// query param endpoints
		var qrEndpoints string

		if o.Endpoints != nil {
			qrEndpoints = *o.Endpoints
		}
		qEndpoints := qrEndpoints
		if qEndpoints != "" {

			if err := r.SetQueryParam("endpoints", qEndpoints); err != nil {
				return err
			}
		}
	}

	if o.Pipeline != nil {

		// query param pipeline
		var qrPipeline string

		if o.Pipeline != nil {
			qrPipeline = *o.Pipeline
		}
		qPipeline := qrPipeline
		if qPipeline != "" {

			if err := r.SetQueryParam("pipeline", qPipeline); err != nil {
				return err
			}
		}
	}

	if o.Pipelines != nil {

		// query param pipelines
		var qrPipelines string

		if o.Pipelines != nil {
			qrPipelines = *o.Pipelines
		}
		qPipelines := qrPipelines
		if qPipelines != "" {

			if err := r.SetQueryParam("pipelines", qPipelines); err != nil {
				return err
			}
		}
	}

	if o.Project != nil {

		// query param project
		var qrProject string

		if o.Project != nil {
			qrProject = *o.Project
		}
		qProject := qrProject
		if qProject != "" {

			if err := r.SetQueryParam("project", qProject); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
