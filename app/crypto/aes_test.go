package crypto

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
    AES_128_KEY []byte
    AES_192_KEY []byte
    AES_256_KEY []byte
)
func init() {
    AES_128_KEY, _ = base64.URLEncoding.DecodeString("f1j6-Lady1lSJNqS67Qwfg==")
    AES_192_KEY, _ = base64.URLEncoding.DecodeString("Os3DsiQiWgQwn1YH30_43dV5Jnes-BaI")
    AES_256_KEY, _ = base64.URLEncoding.DecodeString("PHKy3EVrGkOnAwHfY9BvDgePyWB-jDP2i2YcsguOohQ=")
}
func Benchmark_Test_GeneateAESKey_128(b *testing.B) {
    length := 16
    key := GenerateKey(length)
    // b.Log(base64.URLEncoding.EncodeToString(key))
    assert.Equal(b, length, len(key))
}

func Benchmark_Test_GeneateAESKey_192(b *testing.B) {
    length := 24
    key := GenerateKey(length)
    // b.Log(base64.URLEncoding.EncodeToString(key))
    assert.Equal(b, length, len(key))
}

func Benchmark_Test_GeneateAESKey_256(b *testing.B) {
    length := 32
    key := GenerateKey(length)
    // b.Log(base64.URLEncoding.EncodeToString(key))
    assert.Equal(b, length, len(key))
}

func Benchmark_Test_GeneateAESKey_wrong_key_length(b *testing.B) {
    length := 129
    assert.Panics(b, func() {GenerateKey(length)}, "length must be 16, 24 or 32")
}

func Benchmark_Test_GeneateAESKey_zero_key_length(b *testing.B) {
    length := 0
    assert.Panics(b, func() {GenerateKey(length)}, "length must be 16, 24 or 32")
}

func Benchmark_Test_GeneateAESKey_negative_key_length(b *testing.B) {
    length := -1
    assert.Panics(b, func() {GenerateKey(length)}, "length must be 16, 24 or 32")
}

func Benchmark_Test_Generate_IV(b *testing.B) {
    iv := GenerateIV()
    // b.Log(base64.StdEncoding.EncodeToString(iv))
    assert.Equal(b, 16, len(iv))
}

