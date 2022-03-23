package uuid

import (
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

// ----------------------------------------------------------------------------
// Boilerplate
// ----------------------------------------------------------------------------

const ID = "uuid" // <- this is the name of the plugin for the RR

// Plugin to be registered in the app server
type Plugin struct {
	log *zap.Logger
}

func (p *Plugin) Init(log *zap.Logger) error {
	p.log = log
	return nil
}

func (p *Plugin) Serve() chan error {
	errCh := make(chan error, 1)
	// here should be your code
	// send an error into the channel if occurred
	// you may also pass this channel to other functions with goroutines to send error from them
	return errCh
}
func (p *Plugin) Stop() error {
	// this function will be called by Endure (rr container) to stop the plugin
	return nil
}

// RPC interface implementation, RR will find this interface and automatically expose the RPC endpoint with methods (rpc structure)
func (p *Plugin) RPC() interface{} {
	return &rpc{
		log: p.log,
	}
}

// Name This is not mandatory, but if you implement this interface and provide a plugin name, RR will expose the RPC method of this plugin using this name
func (p *Plugin) Name() string {
	return ID
}

// ----------------------------------------------------------------------------
// RPC
// ----------------------------------------------------------------------------

type rpc struct {
	log *zap.Logger
}

func (r *rpc) Generate(input string, output *string) error {
	u, err := uuid.NewV4()
	if err != nil {
		return err
	}
	*output = u.String()
	return nil
}
