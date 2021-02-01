// Code generated by go-swagger; DO NOT EDIT.

package wireguard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/suquant/wgrest/models"
)

// PeerCorsOKCode is the HTTP code returned for type PeerCorsOK
const PeerCorsOKCode int = 200

/*PeerCorsOK ok

swagger:response peerCorsOK
*/
type PeerCorsOK struct {
	/*

	  Default: []interface {}{"GET", "DELETE"}
	*/
	AccessControlAllowMethods []string `json:"Access-Control-Allow-Methods"`
}

// NewPeerCorsOK creates PeerCorsOK with default headers values
func NewPeerCorsOK() *PeerCorsOK {

	var (
		// initialize headers with default values

		accessControlAllowMethodsDefault = []string{"GET", "DELETE"}
	)

	return &PeerCorsOK{

		AccessControlAllowMethods: accessControlAllowMethodsDefault,
	}
}

// WithAccessControlAllowMethods adds the accessControlAllowMethods to the peer cors o k response
func (o *PeerCorsOK) WithAccessControlAllowMethods(accessControlAllowMethods []string) *PeerCorsOK {
	o.AccessControlAllowMethods = accessControlAllowMethods
	return o
}

// SetAccessControlAllowMethods sets the accessControlAllowMethods to the peer cors o k response
func (o *PeerCorsOK) SetAccessControlAllowMethods(accessControlAllowMethods []string) {
	o.AccessControlAllowMethods = accessControlAllowMethods
}

// WriteResponse to the client
func (o *PeerCorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Access-Control-Allow-Methods

	var accessControlAllowMethodsIR []string
	for _, accessControlAllowMethodsI := range o.AccessControlAllowMethods {
		accessControlAllowMethodsIS := accessControlAllowMethodsI
		if accessControlAllowMethodsIS != "" {
			accessControlAllowMethodsIR = append(accessControlAllowMethodsIR, accessControlAllowMethodsIS)
		}
	}
	accessControlAllowMethods := swag.JoinByFormat(accessControlAllowMethodsIR, "")
	if len(accessControlAllowMethods) > 0 {
		hv := accessControlAllowMethods[0]
		if hv != "" {
			rw.Header().Set("Access-Control-Allow-Methods", hv)
		}
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*PeerCorsDefault error

swagger:response peerCorsDefault
*/
type PeerCorsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPeerCorsDefault creates PeerCorsDefault with default headers values
func NewPeerCorsDefault(code int) *PeerCorsDefault {
	if code <= 0 {
		code = 500
	}

	return &PeerCorsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the peer cors default response
func (o *PeerCorsDefault) WithStatusCode(code int) *PeerCorsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the peer cors default response
func (o *PeerCorsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the peer cors default response
func (o *PeerCorsDefault) WithPayload(payload *models.Error) *PeerCorsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the peer cors default response
func (o *PeerCorsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PeerCorsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
