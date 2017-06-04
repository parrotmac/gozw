// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package versionv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandCommandClassReport cc.CommandID = 0x14

func init() {
	gob.Register(CommandClassReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x86),
		Command:      cc.CommandID(0x14),
		Version:      2,
	}, NewCommandClassReport)
}

func NewCommandClassReport() cc.Command {
	return &CommandClassReport{}
}

// <no value>
type CommandClassReport struct {
	RequestedCommandClass byte

	CommandClassVersion byte
}

func (cmd CommandClassReport) CommandClassID() cc.CommandClassID {
	return 0x86
}

func (cmd CommandClassReport) CommandID() cc.CommandID {
	return CommandCommandClassReport
}

func (cmd CommandClassReport) CommandIDString() string {
	return "VERSION_COMMAND_CLASS_REPORT"
}

func (cmd *CommandClassReport) UnmarshalBinary(data []byte) error {
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

	cmd.RequestedCommandClass = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandClassVersion = payload[i]
	i++

	return nil
}

func (cmd *CommandClassReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.RequestedCommandClass)

	payload = append(payload, cmd.CommandClassVersion)

	return
}
