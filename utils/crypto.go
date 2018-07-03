package utils

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pborman/uuid"
	"math/big"
	//"strconv"
	//"encoding/hex"
	"bytes"
	//"github.com/ethereum/go-ethereum/rlp"
	"encoding/hex"
	"fmt"
)

//生成keystore
func New_keystore(auth string) (address string, store string, privatekey string, err error) {

	//生成keystore
	privateKey, err := crypto.GenerateKey()
	//fmt.Println(privateKey)
	if err != nil {
		//fmt.Println("Can't load private key: %v", err)
		return "", "", "", err
	}
	id := uuid.NewRandom()
	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		PrivateKey: privateKey,
	}
	keyjson, err := keystore.EncryptKey(key, auth, keystore.StandardScryptN, keystore.StandardScryptP)

	return key.Address.String(), string(keyjson), hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)), nil
}

//解锁keystore
func Unlock_keystore(keyjson []byte, auth string) (*keystore.Key, error) {
	return keystore.DecryptKey(keyjson, auth)
}

//签名交易
//amount 单位是最小单位—>微
func Signtx(key *keystore.Key, nonce int, to string, amount *big.Int, gasLimit int, gasPrice int64, contract string, chainid int) ([]byte, error) {
	var data []byte
	//nonce = 2
	if contract != "" { //代币处理
		var err error
		temp := "0000000000000000000000000000000000000000000000000000000000000000"
		value := hex.EncodeToString(amount.Bytes())
		datastr := "a9059cbb" + temp[:24] + to[2:] + temp[0:64-len(value)] + value
		fmt.Println(datastr)
		to = contract
		amount.SetInt64(0)
		data, err = hex.DecodeString(datastr)
		if err != nil {
			return nil, err
		}
	}

	tx := types.NewTransaction(uint64(nonce), common.HexToAddress(to), amount, uint64(gasLimit), big.NewInt((gasPrice)), data)
	//fmt.Println(uint64(nonce),common.HexToAddress(to),to,amount,uint64(gasLimit),big.NewInt((gasPrice)),data)
	//fmt.Println(tx.Data(),tx.Gas(),tx.GasPrice(),tx.Value(), tx.Nonce(),tx.CheckNonce(),chainid,to)
	var sgtx *types.Transaction
	var err error
	if chainid > 0 {

		sgtx, err = types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(chainid))), key.PrivateKey)
	} else {
		sgtx, err = types.SignTx(tx, types.HomesteadSigner{}, key.PrivateKey)
	}
	if err != nil {
		return nil, err
	}
	w := new(bytes.Buffer)
	err = sgtx.EncodeRLP(w)
	if err != nil {
		return nil, err
	}
	//fmt.Println(hex.EncodeToString(w.Bytes()))
	return w.Bytes(), nil
}
