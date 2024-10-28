package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
	"github.com/valek177/chat-server/internal/api/chat"
	"github.com/valek177/chat-server/internal/repository"
	repoMocks "github.com/valek177/chat-server/internal/repository/mocks"
)

func TestCreateChat(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository

	type args struct {
		ctx context.Context
		req *chat_v1.CreateChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id      = int64(342)
		name    = gofakeit.Name()
		userIds = []int64{1, 2, 3}

		repoErr = fmt.Errorf("repo error")

		req = &chat_v1.CreateChatRequest{
			Name:    name,
			UserIds: userIds,
		}

		res = &chat_v1.CreateChatResponse{
			Id: id,
		}
	)

	testsSuccessful := []struct {
		name               string
		args               args
		want               *chat_v1.CreateChatResponse
		err                error
		chatRepositoryMock chatRepositoryMockFunc
	}{
		{
			name: "success case 1",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx, req).Return(id, nil)
				return mock
			},
		},
	}
	testsErrors := []struct {
		name               string
		args               args
		want               *chat_v1.CreateChatResponse
		err                error
		chatRepositoryMock chatRepositoryMockFunc
	}{
		{
			name: "repo error case 1",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  repoErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx, req).Return(0, repoErr)
				return mock
			},
		},
	}

	for _, tt := range testsSuccessful {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepositoryMock := tt.chatRepositoryMock(mc)
			service := chat.NewImplementation(chatRepositoryMock)

			newID, err := service.CreateChat(tt.args.ctx, tt.args.req)

			assert.Nil(t, err)
			assert.Equal(t, tt.want, newID)
		})
	}

	for _, tt := range testsErrors {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepositoryMock := tt.chatRepositoryMock(mc)
			service := chat.NewImplementation(chatRepositoryMock)

			_, err := service.CreateChat(tt.args.ctx, tt.args.req)

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, "repo error")
		})
	}
}
