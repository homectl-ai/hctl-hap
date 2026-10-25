package hap

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/homectl-ai/hctl-hap/tlv8"
)

// iOS lists an accessory's pairings with only Method (5) and State (1):
// it must decode, not fail for the missing Identifier (brutella/hap#67).
func TestListPairingsRequest(t *testing.T) {
	b, _ := hex.DecodeString("000105060101") // as captured from a real controller
	var d pairingsRequest
	if err := tlv8.UnmarshalReader(bytes.NewReader(b), &d); err != nil {
		t.Fatal(err)
	}
	if d.Method != 5 || d.State != 1 || d.Identifier != "" {
		t.Errorf("%+v", d)
	}
}
