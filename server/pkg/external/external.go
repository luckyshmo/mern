package external

import (
	"github.com/luckyshmo/api-example/models/keep"
)

type NoteGetter interface {
	GetAll() (keep.Note, error)
}

type DocsGetter interface {
	GetDocs()
	GetDoc()
}

type ExternalSources struct {
	NoteGetter
	DocsGetter
}

func NewExternalSource(ng NoteGetter) *ExternalSources {
	return &ExternalSources{
		NoteGetter: ng,
		DocsGetter: nil, //TODO
	}
}
