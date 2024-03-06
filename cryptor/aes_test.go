package cryptor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAesEcbDecrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted, _ := AesEcbEncrypt([]byte(data), []byte(key))
	decrypted, _ := AesEcbDecrypt(encrypted, []byte(key))

	assert.Equal(t, data, string(decrypted))
}

func TestAesCbcDecrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesCbcEncrypt([]byte(data), []byte(key))
	decrypted := AesCbcDecrypt(encrypted, []byte(key))

	assert.Equal(t, data, string(decrypted))
}
