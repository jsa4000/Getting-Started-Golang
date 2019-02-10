// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUserByIDHandlerFunc turns a function with the right signature into a get user by Id handler
type GetUserByIDHandlerFunc func(GetUserByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserByIDHandlerFunc) Handle(params GetUserByIDParams) middleware.Responder {
	return fn(params)
}

// GetUserByIDHandler interface for that can handle valid get user by Id params
type GetUserByIDHandler interface {
	Handle(GetUserByIDParams) middleware.Responder
}

// NewGetUserByID creates a new http.Handler for the get user by Id operation
func NewGetUserByID(ctx *middleware.Context, handler GetUserByIDHandler) *GetUserByID {
	return &GetUserByID{Context: ctx, Handler: handler}
}

/*GetUserByID swagger:route GET /users/{id} Users getUserById

Find user by ID

Returns a single user

*/
type GetUserByID struct {
	Context *middleware.Context
	Handler GetUserByIDHandler
}

func (o *GetUserByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
