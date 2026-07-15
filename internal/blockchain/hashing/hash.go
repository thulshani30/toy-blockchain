package hashing

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
)

// CalculateBlockHash computes the SHA-256 hash of a block using
// deterministic serialization of block fields excluding the Hash field.
func CalculateBlockHash(b *block.Block) string {
	data := serializeBlock(b)

	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:])
}

// serializeBlock converts block fields into a deterministic byte sequence.
//
// Hashing order:
// 1. Index
// 2. Timestamp (Unix seconds)
// 3. Merkle Root
// 4. Previous hash
// 5. Nonce
//
// The Hash field is intentionally excluded because it is the value being calculated.
func serializeBlock(b *block.Block) []byte {
	var buf bytes.Buffer

	// Block index
	_ = binary.Write(&buf, binary.BigEndian, int64(b.Index))

	// Timestamp
	_ = binary.Write(&buf, binary.BigEndian, b.Timestamp.Unix())

	// Merkle Root
	writeString(&buf, b.MerkleRoot)

	// Previous hash
	writeString(&buf, b.PreviousHash)

	// Nonce
	_ = binary.Write(&buf, binary.BigEndian, int64(b.Nonce))

	return buf.Bytes()
}

func writeString(buf *bytes.Buffer, s string) {
	_ = binary.Write(buf, binary.BigEndian, uint32(len(s)))
	buf.WriteString(s)
}
