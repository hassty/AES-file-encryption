package main

import "testing"

type encryptionTest struct {
	key       string
	message   string
	encrypted string
}

var encryptionTests = []encryptionTest{
	{
		"cafebabedeadbeef",
		"",
		"921047074bbe2ae05437fd555d6aea0c",
	},
	{
		"cafebabedeadbeef",
		"two words",
		"5e6e1774e67ce01238e9198b71b96e6e",
	},
	{
		"cafebabedeadbeef",
		"one block length",
		"ab424cc5629a0663c13d5f1e4a2a4957921047074bbe2ae05437fd555d6aea0c",
	},
	{
		"cafebabedeadbeef",
		"this is a very long text to test multiple blocks",
		"88cbcb13c70924d443736634fd814cb667c6fe4d22e58f7cd584c4290035695c481cfd71925415a64369457a2d13ed2a921047074bbe2ae05437fd555d6aea0c",
	},
}

func TestEncrypt(t *testing.T) {
	for _, test := range encryptionTests {
		got := Encrypt([]byte(test.key), test.message)
		if got != test.encrypted {
			t.Fatalf(`failed to encrypt message %q
            expected %q
            got %q`,
				test.message, test.encrypted, got)
		}
	}
}

func TestDecrypt(t *testing.T) {
	for _, test := range encryptionTests {
		got := Decrypt([]byte(test.key), test.encrypted)
		if got != test.message {
			t.Fatalf(`failed to decrypt cipher %q
            expected %q
            got %q`,
				test.encrypted, test.message, got)
		}
	}
}
