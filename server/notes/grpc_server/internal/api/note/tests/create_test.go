package tests

import (
	"context"
	"fmt"
	"notes/internal/api/note"
	"notes/internal/model"
	"notes/internal/service"
	serviceMocks "notes/internal/service/mocks"
	desc "notes/pkg/note_v1"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {

	type noteServiceMockFunc func(mc *minimock.Controller) service.NoteService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id      = gofakeit.Int64()
		title   = gofakeit.Animal()
		content = gofakeit.Animal()

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			Info: &desc.NoteInfo{
				Title:   title,
				Content: content,
			},
		}

		info = &model.NoteInfo{
			Title:   title,
			Content: content,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
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
				mock.CreateMock.Expect(ctx, info).Return(id, nil)
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
				mock.CreateMock.Expect(ctx, info).Return(0, serviceErr)
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

			res, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}

}
