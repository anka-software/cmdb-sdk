package table

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
	GetTableItems(params *GetTableItemsParams, opts ...ClientOption) (*GetTableItemsOK, error)
	DeleteRecord(params *DeleteRecordParams, opts ...ClientOption) (*DeleteRecordOK, error)
	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateIdentifyReconcile creates IdentifyReconcile

  Create IdentifyReconcile
*/
func (a *Client) GetTableItems(params *GetTableItemsParams, opts ...ClientOption) (*GetTableItemsOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetTableItemParams()
	}

	op := &runtime.ClientOperation{
		ID:                 "getTableItems",
		Method:             "GET",
		PathPattern:        "/table/{tableName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTableItemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTableItemsOK)
	if ok {
		return success, nil
	}

	msg := fmt.Sprintf("unexpected success response for createIdentifyReconcile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

func (a *Client) DeleteRecord(params *DeleteRecordParams, opts ...ClientOption) (*DeleteRecordOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewDeleteRecordParams()
	}

	op := &runtime.ClientOperation{
		ID:                 "deleteRecord",
		Method:             "DELETE",
		PathPattern:        "/table/{tableName}/{sys_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRecordReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRecordOK)
	if ok {
		return success, nil
	}

	msg := fmt.Sprintf("unexpected success response for createIdentifyReconcile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
