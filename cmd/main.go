package main

import (
	"context"
	"flag"
	"log"
	"net"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
	"github.com/valek177/chat-server/internal/config"
	"github.com/valek177/chat-server/internal/config/env"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	chat_v1.UnimplementedChatV1Server
	pool *pgxpool.Pool
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	chat_v1.RegisterChatV1Server(s, &server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// CreateChat creates new chat with parameters
func (s *server) CreateChat(ctx context.Context, req *chat_v1.CreateChatRequest) (
	*chat_v1.CreateChatResponse, error,
) {
	log.Printf("Create new chat with name %s", req.GetName())

	builderInsert := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(req.GetName()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var chatID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		log.Fatalf("failed to insert chats: %v", err)
	}

	log.Printf("inserted chat with id: %d", chatID)

	builderInsertChatUser := sq.Insert("chat_users").
		PlaceholderFormat(sq.Dollar).
		Columns("user_id", "chat_id")

	for _, id := range req.UserIds {
		builderInsertChatUser = builderInsertChatUser.Values(id, chatID)
	}

	query, args, err = builderInsertChatUser.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return nil, err
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to insert chat users: %v", err)
	}

	return &chat_v1.CreateChatResponse{
		Id: chatID,
	}, nil
}

// DeleteChat removes chat
func (s *server) DeleteChat(ctx context.Context, req *chat_v1.DeleteChatRequest) (
	*emptypb.Empty, error,
) {
	id := req.GetId()
	log.Printf("Delete chat users for chat id %d", id)

	builderDelete := sq.Delete("chat_users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"chat_id": id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to delete chat user: %v", err)
	}

	log.Printf("Delete chat for chat id %d", id)

	builderChatDelete := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err = builderChatDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to delete chat: %v", err)
		return nil, err
	}

	log.Printf("deleted chat with id %d", id)

	return &emptypb.Empty{}, nil
}

// SendMessage sends message to server
func (s *server) SendMessage(_ context.Context, req *chat_v1.SendMessageRequest) (
	*emptypb.Empty, error,
) {
	log.Printf("Send message from %s with text: %s", req.GetFrom(), req.GetText())

	return &emptypb.Empty{}, nil
}
