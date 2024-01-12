package task

import (
	"context"
	"errors"
	"fmt"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/cache"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"go.uber.org/zap"
)

type InfoCache struct {
	taskInfo      *cache.Cache[*model.Task]
	learnedGroup  *cache.Cache[int64]
	reviewedGroup *cache.Cache[int64]
}

type IInfoCache interface {
	SetTaskInfo(ctx context.Context, taskId int64, task *model.Task) error
	GetTaskInfo(ctx context.Context, taskId int64) (*model.Task, error)
	SetLearnedGroup(ctx context.Context, userId, taskId, groupNum int64) error
	GetLearnedGroup(ctx context.Context, userId, taskId int64) (int64, error)
	SetReviewedGroup(ctx context.Context, userId, taskId, groupNum int64) error
	GetReviewedGroup(ctx context.Context, userId, taskId int64) (int64, error)
}

func NewInfoCache() *InfoCache {
	taskInfoCache := cache.New[*model.Task]()
	learnedGroupCache := cache.New[int64]()
	reviewedGroupCache := cache.New[int64]()
	return &InfoCache{
		taskInfo:      taskInfoCache,
		learnedGroup:  learnedGroupCache,
		reviewedGroup: reviewedGroupCache,
	}
}

func (c *InfoCache) SetTaskInfo(ctx context.Context, taskId int64, task *model.Task) error {
	c.taskInfo.Set(fmt.Sprint(taskId), task)
	return nil
}

func (c *InfoCache) GetTaskInfo(ctx context.Context, taskId int64) (*model.Task, error) {
	logger := ctxlog.Extract(ctx)

	task, ok := c.taskInfo.Get(fmt.Sprint(taskId))
	if !ok {
		logger.Warn("get task info from cache failed", zap.Int64("task_id", taskId))
		return nil, errors.New("get task info from cache failed")
	}

	return task, nil
}

func (c *InfoCache) SetLearnedGroup(ctx context.Context, userId, taskId, groupNum int64) error {
	c.learnedGroup.Set(fmt.Sprintf("%d_%d", userId, taskId), groupNum)
	return nil
}

func (c *InfoCache) GetLearnedGroup(ctx context.Context, userId, taskId int64) (int64, error) {
	logger := ctxlog.Extract(ctx)

	groupNum, ok := c.learnedGroup.Get(fmt.Sprintf("%d_%d", userId, taskId))
	if !ok {
		logger.Warn("get learned group from cache failed", zap.Int64("user_id", userId), zap.Int64("task_id", taskId))
		return -1, nil
	}

	return groupNum, nil
}

func (c *InfoCache) SetReviewedGroup(ctx context.Context, userId, taskId, groupNum int64) error {
	c.reviewedGroup.Set(fmt.Sprintf("%d_%d", userId, taskId), groupNum)
	return nil
}

func (c *InfoCache) GetReviewedGroup(ctx context.Context, userId, taskId int64) (int64, error) {
	logger := ctxlog.Extract(ctx)

	groupNum, ok := c.reviewedGroup.Get(fmt.Sprintf("%d_%d", userId, taskId))
	if !ok {
		logger.Warn("get reviewed group from cache failed", zap.Int64("user_id", userId), zap.Int64("task_id", taskId))
		return -1, nil
	}

	return groupNum, nil
}

var _ IInfoCache = (*InfoCache)(nil)
