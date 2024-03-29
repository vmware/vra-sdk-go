// Code generated by go-swagger; DO NOT EDIT.

package capability_tags

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

// NewSearchTagsUsingGETParams creates a new SearchTagsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchTagsUsingGETParams() *SearchTagsUsingGETParams {
	return &SearchTagsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchTagsUsingGETParamsWithTimeout creates a new SearchTagsUsingGETParams object
// with the ability to set a timeout on a request.
func NewSearchTagsUsingGETParamsWithTimeout(timeout time.Duration) *SearchTagsUsingGETParams {
	return &SearchTagsUsingGETParams{
		timeout: timeout,
	}
}

// NewSearchTagsUsingGETParamsWithContext creates a new SearchTagsUsingGETParams object
// with the ability to set a context for a request.
func NewSearchTagsUsingGETParamsWithContext(ctx context.Context) *SearchTagsUsingGETParams {
	return &SearchTagsUsingGETParams{
		Context: ctx,
	}
}

// NewSearchTagsUsingGETParamsWithHTTPClient creates a new SearchTagsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchTagsUsingGETParamsWithHTTPClient(client *http.Client) *SearchTagsUsingGETParams {
	return &SearchTagsUsingGETParams{
		HTTPClient: client,
	}
}

/*
SearchTagsUsingGETParams contains all the parameters to send to the API endpoint

	for the search tags using g e t operation.

	Typically these are written to a http.Request.
*/
type SearchTagsUsingGETParams struct {

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

// WithDefaults hydrates default values in the search tags using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchTagsUsingGETParams) WithDefaults() *SearchTagsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search tags using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchTagsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search tags using get params
func (o *SearchTagsUsingGETParams) WithTimeout(timeout time.Duration) *SearchTagsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search tags using get params
func (o *SearchTagsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search tags using get params
func (o *SearchTagsUsingGETParams) WithContext(ctx context.Context) *SearchTagsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search tags using get params
func (o *SearchTagsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search tags using get params
func (o *SearchTagsUsingGETParams) WithHTTPClient(client *http.Client) *SearchTagsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search tags using get params
func (o *SearchTagsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOffset adds the offset to the search tags using get params
func (o *SearchTagsUsingGETParams) WithOffset(offset *int64) *SearchTagsUsingGETParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search tags using get params
func (o *SearchTagsUsingGETParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPageNumber adds the pageNumber to the search tags using get params
func (o *SearchTagsUsingGETParams) WithPageNumber(pageNumber *int32) *SearchTagsUsingGETParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the search tags using get params
func (o *SearchTagsUsingGETParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the search tags using get params
func (o *SearchTagsUsingGETParams) WithPageSize(pageSize *int32) *SearchTagsUsingGETParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the search tags using get params
func (o *SearchTagsUsingGETParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPaged adds the paged to the search tags using get params
func (o *SearchTagsUsingGETParams) WithPaged(paged *bool) *SearchTagsUsingGETParams {
	o.SetPaged(paged)
	return o
}

// SetPaged adds the paged to the search tags using get params
func (o *SearchTagsUsingGETParams) SetPaged(paged *bool) {
	o.Paged = paged
}

// WithSortSorted adds the sortSorted to the search tags using get params
func (o *SearchTagsUsingGETParams) WithSortSorted(sortSorted *bool) *SearchTagsUsingGETParams {
	o.SetSortSorted(sortSorted)
	return o
}

// SetSortSorted adds the sortSorted to the search tags using get params
func (o *SearchTagsUsingGETParams) SetSortSorted(sortSorted *bool) {
	o.SortSorted = sortSorted
}

// WithSortUnsorted adds the sortUnsorted to the search tags using get params
func (o *SearchTagsUsingGETParams) WithSortUnsorted(sortUnsorted *bool) *SearchTagsUsingGETParams {
	o.SetSortUnsorted(sortUnsorted)
	return o
}

// SetSortUnsorted adds the sortUnsorted to the search tags using get params
func (o *SearchTagsUsingGETParams) SetSortUnsorted(sortUnsorted *bool) {
	o.SortUnsorted = sortUnsorted
}

// WithUnpaged adds the unpaged to the search tags using get params
func (o *SearchTagsUsingGETParams) WithUnpaged(unpaged *bool) *SearchTagsUsingGETParams {
	o.SetUnpaged(unpaged)
	return o
}

// SetUnpaged adds the unpaged to the search tags using get params
func (o *SearchTagsUsingGETParams) SetUnpaged(unpaged *bool) {
	o.Unpaged = unpaged
}

// WriteToRequest writes these params to a swagger request
func (o *SearchTagsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
