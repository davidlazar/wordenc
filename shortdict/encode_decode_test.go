package shortdict

import "testing"

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
