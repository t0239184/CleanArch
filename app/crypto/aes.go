package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// /**
//  * Generate Key with length
//  * @param length 128(16 bytes), 192(24 bytes), 256(32 bytes)
//  * @return key
//  */
// func GenerateKey(length int) []byte {
// 	if !(length == 16 || length == 24 || length == 32) {
// 		panic("length must be 16, 24 or 32")
// 	}
// 	key := make([]byte, length)
// 	if _, err := rand.Read(key); err != nil {
// 		panic(err)
// 	}
// 	return key
// }

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS7UnPadding(origData []byte) []byte {
	unPadding := int(origData[len(origData)-1])
	return origData[:(len(origData) - unPadding)]
}

func PKCS7Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	var paddingText []byte
	if padding == 0 {
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}
	return append(text, paddingText...)
}

/**
 * Generate Initialization Vector (IV)
 * return iv []byte length:16
 */
func GenerateIV() []byte {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	return iv
}

/**
 * Encrypt with Electronic Codebook Book (ECB)
 * @param plainText
 * @param key
 * @param paddingFunc
 * @return cipherText string length:24
 */
func EncryptWithECB(plainText, key []byte, paddingFunc func([]byte, int) []byte) []byte {
	block, _ := aes.NewCipher(key)
	plainText = paddingFunc(plainText, block.BlockSize())
	decrypted := make([]byte, len(plainText))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(plainText); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], plainText[bs:be])
	}
	return decrypted
}

/**
 * Decrypt with Electronic Codebook Book (ECB)
 * @param cipherText
 * @param key
 * @param unPaddingFunc
 * @return plainText
 */
func DecryptWithECB(cipherText, key []byte, unPaddingFunc func([]byte) []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(cipherText))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(cipherText); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], cipherText[bs:be])
	}

	return unPaddingFunc(decrypted)
}

/**
 * Encrypt with Cipher Block Chaining (CBC)
 * @param plainText
 * @param key
 * @param iv
 * @param paddingFunc
 * @return cipherText string length:24
 */
func EncryptWithCBC(plainText, key, iv []byte, paddingFunc func([]byte, int) []byte) []byte {
	block, _ := aes.NewCipher(key)
	padText := paddingFunc(plainText, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	result := make([]byte, len(padText))
	blockMode.CryptBlocks(result, padText)
	return result
}

/**
 * Decrypt with Cipher Block Chaining (CBC)
 * @param cipherText
 * @param key
 * @param iv
 * @param unpaddingFunc
 * @return plainText
 */
func DecryptWithCBC(cipherText, key, iv []byte, unPaddingFunc func([]byte) []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, iv)
	result := make([]byte, len(cipherText))
	blockMode.CryptBlocks(result, cipherText)
	result = unPaddingFunc(result)
	return result
}

/**
 * Encrypt with Counter (CTR)
 * @param plainText
 * @param key
 * @param iv
 * @return cipherText
 */
func EncryptWithCTR(plainText, key, iv []byte, paddingFunc func([]byte, int) []byte) []byte {
	padText := paddingFunc(plainText, aes.BlockSize)
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCTR(block, iv)
	message := make([]byte, len(padText))
	blockMode.XORKeyStream(message, padText)
	return message
}

/**
 * Decrypt with Cipher Feedback (CFB)
 * @param plainText
 * @param key
 * @param iv
 * @return plainText
 */
func DecryptWithCTR(cipherText, key, iv []byte, unPaddingFunc func([]byte) []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCTR(block, iv)
	message := make([]byte, len(cipherText))
	blockMode.XORKeyStream(message, cipherText)
	return unPaddingFunc(message)
}

/*Cipher FeedBack (CFB)*/
func EncryptWithCFB(plainText, key, iv []byte, paddingFunc func([]byte, int) []byte) []byte {
	block, _ := aes.NewCipher(key)
	padText := paddingFunc(plainText, block.BlockSize())
	blockMode := cipher.NewCFBEncrypter(block, iv)
	message := make([]byte, len(padText))
	blockMode.XORKeyStream(message, padText)
	return message
}

func DecryptWithCFB(cipherText, key, iv []byte, unPaddingFunc func([]byte) []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCFBDecrypter(block, iv)
	message := make([]byte, len(cipherText))
	blockMode.XORKeyStream(message, cipherText)
	return unPaddingFunc(message)
}

/*Output FeedBack (OFB)*/
func EncryptWithOFB(plainText, key []byte, iv []byte, paddingFunc func([]byte, int) []byte) []byte {
	block, _ := aes.NewCipher(key)
	padText := paddingFunc(plainText, block.BlockSize())
	blockMode := cipher.NewOFB(block, iv)
	message := make([]byte, len(padText))
	blockMode.XORKeyStream(message, padText)
	return message
}

func DecryptWithOFB(cipherText, key, iv []byte, unPaddingFunc func([]byte) []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewOFB(block, iv)
	message := make([]byte, len(cipherText))
	blockMode.XORKeyStream(message, cipherText)
	return unPaddingFunc(message)
}

//CCM：Counter with CBC-MAC

//GCM：Galois/Counter Mode
