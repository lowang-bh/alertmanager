// Code generated by go-swagger; DO NOT EDIT.

package receiver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new receiver API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for receiver API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetReceivers Get list of all receivers (name of notification integrations)
*/
func (a *Client) GetReceivers(params *GetReceiversParams) (*GetReceiversOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReceiversParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReceivers",
		Method:             "GET",
		PathPattern:        "/receivers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReceiversReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReceiversOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}