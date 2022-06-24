package crypto

type RSA interface {
    GenerateRSAKeyPair(bits int) (publicKey, privateKey []byte, err error)
    /*.PEM*/
    LoadKeyFromPem(pem []byte) ([]byte, error)
    ConvertKeyToPem(key []byte) ([]byte, error)
    /*.P12*/
    ConvertCertAndPriKeyToP12(cert, priKey []byte, password string) ([]byte, error)
    ExtractCertAndPriKeyFromP12(p12 []byte, password string) ([]byte, []byte, error)

    /*Cipher*/
    EncryptWithRSA(publicKey, plainText []byte) (cipherText []byte, err error)
    DecryptWithRSA(privateKey, cipherText []byte) (plainText []byte, err error)
}