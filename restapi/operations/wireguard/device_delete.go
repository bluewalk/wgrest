// Code generated by go-swagger; DO NOT EDIT.

package wireguard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeviceDeleteHandlerFunc turns a function with the right signature into a device delete handler
type DeviceDeleteHandlerFunc func(DeviceDeleteParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeviceDeleteHandlerFunc) Handle(params DeviceDeleteParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeviceDeleteHandler interface for that can handle valid device delete params
type DeviceDeleteHandler interface {
	Handle(DeviceDeleteParams, interface{}) middleware.Responder
}

// NewDeviceDelete creates a new http.Handler for the device delete operation
func NewDeviceDelete(ctx *middleware.Context, handler DeviceDeleteHandler) *DeviceDelete {
	return &DeviceDelete{Context: ctx, Handler: handler}
}

/* DeviceDelete swagger:route DELETE /devices/{dev} wireguard deviceDelete

delete wireguard interface

*/
type DeviceDelete struct {
	Context *middleware.Context
	Handler DeviceDeleteHandler
}

func (o *DeviceDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeviceDeleteParams()
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
