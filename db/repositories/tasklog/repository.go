package tasklog

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/db/query"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	taskLog query.ITaskLogDo
}

type IRepository interface {
	AddLearnStatus(ctx context.Context, taskCardRecordId int64, status string) error
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		taskLog: query.TaskLog.WithContext(db.Statement.Context),
	}
}

func (r *Repository) AddLearnStatus(ctx context.Context, taskCardRecordId int64, status string) error {
	logger := ctxlog.Extract(ctx)

	taskLog := &model.TaskLog{
		TaskCardRecordId: taskCardRecordId,
		LearnStatus:      status,
	}
	if err := r.taskLog.Create(taskLog); err != nil {
		logger.Error("add learn status failed", zap.Error(err),
			zap.Int64("task_card_record_id", taskCardRecordId), zap.String("status", status))
		return err
	}

	return nil
}

var _ IRepository = (*Repository)(nil)
