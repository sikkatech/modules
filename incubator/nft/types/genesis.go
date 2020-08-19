package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// GenesisState is the state that must be provided at genesis.
type GenesisState struct {
	Owners      []Owner     `json:"owners"`
	Collections Collections `json:"collections"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(owners []Owner, collections Collections) GenesisState {
	return GenesisState{
		Owners:      owners,
		Collections: collections,
	}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState([]Owner{}, NewCollections())
}

// ValidateGenesis performs basic validation of nfts genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	for _, Owner := range data.Owners {
		if Owner.Address.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address cannot be empty")
		}
	}
	return nil
}
