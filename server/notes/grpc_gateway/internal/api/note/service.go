package note

import (
	"gateway/internal/service"
	desc "gateway/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	NoteService service.NoteService
}

func NewImplementation(noteService service.NoteService) *Implementation {
	return &Implementation{
		NoteService: noteService,
	}
}
