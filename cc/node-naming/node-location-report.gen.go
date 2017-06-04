// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package nodenaming

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandNodeLocationReport cc.CommandID = 0x06

func init() {
	gob.Register(NodeLocationReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x77),
		Command:      cc.CommandID(0x06),
		Version:      1,
	}, NewNodeLocationReport)
}

func NewNodeLocationReport() cc.Command {
	return &NodeLocationReport{}
}

// <no value>
type NodeLocationReport struct {
	Level struct {
		CharPresentation byte
	}

	NodeLocationChar string
}

func (cmd NodeLocationReport) CommandClassID() cc.CommandClassID {
	return 0x77
}

func (cmd NodeLocationReport) CommandID() cc.CommandID {
	return CommandNodeLocationReport
}

func (cmd NodeLocationReport) CommandIDString() string {
	return "NODE_NAMING_NODE_LOCATION_REPORT"
}

func (cmd *NodeLocationReport) UnmarshalBinary(data []byte) error {
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

	cmd.Level.CharPresentation = (payload[i] & 0x07)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeLocationChar = string(payload[i : i+16])

	i += 16

	return nil
}

func (cmd *NodeLocationReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Level.CharPresentation) & byte(0x07)

		payload = append(payload, val)
	}

	if paramLen := len(cmd.NodeLocationChar); paramLen > 16 {
		return nil, errors.New("Length overflow in array parameter NodeLocationChar")
	}

	payload = append(payload, []byte(cmd.NodeLocationChar)...)

	return
}
