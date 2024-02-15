package tests

import (
	"context"
	"fmt"
	"gateway/internal/api/note"
	"gateway/internal/model"
	"gateway/internal/service"
	serviceMocks "gateway/internal/service/mocks"
	desc "gateway/pkg/note_v1"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {

	type gatewayerviceMockFunc func(mc *minimock.Controller) service.NoteService

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
		name              string
		args              args
		want              *desc.CreateResponse
		err               error
		gatewayerviceMock gatewayerviceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			gatewayerviceMock: func(mc *minimock.Controller) service.NoteService {
				mock := serviceMocks.NewgatewayerviceMock(mc)
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
			gatewayerviceMock: func(mc *minimock.Controller) service.NoteService {
				mock := serviceMocks.NewgatewayerviceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gatewayerviceMock := tt.gatewayerviceMock(mc)
			api := note.NewImplementation(gatewayerviceMock)

			res, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}

}
