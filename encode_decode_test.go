package wordenc

import (
	"encoding/hex"
	"math/rand"
	"testing"
)

func TestVector(t *testing.T) {
	entropy, _ := hex.DecodeString("066dca1a2bb7e8a1db2832148ce9933eea0f3ac9548d793112d9a95c9407efad")
	expected := "all hour make first leader extend hole alien behind guard gospel lava path output census museum junior mass reopen famous sing advance salt parade"
	s := EncodeToString(entropy)
	if s != expected {
		t.Fatalf("want %q, got %q", expected, s)
	}
}

func roundtrip(t *testing.T, data []byte) {
	encoded := EncodeToString(data)
	decoded, err := DecodeString(encoded, len(data))
	if err != nil {
		t.Errorf("decoding %v gives error %v", data, err)
		return
	}
	incorrect := false
	if len(data) != len(decoded) {
		incorrect = true
	} else {
		for i := 0; i < len(data); i++ {
			if data[i] != decoded[i] {
				incorrect = true
			}
		}
	}
	if incorrect {
		t.Errorf("encoded/decoded %v results in %v (%s)",
			data,
			decoded,
			encoded)
	}
}

func TestEncodeDecodeEmpty(t *testing.T) {
	roundtrip(t, []byte{})
}

func TestEncodeDecodeSingleByte(t *testing.T) {
	roundtrip(t, []byte{3})
	roundtrip(t, []byte{255})
	roundtrip(t, []byte{0})
}

func TestEncodeDecodeTwoBytes(t *testing.T) {
	roundtrip(t, []byte{2, 3})
	roundtrip(t, []byte{2, 4})
	roundtrip(t, []byte{128, 197})
}

func TestEncodeDecodeMultipleBytes(t *testing.T) {
	roundtrip(t, []byte{2, 3, 0})
	roundtrip(t, []byte{0, 2, 3})
	roundtrip(t, []byte{123, 104, 12, 128})
	roundtrip(t, []byte{123, 104, 12, 86})
	roundtrip(t, []byte{123, 4, 104, 12, 86})
	roundtrip(t, []byte{32, 107, 65, 12, 204, 198})
	roundtrip(t, []byte{123, 104, 12, 86, 100, 0})
	roundtrip(t, []byte{123, 104, 12, 255, 86, 100, 0})
}

func BenchmarkEncoding32Bytes(b *testing.B) {
	data := make([]byte, 32)
	for i := range data {
		data[i] = byte(rand.Int() & (1<<8 - 1))
	}
	for i := 0; i < b.N; i++ {
		EncodeToString(data)
	}
}

func BenchmarkDecoding32Bytes(b *testing.B) {
	// Setup a bunch of encoded strings
	encoded := make([]string, b.N/10+1)
	for i := range encoded {
		data := make([]byte, 32)
		for i := range data {
			data[i] = byte(rand.Int() & (1<<8 - 1))
		}
		encoded[i] = EncodeToString(data)
	}
	b.ResetTimer()

	// Decode them each about 10 times
	for i := 0; i < b.N; i++ {
		_, err := DecodeString(encoded[i%len(encoded)], 32)
		if err != nil {
			b.Error(err)
		}
	}
}
