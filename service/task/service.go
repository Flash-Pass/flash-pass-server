package task

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/db"
	"time"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/Flash-Pass/flash-pass-server/internal/utils/timeutils"
	"go.uber.org/zap"
)

type Service struct {
	cardRepo           ICardRepository
	taskRepo           ITaskRepository
	taskCardRecordRepo ITaskCardRecordRepository
	taskLogRepo        ITaskLogRepository
	bookRepo           IBookRepository
	bookCardRepo       IBookCardRepository
	planRepo           IPlanRepository
	taskInfoCache      ITaskInfoCache
}

func NewService(cardRepo ICardRepository, taskRepo ITaskRepository, taskCardRecordRepo ITaskCardRecordRepository,
	taskLogRepo ITaskLogRepository, bookRepo IBookRepository, bookCardRepo IBookCardRepository, planRepo IPlanRepository,
	taskInfoCache ITaskInfoCache) *Service {
	return &Service{
		cardRepo:           cardRepo,
		taskRepo:           taskRepo,
		taskCardRecordRepo: taskCardRecordRepo,
		taskLogRepo:        taskLogRepo,
		bookRepo:           bookRepo,
		bookCardRepo:       bookCardRepo,
		planRepo:           planRepo,
		taskInfoCache:      taskInfoCache,
	}
}

type IService interface {
	CreateTask(ctx context.Context, planId, bookId, userId int64, name string) (*model.Task, error)
	Active(ctx context.Context, taskId int64, isActive bool) (*model.Task, error)
	DeleteTask(ctx context.Context, taskId int64) error
	GetTaskList(ctx context.Context, userId int64) ([]*model.Task, error)
	GetTaskListByIsActive(ctx context.Context, userId int64, isActive bool) ([]*model.Task, error)
	Feed(ctx context.Context, userId, taskId int64) ([]*model.TaskCardRecord, error)
	AddLearnStatus(ctx context.Context, taskCardRecordId int64, status string) error
	Update(ctx context.Context, taskId int64, taskName string, isActive bool) (*model.Task, error)
}

type ICardRepository interface {
	GetById(ctx context.Context, cardId int64) (*model.Card, error)
	GetListByIds(ctx context.Context, cardIds []int64) ([]*model.Card, error)
}

type IBookRepository interface {
	GetById(ctx context.Context, bookId int64) (*model.Book, error)
}

type IBookCardRepository interface {
	CountByBookId(ctx context.Context, bookId int64) (int64, error)
}

type IPlanRepository interface {
	GetPlanByTaskId(ctx context.Context, taskId int64) (*model.Plan, error)
}

type ITaskRepository interface {
	Create(ctx context.Context, planId, bookId, userId int64, name string) (*model.Task, error)
	GetById(ctx context.Context, taskId int64) (*model.Task, error)
	Active(ctx context.Context, taskId int64, isActive bool) (*model.Task, error)
	UpdateTaskName(ctx context.Context, taskId int64, name string) (*model.Task, error)
	DeleteTask(ctx context.Context, taskId int64) error
	GetTaskList(ctx context.Context, userId int64) ([]*model.Task, error)
	GetTaskListByIsActive(ctx context.Context, userId int64, isActive bool) ([]*model.Task, error)
}

type ITaskCardRecordRepository interface {
	GetById(ctx context.Context, taskCardRecordId int64) (*model.TaskCardRecord, error)
	Feed(ctx context.Context, userId int64) ([]*model.TaskCardRecord, error)
	GetRecordByTimestamp(ctx context.Context, taskId, timestamp, groupNum, groupSize int64) ([]*model.TaskCardRecord, error)
	CountRecordByTimestamp(ctx context.Context, taskId, timestamp int64) (int64, error)
	Generate(ctx context.Context, taskId, userId, timestamp int64) error
}

type ITaskLogRepository interface {
	AddLearnStatus(ctx context.Context, taskCardRecordId int64, status string) error
}

type ITaskInfoCache interface {
	SetTaskInfo(ctx context.Context, taskId int64, task *model.Task) error
	GetTaskInfo(ctx context.Context, taskId int64) (*model.Task, error)
	SetLearnedGroup(ctx context.Context, userId, taskId, groupNum int64) error
	GetLearnedGroup(ctx context.Context, userId, taskId int64) (int64, error)
	SetReviewedGroup(ctx context.Context, userId, taskId, groupNum int64) error
	GetReviewedGroup(ctx context.Context, userId, taskId int64) (int64, error)
}

func (s *Service) CreateTask(ctx context.Context, planId, bookId, userId int64, name string) (*model.Task, error) {
	return s.taskRepo.Create(ctx, planId, bookId, userId, name)
}

