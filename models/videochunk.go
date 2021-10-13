package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
)

type VideoChunk struct {
	Index		int		`bson:"index" json:"index"`
	Hash 		string	`bson:"hash" json:"hash"`
	FileHash	string	`bson:"file_hash" json:"file_hash"`
}

var (
	FileChunkNotMatch = errors.New("chunk's hash or size does not match our record")
	ChunkNotOnDisk    = errors.New("chunk is not on the disk")
)

// NewVideoChunk constructor
func NewVideoChunk(index int, hash string, fileHash string) *VideoChunk {
	return &VideoChunk{
		Index:index,
		Hash: hash,
		FileHash: fileHash,
	}
}

// GetChunkFilePath get path
func (vc *VideoChunk) GetChunkFilePath() string {
	return fmt.Sprintf("storage/%s/%s", vc.FileHash, vc.Hash)
}

// StoreToDisk func to  store to disk
func(vc *VideoChunk)StoreToDisk(content []byte) error {
	// verify MD5 checksum
	sum := md5.Sum(content)
	var tmp  [16]byte
	copy(tmp, sum)
	if hex.EncodeToString() != vc.Hash {
		return FileChunkNotMatch
	}

}
