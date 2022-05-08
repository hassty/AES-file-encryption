package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
)

const BlockSize int = 16

type Block = []byte

func EncryptCBC(key, iv, msg []byte) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blocks := padPKCS7(splitIntoBlocks(msg), msg)

	var encrypted []byte
	if len(iv) != BlockSize {
		return "", fmt.Errorf("initialization vector must be %d bytes long", BlockSize)
	}
	cbc := cipher.NewCBCEncrypter(c, iv)
	for _, block := range blocks {
		encryptedBlock := make(Block, BlockSize)
		cbc.CryptBlocks(encryptedBlock, block)
		encrypted = append(encrypted, encryptedBlock...)
	}

	return hex.EncodeToString(encrypted), nil
}

func DecryptCBC(key, iv, ciphertext []byte) (string, error) {
	decoded := make([]byte, len(ciphertext))
	bytes, err := hex.Decode(decoded, ciphertext)
	decoded = decoded[:bytes]
	if err != nil {
		if _, ok := err.(hex.InvalidByteError); ok {
			return "", errors.New("cipher contains non-hex value")
		}
		return "", err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(iv) != BlockSize {
		return "", fmt.Errorf("initialization vector must be %d bytes long", BlockSize)
	}
	var decrypted []byte
	cbc := cipher.NewCBCDecrypter(c, iv)
	blocks := splitIntoBlocks(decoded)
	for _, block := range blocks {
		decryptedBlock := make(Block, BlockSize)
		cbc.CryptBlocks(decryptedBlock, block)
		decrypted = append(decrypted, decryptedBlock...)
	}

	padding := int(decrypted[len(decrypted)-1])
	s := string(decrypted[:len(decrypted)-padding])

	return s, nil

}

func splitIntoBlocks(s []byte) []Block {
	size := len(s) / BlockSize
	blocks := make([]Block, size)

	for i := 0; i < size; i++ {
		blocks[i] = Block(s[i*BlockSize : (i+1)*BlockSize])
	}

	return blocks
}

func padPKCS7(b []Block, s []byte) []Block {
	remainingSize := len(s) % BlockSize
	if remainingSize != 0 {
		lastBlock := Block(s[len(s)-remainingSize:])
		padding := BlockSize - len(lastBlock)
		for i := 0; i < padding; i++ {
			lastBlock = append(lastBlock, byte(padding))
		}
		b = append(b, lastBlock)
	}

	if len(s) == 0 || remainingSize == 0 {
		paddingBlock := make(Block, BlockSize)
		for i := 0; i < BlockSize; i++ {
			paddingBlock[i] = byte(BlockSize)
		}
		b = append(b, paddingBlock)
	}

	return b
}
