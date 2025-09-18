package benchmark

import (
	"encoding/binary"
	"errors"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/dgraph-io/ristretto/v2/z"
	"os"
)

type MmapTxFile struct {
	file     *os.File
	mmapFile *z.MmapFile
	offset   int
}

func NewMmapTxFile(fileName string, create bool) (*MmapTxFile, error) {
	var file *os.File
	var err error
	if len(fileName) == 0 {
		file, err = os.CreateTemp("", "benchmarktxs")
	} else {
		flag := os.O_RDWR
		if create {
			flag = flag | os.O_CREATE | os.O_TRUNC
		} else {
			flag = flag | os.O_APPEND
		}
		file, err = os.OpenFile(fileName, flag, 0666)
	}
	if err != nil {
		return nil, err
	}
	file.Name()
	return &MmapTxFile{
		file: file,
	}, nil
}
func (m *MmapTxFile) Truncate() error {
	if m.mmapFile != nil {
		return errors.New("had mmap")
	}
	return m.file.Truncate(0)
}
func (m *MmapTxFile) Mmap() error {
	if m.mmapFile != nil {
		return errors.New("had mmap")
	}
	var err error
	stat, err := m.file.Stat()
	if err != nil {
		return err
	}
	m.mmapFile, err = z.OpenMmapFile(m.file.Name(), os.O_RDONLY, int(stat.Size()))
	if err != nil {
		return err
	}
	return nil
}

func (m *MmapTxFile) UnMmap() error {
	var err error
	if m.mmapFile != nil {
		err = m.mmapFile.Close(-1)
		m.mmapFile = nil
	}
	return err
}

type Batch struct {
	Addr common.Address
	Txs  []*types.Transaction
}

func (m *MmapTxFile) WriteTxs(addr common.Address, txs []*types.Transaction) error {
	buf, err := rlp.EncodeToBytes(&Batch{
		Addr: addr,
		Txs:  txs,
	})
	if err != nil {
		return err
	}
	var length [4]byte
	binary.BigEndian.PutUint32(length[:], uint32(len(buf)))
	_, err = m.file.Write(length[:])
	if err != nil {
		return err
	}
	_, err = m.file.Write(buf)
	return err
}

func (m *MmapTxFile) ReadTxs() (common.Address, []*types.Transaction, error) {
	buf, err := m.mmapFile.Bytes(m.offset, 4)
	if err != nil {
		return common.Address{}, nil, err
	}

	length := binary.BigEndian.Uint32(buf)
	buf, err = m.mmapFile.Bytes(m.offset+4, int(length))
	if err != nil {
		return common.Address{}, nil, err
	}

	var batch Batch
	if err := rlp.DecodeBytes(buf, &batch); err != nil {
		return common.Address{}, nil, err
	}
	m.offset += 4 + int(length)
	return batch.Addr, batch.Txs, nil
}
