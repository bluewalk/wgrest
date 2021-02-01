// Code generated by go-swagger; DO NOT EDIT.

package wireguard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PeerGetHandlerFunc turns a function with the right signature into a peer get handler
type PeerGetHandlerFunc func(PeerGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PeerGetHandlerFunc) Handle(params PeerGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PeerGetHandler interface for that can handle valid peer get params
type PeerGetHandler interface {
	Handle(PeerGetParams, interface{}) middleware.Responder
}

// NewPeerGet creates a new http.Handler for the peer get operation
func NewPeerGet(ctx *middleware.Context, handler PeerGetHandler) *PeerGet {
	return &PeerGet{Context: ctx, Handler: handler}
}

/* PeerGet swagger:route GET /devices/{dev}/peers/{peer_id} wireguard peerGet

wireguard peer's detail

*/
type PeerGet struct {
	Context *middleware.Context
	Handler PeerGetHandler
}

func (o *PeerGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPeerGetParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
