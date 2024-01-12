package taskcardrecord

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/db/query"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/Flash-Pass/flash-pass-server/internal/utils/timeutils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	taskCardRecord query.ITaskCardRecordDo
	card           query.ICardDo
	task           query.ITaskDo
	bookCard       query.IBookCardDo
}

type IRepository interface {
	GetById(ctx context.Context, taskCardRecordId int64) (*model.TaskCardRecord, error)
	Feed(ctx context.Context, userId int64) ([]*model.TaskCardRecord, error)
	GetRecordByTimestamp(ctx context.Context, taskId, timestamp, groupNum, groupSize int64) ([]*model.TaskCardRecord, error)
	CountRecordByTimestamp(ctx context.Context, taskId, timestamp int64) (int64, error)
	Generate(ctx context.Context, taskId, userId, timestamp int64) error
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		taskCardRecord: query.TaskCardRecord.WithContext(db.Statement.Context),
	}
}

func (r *Repository) GetById(ctx context.Context, taskCardRecordId int64) (*model.TaskCardRecord, error) {
	logger := ctxlog.Extract(ctx)

	taskCardRecord, err := r.taskCardRecord.WithContext(ctx).Where(query.TaskCardRecord.Id.Eq(taskCardRecordId)).First()
	if err != nil {
		logger.Error("get task card record by id failed", zap.Error(err), zap.Int64("task_card_record_id", taskCardRecordId))
		return nil, err
	}

	return taskCardRecord, nil
}

func (r *Repository) Feed(ctx context.Context, userId int64) ([]*model.TaskCardRecord, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetRecordByTimestamp(ctx context.Context, taskId, timestamp, groupNum, groupSize int64) ([]*model.TaskCardRecord, error) {
	logger := ctxlog.Extract(ctx)

	dayStart := timeutils.GetStartTimestampOfDay(timestamp)
	dayEnd := timeutils.GetEndTimestampOfDay(timestamp)

	if groupNum < 0 || groupSize < 0 {
		list, err := r.taskCardRecord.WithContext(ctx).Where(
			query.TaskCardRecord.TaskId.Eq(taskId),
			query.TaskCardRecord.Timestamp.Between(dayStart, dayEnd),
		).Find()

		if err != nil {
			logger.Error("get task card record by timestamp failed", zap.Error(err),
				zap.Int64("task_id", taskId), zap.Int64("timestamp", timestamp))
			return nil, err
		}

		return list, nil
	} else {
		list, err := r.taskCardRecord.WithContext(ctx).Where(
			query.TaskCardRecord.TaskId.Eq(taskId),
			query.TaskCardRecord.Timestamp.Between(dayStart, dayEnd),
		).Limit(int(groupSize)).Offset(int(groupNum * groupSize)).Find()

		if err != nil {
			logger.Error("get task card record by timestamp failed", zap.Error(err),
				zap.Int64("task_id", taskId), zap.Int64("timestamp", timestamp))
			return nil, err
		}

		return list, nil
	}
}

func (r *Repository) CountRecordByTimestamp(ctx context.Context, taskId, timestamp int64) (int64, error) {
	logger := ctxlog.Extract(ctx)

	dayStart := timeutils.GetStartTimestampOfDay(timestamp)
	dayEnd := timeutils.GetEndTimestampOfDay(timestamp)

	count, err := r.taskCardRecord.WithContext(ctx).Where(
		query.TaskCardRecord.TaskId.Eq(taskId),
		query.TaskCardRecord.Timestamp.Between(dayStart, dayEnd),
	).Count()

	if err != nil {
		logger.Error("count task card record by timestamp failed", zap.Error(err),
			zap.Int64("task_id", taskId), zap.Int64("timestamp", timestamp))
		return 0, err
	}

	return count, nil
}

func (r *Repository) Generate(ctx context.Context, taskId, userId, timestamp int64) error {
	logger := ctxlog.Extract(ctx)

	task, err := r.task.WithContext(ctx).Where(query.Task.Id.Eq(taskId)).First()
	if err != nil {
		logger.Error("get task failed", zap.Error(err), zap.Int64("task_id", taskId))
		return err
	}

	cardList, err := r.bookCard.WithContext(ctx).Where(
		query.BookCard.BookId.Eq(task.BookId),
	).Find()
	if err != nil {
		logger.Error("get book card list failed", zap.Error(err), zap.Int64("book_id", task.BookId))
		return err
	}

	var taskCardRecordList []*model.TaskCardRecord
	for _, item := range cardList {
		taskCardRecordList = append(taskCardRecordList, &model.TaskCardRecord{
			TaskId:    taskId,
			CardId:    item.CardId,
			UserId:    userId,
			Timestamp: timestamp,
			IsViewed:  false,
		})
	}

	if err := r.taskCardRecord.WithContext(ctx).Create(taskCardRecordList...); err != nil {
		logger.Error("create task card record failed", zap.Error(err), zap.Int64("task_id", taskId))
		return err
	}

	return nil
}

var _ IRepository = (*Repository)(nil)
