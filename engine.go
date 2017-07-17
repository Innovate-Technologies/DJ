package main

import (
	"io"

	"github.com/chuckpreslar/emission"
	"github.com/innovate-technologies/DJ/data"
)

// Engine is the interface an engine has to comply to
type Engine interface {
	Start(*emission.Emitter)
	PutQueue([]data.Song)
	Skip()
	PipeLive(io.Reader)
	SetMetadata(string) // TO DO: extend
}
