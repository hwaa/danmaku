package models

import (
	"fmt"
	"sync"
)

type (
	VideoFileState	int

	//Video storages on the server
	Video struct {
		Title		string
		FileName	string
		Description string
		VideoURL	string
		AvatarImg	string
		Hash		string
		Chunks		[]VideoChunk
		Size		string
		State		VideoFileState
		mu 			sync.Mutex
	}

	//VideoMetaResponse a brief data to description a video
	VideoMetaResponse struct {
		Title		string
		Description string
		AvatarImg	string
		VideoURL	string
	}
)

const (
	Inprogress VideoFileState = iota	//video chunks is not complete received
	Merged								//All video files is received
	complete							//All video chunks is merged to a single file and Checked MD5
)
//NewVideo a constructor to make new video
func NewVideo(title string, fileName string, description string, hash string, chunks []Videochunk, size uint64) *Video {
	return &Video{
		Title: title,
		FileName: fileName,
		Description: description,
		VideoURL: fmt.Sprintf("storage/#{hash}/#{fileName}"),
		Hash: hash,
		Chunks: chunks,

	}
}