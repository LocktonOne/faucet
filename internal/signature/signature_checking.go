package signature

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strings"
)

func CheckSignature(hash, signatureS, addr string) (bool, error) {
	addrR, err := RecoverAddress(hash, signatureS)
	if err != nil {
		return false, errors.Wrap(err, "failed to recover address")
	}

	addr = strings.ToLower(addr)
	addrRS := strings.ToLower(addrR.String())

	if strings.Compare(addr, addrRS) == 0 {
		return false, nil
	}

	return true, nil
}

func RecoverAddress(hash, signatureS string) (common.Address, error) {
	sig, err := hexutil.Decode(signatureS)
	data := []byte(hash)

	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes long")
	}
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] -= 27 // Transform yellow paper V from 27/28 to 0/1

	rpk, err := crypto.Ecrecover(signHash(data), sig)
	if err != nil {
		return common.Address{}, err
	}
	pubKey, err := crypto.UnmarshalPubkey(rpk)
	if err != nil {
		return common.Address{}, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return recoveredAddr, nil
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
