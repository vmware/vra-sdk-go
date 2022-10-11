// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Operation Operation
//
// swagger:model Operation
type Operation struct {

	// action
	// Enum: [GET POST PATCH PUT DELETE OPTIONS]
	Action string `json:"action,omitempty"`

	// authorization context
	AuthorizationContext *AuthorizationContext `json:"authorizationContext,omitempty"`

	// body raw
	BodyRaw interface{} `json:"bodyRaw,omitempty"`

	// commit
	Commit bool `json:"commit,omitempty"`

	// completion
	Completion CompletionHandler `json:"completion,omitempty"`

	// connection sharing
	ConnectionSharing bool `json:"connectionSharing,omitempty"`

	// connection tag
	ConnectionTag string `json:"connectionTag,omitempty"`

	// content length
	ContentLength int64 `json:"contentLength,omitempty"`

	// content type
	ContentType string `json:"contentType,omitempty"`

	// context Id
	ContextID string `json:"contextId,omitempty"`

	// cookies
	Cookies map[string]string `json:"cookies,omitempty"`

	// error response body
	ErrorResponseBody *ServiceErrorResponse `json:"errorResponseBody,omitempty"`

	// expiration micros utc
	ExpirationMicrosUtc int64 `json:"expirationMicrosUtc,omitempty"`

	// failure logging disabled
	FailureLoggingDisabled bool `json:"failureLoggingDisabled,omitempty"`

	// forwarded
	Forwarded bool `json:"forwarded,omitempty"`

	// forwarding disabled
	ForwardingDisabled bool `json:"forwardingDisabled,omitempty"`

	// from replication
	FromReplication bool `json:"fromReplication,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// keep alive
	KeepAlive bool `json:"keepAlive,omitempty"`

	// notification
	Notification bool `json:"notification,omitempty"`

	// notification disabled
	NotificationDisabled bool `json:"notificationDisabled,omitempty"`

	// options
	Options []string `json:"options"`

	// peer certificate chain
	PeerCertificateChain []*X509Certificate `json:"peerCertificateChain"`

	// peer principal
	PeerPrincipal *Principal `json:"peerPrincipal,omitempty"`

	// referer
	Referer *URI `json:"referer,omitempty"`

	// referer as string
	RefererAsString string `json:"refererAsString,omitempty"`

	// remote
	Remote bool `json:"remote,omitempty"`

	// replication disabled
	ReplicationDisabled bool `json:"replicationDisabled,omitempty"`

	// request headers
	RequestHeaders map[string]string `json:"requestHeaders,omitempty"`

	// response headers
	ResponseHeaders map[string]string `json:"responseHeaders,omitempty"`

	// retries remaining
	RetriesRemaining int32 `json:"retriesRemaining,omitempty"`

	// retry count
	RetryCount int32 `json:"retryCount,omitempty"`

	// socket context
	SocketContext *SocketContext `json:"socketContext,omitempty"`

	// status code
	StatusCode int32 `json:"statusCode,omitempty"`

	// synchronize
	Synchronize bool `json:"synchronize,omitempty"`

	// synchronize owner
	SynchronizeOwner bool `json:"synchronizeOwner,omitempty"`

	// synchronize peer
	SynchronizePeer bool `json:"synchronizePeer,omitempty"`

	// target replicated
	TargetReplicated bool `json:"targetReplicated,omitempty"`

	// transaction Id
	TransactionID string `json:"transactionId,omitempty"`

	// update
	Update bool `json:"update,omitempty"`

	// uri
	URI *URI `json:"uri,omitempty"`

	// within transaction
	WithinTransaction bool `json:"withinTransaction,omitempty"`
}

// Validate validates this operation
func (m *Operation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorizationContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorResponseBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerCertificateChain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerPrincipal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSocketContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var operationTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GET","POST","PATCH","PUT","DELETE","OPTIONS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operationTypeActionPropEnum = append(operationTypeActionPropEnum, v)
	}
}

const (

	// OperationActionGET captures enum value "GET"
	OperationActionGET string = "GET"

	// OperationActionPOST captures enum value "POST"
	OperationActionPOST string = "POST"

	// OperationActionPATCH captures enum value "PATCH"
	OperationActionPATCH string = "PATCH"

	// OperationActionPUT captures enum value "PUT"
	OperationActionPUT string = "PUT"

	// OperationActionDELETE captures enum value "DELETE"
	OperationActionDELETE string = "DELETE"

	// OperationActionOPTIONS captures enum value "OPTIONS"
	OperationActionOPTIONS string = "OPTIONS"
)

// prop value enum
func (m *Operation) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, operationTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Operation) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *Operation) validateAuthorizationContext(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthorizationContext) { // not required
		return nil
	}

	if m.AuthorizationContext != nil {
		if err := m.AuthorizationContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authorizationContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authorizationContext")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) validateErrorResponseBody(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrorResponseBody) { // not required
		return nil
	}

	if m.ErrorResponseBody != nil {
		if err := m.ErrorResponseBody.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorResponseBody")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorResponseBody")
			}
			return err
		}
	}

	return nil
}

var operationOptionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONNECTION_SHARING","KEEP_ALIVE","REPLICATED","FORWARDED","REPLICATION_DISABLED","CLONING_DISABLED","NOTIFICATION_DISABLED","REPLICATED_TARGET","FAILURE_LOGGING_DISABLED","REMOTE","RATE_LIMITED","SOCKET_ACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operationOptionsItemsEnum = append(operationOptionsItemsEnum, v)
	}
}

func (m *Operation) validateOptionsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, operationOptionsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Operation) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Options) { // not required
		return nil
	}

	for i := 0; i < len(m.Options); i++ {

		// value enum
		if err := m.validateOptionsItemsEnum("options"+"."+strconv.Itoa(i), "body", m.Options[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *Operation) validatePeerCertificateChain(formats strfmt.Registry) error {
	if swag.IsZero(m.PeerCertificateChain) { // not required
		return nil
	}

	for i := 0; i < len(m.PeerCertificateChain); i++ {
		if swag.IsZero(m.PeerCertificateChain[i]) { // not required
			continue
		}

		if m.PeerCertificateChain[i] != nil {
			if err := m.PeerCertificateChain[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peerCertificateChain" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("peerCertificateChain" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Operation) validatePeerPrincipal(formats strfmt.Registry) error {
	if swag.IsZero(m.PeerPrincipal) { // not required
		return nil
	}

	if m.PeerPrincipal != nil {
		if err := m.PeerPrincipal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peerPrincipal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peerPrincipal")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) validateReferer(formats strfmt.Registry) error {
	if swag.IsZero(m.Referer) { // not required
		return nil
	}

	if m.Referer != nil {
		if err := m.Referer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("referer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("referer")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) validateSocketContext(formats strfmt.Registry) error {
	if swag.IsZero(m.SocketContext) { // not required
		return nil
	}

	if m.SocketContext != nil {
		if err := m.SocketContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("socketContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("socketContext")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) validateURI(formats strfmt.Registry) error {
	if swag.IsZero(m.URI) { // not required
		return nil
	}

	if m.URI != nil {
		if err := m.URI.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uri")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uri")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this operation based on the context it is used
func (m *Operation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthorizationContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrorResponseBody(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeerCertificateChain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeerPrincipal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReferer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSocketContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Operation) contextValidateAuthorizationContext(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthorizationContext != nil {
		if err := m.AuthorizationContext.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authorizationContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authorizationContext")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) contextValidateErrorResponseBody(ctx context.Context, formats strfmt.Registry) error {

	if m.ErrorResponseBody != nil {
		if err := m.ErrorResponseBody.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorResponseBody")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorResponseBody")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) contextValidatePeerCertificateChain(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PeerCertificateChain); i++ {

		if m.PeerCertificateChain[i] != nil {
			if err := m.PeerCertificateChain[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peerCertificateChain" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("peerCertificateChain" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Operation) contextValidatePeerPrincipal(ctx context.Context, formats strfmt.Registry) error {

	if m.PeerPrincipal != nil {
		if err := m.PeerPrincipal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peerPrincipal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peerPrincipal")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) contextValidateReferer(ctx context.Context, formats strfmt.Registry) error {

	if m.Referer != nil {
		if err := m.Referer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("referer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("referer")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) contextValidateSocketContext(ctx context.Context, formats strfmt.Registry) error {

	if m.SocketContext != nil {
		if err := m.SocketContext.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("socketContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("socketContext")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) contextValidateURI(ctx context.Context, formats strfmt.Registry) error {

	if m.URI != nil {
		if err := m.URI.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uri")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uri")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Operation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Operation) UnmarshalBinary(b []byte) error {
	var res Operation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}