package main

import (
	"crypto/aes"
	"encoding/hex"
)

const BlockSize int = 16

type Block = []byte

func splitIntoBlocks(s string) []Block {
	size := len(s) / BlockSize
	blocks := make([]Block, size)

	for i := 0; i < size; i++ {
		blocks[i] = Block(s[i*BlockSize : (i+1)*BlockSize])
	}

	return blocks
}

func padPKCS7(b []Block, s string) []Block {
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

func Encrypt(key []byte, msg string) string {
	c, err := aes.NewCipher(key)
	CheckError(err)

	var encrypted []byte

	blocks := padPKCS7(splitIntoBlocks(msg), msg)

	for _, block := range blocks {
		encryptedBlock := make(Block, BlockSize)
		c.Encrypt(encryptedBlock, block)
		encrypted = append(encrypted, encryptedBlock...)
	}

	return hex.EncodeToString(encrypted)
}

func Decrypt(key []byte, cipher string) string {
	ciphertext, err := hex.DecodeString(cipher)
	CheckError(err)

	c, err := aes.NewCipher(key)
	CheckError(err)

	var decrypted []byte

	blocks := splitIntoBlocks(string(ciphertext))
	for _, block := range blocks {
		decryptedBlock := make(Block, BlockSize)
		c.Decrypt(decryptedBlock, block)
		decrypted = append(decrypted, decryptedBlock...)
	}

	padding := int(decrypted[len(decrypted)-1])
	s := string(decrypted[:len(decrypted)-padding])

	return s
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
