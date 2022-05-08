package aes

import "testing"

type cbcTest struct {
	key       string
	iv        string
	message   string
	encrypted string
}

var cbcTests = []cbcTest{
	{
		"cafebabedeadbeef",
		"deadfa11feedc0de",
		"",
		"32197df52dc8def99c4e1d7519e42fb4",
	},
	{
		"cafebabedeadbeef",
		"deadfa11feedc0de",
		"two words",
		"1cd722924b0931d6473421ece9ec8ccb",
	},
	{
		"cafebabedeadbeef",
		"deadfa11feedc0de",
		"one block length",
		"fda1f0c669f7fe1fbf6c94c4aca48cedc68a491a8d177850f3af828a55afa5d6",
	},
	{
		"cafebabedeadbeef",
		"deadfa11feedc0de",
		"this is a very long text to test multiple blocks",
		"01d3546248c9eea31e3289f082da6b34831befe37a394ecc46e1d047eaaf38aa68f71b348b9f4b65e629949ad462eed29734c3d050a6b687d102d2509db02bef",
	},
}

func TestEncryptCBC(t *testing.T) {
	for _, test := range cbcTests {
		got, _ := EncryptCBC([]byte(test.key), []byte(test.iv), []byte(test.message))
		if got != test.encrypted {
			t.Fatalf(`failed to encrypt message %q
            expected %q
            got %q`,
				test.message, test.encrypted, got)
		}
	}
}

func TestDecryptCBC(t *testing.T) {
	for _, test := range cbcTests {
		got, _ := DecryptCBC([]byte(test.key), []byte(test.iv), []byte(test.encrypted))
		if got != test.message {
			t.Fatalf(`failed to decrypt cipher %q
            expected %q
            got %q`,
				test.encrypted, test.message, got)
		}
	}
}
