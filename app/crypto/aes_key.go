package crypto

import "crypto/rand"

/**
 * Generate Key with length
 * @param length 128(16 bytes), 192(24 bytes), 256(32 bytes)
 * @return key
 */
 func GenerateKey(length int) []byte {
	if !(length == 16 || length == 24 || length == 32) {
		panic("length must be 16, 24 or 32")
	}
	key := make([]byte, length)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	return key
}