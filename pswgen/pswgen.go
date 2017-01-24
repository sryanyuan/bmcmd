package pswgen

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

var serialSeed uint64
var lastGenTime int64
var serialSeedMutex sync.Mutex

func getSerialSeed(tm int64) uint64 {
	serialSeedMutex.Lock()
	defer serialSeedMutex.Unlock()

	if 0 == lastGenTime {
		// first time to generate serial
		serialSeed = 0
	} else {
		if tm-lastGenTime != 0 {
			// intialize serial
			serialSeed = 0
		} else {
			serialSeed++
		}
	}
	lastGenTime = tm

	return serialSeed
}

// GenSerial get a new serial id
func GenSerial() string {
	tm := time.Now().Unix()
	serial := fmt.Sprintf("%d%07d", tm, getSerialSeed(tm))
	return serial
}

// GenPassword generate a 8 bits password for serial
func GenPassword(serial, key string) string {
	md5Source := fmt.Sprintf("%s_+@%s", serial, key)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(md5Source))
	md5Sum := md5Ctx.Sum(nil)
	md5Result := hex.EncodeToString(md5Sum)

	// get 5-12 fields
	return md5Result[4:12]
}

// VerifyPassword verify if a password for serial is valid
func VerifyPassword(serial, key, password string) bool {
	return GenPassword(serial, key) == password
}
