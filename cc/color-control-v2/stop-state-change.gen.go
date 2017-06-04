// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package colorcontrolv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandStopStateChange cc.CommandID = 0x07

func init() {
	gob.Register(StopStateChange{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x33),
		Command:      cc.CommandID(0x07),
		Version:      2,
	}, NewStopStateChange)
}

func NewStopStateChange() cc.Command {
	return &StopStateChange{}
}

// <no value>
type StopStateChange struct {
	CapabilityId byte
}

func (cmd StopStateChange) CommandClassID() cc.CommandClassID {
	return 0x33
}

func (cmd StopStateChange) CommandID() cc.CommandID {
	return CommandStopStateChange
}

func (cmd StopStateChange) CommandIDString() string {
	return "STOP_STATE_CHANGE"
}

func (cmd *StopStateChange) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CapabilityId = payload[i]
	i++

	return nil
}

func (cmd *StopStateChange) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.CapabilityId)

	return
}
