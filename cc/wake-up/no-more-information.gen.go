// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package wakeup

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandNoMoreInformation cc.CommandID = 0x08

func init() {
	gob.Register(NoMoreInformation{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x84),
		Command:      cc.CommandID(0x08),
		Version:      1,
	}, NewNoMoreInformation)
}

func NewNoMoreInformation() cc.Command {
	return &NoMoreInformation{}
}

// <no value>
type NoMoreInformation struct {
}

func (cmd NoMoreInformation) CommandClassID() cc.CommandClassID {
	return 0x84
}

func (cmd NoMoreInformation) CommandID() cc.CommandID {
	return CommandNoMoreInformation
}

func (cmd NoMoreInformation) CommandIDString() string {
	return "WAKE_UP_NO_MORE_INFORMATION"
}

func (cmd *NoMoreInformation) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *NoMoreInformation) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
