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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetDeploymentFiltersUsingGET2Params creates a new GetDeploymentFiltersUsingGET2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentFiltersUsingGET2Params() *GetDeploymentFiltersUsingGET2Params {
	return &GetDeploymentFiltersUsingGET2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentFiltersUsingGET2ParamsWithTimeout creates a new GetDeploymentFiltersUsingGET2Params object
// with the ability to set a timeout on a request.
func NewGetDeploymentFiltersUsingGET2ParamsWithTimeout(timeout time.Duration) *GetDeploymentFiltersUsingGET2Params {
	return &GetDeploymentFiltersUsingGET2Params{
		timeout: timeout,
	}
}

// NewGetDeploymentFiltersUsingGET2ParamsWithContext creates a new GetDeploymentFiltersUsingGET2Params object
// with the ability to set a context for a request.
func NewGetDeploymentFiltersUsingGET2ParamsWithContext(ctx context.Context) *GetDeploymentFiltersUsingGET2Params {
	return &GetDeploymentFiltersUsingGET2Params{
		Context: ctx,
	}
}

// NewGetDeploymentFiltersUsingGET2ParamsWithHTTPClient creates a new GetDeploymentFiltersUsingGET2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentFiltersUsingGET2ParamsWithHTTPClient(client *http.Client) *GetDeploymentFiltersUsingGET2Params {
	return &GetDeploymentFiltersUsingGET2Params{
		HTTPClient: client,
	}
}

/*
GetDeploymentFiltersUsingGET2Params contains all the parameters to send to the API endpoint

	for the get deployment filters using g e t 2 operation.

	Typically these are written to a http.Request.
*/
type GetDeploymentFiltersUsingGET2Params struct {

	// ISO3Country.
	ISO3Country *string

	// ISO3Language.
	ISO3Language *string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	// Country.
	Country *string

	// DisplayCountry.
	DisplayCountry *string

	// DisplayLanguage.
	DisplayLanguage *string

	// DisplayName.
	DisplayName *string

	// DisplayScript.
	DisplayScript *string

	// DisplayVariant.
	DisplayVariant *string

	// Language.
	Language *string

	// Script.
	Script *string

	// UnicodeLocaleAttributes.
	UnicodeLocaleAttributes []string

	// UnicodeLocaleKeys.
	UnicodeLocaleKeys []string

	// Variant.
	Variant *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deployment filters using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentFiltersUsingGET2Params) WithDefaults() *GetDeploymentFiltersUsingGET2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployment filters using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentFiltersUsingGET2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithTimeout(timeout time.Duration) *GetDeploymentFiltersUsingGET2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithContext(ctx context.Context) *GetDeploymentFiltersUsingGET2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithHTTPClient(client *http.Client) *GetDeploymentFiltersUsingGET2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithISO3Country adds the iSO3Country to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithISO3Country(iSO3Country *string) *GetDeploymentFiltersUsingGET2Params {
	o.SetISO3Country(iSO3Country)
	return o
}

// SetISO3Country adds the iSO3Country to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetISO3Country(iSO3Country *string) {
	o.ISO3Country = iSO3Country
}

// WithISO3Language adds the iSO3Language to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithISO3Language(iSO3Language *string) *GetDeploymentFiltersUsingGET2Params {
	o.SetISO3Language(iSO3Language)
	return o
}

// SetISO3Language adds the iSO3Language to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetISO3Language(iSO3Language *string) {
	o.ISO3Language = iSO3Language
}

// WithAPIVersion adds the aPIVersion to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithAPIVersion(aPIVersion *string) *GetDeploymentFiltersUsingGET2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithCountry adds the country to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithCountry(country *string) *GetDeploymentFiltersUsingGET2Params {
	o.SetCountry(country)
	return o
}

// SetCountry adds the country to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetCountry(country *string) {
	o.Country = country
}

// WithDisplayCountry adds the displayCountry to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithDisplayCountry(displayCountry *string) *GetDeploymentFiltersUsingGET2Params {
	o.SetDisplayCountry(displayCountry)
	return o
}

// SetDisplayCountry adds the displayCountry to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetDisplayCountry(displayCountry *string) {
	o.DisplayCountry = displayCountry
}

// WithDisplayLanguage adds the displayLanguage to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithDisplayLanguage(displayLanguage *string) *GetDeploymentFiltersUsingGET2Params {
	o.SetDisplayLanguage(displayLanguage)
	return o
}

// SetDisplayLanguage adds the displayLanguage to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetDisplayLanguage(displayLanguage *string) {
	o.DisplayLanguage = displayLanguage
}

// WithDisplayName adds the displayName to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithDisplayName(displayName *string) *GetDeploymentFiltersUsingGET2Params {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithDisplayScript adds the displayScript to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithDisplayScript(displayScript *string) *GetDeploymentFiltersUsingGET2Params {
	o.SetDisplayScript(displayScript)
	return o
}

