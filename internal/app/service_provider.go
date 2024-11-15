package app

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/valek177/chat-server/internal/api/chat"
	"github.com/valek177/chat-server/internal/client"
	"github.com/valek177/chat-server/internal/config"
	"github.com/valek177/chat-server/internal/config/env"
	"github.com/valek177/chat-server/internal/repository"
	chatRepository "github.com/valek177/chat-server/internal/repository/chat"
	logRepo "github.com/valek177/chat-server/internal/repository/log"
	"github.com/valek177/chat-server/internal/service"
	chatService "github.com/valek177/chat-server/internal/service/chat"
	"github.com/valek177/platform-common/pkg/client/db"
	"github.com/valek177/platform-common/pkg/client/db/pg"
	"github.com/valek177/platform-common/pkg/client/db/transaction"
	"github.com/valek177/platform-common/pkg/closer"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager      db.TxManager
	chatRepository repository.ChatRepository
	logRepository  repository.LogRepository

	authClient client.AuthClient

	authConn *grpc.ClientConn

	chatService service.ChatService

	chatImpl *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig returns new PGConfig
func (s *serviceProvider) PGConfig() (config.PGConfig, error) {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			return nil, err
		}

		s.pgConfig = cfg
	}

	return s.pgConfig, nil
}

// GRPCConfig returns new GRPCConfig
func (s *serviceProvider) GRPCConfig() (config.GRPCConfig, error) {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			return nil, err
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig, nil
}

// DBClient returns new db client
func (s *serviceProvider) DBClient(ctx context.Context) (db.Client, error) {
	if s.dbClient == nil {
		pgConfig, err := s.PGConfig()
		if err != nil {
			return nil, err
		}
		cl, err := pg.New(ctx, pgConfig.DSN())
		if err != nil {
			return nil, err
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			return nil, err
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient, nil
}

// TxManager returns new db TxManager
func (s *serviceProvider) TxManager(ctx context.Context) (db.TxManager, error) {
	if s.txManager == nil {
		dbClient, err := s.DBClient(ctx)
		if err != nil {
			return nil, err
		}
		s.txManager = transaction.NewTransactionManager(dbClient.DB())
	}

	return s.txManager, nil
}

// ChatRepository returns new ChatRepository
func (s *serviceProvider) ChatRepository(ctx context.Context) (repository.ChatRepository, error) {
	if s.chatRepository == nil {
		dbClient, err := s.DBClient(ctx)
		if err != nil {
			return nil, err
		}
		s.chatRepository = chatRepository.NewRepository(dbClient)
	}

	return s.chatRepository, nil
}

// LogRepository returns new LogRepository
func (s *serviceProvider) LogRepository(ctx context.Context) (repository.LogRepository, error) {
	if s.logRepository == nil {
		dbClient, err := s.DBClient(ctx)
		if err != nil {
			return nil, err
		}
		s.logRepository = logRepo.NewRepository(dbClient)
	}

	return s.logRepository, nil
}

// ChatService returns new ChatService
func (s *serviceProvider) ChatService(ctx context.Context) (service.ChatService, error) {
	if s.chatService == nil {
		chatRepo, err := s.ChatRepository(ctx)
		if err != nil {
			return nil, err
		}
		logRepo, err := s.LogRepository(ctx)
		if err != nil {
			return nil, err
		}
		txManager, err := s.TxManager(ctx)
		if err != nil {
			return nil, err
		}
		s.chatService = chatService.NewService(
			chatRepo, logRepo, txManager,
		)
	}

	return s.chatService, nil
}

// ChatImpl returns new Chat Service implementation
func (s *serviceProvider) ChatImpl(ctx context.Context) (*chat.Implementation, error) {
	if s.chatImpl == nil {
		chatServ, err := s.ChatService(ctx)
		if err != nil {
			return nil, err
		}
		s.chatImpl = chat.NewImplementation(chatServ)
	}

	return s.chatImpl, nil
}

func (s *serviceProvider) AuthClient() (client.AuthClient, error) {
	if s.authClient == nil {
		authConn, err := s.AuthConnection()
		if err != nil {
			return nil, err
		}
		s.authClient = client.NewAuthClient(authConn)
	}

	return s.authClient, nil
}

func (s *serviceProvider) AuthConnection() (*grpc.ClientConn, error) {
	if s.authConn == nil {
		var err error
		creds := insecure.NewCredentials()
		conn, err := grpc.NewClient("127.0.0.1:50061", grpc.WithTransportCredentials(creds))
		if err != nil {
			return nil, err
		}

		closer.Add(conn.Close)

		s.authConn = conn
	}

	return s.authConn, nil
}
