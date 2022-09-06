package cmdb

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new IdentifyReconcile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for IdentifyReconcile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods

type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateIdentifyReconcile(params *CreateIdentifyReconcileParams, opts ...ClientOption) (*CreateIdentifyReconcileCreated, error)
	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateIdentifyReconcile creates IdentifyReconcile

  Create IdentifyReconcile
*/
func (a *Client) CreateIdentifyReconcile(params *CreateIdentifyReconcileParams, opts ...ClientOption) (*CreateIdentifyReconcileCreated, error) {
	// TODO: Validate the params before sending

	fmt.Println("Running1")
	if params == nil {
		params = NewCreateIdentifyReconcileParams()
	}
	fmt.Println("Running2")
	op := &runtime.ClientOperation{
		ID:                 "createIdentify",
		Method:             "POST",
		PathPattern:        "/identifyreconcile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateIdentifyReconcileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	fmt.Println("Running3")
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateIdentifyReconcileCreated)
	if ok {
		return success, nil
	}

	msg := fmt.Sprintf("unexpected success response for createIdentifyReconcile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	fmt.Println(msg)
	panic(msg)

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
