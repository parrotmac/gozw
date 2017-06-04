// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package wakeupv2

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandNotification cc.CommandID = 0x07

func init() {
	gob.Register(Notification{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x84),
		Command:      cc.CommandID(0x07),
		Version:      2,
	}, NewNotification)
}

func NewNotification() cc.Command {
	return &Notification{}
}

// <no value>
type Notification struct {
}

func (cmd Notification) CommandClassID() cc.CommandClassID {
	return 0x84
}

func (cmd Notification) CommandID() cc.CommandID {
	return CommandNotification
}

func (cmd Notification) CommandIDString() string {
	return "WAKE_UP_NOTIFICATION"
}

func (cmd *Notification) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *Notification) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
