package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
	"github.com/valek177/chat-server/internal/client/db"
	dbMocks "github.com/valek177/chat-server/internal/client/db/mocks"
	"github.com/valek177/chat-server/internal/model"
	"github.com/valek177/chat-server/internal/repository"
	repoMocks "github.com/valek177/chat-server/internal/repository/mocks"
	"github.com/valek177/chat-server/internal/service/chat"
)

func TestDeleteChat(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type logRepositoryMockFunc func(mc *minimock.Controller) repository.LogRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *chat_v1.DeleteChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = int64(342)
		// recordId = int64(200)

		repoErr = fmt.Errorf("repo error")

		req = &chat_v1.DeleteChatRequest{
			Id: id,
		}

		// record = &model.Record{
		// 	ID:     recordId,
		// 	ChatID: id,
		// 	Action: "delete",
		// }

		res = &emptypb.Empty{}
	)

	chatRepoFunc := func(mc *minimock.Controller) repository.ChatRepository {
		mock := repoMocks.NewChatRepositoryMock(mc)
		mock.DeleteChatMock.Expect(ctx, id).Return(nil)
		return mock
	}

	logRepoFunc := func(mc *minimock.Controller) repository.LogRepository {
		mock := repoMocks.NewLogRepositoryMock(mc)
		mock.CreateRecordMock.Set(func(ctx context.Context, model *model.Record) (int64, error) {
			return 0, nil
		})
		return mock
	}

	txManagerFunc := func(mc *minimock.Controller) db.TxManager {
		mock := dbMocks.NewTxManagerMock(mc)
		mock.ReadCommittedMock.
			Set(func(ctx context.Context, f db.Handler) error { return f(ctx) })
		return mock
	}

	testsSuccessful := []struct {
		name               string
		args               args
		want               *emptypb.Empty
		err                error
		chatRepositoryMock chatRepositoryMockFunc
		logRepositoryMock  logRepositoryMockFunc
		txManagerMock      txManagerMockFunc
	}{
		{
			name: "success case 1",
			args: args{
				ctx: ctx,
				req: req,
			},
			want:               res,
			err:                nil,
			chatRepositoryMock: chatRepoFunc,
			logRepositoryMock:  logRepoFunc,
			txManagerMock:      txManagerFunc,
		},
	}

	testsErrors := []struct {
		name               string
		args               args
		want               *emptypb.Empty
		err                error
		chatRepositoryMock chatRepositoryMockFunc
		logRepositoryMock  logRepositoryMockFunc
		txManagerMock      txManagerMockFunc
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
				mock.DeleteChatMock.Expect(ctx, id).Return(repoErr)
				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repoMocks.NewLogRepositoryMock(mc)
				return mock
			},
			txManagerMock: txManagerFunc,
		},
	}

	for _, tt := range testsSuccessful {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepositoryMock := tt.chatRepositoryMock(mc)
			logRepositoryMock := tt.logRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(mc)

			service := chat.NewService(
				chatRepositoryMock, logRepositoryMock, txManagerMock,
			)

			err := service.DeleteChat(tt.args.ctx, tt.args.req.Id)

			assert.Nil(t, err)
		})
	}

	for _, tt := range testsErrors {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepositoryMock := tt.chatRepositoryMock(mc)
			logRepositoryMock := tt.logRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(mc)

			service := chat.NewService(
				chatRepositoryMock, logRepositoryMock, txManagerMock,
			)

			err := service.DeleteChat(tt.args.ctx, tt.args.req.Id)
			assert.NotNil(t, err)
			assert.ErrorContains(t, err, "repo error")
		})
	}
}
