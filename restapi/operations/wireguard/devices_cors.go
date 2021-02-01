// Code generated by go-swagger; DO NOT EDIT.

package wireguard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DevicesCorsHandlerFunc turns a function with the right signature into a devices cors handler
type DevicesCorsHandlerFunc func(DevicesCorsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DevicesCorsHandlerFunc) Handle(params DevicesCorsParams) middleware.Responder {
	return fn(params)
}

// DevicesCorsHandler interface for that can handle valid devices cors params
type DevicesCorsHandler interface {
	Handle(DevicesCorsParams) middleware.Responder
}

// NewDevicesCors creates a new http.Handler for the devices cors operation
func NewDevicesCors(ctx *middleware.Context, handler DevicesCorsHandler) *DevicesCors {
	return &DevicesCors{Context: ctx, Handler: handler}
}

/* DevicesCors swagger:route OPTIONS /devices/ wireguard devicesCors

CORS wireguard device

*/
type DevicesCors struct {
	Context *middleware.Context
	Handler DevicesCorsHandler
}

func (o *DevicesCors) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDevicesCorsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