func (s *Service) Active(ctx context.Context, taskId int64, isActive bool) (*model.Task, error) {
	return s.taskRepo.Active(ctx, taskId, isActive)
}

func (s *Service) DeleteTask(ctx context.Context, taskId int64) error {
	return s.taskRepo.DeleteTask(ctx, taskId)
}

func (s *Service) GetTaskList(ctx context.Context, userId int64) ([]*model.Task, error) {
	return s.taskRepo.GetTaskList(ctx, userId)
}

func (s *Service) GetTaskListByIsActive(ctx context.Context, userId int64, isActive bool) ([]*model.Task, error) {
	return s.taskRepo.GetTaskListByIsActive(ctx, userId, isActive)
}

// Feed
// 1. 从Context中获取userId和taskId
// 2. 根据taskId读取task，根据当天日期获得要复习的日期，存入缓存(用户复习任务日期)
// 3. 根据taskId获取当日所有的taskCardRecord记录，获得记录条数; 根据task所绑定的bookId获取包含的card的总数，判断是否已经生
// 成全部的taskCardRecord; 若未生成全部的taskCardRecord，则调用生成taskCardRecord的方法
// 4. 从taskCardRecord中根据taskId和当日日期、复习日期获取全部的记录，并标记学习性质（新学，复习）存入缓存（用户要学习卡片缓存）
// 5. 从缓存中获取一组数据，传递给用户
func (s *Service) Feed(ctx context.Context, userId, taskId int64) ([]*model.TaskCardRecord, error) {
	logger := ctxlog.Extract(ctx)

	var task *model.Task
	task, err := s.taskInfoCache.GetTaskInfo(ctx, taskId)
	if err != nil {
		logger.Warn("get task info from cache failed",
			zap.Error(err), zap.Int64("user_id", userId), zap.Int64("task_id", taskId))
		task, err = s.taskRepo.GetById(ctx, taskId)
		if err != nil {
			return nil, err
		}
		if err := s.taskInfoCache.SetTaskInfo(ctx, taskId, task); err != nil {
			logger.Warn("set task info to cache failed",
				zap.Error(err), zap.Int64("user_id", userId), zap.Int64("task_id", taskId))
		}
	}

	plan, err := s.planRepo.GetPlanByTaskId(ctx, taskId)
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	reviewTs := timeutils.TimestampLag(now, int64(plan.CycleSize*24*60*60*1000))

	taskRecordCount, err := s.taskCardRecordRepo.CountRecordByTimestamp(ctx, taskId, now)
	if err != nil {
		return nil, err
	}
	bookCardCount, err := s.bookCardRepo.CountByBookId(ctx, task.BookId)
	if err != nil {
		return nil, err
	}

	if taskRecordCount != bookCardCount {
		logger.Warn("not enough taskCardRecord, generate taskCardRecord", zap.Int64("user_id", userId), zap.Int64("task_id", taskId))
		if err := s.taskCardRecordRepo.Generate(ctx, taskId, userId, now); err != nil {
			return nil, err
		}
	}

	learnedGroupNum, _ := s.taskInfoCache.GetLearnedGroup(ctx, userId, taskId)

	var taskCardRecordList []*model.TaskCardRecord
	if learnedGroupNum*int64(plan.GroupSize) < taskRecordCount {
		taskCardRecordList, err = s.taskCardRecordRepo.GetRecordByTimestamp(ctx, taskId, now, learnedGroupNum+1, int64(plan.GroupSize))
		if err != nil {
			return nil, err
		}
		s.taskInfoCache.SetLearnedGroup(ctx, userId, taskId, learnedGroupNum+1)
	} else {
		reviewedGroupNum, _ := s.taskInfoCache.GetReviewedGroup(ctx, userId, taskId)
		taskCardRecordList, err = s.taskCardRecordRepo.GetRecordByTimestamp(ctx, taskId, reviewTs, reviewedGroupNum+1, int64(plan.GroupSize))
		if err != nil {
			return nil, err
		}
		s.taskInfoCache.SetReviewedGroup(ctx, userId, taskId, reviewedGroupNum+1)
	}

	return taskCardRecordList, nil
}

func (s *Service) AddLearnStatus(ctx context.Context, taskCardRecordId int64, status string) error {
	return s.taskLogRepo.AddLearnStatus(ctx, taskCardRecordId, status)
}

func (s *Service) Update(ctx context.Context, taskId int64, taskName string, isActive bool) (task *model.Task, err error) {
	ctx, persist := db.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	task, err = s.taskRepo.UpdateTaskName(ctx, taskId, taskName)
	if err != nil {
		return nil, err
	}

	task, err = s.taskRepo.Active(ctx, taskId, isActive)
	if err != nil {
		return nil, err
	}

	return task, nil
}

var _ IService = (*Service)(nil)
