package tests

import "testing"

// import (
// 	"context"
// 	"fmt"
// 	"gateway/internal/model"
// 	"gateway/internal/repository"
// 	repoMocks "gateway/internal/repository/mocks"
// 	"gateway/internal/service/note"
// 	"testing"

// 	"github.com/brianvoe/gofakeit/v6"
// 	"github.com/gojuno/minimock/v3"

// 	"github.com/stretchr/testify/require"
// )

func TestCreate(t *testing.T) {
	// t.Parallel()
	// type noteRepositoryMockFunc func(mc *minimock.Controller) repository.NoteRepository

	// type args struct {
	// 	ctx context.Context
	// 	req *model.NoteInfo
	// }

	// var (
	// 	ctx = context.Background()
	// 	mc  = minimock.NewController(t)

	// 	id      = gofakeit.Int64()
	// 	title   = gofakeit.Animal()
	// 	content = gofakeit.Animal()

	// 	repoErr = fmt.Errorf("repo error")

	// 	req = &model.NoteInfo{
	// 		Title:   title,
	// 		Content: content,
	// 	}
	// )

	// tests := []struct {
	// 	name               string
	// 	args               args
	// 	want               int64
	// 	err                error
	// 	noteRepositoryMock noteRepositoryMockFunc
	// }{
	// 	{
	// 		name: "success case",
	// 		args: args{
	// 			ctx: ctx,
	// 			req: req,
	// 		},
	// 		want: id,
	// 		err:  nil,
	// 		noteRepositoryMock: func(mc *minimock.Controller) repository.NoteRepository {
	// 			mock := repoMocks.NewNoteRepositoryMock(mc)
	// 			mock.CreateMock.Expect(ctx, req).Return(id, nil)
	// 			return mock
	// 		},
	// 	},
	// 	{
	// 		name: "fail case",
	// 		args: args{
	// 			ctx: ctx,
	// 			req: req,
	// 		},
	// 		want: 0,
	// 		err:  repoErr,
	// 		noteRepositoryMock: func(mc *minimock.Controller) repository.NoteRepository {
	// 			mock := repoMocks.NewNoteRepositoryMock(mc)
	// 			mock.CreateMock.Expect(ctx, req).Return(0, repoErr)
	// 			return mock
	// 		},
	// 	},
	// }

	// for _, tt := range tests {
	// 	tt := tt
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		t.Parallel()

	// 		noteRepoMock := tt.noteRepositoryMock(mc)
	// 		service := note.NewMockService(noteRepoMock)

	// 		res, err := service.Create(tt.args.ctx, tt.args.req)

	// 		fmt.Println(res)

	// 		require.Equal(t, tt.err, err)
	// 		require.Equal(t, tt.want, res)
	// 	})
	// }
}
