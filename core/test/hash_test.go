package test

import (
	"testing"

	"github.com/danduggan98/BrowserCoin-Go/core/src/lib"
)

// Test hash function on individual strings
func TestHash(t *testing.T) {

	// Test #1 - single word
	helloHash := lib.HashString("hello")
	helloHashString := lib.StringFromHash(helloHash)
	helloHashExpected := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"

	if (helloHashString != helloHashExpected) {
		t.Errorf("Hash of \"hello\" should be \"%v\", not \"%v\"", helloHashExpected, helloHashString)
	}

	// Test #2 - full sentence
	sentence := "This is a full sentence with lots of words and spaces and stuff."
	sentenceHash := lib.HashString(sentence)
	sentenceHashString := lib.StringFromHash(sentenceHash)
	sentenceHashExpected := "8e24a56a25945f3bab120c98103c1ca3bd378bfa9cf289f61cbba9e9b0277e80"

	if (sentenceHashString != sentenceHashExpected) {
		t.Errorf("Hash of the full sentence should be \"%v\", not \"%v\"", sentenceHashExpected, sentenceHashString)
	}
}
