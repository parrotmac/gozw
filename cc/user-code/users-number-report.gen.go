// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package usercode

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandUsersNumberReport cc.CommandID = 0x05

func init() {
	gob.Register(UsersNumberReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x63),
		Command:      cc.CommandID(0x05),
		Version:      1,
	}, NewUsersNumberReport)
}

func NewUsersNumberReport() cc.Command {
	return &UsersNumberReport{}
}

// <no value>
type UsersNumberReport struct {
	SupportedUsers byte
}

func (cmd UsersNumberReport) CommandClassID() cc.CommandClassID {
	return 0x63
}

func (cmd UsersNumberReport) CommandID() cc.CommandID {
	return CommandUsersNumberReport
}

func (cmd UsersNumberReport) CommandIDString() string {
	return "USERS_NUMBER_REPORT"
}

func (cmd *UsersNumberReport) UnmarshalBinary(data []byte) error {
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

	cmd.SupportedUsers = payload[i]
	i++

	return nil
}

func (cmd *UsersNumberReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SupportedUsers)

	return
}
