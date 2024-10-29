package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
	"github.com/valek177/chat-server/internal/api/chat"
	"github.com/valek177/chat-server/internal/service"
	serviceMocks "github.com/valek177/chat-server/internal/service/mocks"
)

func TestDeleteChat(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *chat_v1.DeleteChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = int64(342)

		serviceErr = fmt.Errorf("service error")

		req = &chat_v1.DeleteChatRequest{
			Id: id,
		}

		res = &emptypb.Empty{}
	)

	testsSuccessful := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.DeleteChatMock.Expect(ctx, id).Return(nil)
				return mock
			},
		},
	}
	testsErrors := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "service error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.DeleteChatMock.Expect(ctx, id).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range testsSuccessful {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewImplementation(chatServiceMock)

			res, err := api.DeleteChat(tt.args.ctx, tt.args.req)

			assert.Nil(t, err)
			assert.Equal(t, tt.want, res)
		})
	}

	for _, tt := range testsErrors {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewImplementation(chatServiceMock)

			_, err := api.DeleteChat(tt.args.ctx, tt.args.req)

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, "service error")
		})
	}
}
