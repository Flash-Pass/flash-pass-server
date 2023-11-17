package bookcard

import (
	"github.com/Flash-Pass/flash-pass-server/db/query"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	bookCard query.IBookCardDo
}

type IRepository interface {
	CountByBookId(ctx *gin.Context, bookId int64) (int64, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		bookCard: query.BookCard.WithContext(db.Statement.Context),
	}
}

func (r *Repository) CountByBookId(ctx *gin.Context, bookId int64) (int64, error) {
	logger := ctxlog.GetLogger(ctx)

	count, err := r.bookCard.WithContext(ctx).Where(query.BookCard.BookId.Eq(bookId)).Count()
	if err != nil {
		logger.Error("count book card by book id defeat", zap.Error(err), zap.Int64("book_id", bookId))
		return 0, err
	}

	return count, nil
}

var _ IRepository = (*Repository)(nil)