// SetDisplayScript adds the displayScript to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetDisplayScript(displayScript *string) {
	o.DisplayScript = displayScript
}

// WithDisplayVariant adds the displayVariant to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithDisplayVariant(displayVariant *string) *GetDeploymentFiltersUsingGET2Params {
	o.SetDisplayVariant(displayVariant)
	return o
}

// SetDisplayVariant adds the displayVariant to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetDisplayVariant(displayVariant *string) {
	o.DisplayVariant = displayVariant
}

// WithLanguage adds the language to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithLanguage(language *string) *GetDeploymentFiltersUsingGET2Params {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetLanguage(language *string) {
	o.Language = language
}

// WithScript adds the script to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithScript(script *string) *GetDeploymentFiltersUsingGET2Params {
	o.SetScript(script)
	return o
}

// SetScript adds the script to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetScript(script *string) {
	o.Script = script
}

// WithUnicodeLocaleAttributes adds the unicodeLocaleAttributes to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithUnicodeLocaleAttributes(unicodeLocaleAttributes []string) *GetDeploymentFiltersUsingGET2Params {
	o.SetUnicodeLocaleAttributes(unicodeLocaleAttributes)
	return o
}

// SetUnicodeLocaleAttributes adds the unicodeLocaleAttributes to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetUnicodeLocaleAttributes(unicodeLocaleAttributes []string) {
	o.UnicodeLocaleAttributes = unicodeLocaleAttributes
}

// WithUnicodeLocaleKeys adds the unicodeLocaleKeys to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithUnicodeLocaleKeys(unicodeLocaleKeys []string) *GetDeploymentFiltersUsingGET2Params {
	o.SetUnicodeLocaleKeys(unicodeLocaleKeys)
	return o
}

// SetUnicodeLocaleKeys adds the unicodeLocaleKeys to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetUnicodeLocaleKeys(unicodeLocaleKeys []string) {
	o.UnicodeLocaleKeys = unicodeLocaleKeys
}

// WithVariant adds the variant to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) WithVariant(variant *string) *GetDeploymentFiltersUsingGET2Params {
	o.SetVariant(variant)
	return o
}

// SetVariant adds the variant to the get deployment filters using g e t 2 params
func (o *GetDeploymentFiltersUsingGET2Params) SetVariant(variant *string) {
	o.Variant = variant
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentFiltersUsingGET2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.UnicodeLocaleAttributes != nil {

		// binding items for unicodeLocaleAttributes
		joinedUnicodeLocaleAttributes := o.bindParamUnicodeLocaleAttributes(reg)

		// query array param unicodeLocaleAttributes
		if err := r.SetQueryParam("unicodeLocaleAttributes", joinedUnicodeLocaleAttributes...); err != nil {
			return err
		}
	}

	if o.UnicodeLocaleKeys != nil {

		// binding items for unicodeLocaleKeys
		joinedUnicodeLocaleKeys := o.bindParamUnicodeLocaleKeys(reg)

		// query array param unicodeLocaleKeys
		if err := r.SetQueryParam("unicodeLocaleKeys", joinedUnicodeLocaleKeys...); err != nil {
			return err
		}
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

// bindParamGetDeploymentFiltersUsingGET2 binds the parameter unicodeLocaleAttributes
func (o *GetDeploymentFiltersUsingGET2Params) bindParamUnicodeLocaleAttributes(formats strfmt.Registry) []string {
	unicodeLocaleAttributesIR := o.UnicodeLocaleAttributes

	var unicodeLocaleAttributesIC []string
	for _, unicodeLocaleAttributesIIR := range unicodeLocaleAttributesIR { // explode []string

		unicodeLocaleAttributesIIV := unicodeLocaleAttributesIIR // string as string
		unicodeLocaleAttributesIC = append(unicodeLocaleAttributesIC, unicodeLocaleAttributesIIV)
	}

	// items.CollectionFormat: "multi"
	unicodeLocaleAttributesIS := swag.JoinByFormat(unicodeLocaleAttributesIC, "multi")

	return unicodeLocaleAttributesIS
}

// bindParamGetDeploymentFiltersUsingGET2 binds the parameter unicodeLocaleKeys
func (o *GetDeploymentFiltersUsingGET2Params) bindParamUnicodeLocaleKeys(formats strfmt.Registry) []string {
	unicodeLocaleKeysIR := o.UnicodeLocaleKeys

	var unicodeLocaleKeysIC []string
	for _, unicodeLocaleKeysIIR := range unicodeLocaleKeysIR { // explode []string

		unicodeLocaleKeysIIV := unicodeLocaleKeysIIR // string as string
		unicodeLocaleKeysIC = append(unicodeLocaleKeysIC, unicodeLocaleKeysIIV)
	}

	// items.CollectionFormat: "multi"
	unicodeLocaleKeysIS := swag.JoinByFormat(unicodeLocaleKeysIC, "multi")

	return unicodeLocaleKeysIS
}
