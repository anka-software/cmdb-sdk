package client

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/anka-software/cmdb-sdk/pkg/client/cmdb"
	"github.com/anka-software/cmdb-sdk/pkg/client/cmdb_meta"
	"github.com/anka-software/cmdb-sdk/pkg/client/table"
)

// Default servicenow API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "dev125776.service-now.com/api/now"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new service now  API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *MainClient {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new service now  API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *MainClient {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new service now  API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *MainClient {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(MainClient)
	cli.Transport = transport

	cli.Cmdb = cmdb.New(transport, formats)
	cli.Table = table.New(transport, formats)
	cli.CmdbMeta = cmdb_meta.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// MulticloudIaaS is a client for service now  API
// TODO: Client isimleri ayarlanacak orneginz cmdb: idenrecon
type MainClient struct {
	Cmdb      cmdb.ClientService
	CmdbMeta  cmdb_meta.ClientService
	Table     table.ClientService
	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *MainClient) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Table.SetTransport(transport)
	c.Cmdb.SetTransport(transport)
	c.CmdbMeta.SetTransport(transport)

}
