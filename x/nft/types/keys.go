package types

import (
	"bytes"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "nft"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_nft"

	// this line is used by starport scaffolding # ibc/keys/name
)

var (
	PrefixNFT                     = []byte{0x01}
	PrefixOwners                  = []byte{0x02} // key for a owner
	PrefixCollection              = []byte{0x03} // key for balance of NFTs held by the denom
	PrefixDenom                   = []byte{0x04} // key for denom of the nft
	PrefixDenomName               = []byte{0x05} // key for denom name of the nft
	PrefixApprovedAddresses       = []byte{0x06} // key for globally approved operator addresses
	PrefixCollectionTotalNftCount = []byte{0x07} // key for total nft count
	PrefixDenomSymbol             = []byte{0x08} // key for denom name of the nft
	PrefixNFTLockOwner            = []byte{0x09} // key for owner of nft lock
	delimiter                     = []byte("/")
)

// SplitKeyOwner return the address,denom,id from the key of stored owner
func SplitKeyOwner(key []byte) (address sdk.AccAddress, denomID, tokenID string, err error) {
	key = key[len(PrefixOwners)+len(delimiter):]
	keys := bytes.Split(key, delimiter)
	if len(keys) != 3 {
		return address, denomID, tokenID, errors.New("wrong KeyOwner")
	}

	address, err = sdk.AccAddressFromBech32(string(keys[0]))

	if err != nil {
		return address, denomID, tokenID, err
	}

	denomID = string(keys[1])
	tokenID = string(keys[2])

	return
}

func SplitKeyDenom(key []byte) (denomID, tokenID string, err error) {
	keys := bytes.Split(key, delimiter)
	if len(keys) != 2 {
		return denomID, tokenID, errors.New("wrong KeyOwner")
	}

	denomID = string(keys[0])
	tokenID = string(keys[1])
	return
}

// KeyOwner gets the key of a collection owned by an account address
func KeyOwner(address sdk.AccAddress, denomID, tokenID string) []byte {
	key := append(PrefixOwners, delimiter...)
	if address != nil {
		key = append(key, []byte(address.String())...)
		key = append(key, delimiter...)
	}

	if address != nil && len(denomID) > 0 {
		key = append(key, []byte(denomID)...)
		key = append(key, delimiter...)
	}

	if address != nil && len(denomID) > 0 && len(tokenID) > 0 {
		key = append(key, []byte(tokenID)...)
	}
	return key
}

// KeyNFT gets the key of nft stored by denom and id
func KeyNFT(denomID, tokenID string) []byte {
	key := append(PrefixNFT, delimiter...)
	if len(denomID) > 0 {
		key = append(key, []byte(denomID)...)
		key = append(key, delimiter...)
	}

	if len(denomID) > 0 && len(tokenID) > 0 {
		key = append(key, []byte(tokenID)...)
	}
	return key
}

// KeyCollection gets the storeKey by the collection
func KeyCollection(denomID string) []byte {
	key := append(PrefixCollection, delimiter...)
	return append(key, []byte(denomID)...)
}

// KeyDenomID gets the storeKey by the denom id
func KeyDenomID(id string) []byte {
	key := append(PrefixDenom, delimiter...)
	return append(key, []byte(id)...)
}

// KeyDenomName gets the storeKey by the denom name
func KeyDenomName(name string) []byte {
	key := append(PrefixDenomName, delimiter...)
	return append(key, []byte(name)...)
}

// KeyDenomSymbol gets the storeKey by the symbol
func KeyDenomSymbol(name string) []byte {
	key := append(PrefixDenomSymbol, delimiter...)
	return append(key, []byte(name)...)
}

// KeyApprovedAddresses  gets the key of an approved address
func KeyApprovedAddresses(address string) []byte {
	key := append(PrefixApprovedAddresses, delimiter...)
	return append(key, []byte(address)...)
}

// KeyCollectionTotalNfts gets the storeKey by the collection for all the minted nfts for it
func KeyCollectionTotalNfts(denomID string) []byte {
	key := append(PrefixCollectionTotalNftCount, delimiter...)
	return append(key, []byte(denomID)...)
}

func KeyNFTLockOwner(denomID, tokenID string) []byte {
	key := append(PrefixNFTLockOwner, delimiter...)
	key = append(key, []byte(denomID)...)
	key = append(key, delimiter...)
	return append(key, []byte(tokenID)...)
}
