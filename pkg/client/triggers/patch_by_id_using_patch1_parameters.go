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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewPatchByIDUsingPATCH1Params creates a new PatchByIDUsingPATCH1Params object
// with the default values initialized.
func NewPatchByIDUsingPATCH1Params() *PatchByIDUsingPATCH1Params {
	var ()
	return &PatchByIDUsingPATCH1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchByIDUsingPATCH1ParamsWithTimeout creates a new PatchByIDUsingPATCH1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchByIDUsingPATCH1ParamsWithTimeout(timeout time.Duration) *PatchByIDUsingPATCH1Params {
	var ()
	return &PatchByIDUsingPATCH1Params{

		timeout: timeout,
	}
}

// NewPatchByIDUsingPATCH1ParamsWithContext creates a new PatchByIDUsingPATCH1Params object
// with the default values initialized, and the ability to set a context for a request
func NewPatchByIDUsingPATCH1ParamsWithContext(ctx context.Context) *PatchByIDUsingPATCH1Params {
	var ()
	return &PatchByIDUsingPATCH1Params{

		Context: ctx,
	}
}

// NewPatchByIDUsingPATCH1ParamsWithHTTPClient creates a new PatchByIDUsingPATCH1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchByIDUsingPATCH1ParamsWithHTTPClient(client *http.Client) *PatchByIDUsingPATCH1Params {
	var ()
	return &PatchByIDUsingPATCH1Params{
		HTTPClient: client,
	}
}

/*PatchByIDUsingPATCH1Params contains all the parameters to send to the API endpoint
for the patch by Id using p a t c h 1 operation typically these are written to a http.Request
*/
type PatchByIDUsingPATCH1Params struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*GerritListenerPatch
	  gerritListenerPatch

	*/
	GerritListenerPatch models.GerritListenerPatch
	/*ID
	  id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch by Id using p a t c h 1 params
func (o *PatchByIDUsingPATCH1Params) WithTimeout(timeout time.Duration) *PatchByIDUsingPATCH1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch by Id using p a t c h 1 params
func (o *PatchByIDUsingPATCH1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch by Id using p a t c h 1 params
func (o *PatchByIDUsingPATCH1Params) WithContext(ctx context.Context) *PatchByIDUsingPATCH1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch by Id using p a t c h 1 params
func (o *PatchByIDUsingPATCH1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch by Id using p a t c h 1 params
func (o *PatchByIDUsingPATCH1Params) WithHTTPClient(client *http.Client) *PatchByIDUsingPATCH1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch by Id using p a t c h 1 params
func (o *PatchByIDUsingPATCH1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the patch by Id using p a t c h 1 params
func (o *PatchByIDUsingPATCH1Params) WithAPIVersion(aPIVersion string) *PatchByIDUsingPATCH1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the patch by Id using p a t c h 1 params
func (o *PatchByIDUsingPATCH1Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithGerritListenerPatch adds the gerritListenerPatch to the patch by Id using p a t c h 1 params
func (o *PatchByIDUsingPATCH1Params) WithGerritListenerPatch(gerritListenerPatch models.GerritListenerPatch) *PatchByIDUsingPATCH1Params {
	o.SetGerritListenerPatch(gerritListenerPatch)
	return o
}

// SetGerritListenerPatch adds the gerritListenerPatch to the patch by Id using p a t c h 1 params
func (o *PatchByIDUsingPATCH1Params) SetGerritListenerPatch(gerritListenerPatch models.GerritListenerPatch) {
	o.GerritListenerPatch = gerritListenerPatch
}

// WithID adds the id to the patch by Id using p a t c h 1 params
func (o *PatchByIDUsingPATCH1Params) WithID(id string) *PatchByIDUsingPATCH1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch by Id using p a t c h 1 params
func (o *PatchByIDUsingPATCH1Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchByIDUsingPATCH1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.GerritListenerPatch); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
