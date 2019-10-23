// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDeploymentFiltersUsingGETParams creates a new GetDeploymentFiltersUsingGETParams object
// with the default values initialized.
func NewGetDeploymentFiltersUsingGETParams() *GetDeploymentFiltersUsingGETParams {
	var ()
	return &GetDeploymentFiltersUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentFiltersUsingGETParamsWithTimeout creates a new GetDeploymentFiltersUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentFiltersUsingGETParamsWithTimeout(timeout time.Duration) *GetDeploymentFiltersUsingGETParams {
	var ()
	return &GetDeploymentFiltersUsingGETParams{

		timeout: timeout,
	}
}

// NewGetDeploymentFiltersUsingGETParamsWithContext creates a new GetDeploymentFiltersUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentFiltersUsingGETParamsWithContext(ctx context.Context) *GetDeploymentFiltersUsingGETParams {
	var ()
	return &GetDeploymentFiltersUsingGETParams{

		Context: ctx,
	}
}

// NewGetDeploymentFiltersUsingGETParamsWithHTTPClient creates a new GetDeploymentFiltersUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentFiltersUsingGETParamsWithHTTPClient(client *http.Client) *GetDeploymentFiltersUsingGETParams {
	var ()
	return &GetDeploymentFiltersUsingGETParams{
		HTTPClient: client,
	}
}

/*GetDeploymentFiltersUsingGETParams contains all the parameters to send to the API endpoint
for the get deployment filters using g e t operation typically these are written to a http.Request
*/
type GetDeploymentFiltersUsingGETParams struct {

	/*ISO3Country*/
	ISO3Country *string
	/*ISO3Language*/
	ISO3Language *string
	/*Country*/
	Country *string
	/*DisplayCountry*/
	DisplayCountry *string
	/*DisplayLanguage*/
	DisplayLanguage *string
	/*DisplayName*/
	DisplayName *string
	/*DisplayScript*/
	DisplayScript *string
	/*DisplayVariant*/
	DisplayVariant *string
	/*Language*/
	Language *string
	/*Script*/
	Script *string
	/*UnicodeLocaleAttributes*/
	UnicodeLocaleAttributes []string
	/*UnicodeLocaleKeys*/
	UnicodeLocaleKeys []string
	/*Variant*/
	Variant *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithTimeout(timeout time.Duration) *GetDeploymentFiltersUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithContext(ctx context.Context) *GetDeploymentFiltersUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithHTTPClient(client *http.Client) *GetDeploymentFiltersUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithISO3Country adds the iSO3Country to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithISO3Country(iSO3Country *string) *GetDeploymentFiltersUsingGETParams {
	o.SetISO3Country(iSO3Country)
	return o
}

// SetISO3Country adds the iSO3Country to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetISO3Country(iSO3Country *string) {
	o.ISO3Country = iSO3Country
}

// WithISO3Language adds the iSO3Language to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithISO3Language(iSO3Language *string) *GetDeploymentFiltersUsingGETParams {
	o.SetISO3Language(iSO3Language)
	return o
}

// SetISO3Language adds the iSO3Language to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetISO3Language(iSO3Language *string) {
	o.ISO3Language = iSO3Language
}

// WithCountry adds the country to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithCountry(country *string) *GetDeploymentFiltersUsingGETParams {
	o.SetCountry(country)
	return o
}

// SetCountry adds the country to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetCountry(country *string) {
	o.Country = country
}

// WithDisplayCountry adds the displayCountry to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithDisplayCountry(displayCountry *string) *GetDeploymentFiltersUsingGETParams {
	o.SetDisplayCountry(displayCountry)
	return o
}

// SetDisplayCountry adds the displayCountry to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetDisplayCountry(displayCountry *string) {
	o.DisplayCountry = displayCountry
}

// WithDisplayLanguage adds the displayLanguage to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithDisplayLanguage(displayLanguage *string) *GetDeploymentFiltersUsingGETParams {
	o.SetDisplayLanguage(displayLanguage)
	return o
}

// SetDisplayLanguage adds the displayLanguage to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetDisplayLanguage(displayLanguage *string) {
	o.DisplayLanguage = displayLanguage
}

// WithDisplayName adds the displayName to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithDisplayName(displayName *string) *GetDeploymentFiltersUsingGETParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithDisplayScript adds the displayScript to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithDisplayScript(displayScript *string) *GetDeploymentFiltersUsingGETParams {
	o.SetDisplayScript(displayScript)
	return o
}

// SetDisplayScript adds the displayScript to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetDisplayScript(displayScript *string) {
	o.DisplayScript = displayScript
}

// WithDisplayVariant adds the displayVariant to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithDisplayVariant(displayVariant *string) *GetDeploymentFiltersUsingGETParams {
	o.SetDisplayVariant(displayVariant)
	return o
}

// SetDisplayVariant adds the displayVariant to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetDisplayVariant(displayVariant *string) {
	o.DisplayVariant = displayVariant
}

// WithLanguage adds the language to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithLanguage(language *string) *GetDeploymentFiltersUsingGETParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetLanguage(language *string) {
	o.Language = language
}

// WithScript adds the script to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithScript(script *string) *GetDeploymentFiltersUsingGETParams {
	o.SetScript(script)
	return o
}

// SetScript adds the script to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetScript(script *string) {
	o.Script = script
}

// WithUnicodeLocaleAttributes adds the unicodeLocaleAttributes to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithUnicodeLocaleAttributes(unicodeLocaleAttributes []string) *GetDeploymentFiltersUsingGETParams {
	o.SetUnicodeLocaleAttributes(unicodeLocaleAttributes)
	return o
}

