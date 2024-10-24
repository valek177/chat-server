package auth

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
	"github.com/valek177/chat-server/internal/client/db"
	"github.com/valek177/chat-server/internal/repository"
)

const (
	tableChatsName      = "chats"
	tableChatsUsersName = "chat_users"

	idColumn     = "id"
	chatIdColumn = "chat_id"
	userIdColumn = "user_id"
	nameColumn   = "name"
)

type repo struct {
	db db.Client
}

// NewRepository creates new repository object
func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

// CreateChat creates new chat with specified parameters
func (r *repo) CreateChat(ctx context.Context, req *chat_v1.CreateChatRequest) (int64, error) {
	chatID, err := r.createNewChat(ctx, req)
	if err != nil {
		return 0, err
	}

	err = r.createChatUsers(ctx, req, chatID)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

// DeleteChat removes chat
func (r *repo) DeleteChat(ctx context.Context, id int64) error {
	err := r.deleteChatUsers(ctx, id)
	if err != nil {
		return err
	}

	err = r.deleteChat(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) createNewChat(ctx context.Context, req *chat_v1.CreateChatRequest) (int64, error) {
	builderInsert := sq.Insert(tableChatsName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn).
		Values(req.Name).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.CreateChat",
		QueryRaw: query,
	}

	var chatID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

func (r *repo) createChatUsers(ctx context.Context, req *chat_v1.CreateChatRequest, chatID int64) error {
	builderInsertChatUser := sq.Insert(tableChatsUsersName).
		PlaceholderFormat(sq.Dollar).
		Columns(userIdColumn, chatIdColumn)

	for _, id := range req.UserIds {
		builderInsertChatUser = builderInsertChatUser.Values(id, chatID)
	}

	query, args, err := builderInsertChatUser.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.CreateChatUser",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) deleteChatUsers(ctx context.Context, id int64) error {
	builderDelete := sq.Delete("chat_users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIdColumn: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.DeleteChatUsers",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) deleteChat(ctx context.Context, id int64) error {
	builderChatDelete := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builderChatDelete.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.DeleteChat",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
