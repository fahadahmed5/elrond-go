package pubkeyConverter

import (
	"fmt"

	"github.com/ElrondNetwork/elrond-go/data/state"
	"github.com/btcsuite/btcutil/bech32"
)

type config struct {
	prefix   string
	fromBits byte
	toBits   byte
	pad      bool
}

var bech32Config = config{
	prefix:   "erd",
	fromBits: byte(8),
	toBits:   byte(5),
	pad:      true,
}

// bech32PubkeyConverter encodes or decodes provided public key as/from bech32 format
type bech32PubkeyConverter struct {
	len int
}

// NewBech32PubkeyConverter returns a bech32PubkeyConverter instance
func NewBech32PubkeyConverter(addressLen int) (*bech32PubkeyConverter, error) {
	if addressLen < 1 {
		return nil, fmt.Errorf("%w when creating hex address converter, addressLen should have been greater than 0",
			state.ErrInvalidAddressLength)
	}
	if addressLen%2 == 1 {
		return nil, fmt.Errorf("%w when creating hex address converter, addressLen should have been an even number",
			state.ErrInvalidAddressLength)
	}

	return &bech32PubkeyConverter{
		len: addressLen,
	}, nil
}

// Len returns the decoded address length
func (bpc *bech32PubkeyConverter) Len() int {
	return bpc.len
}

// Bytes converts the provided public key string as bech32 decoded bytes
func (bpc *bech32PubkeyConverter) Bytes(humanReadable string) ([]byte, error) {
	decodedPrefix, buff, err := bech32.Decode(humanReadable)
	if err != nil {
		return nil, err
	}
	if decodedPrefix != bech32Config.prefix {
		return nil, state.ErrInvalidErdAddress
	}

	// warning: mind the order of the parameters, those should be inverted
	decodedBytes, err := bech32.ConvertBits(buff, bech32Config.toBits, bech32Config.fromBits, !bech32Config.pad)
	if err != nil {
		return nil, state.ErrBech32ConvertError
	}

	if len(decodedBytes) != bpc.len {
		return nil, fmt.Errorf("%w when converting to address, expected length %d, received %d",
			state.ErrWrongSize, bpc.len, len(decodedBytes))
	}

	return decodedBytes, nil
}

// String converts the provided bytes in a bech32 form
func (bpc *bech32PubkeyConverter) String(pkBytes []byte) (string, error) {
	conv, err := bech32.ConvertBits(pkBytes, 8, 5, bech32Config.pad)
	if err != nil {
		return "", err
	}

	return bech32.Encode(bech32Config.prefix, conv)
}

// CreateAddressFromString creates an address container based on the provided string
func (bpc *bech32PubkeyConverter) CreateAddressFromString(humanReadable string) (state.AddressContainer, error) {
	buff, err := bpc.Bytes(humanReadable)
	if err != nil {
		return nil, err
	}

	return state.NewAddress(buff), nil
}

// CreateAddressFromBytes creates an address container based on the provided public key bytes
func (bpc *bech32PubkeyConverter) CreateAddressFromBytes(pkBytes []byte) (state.AddressContainer, error) {
	if len(pkBytes) != bpc.len {
		return nil, fmt.Errorf("%w when converting to address, expected length %d, received %d",
			state.ErrWrongSize, bpc.len, len(pkBytes))
	}

	return state.NewAddress(pkBytes), nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (bpc *bech32PubkeyConverter) IsInterfaceNil() bool {
	return bpc == nil
}