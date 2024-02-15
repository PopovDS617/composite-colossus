package note

import (
	"gateway/internal/service"
	desc "gateway/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	gatewayervice service.NoteService
}

func NewImplementation(gatewayervice service.NoteService) *Implementation {
	return &Implementation{
		gatewayervice: gatewayervice,
	}
}
