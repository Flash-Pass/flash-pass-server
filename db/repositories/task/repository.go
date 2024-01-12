package task

import (
	"context"
	"errors"
	"github.com/Flash-Pass/flash-pass-server/db"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/db/query"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	task query.ITaskDo
}

type IRepository interface {
	Create(ctx context.Context, planId, bookId, userId int64, name string) (*model.Task, error)
	GetById(ctx context.Context, taskId int64) (*model.Task, error)
	Active(ctx context.Context, taskId int64, isActive bool) (*model.Task, error)
	UpdateTaskName(ctx context.Context, taskId int64, name string) (*model.Task, error)
	DeleteTask(ctx context.Context, taskId int64) error
	GetTaskList(ctx context.Context, userId int64) ([]*model.Task, error)
	GetTaskListByIsActive(ctx context.Context, userId int64, isActive bool) ([]*model.Task, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		task: query.Task.WithContext(db.Statement.Context),
	}
}

func (r *Repository) Create(ctx context.Context, planId, bookId, userId int64, name string) (*model.Task, error) {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	newTask := &model.Task{
		PlanId:    planId,
		BookId:    bookId,
		Name:      name,
		CreatedBy: userId,
	}

	if err := tx.Task.Create(newTask); err != nil {
		logger.Error("create task failed", zap.Error(err),
			zap.Int64("plan_id", planId), zap.Int64("book_id", bookId), zap.String("name", name))
		return nil, err
	}

	return newTask, nil
}

func (r *Repository) GetById(ctx context.Context, taskId int64) (*model.Task, error) {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	task, err := tx.Task.WithContext(ctx).Where(query.Task.Id.Eq(taskId)).First()
	if err != nil {
		logger.Error("get task by id failed", zap.Error(err), zap.Int64("task_id", taskId))
		return nil, err
	}

	return task, nil
}

func (r *Repository) Active(ctx context.Context, taskId int64, isActive bool) (*model.Task, error) {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	updateInfo, err := tx.Task.WithContext(ctx).Where(query.Task.Id.Eq(taskId)).Update(query.Task.IsActive, isActive)
	if err != nil {
		logger.Error("update task failed", zap.Error(err), zap.Int64("task_id", taskId), zap.Bool("is_active", isActive))
		return nil, err
	}
	if updateInfo.RowsAffected == 0 {
		logger.Warn("no task updated", zap.Int64("task_id", taskId), zap.Bool("is_active", isActive))
		return nil, nil
	}

	return r.GetById(ctx, taskId)
}

func (r *Repository) UpdateTaskName(ctx context.Context, taskId int64, name string) (*model.Task, error) {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	updateInfo, err := tx.Task.WithContext(ctx).Where(query.Task.Id.Eq(taskId)).Update(query.Task.Name, name)
	if err != nil {
		logger.Error("update task failed", zap.Error(err), zap.Int64("task_id", taskId), zap.String("name", name))
		return nil, err
	}
	if updateInfo.RowsAffected == 0 {
		logger.Warn("no task updated", zap.Int64("task_id", taskId), zap.String("name", name))
		return nil, nil
	}

	return r.GetById(ctx, taskId)
}

func (r *Repository) DeleteTask(ctx context.Context, taskId int64) error {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	updateInfo, err := tx.Task.WithContext(ctx).Where(query.Task.Id.Eq(taskId)).Update(query.Task.IsDeleted, true)
	if err != nil {
		logger.Error("update task failed", zap.Error(err), zap.Int64("task_id", taskId))
		return err
	}
	if updateInfo.RowsAffected == 0 {
		logger.Warn("no task updated", zap.Int64("task_id", taskId))
		return errors.New("no task updated")
	}

	return nil
}

func (r *Repository) GetTaskList(ctx context.Context, userId int64) ([]*model.Task, error) {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	list, err := tx.Task.WithContext(ctx).Where(query.Task.CreatedBy.Eq(userId)).Find()
	if err != nil {
		logger.Error("get task list failed", zap.Error(err))
		return nil, err
	}

	return list, nil
}

func (r *Repository) GetTaskListByIsActive(ctx context.Context, userId int64, isActive bool) ([]*model.Task, error) {
	tx := db.Tx(ctx)
	logger := ctxlog.Extract(ctx)

	list, err := tx.Task.WithContext(ctx).Where(
		query.Task.IsActive.Is(isActive),
		query.Task.CreatedBy.Eq(userId),
	).Find()
	if err != nil {
		logger.Error("get task list failed", zap.Error(err), zap.Bool("is_active", isActive))
		return nil, err
	}

	return list, nil
}

var _ IRepository = (*Repository)(nil)
