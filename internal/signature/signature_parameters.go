package signature

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Parameters struct {
	R string `json:"r"`
	S string `json:"s"`
	V int    `json:"v"`
}

func (p Parameters) ToGeth() ([32]byte, [32]byte, uint8, error) {
	var (
		r [32]byte
		s [32]byte
		v uint8
	)

	rS, err := hexutil.Decode(p.R)
	if err != nil {
		return r, s, v, err
	}

	sS, err := hexutil.Decode(p.S)
	if err != nil {
		return r, s, v, err
	}

	v = uint8(p.V)
	copy(r[:], rS[:])
	copy(s[:], sS[:])
	return r, s, v, nil
}

func ParseSignatureParameters(signature []byte) (*Parameters, error) {
	if len(signature) != 65 {
		return nil, errors.New("bad signature")
	}
	params := Parameters{}

	params.R = hexutil.Encode(signature[:32])
	params.S = hexutil.Encode(signature[32:64])
	params.V = 27 + int(signature[64])

	return &params, nil
}
