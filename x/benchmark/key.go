package benchmark

import (
	"crypto/ecdsa"
	"encoding/binary"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
)

var rand Reader

type Reader struct {
	data []byte
	num  uint64
}

func (r *Reader) Init() {
	r.add()
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if len(p) > 1 {
		r.add()
	}
	for i := 0; i < len(p); i++ {
		p[i] = r.data[i]
	}
	return len(p), nil
}

func (r *Reader) add() {
	data := make([]byte, 32, 32)
	binary.BigEndian.PutUint64(data, r.num)
	r.data = data
	r.num++
}

func GenKeys(n int) ([]*ecdsa.PrivateKey, error) {
	r := &Reader{}
	r.Init()
	var keys []*ecdsa.PrivateKey
	for i := 0; i < n; i++ {
		sk, err := ecdsa.GenerateKey(crypto.S256(), r)
		if err != nil {
			return nil, err
		}
		keys = append(keys, sk)
	}
	return keys, nil
}
