package rhdiff

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestSignature(t *testing.T) {

	var buf bytes.Buffer
	reader := strings.NewReader("be happy")
	sig, err := Signature(reader, &buf, uint32(16), uint32(16))

	if err != nil {
		t.Fatalf("Signature creation failed, %s", err)
	}

	// Load signature, compare to original
	sig_cpy, err := LoadSigFile(&buf, uint32(16))
	if err != nil {
		t.Fatalf("Loading signature file failed, %s", err)
	}

	if sig_cpy.chunkLen != sig.chunkLen {
		t.Errorf("Chunks missmatch, expected %d got %d", sig.chunkLen, sig_cpy.chunkLen)
	}
	if !reflect.DeepEqual(sig_cpy.hash, sig.hash) {
		t.Errorf("Hash missmatch, expected %v got %v", sig.hash, sig_cpy.hash)
	}
	if !reflect.DeepEqual(sig_cpy.strongHash, sig.strongHash) {
		t.Errorf("Strong Hash missmatch, expected %v, got %v", sig.strongHash, sig_cpy.strongHash)
	}
}