func Benchmark_Test_ECB_128(b *testing.B) {
    data := []byte("message")
    cipherTextByte := EncryptWithECB(data, AES_128_KEY, PKCS7Padding)
    decryptedTextByte := DecryptWithECB(cipherTextByte, AES_128_KEY, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_ECB_192(b *testing.B) {
    data := []byte("message")
    cipherTextByte := EncryptWithECB(data, AES_192_KEY, PKCS7Padding)
    decryptedTextByte := DecryptWithECB(cipherTextByte, AES_192_KEY, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_ECB_256(b *testing.B) {
    data := []byte("message")
    cipherTextByte := EncryptWithECB(data, AES_256_KEY, PKCS7Padding)
    decryptedTextByte := DecryptWithECB(cipherTextByte, AES_256_KEY, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_CBC_128(b *testing.B) {
    iv, _:= base64.StdEncoding.DecodeString("ibv6AGlPdzF0dJXUCrbBYA==")
    data := []byte("message")
    cipherTextByte := EncryptWithCBC(data, AES_128_KEY, iv, PKCS7Padding)
    decryptedTextByte := DecryptWithCBC(cipherTextByte, AES_128_KEY, iv, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_CBC_192(b *testing.B) {
    iv, _:= base64.StdEncoding.DecodeString("ibv6AGlPdzF0dJXUCrbBYA==")
    data := []byte("message")
    cipherTextByte := EncryptWithCBC(data, AES_192_KEY, iv, PKCS7Padding)
    decryptedTextByte := DecryptWithCBC(cipherTextByte, AES_192_KEY, iv, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_CBC_256(b *testing.B) {
    iv, _:= base64.StdEncoding.DecodeString("ibv6AGlPdzF0dJXUCrbBYA==")
    data := []byte("message")
    cipherTextByte := EncryptWithCBC(data, AES_256_KEY, iv, PKCS7Padding)
    decryptedTextByte := DecryptWithCBC(cipherTextByte, AES_256_KEY, iv, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_CTR_128(b *testing.B) {
    iv, _:= base64.StdEncoding.DecodeString("ibv6AGlPdzF0dJXUCrbBYA==")
    data := []byte("message")
    cipherTextByte := EncryptWithCTR(data, AES_128_KEY, iv, PKCS7Padding)
    decryptedTextByte := DecryptWithCTR(cipherTextByte, AES_128_KEY, iv, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_CTR_192(b *testing.B) {
    iv, _:= base64.StdEncoding.DecodeString("ibv6AGlPdzF0dJXUCrbBYA==")
    data := []byte("message")
    cipherTextByte := EncryptWithCTR(data, AES_192_KEY, iv, PKCS7Padding)
    decryptedTextByte := DecryptWithCTR(cipherTextByte, AES_192_KEY, iv, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_CTR_256(b *testing.B) {
    iv, _:= base64.StdEncoding.DecodeString("ibv6AGlPdzF0dJXUCrbBYA==")
    data := []byte("message")
    cipherTextByte := EncryptWithCTR(data, AES_256_KEY, iv, PKCS7Padding)
    decryptedTextByte := DecryptWithCTR(cipherTextByte, AES_256_KEY, iv, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_CFB_128(b *testing.B) {
    iv, _:= base64.StdEncoding.DecodeString("ibv6AGlPdzF0dJXUCrbBYA==")
    data := []byte("message")
    cipherTextByte := EncryptWithCFB(data, AES_128_KEY, iv, PKCS7Padding)
    decryptedTextByte := DecryptWithCFB(cipherTextByte, AES_128_KEY, iv, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_CFB_192(b *testing.B) {
    iv, _:= base64.StdEncoding.DecodeString("ibv6AGlPdzF0dJXUCrbBYA==")
    data := []byte("message")
    cipherTextByte := EncryptWithCFB(data, AES_192_KEY, iv, PKCS7Padding)
    decryptedTextByte := DecryptWithCFB(cipherTextByte, AES_192_KEY, iv, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_CFB_256(b *testing.B) {
    iv, _:= base64.StdEncoding.DecodeString("ibv6AGlPdzF0dJXUCrbBYA==")
    data := []byte("message")
    cipherTextByte := EncryptWithCFB(data, AES_256_KEY, iv, PKCS7Padding)
    decryptedTextByte := DecryptWithCFB(cipherTextByte, AES_256_KEY, iv, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_OFB_128(b *testing.B) {
    iv, _:= base64.StdEncoding.DecodeString("ibv6AGlPdzF0dJXUCrbBYA==")
    data := []byte("message")
    cipherTextByte := EncryptWithOFB(data, AES_128_KEY, iv, PKCS7Padding)
    decryptedTextByte := DecryptWithOFB(cipherTextByte, AES_128_KEY, iv, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_OFB_192(b *testing.B) {
    iv, _:= base64.StdEncoding.DecodeString("ibv6AGlPdzF0dJXUCrbBYA==")
    data := []byte("message")
    cipherTextByte := EncryptWithOFB(data, AES_192_KEY, iv, PKCS7Padding)
    decryptedTextByte := DecryptWithOFB(cipherTextByte, AES_192_KEY, iv, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}

func Benchmark_Test_OFB_256(b *testing.B) {
    iv, _:= base64.StdEncoding.DecodeString("ibv6AGlPdzF0dJXUCrbBYA==")
    data := []byte("message")
    cipherTextByte := EncryptWithOFB(data, AES_256_KEY, iv, PKCS7Padding)
    decryptedTextByte := DecryptWithOFB(cipherTextByte, AES_256_KEY, iv, PKCS7UnPadding)
    // cipherText := base64.URLEncoding.EncodeToString(cipherTextByte)
    // b.Log(cipherText, len(cipherText), string(decryptedTextByte))
    assert.Equal(b, string(data), string(decryptedTextByte))
}


