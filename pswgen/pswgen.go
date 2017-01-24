package pswgen

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

var serialSeed uint64
var lastGenDay string
var serialSeedMutex sync.Mutex

func getSerialSeed(day string) uint64 {
	serialSeedMutex.Lock()
	defer serialSeedMutex.Unlock()

	// another day , intialize serial
	if day != lastGenDay {
		serialSeed = 0
		lastGenDay = day
	} else {
		serialSeed++
	}

	return serialSeed
}

// GenSerial get a new serial id
func GenSerial() string {
	day := time.Now().Format("060102")
	serial := fmt.Sprintf("%s%04d", day, getSerialSeed(day))
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
