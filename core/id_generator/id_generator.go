package id_generator

import (
	"log"
	"math"
	"sync"
	"time"
)

var (
	serverID int64
	sequence int64
	unixTime int64
	mux      sync.Mutex
)

const (
	serverBitCount   = 8
	sequenceBitCount = 24
	serverShift      = sequenceBitCount
	timeShift        = sequenceBitCount + serverBitCount
	sequenceMax      = (1 << sequenceBitCount) - 1
)

func Init(initServerID int) {
	if math.MaxInt8 < initServerID {
		log.Fatal("max server id is ", math.MaxInt8)
	}

	serverID = int64(initServerID)
	unixTime = time.Now().Unix()
}

func Generate() int64 {
	mux.Lock()
	defer mux.Unlock()

	curTime := time.Now().Unix()
	if unixTime != curTime {
		sequence = 0
		unixTime = curTime
	}

	id := unixTime<<timeShift |
		serverID<<serverShift |
		sequence

	sequence++

	if sequenceMax < sequence {
		time.Sleep(time.Second)
		return Generate()
	}

	return id
}
