package example1

type Encryptable interface {
	Encrypt(key []byte) []byte
	Decryption(key []byte) []byte
}
