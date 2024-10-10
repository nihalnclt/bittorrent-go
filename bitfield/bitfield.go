package bitfield

// A Bitfield represents the pieces that a peer has
type Bitfield []byte

// 0 1 0 1 0 0 1 1 - 1 byte (each bit have single piece available or not)
// 5 / 8 -> 0
// 5 % 8 -> 5
// This means the piece is in first byte's 5th position
func (bf Bitfield) HasPiece(index int) bool {
	byteIndex := index / 8
	offset := index % 8
	if byteIndex < 0 || byteIndex >= len(bf) {
		return false
	}
	// shifting bit to right for coming offset position at end
	// then compare with 00000001
	return bf[byteIndex]>>(7-offset)&1 != 0
}

func (bf Bitfield) SetPiece(index int) {
	byteIndex := index / 8
	offset := index % 8

	// silently discard invalid bounded index
	if byteIndex < 0 || byteIndex > len(bf) {
		return
	}
	bf[byteIndex] |= 1 << (7 - offset)
}