// SetUnicodeLocaleAttributes adds the unicodeLocaleAttributes to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetUnicodeLocaleAttributes(unicodeLocaleAttributes []string) {
	o.UnicodeLocaleAttributes = unicodeLocaleAttributes
}

// WithUnicodeLocaleKeys adds the unicodeLocaleKeys to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithUnicodeLocaleKeys(unicodeLocaleKeys []string) *GetDeploymentFiltersUsingGETParams {
	o.SetUnicodeLocaleKeys(unicodeLocaleKeys)
	return o
}

// SetUnicodeLocaleKeys adds the unicodeLocaleKeys to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetUnicodeLocaleKeys(unicodeLocaleKeys []string) {
	o.UnicodeLocaleKeys = unicodeLocaleKeys
}

// WithVariant adds the variant to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) WithVariant(variant *string) *GetDeploymentFiltersUsingGETParams {
	o.SetVariant(variant)
	return o
}

// SetVariant adds the variant to the get deployment filters using get params
func (o *GetDeploymentFiltersUsingGETParams) SetVariant(variant *string) {
	o.Variant = variant
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentFiltersUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ISO3Country != nil {

		// query param ISO3Country
		var qrISO3Country string
		if o.ISO3Country != nil {
			qrISO3Country = *o.ISO3Country
		}
		qISO3Country := qrISO3Country
		if qISO3Country != "" {
			if err := r.SetQueryParam("ISO3Country", qISO3Country); err != nil {
				return err
			}
		}

	}

	if o.ISO3Language != nil {

		// query param ISO3Language
		var qrISO3Language string
		if o.ISO3Language != nil {
			qrISO3Language = *o.ISO3Language
		}
		qISO3Language := qrISO3Language
		if qISO3Language != "" {
			if err := r.SetQueryParam("ISO3Language", qISO3Language); err != nil {
				return err
			}
		}

	}

	if o.Country != nil {

		// query param country
		var qrCountry string
		if o.Country != nil {
			qrCountry = *o.Country
		}
		qCountry := qrCountry
		if qCountry != "" {
			if err := r.SetQueryParam("country", qCountry); err != nil {
				return err
			}
		}

	}

	if o.DisplayCountry != nil {

		// query param displayCountry
		var qrDisplayCountry string
		if o.DisplayCountry != nil {
			qrDisplayCountry = *o.DisplayCountry
		}
		qDisplayCountry := qrDisplayCountry
		if qDisplayCountry != "" {
			if err := r.SetQueryParam("displayCountry", qDisplayCountry); err != nil {
				return err
			}
		}

	}

	if o.DisplayLanguage != nil {

		// query param displayLanguage
		var qrDisplayLanguage string
		if o.DisplayLanguage != nil {
			qrDisplayLanguage = *o.DisplayLanguage
		}
		qDisplayLanguage := qrDisplayLanguage
		if qDisplayLanguage != "" {
			if err := r.SetQueryParam("displayLanguage", qDisplayLanguage); err != nil {
				return err
			}
		}

	}

	if o.DisplayName != nil {

		// query param displayName
		var qrDisplayName string
		if o.DisplayName != nil {
			qrDisplayName = *o.DisplayName
		}
		qDisplayName := qrDisplayName
		if qDisplayName != "" {
			if err := r.SetQueryParam("displayName", qDisplayName); err != nil {
				return err
			}
		}

	}

	if o.DisplayScript != nil {

		// query param displayScript
		var qrDisplayScript string
		if o.DisplayScript != nil {
			qrDisplayScript = *o.DisplayScript
		}
		qDisplayScript := qrDisplayScript
		if qDisplayScript != "" {
			if err := r.SetQueryParam("displayScript", qDisplayScript); err != nil {
				return err
			}
		}

	}

	if o.DisplayVariant != nil {

		// query param displayVariant
		var qrDisplayVariant string
		if o.DisplayVariant != nil {
			qrDisplayVariant = *o.DisplayVariant
		}
		qDisplayVariant := qrDisplayVariant
		if qDisplayVariant != "" {
			if err := r.SetQueryParam("displayVariant", qDisplayVariant); err != nil {
				return err
			}
		}

	}

	if o.Language != nil {

		// query param language
		var qrLanguage string
		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {
			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}

	}

	if o.Script != nil {

		// query param script
		var qrScript string
		if o.Script != nil {
			qrScript = *o.Script
		}
		qScript := qrScript
		if qScript != "" {
			if err := r.SetQueryParam("script", qScript); err != nil {
				return err
			}
		}

	}

	valuesUnicodeLocaleAttributes := o.UnicodeLocaleAttributes

	joinedUnicodeLocaleAttributes := swag.JoinByFormat(valuesUnicodeLocaleAttributes, "multi")
	// query array param unicodeLocaleAttributes
	if err := r.SetQueryParam("unicodeLocaleAttributes", joinedUnicodeLocaleAttributes...); err != nil {
		return err
	}

	valuesUnicodeLocaleKeys := o.UnicodeLocaleKeys

	joinedUnicodeLocaleKeys := swag.JoinByFormat(valuesUnicodeLocaleKeys, "multi")
	// query array param unicodeLocaleKeys
	if err := r.SetQueryParam("unicodeLocaleKeys", joinedUnicodeLocaleKeys...); err != nil {
		return err
	}

	if o.Variant != nil {

		// query param variant
		var qrVariant string
		if o.Variant != nil {
			qrVariant = *o.Variant
		}
		qVariant := qrVariant
		if qVariant != "" {
			if err := r.SetQueryParam("variant", qVariant); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
