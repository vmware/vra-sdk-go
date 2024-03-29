// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_clusters

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

// NewOnboardUsingPOSTParams creates a new OnboardUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOnboardUsingPOSTParams() *OnboardUsingPOSTParams {
	return &OnboardUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOnboardUsingPOSTParamsWithTimeout creates a new OnboardUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewOnboardUsingPOSTParamsWithTimeout(timeout time.Duration) *OnboardUsingPOSTParams {
	return &OnboardUsingPOSTParams{
		timeout: timeout,
	}
}

// NewOnboardUsingPOSTParamsWithContext creates a new OnboardUsingPOSTParams object
// with the ability to set a context for a request.
func NewOnboardUsingPOSTParamsWithContext(ctx context.Context) *OnboardUsingPOSTParams {
	return &OnboardUsingPOSTParams{
		Context: ctx,
	}
}

// NewOnboardUsingPOSTParamsWithHTTPClient creates a new OnboardUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewOnboardUsingPOSTParamsWithHTTPClient(client *http.Client) *OnboardUsingPOSTParams {
	return &OnboardUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
OnboardUsingPOSTParams contains all the parameters to send to the API endpoint

	for the onboard using p o s t operation.

	Typically these are written to a http.Request.
*/
type OnboardUsingPOSTParams struct {

	/* Cluster.

	   cluster
	*/
	Cluster *models.K8SCluster

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the onboard using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OnboardUsingPOSTParams) WithDefaults() *OnboardUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the onboard using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OnboardUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the onboard using p o s t params
func (o *OnboardUsingPOSTParams) WithTimeout(timeout time.Duration) *OnboardUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the onboard using p o s t params
func (o *OnboardUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the onboard using p o s t params
func (o *OnboardUsingPOSTParams) WithContext(ctx context.Context) *OnboardUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the onboard using p o s t params
func (o *OnboardUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the onboard using p o s t params
func (o *OnboardUsingPOSTParams) WithHTTPClient(client *http.Client) *OnboardUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the onboard using p o s t params
func (o *OnboardUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the onboard using p o s t params
func (o *OnboardUsingPOSTParams) WithCluster(cluster *models.K8SCluster) *OnboardUsingPOSTParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the onboard using p o s t params
func (o *OnboardUsingPOSTParams) SetCluster(cluster *models.K8SCluster) {
	o.Cluster = cluster
}

// WriteToRequest writes these params to a swagger request
func (o *OnboardUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Cluster != nil {
		if err := r.SetBodyParam(o.Cluster); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
