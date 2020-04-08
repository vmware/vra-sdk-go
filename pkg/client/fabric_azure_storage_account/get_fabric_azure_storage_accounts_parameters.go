// Code generated by go-swagger; DO NOT EDIT.

package fabric_azure_storage_account

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
)

// NewGetFabricAzureStorageAccountsParams creates a new GetFabricAzureStorageAccountsParams object
// with the default values initialized.
func NewGetFabricAzureStorageAccountsParams() *GetFabricAzureStorageAccountsParams {
	var ()
	return &GetFabricAzureStorageAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFabricAzureStorageAccountsParamsWithTimeout creates a new GetFabricAzureStorageAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFabricAzureStorageAccountsParamsWithTimeout(timeout time.Duration) *GetFabricAzureStorageAccountsParams {
	var ()
	return &GetFabricAzureStorageAccountsParams{

		timeout: timeout,
	}
}

// NewGetFabricAzureStorageAccountsParamsWithContext creates a new GetFabricAzureStorageAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFabricAzureStorageAccountsParamsWithContext(ctx context.Context) *GetFabricAzureStorageAccountsParams {
	var ()
	return &GetFabricAzureStorageAccountsParams{

		Context: ctx,
	}
}

// NewGetFabricAzureStorageAccountsParamsWithHTTPClient creates a new GetFabricAzureStorageAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFabricAzureStorageAccountsParamsWithHTTPClient(client *http.Client) *GetFabricAzureStorageAccountsParams {
	var ()
	return &GetFabricAzureStorageAccountsParams{
		HTTPClient: client,
	}
}

/*GetFabricAzureStorageAccountsParams contains all the parameters to send to the API endpoint
for the get fabric azure storage accounts operation typically these are written to a http.Request
*/
type GetFabricAzureStorageAccountsParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fabric azure storage accounts params
func (o *GetFabricAzureStorageAccountsParams) WithTimeout(timeout time.Duration) *GetFabricAzureStorageAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fabric azure storage accounts params
func (o *GetFabricAzureStorageAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fabric azure storage accounts params
func (o *GetFabricAzureStorageAccountsParams) WithContext(ctx context.Context) *GetFabricAzureStorageAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fabric azure storage accounts params
func (o *GetFabricAzureStorageAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fabric azure storage accounts params
func (o *GetFabricAzureStorageAccountsParams) WithHTTPClient(client *http.Client) *GetFabricAzureStorageAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fabric azure storage accounts params
func (o *GetFabricAzureStorageAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get fabric azure storage accounts params
func (o *GetFabricAzureStorageAccountsParams) WithAPIVersion(aPIVersion *string) *GetFabricAzureStorageAccountsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get fabric azure storage accounts params
func (o *GetFabricAzureStorageAccountsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetFabricAzureStorageAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
