package service

import (
	"crypto/sha256"
	"encoding/binary"
	"sync"
	"time"
)

type LockManager struct {
	shards []*sync.Mutex
	count uint32
}

func NewLockManager(count uint32) *LockManager{
	locks := make([]*sync.Mutex, count)
	for i := uint32(0); i < count; i++{
		locks[i] = &sync.Mutex{}
	}
	return &LockManager{
		shards: locks,
		count: count,
	}
}


func(lm *LockManager) GetLock(trancheVaultAddress string , date time.Time) *sync.Mutex {
	// take the first 5 elements of the address string -> convert them to a number
	address := trancheVaultAddress[:8]
	date_data := date.Format("2006-01-02")

	//combine this data to get a unique key
	data := []byte(address + date_data)
	//create a hash
	hash := sha256.Sum256(data)

	//get the first 4 digits from it
	val := binary.BigEndian.Uint32(hash[:4])

	//use it to distribute the keys 
	return lm.shards[val % lm.count]

}
