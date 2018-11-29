// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new agents API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for agents API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddMySqldExporter adds my sqld exporter adds mysqld exporter agent
*/
func (a *Client) AddMySqldExporter(params *AddMySqldExporterParams) (*AddMySqldExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMySqldExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddMySQLdExporter",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Agents/AddMySQLdExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddMySqldExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddMySqldExporterOK), nil

}

/*
AddNodeExporter adds node exporter adds node exporter agent
*/
func (a *Client) AddNodeExporter(params *AddNodeExporterParams) (*AddNodeExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddNodeExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddNodeExporter",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Agents/AddNodeExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddNodeExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddNodeExporterOK), nil

}

/*
AddPMMAgent adds PMM agent adds pmm agent agent
*/
func (a *Client) AddPMMAgent(params *AddPMMAgentParams) (*AddPMMAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPMMAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddPMMAgent",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Agents/AddPMMAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddPMMAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddPMMAgentOK), nil

}

/*
DisableAgent disables agent disables and stops agent
*/
func (a *Client) DisableAgent(params *DisableAgentParams) (*DisableAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DisableAgent",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Agents/DisableAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisableAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DisableAgentOK), nil

}

/*
EnableAgent enables agent enables and starts agent
*/
func (a *Client) EnableAgent(params *EnableAgentParams) (*EnableAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EnableAgent",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Agents/EnableAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EnableAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EnableAgentOK), nil

}

/*
GetAgent gets agent returns a single agent by ID
*/
func (a *Client) GetAgent(params *GetAgentParams) (*GetAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAgent",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Agents/GetAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAgentOK), nil

}

/*
ListAgents lists agents returns a list of all agents
*/
func (a *Client) ListAgents(params *ListAgentsParams) (*ListAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAgentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListAgents",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Agents/ListAgents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAgentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAgentsOK), nil

}

/*
RemoveAgent removes agent removes agent
*/
func (a *Client) RemoveAgent(params *RemoveAgentParams) (*RemoveAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemoveAgent",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Agents/RemoveAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveAgentOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}