package utils

import (
	"fmt"
	rand2 "math/rand"
	"testing"
)

func TestNewArgon2ID(t *testing.T) {

	argon2id := NewArgon2ID()

	plainTexts := make([]string, 100)
	for i := 0; i < 100; i++ {
		plainTexts[i] = randomString(int(rand2.Uint32()) % 255)
	}

	for _, plainText := range plainTexts {

		hash, hashErr := argon2id.Hash(plainText)

		if hashErr != nil {
			t.Errorf("Error producing hash for plaintext %s: %s", plainText, hashErr.Error())
		}

		res, verErr := argon2id.Verify(plainText, hash)
		if verErr != nil {
			t.Errorf("Error verifying plainText with hash %s %s: %s", plainText, hash, hashErr.Error())
		}

		if !res {
			t.Errorf("plainText not verified against hash %s %s", plainText, hash)
		}

	}

}

func TestRandomString(t *testing.T) {

	tests := 100
	var randomStrings = make(map[string]string)

	for i := 0; i < tests; i++ {
		rs := randomString(int(rand2.Uint32()) % 255)
		randomStrings[rs] = ""
		fmt.Println(rs)
	}

	entries := len(randomStrings)
	if entries != tests {
		t.Errorf("Tests: %d\t Entries: %d", tests, entries)
	}

}
