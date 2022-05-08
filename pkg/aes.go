package aes

import (
	"crypto/aes"
	"encoding/hex"
	"errors"
)

const BlockSize int = 16

type Block = []byte

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

func Encrypt(key []byte, msg []byte) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	var encrypted []byte

	blocks := padPKCS7(splitIntoBlocks(msg), msg)

	for _, block := range blocks {
		encryptedBlock := make(Block, BlockSize)
		c.Encrypt(encryptedBlock, block)
		encrypted = append(encrypted, encryptedBlock...)
	}

	return hex.EncodeToString(encrypted), nil
}

func Decrypt(key []byte, cipher []byte) (string, error) {
	decoded := make([]byte, len(cipher))
	bytes, err := hex.Decode(decoded, cipher)
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

	var decrypted []byte

	blocks := splitIntoBlocks(decoded)
	for _, block := range blocks {
		decryptedBlock := make(Block, BlockSize)
		c.Decrypt(decryptedBlock, block)
		decrypted = append(decrypted, decryptedBlock...)
	}

	padding := int(decrypted[len(decrypted)-1])
	s := string(decrypted[:len(decrypted)-padding])

	return s, nil
}
