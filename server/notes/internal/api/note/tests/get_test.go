package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"notes/internal/api/note"
	"notes/internal/model"
	"notes/internal/service"
	serviceMocks "notes/internal/service/mocks"
	desc "notes/pkg/note_v1"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type noteServiceMockFunc func(mc *minimock.Controller) service.NoteService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		title     = gofakeit.Animal()
		content   = gofakeit.Animal()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetRequest{
			Id: id,
		}

		serviceRes = &model.Note{
			ID: id,
			Info: model.NoteInfo{
				Title:   title,
				Content: content,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}

		res = &desc.GetResponse{
			Note: &desc.Note{
				Id: id,
				Info: &desc.NoteInfo{
					Title:   title,
					Content: content,
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		noteServiceMock noteServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			noteServiceMock: func(mc *minimock.Controller) service.NoteService {
				mock := serviceMocks.NewNoteServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(serviceRes, nil)
				return mock
			},
		},
		{
			name: "fail case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			noteServiceMock: func(mc *minimock.Controller) service.NoteService {
				mock := serviceMocks.NewNoteServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			noteServiceMock := tt.noteServiceMock(mc)
			api := note.NewImplementation(noteServiceMock)

			newID, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
