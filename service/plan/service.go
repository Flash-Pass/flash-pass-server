package plan

import (
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/snowflake"
	"github.com/gin-gonic/gin"
)

type Service struct {
	planRepo        Repository
	snowflakeHandle snowflake.IHandle
}

type Repository interface {
	Create(ctx *gin.Context, plan *model.Plan) (*model.Plan, error)
	IsPlanBelongToUser(ctx *gin.Context, planId, userId uint64) (bool, error)
	Get(ctx *gin.Context, planId uint64) (*model.Plan, error)
	Update(ctx *gin.Context, plan *model.Plan) (*model.Plan, error)
	Delete(ctx *gin.Context, planId uint64) error
	GetList(ctx *gin.Context, userId uint64) ([]*model.Plan, error)
}

type IService interface {
	Create(ctx *gin.Context, plan *model.Plan) (*model.Plan, error)
	IsPlanBelongToUser(ctx *gin.Context, planId, userId uint64) (bool, error)
	Get(ctx *gin.Context, planId uint64) (*model.Plan, error)
	Update(ctx *gin.Context, plan *model.Plan) (*model.Plan, error)
	Delete(ctx *gin.Context, planId uint64) error
	GetList(ctx *gin.Context, userId uint64) ([]*model.Plan, error)
}

func NewService(planRepo Repository, snowflakeHandle snowflake.IHandle) *Service {
	return &Service{
		planRepo:        planRepo,
		snowflakeHandle: snowflakeHandle,
	}
}

func (s *Service) Create(ctx *gin.Context, plan *model.Plan) (*model.Plan, error) {
	plan.Id = s.snowflakeHandle.GetUInt64Id()
	return s.planRepo.Create(ctx, plan)
}

func (s *Service) IsPlanBelongToUser(ctx *gin.Context, planId, userId uint64) (bool, error) {
	return s.planRepo.IsPlanBelongToUser(ctx, planId, userId)
}

func (s *Service) Get(ctx *gin.Context, planId uint64) (*model.Plan, error) {
	return s.planRepo.Get(ctx, planId)
}

func (s *Service) Update(ctx *gin.Context, plan *model.Plan) (*model.Plan, error) {
	return s.planRepo.Update(ctx, plan)
}

func (s *Service) Delete(ctx *gin.Context, planId uint64) error {
	return s.planRepo.Delete(ctx, planId)
}

func (s *Service) GetList(ctx *gin.Context, userId uint64) ([]*model.Plan, error) {
	return s.planRepo.GetList(ctx, userId)
}

var _ IService = (*Service)(nil)
