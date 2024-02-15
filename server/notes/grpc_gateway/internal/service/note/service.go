package note

import (
	"gateway/internal/client/db"
	"gateway/internal/repository"
	"gateway/internal/service"
)

type serv struct {
	noteRepository repository.NoteRepository
	txManager      db.TxManager
}

func NewService(
	noteRepository repository.NoteRepository,
	txManager db.TxManager,
) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
		txManager:      txManager,
	}
}

func NewMockService(deps ...interface{}) service.NoteService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.NoteRepository:
			srv.noteRepository = s
		}
	}

	return &srv
}
