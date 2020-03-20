package waitformore

import (
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"math/rand"
)

var seededRand = rand.New(rand.NewSource(generateSeed()))

const allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ123456789"

func generateSeed() int64 {
	// this is the simple one
	// return time.Now().UnixNano()

	// we use secure random so that two parallel running servers will never have the same seed

	b := make([]byte, 8)
	n, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	if n != 8 {
		panic(errors.New("not enough random"))
	}
	i := int64(binary.LittleEndian.Uint64(b))
	return i
}

func GenerateUserKey() string {
	return RandomStringWithCharset(4, allowedChars)
}

func GenerateAdminKey() string {
	return RandomStringWithCharset(16, allowedChars)
}

func RandomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
