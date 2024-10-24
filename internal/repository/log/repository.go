package log

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/valek177/chat-server/internal/client/db"
	"github.com/valek177/chat-server/internal/model"
	"github.com/valek177/chat-server/internal/repository"
)

const (
	tableName = "chats_log"

	idColumn        = "id"
	chatIdColumn    = "chat_id"
	actionColumn    = "action"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

// NewRepository creates new log repository
func NewRepository(db db.Client) repository.LogRepository {
	return &repo{db: db}
}

// CreateRecord creates new record in chats log table
func (r *repo) CreateRecord(ctx context.Context, record *model.Record) (int64, error) {
	builderInsert := sq.Insert(tableName).
		Columns(chatIdColumn, actionColumn, createdAtColumn).
		Values(record.ChatID, record.Action, time.Now()).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "log_repository.CreateRecord",
		QueryRaw: query,
	}

	var recordId int64
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&recordId); err != nil {
		return 0, err
	}

	return recordId, nil
}
