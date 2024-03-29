// Code generated by go-swagger; DO NOT EDIT.

package supervisor_clusters

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

// NewListClustersOnEndpointUsingGETParams creates a new ListClustersOnEndpointUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListClustersOnEndpointUsingGETParams() *ListClustersOnEndpointUsingGETParams {
	return &ListClustersOnEndpointUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListClustersOnEndpointUsingGETParamsWithTimeout creates a new ListClustersOnEndpointUsingGETParams object
// with the ability to set a timeout on a request.
func NewListClustersOnEndpointUsingGETParamsWithTimeout(timeout time.Duration) *ListClustersOnEndpointUsingGETParams {
	return &ListClustersOnEndpointUsingGETParams{
		timeout: timeout,
	}
}

// NewListClustersOnEndpointUsingGETParamsWithContext creates a new ListClustersOnEndpointUsingGETParams object
// with the ability to set a context for a request.
func NewListClustersOnEndpointUsingGETParamsWithContext(ctx context.Context) *ListClustersOnEndpointUsingGETParams {
	return &ListClustersOnEndpointUsingGETParams{
		Context: ctx,
	}
}

// NewListClustersOnEndpointUsingGETParamsWithHTTPClient creates a new ListClustersOnEndpointUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewListClustersOnEndpointUsingGETParamsWithHTTPClient(client *http.Client) *ListClustersOnEndpointUsingGETParams {
	return &ListClustersOnEndpointUsingGETParams{
		HTTPClient: client,
	}
}

/*
ListClustersOnEndpointUsingGETParams contains all the parameters to send to the API endpoint

	for the list clusters on endpoint using g e t operation.

	Typically these are written to a http.Request.
*/
type ListClustersOnEndpointUsingGETParams struct {

	/* EndpointSelfLinkID.

	   endpointSelfLinkId
	*/
	EndpointSelfLinkID string

	// Offset.
	//
	// Format: int64
	Offset *int64

	// PageNumber.
	//
	// Format: int32
	PageNumber *int32

	// PageSize.
	//
	// Format: int32
	PageSize *int32

	// Paged.
	Paged *bool

	/* Registered.

	   registered
	*/
	Registered *string

	// SortSorted.
	SortSorted *bool

	// SortUnsorted.
	SortUnsorted *bool

	// Unpaged.
	Unpaged *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list clusters on endpoint using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListClustersOnEndpointUsingGETParams) WithDefaults() *ListClustersOnEndpointUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list clusters on endpoint using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListClustersOnEndpointUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) WithTimeout(timeout time.Duration) *ListClustersOnEndpointUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) WithContext(ctx context.Context) *ListClustersOnEndpointUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) WithHTTPClient(client *http.Client) *ListClustersOnEndpointUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpointSelfLinkID adds the endpointSelfLinkID to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) WithEndpointSelfLinkID(endpointSelfLinkID string) *ListClustersOnEndpointUsingGETParams {
	o.SetEndpointSelfLinkID(endpointSelfLinkID)
	return o
}

// SetEndpointSelfLinkID adds the endpointSelfLinkId to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) SetEndpointSelfLinkID(endpointSelfLinkID string) {
	o.EndpointSelfLinkID = endpointSelfLinkID
}

// WithOffset adds the offset to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) WithOffset(offset *int64) *ListClustersOnEndpointUsingGETParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPageNumber adds the pageNumber to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) WithPageNumber(pageNumber *int32) *ListClustersOnEndpointUsingGETParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) WithPageSize(pageSize *int32) *ListClustersOnEndpointUsingGETParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPaged adds the paged to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) WithPaged(paged *bool) *ListClustersOnEndpointUsingGETParams {
	o.SetPaged(paged)
	return o
}

// SetPaged adds the paged to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) SetPaged(paged *bool) {
	o.Paged = paged
}

// WithRegistered adds the registered to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) WithRegistered(registered *string) *ListClustersOnEndpointUsingGETParams {
	o.SetRegistered(registered)
	return o
}

// SetRegistered adds the registered to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) SetRegistered(registered *string) {
	o.Registered = registered
}

// WithSortSorted adds the sortSorted to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) WithSortSorted(sortSorted *bool) *ListClustersOnEndpointUsingGETParams {
	o.SetSortSorted(sortSorted)
	return o
}

// SetSortSorted adds the sortSorted to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) SetSortSorted(sortSorted *bool) {
	o.SortSorted = sortSorted
}

// WithSortUnsorted adds the sortUnsorted to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) WithSortUnsorted(sortUnsorted *bool) *ListClustersOnEndpointUsingGETParams {
	o.SetSortUnsorted(sortUnsorted)
	return o
}

// SetSortUnsorted adds the sortUnsorted to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) SetSortUnsorted(sortUnsorted *bool) {
	o.SortUnsorted = sortUnsorted
}

// WithUnpaged adds the unpaged to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) WithUnpaged(unpaged *bool) *ListClustersOnEndpointUsingGETParams {
	o.SetUnpaged(unpaged)
	return o
}

// SetUnpaged adds the unpaged to the list clusters on endpoint using get params
func (o *ListClustersOnEndpointUsingGETParams) SetUnpaged(unpaged *bool) {
	o.Unpaged = unpaged
}

// WriteToRequest writes these params to a swagger request
func (o *ListClustersOnEndpointUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param endpointSelfLinkId
	if err := r.SetPathParam("endpointSelfLinkId", o.EndpointSelfLinkID); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Paged != nil {

		// query param paged
		var qrPaged bool

		if o.Paged != nil {
			qrPaged = *o.Paged
		}
		qPaged := swag.FormatBool(qrPaged)
		if qPaged != "" {

			if err := r.SetQueryParam("paged", qPaged); err != nil {
				return err
			}
		}
	}

	if o.Registered != nil {

		// query param registered
		var qrRegistered string

		if o.Registered != nil {
			qrRegistered = *o.Registered
		}
		qRegistered := qrRegistered
		if qRegistered != "" {

			if err := r.SetQueryParam("registered", qRegistered); err != nil {
				return err
			}
		}
	}

	if o.SortSorted != nil {

		// query param sort.sorted
		var qrSortSorted bool

		if o.SortSorted != nil {
			qrSortSorted = *o.SortSorted
		}
		qSortSorted := swag.FormatBool(qrSortSorted)
		if qSortSorted != "" {

			if err := r.SetQueryParam("sort.sorted", qSortSorted); err != nil {
				return err
			}
		}
	}

	if o.SortUnsorted != nil {

		// query param sort.unsorted
		var qrSortUnsorted bool

		if o.SortUnsorted != nil {
			qrSortUnsorted = *o.SortUnsorted
		}
		qSortUnsorted := swag.FormatBool(qrSortUnsorted)
		if qSortUnsorted != "" {

			if err := r.SetQueryParam("sort.unsorted", qSortUnsorted); err != nil {
				return err
			}
		}
	}

	if o.Unpaged != nil {

		// query param unpaged
		var qrUnpaged bool

		if o.Unpaged != nil {
			qrUnpaged = *o.Unpaged
		}
		qUnpaged := swag.FormatBool(qrUnpaged)
		if qUnpaged != "" {

			if err := r.SetQueryParam("unpaged", qUnpaged); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
