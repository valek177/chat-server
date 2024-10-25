package converter

import (
	"time"

	"github.com/valek177/chat-server/internal/model"
)

// ToRecordRepoFromService converts params to Record model
func ToRecordRepoFromService(chatID int64, action string) *model.Record {
	return &model.Record{
		ChatID:    chatID,
		CreatedAt: time.Now(),
		Action:    action,
	}
}
