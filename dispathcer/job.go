package models

// Payload struct of represents the one load
type Payload struct {
}

// Job struct of represents the job to be run
type Job struct {
	Payload Payload
}

// RPCCall func of remote rpc call
func (p *Payload) RPCCall() error {
	return nil
}
