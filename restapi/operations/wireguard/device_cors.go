// Code generated by go-swagger; DO NOT EDIT.

package wireguard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeviceCorsHandlerFunc turns a function with the right signature into a device cors handler
type DeviceCorsHandlerFunc func(DeviceCorsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeviceCorsHandlerFunc) Handle(params DeviceCorsParams) middleware.Responder {
	return fn(params)
}

// DeviceCorsHandler interface for that can handle valid device cors params
type DeviceCorsHandler interface {
	Handle(DeviceCorsParams) middleware.Responder
}

// NewDeviceCors creates a new http.Handler for the device cors operation
func NewDeviceCors(ctx *middleware.Context, handler DeviceCorsHandler) *DeviceCors {
	return &DeviceCors{Context: ctx, Handler: handler}
}

/* DeviceCors swagger:route OPTIONS /devices/{dev} wireguard deviceCors

CORS wireguard device

*/
type DeviceCors struct {
	Context *middleware.Context
	Handler DeviceCorsHandler
}

func (o *DeviceCors) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeviceCorsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
