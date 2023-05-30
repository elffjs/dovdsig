package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var realAddr = common.HexToAddress("0x4E0a9405D6818D84A1675FAaBDf10D310a23a1f5")
var sig = common.FromHex("0xb9d496f691ff689f4eeb3a0793acd30fa76f57f07271056b05ebbc12c78b54f30c849cbdc1e70db53affd599bccaf181758ca882b53913a4f0b22cc7bc7b30651b")

func main() {
	msg := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"MintVirtualDeviceSign": []apitypes.Type{
				{Name: "integrationNode", Type: "uint256"},
				{Name: "vehicleNode", Type: "uint256"},
			},
		},
		PrimaryType: "MintVirtualDeviceSign",
		Domain: apitypes.TypedDataDomain{
			Name:              "DIMO",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(80001),
			VerifyingContract: "0x4De1bCf2B7E851E31216fC07989caA902A604784",
		},
		Message: apitypes.TypedDataMessage{
			"integrationNode": big.NewInt(1),
			"vehicleNode":     big.NewInt(57),
		},
	}

	hash, _, err := apitypes.TypedDataAndHash(msg)
	if err != nil {
		panic(err)
	}

	sig[64] -= 27

	pub, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		panic(err)
	}

	pubRaw, err := crypto.UnmarshalPubkey(pub)
	if err != nil {
		panic(err)
	}

	addr := crypto.PubkeyToAddress(*pubRaw)
	fmt.Println("Outcome:", addr, addr == realAddr)
}
