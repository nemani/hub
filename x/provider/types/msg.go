package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hub "github.com/sentinel-official/hub/v0.5/types"
)

var (
	_ sdk.Msg = (*MsgRegister)(nil)
	_ sdk.Msg = (*MsgUpdate)(nil)
)

// MsgRegister is for registering a provider.
type MsgRegister struct {
	From        sdk.AccAddress `json:"from"`
	Name        string         `json:"name"`
	Identity    string         `json:"identity,omitempty"`
	Website     string         `json:"website,omitempty"`
	Description string         `json:"description,omitempty"`
}

func NewMsgRegister(from sdk.AccAddress, name, identity, website, description string) MsgRegister {
	return MsgRegister{
		From:        from,
		Name:        name,
		Identity:    identity,
		Website:     website,
		Description: description,
	}
}

func (m MsgRegister) Route() string {
	return RouterKey
}

func (m MsgRegister) Type() string {
	return fmt.Sprintf("%s:register", ModuleName)
}

func (m MsgRegister) ValidateBasic() error {
	if m.From == nil || m.From.Empty() {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Name length should be between 1 and 64
	if len(m.Name) == 0 || len(m.Name) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "name")
	}

	// Identity length should be between 0 and 64
	if len(m.Identity) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "identity")
	}

	// Website length should be between 0 and 64
	if len(m.Website) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "website")
	}

	// Description length should be between 0 and 256
	if len(m.Description) > 256 {
		return errors.Wrapf(ErrorInvalidField, "%s", "description")
	}

	return nil
}

func (m MsgRegister) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRegister) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// MsgUpdate is for updating a provider.
type MsgUpdate struct {
	From        hub.ProvAddress `json:"from"`
	Name        string          `json:"name,omitempty"`
	Identity    string          `json:"identity,omitempty"`
	Website     string          `json:"website,omitempty"`
	Description string          `json:"description,omitempty"`
}

func NewMsgUpdate(from hub.ProvAddress, name, identity, website, description string) MsgUpdate {
	return MsgUpdate{
		From:        from,
		Name:        name,
		Identity:    identity,
		Website:     website,
		Description: description,
	}
}

func (m MsgUpdate) Route() string {
	return RouterKey
}

func (m MsgUpdate) Type() string {
	return fmt.Sprintf("%s:update", ModuleName)
}

func (m MsgUpdate) ValidateBasic() error {
	if m.From == nil || m.From.Empty() {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Name length should be between 0 and 64
	if len(m.Name) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "name")
	}

	// Identity length should be between 0 and 64
	if len(m.Identity) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "identity")
	}

	// Website length should be between 0 and 64
	if len(m.Website) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "website")
	}

	// Description length should be between 0 and 256
	if len(m.Description) > 256 {
		return errors.Wrapf(ErrorInvalidField, "%s", "description")
	}

	return nil
}

func (m MsgUpdate) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpdate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}
