package db

import (
	"context"
	"github.com/Flash-Pass/flash-pass-server/db/query"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"go.uber.org/zap"
	"runtime/debug"
	"strings"
)

type txKey struct{}

func WithTx(ctx context.Context, tx *query.QueryTx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func Tx(ctx context.Context) *query.QueryTx {
	v := ctx.Value(txKey{})
	if v == nil {
		return nil
	}

	tx, ok := v.(*query.QueryTx)
	if !ok {
		return nil
	}
	return tx
}

func persist(ctx context.Context, err error) {
	tx := Tx(ctx)
	logger := ctxlog.Extract(ctx)
	caller := strings.Split(string(debug.Stack()), "\n")

	if err == nil {
		if err = tx.Commit(); err == nil {
			logger.Info("success commit transaction")
			return
		}

		logger.Error("failed commit transaction", zap.Error(err))
	}

	logger.Error("rolling back failed transaction", zap.Any("caller", caller), zap.Error(err))

	if rollBackErr := tx.Rollback(); rollBackErr != nil {
		logger.Error("failed rollback transaction", zap.Error(rollBackErr))
	}

	logger.Error("rolled back failed transaction", zap.Any("caller", caller), zap.Error(err))
}

func ReplaceTxPersist(ctx context.Context) (context.Context, func(err error)) {
	return withTxPersist(ctx, true)
}

func WithTXPersist(ctx context.Context) (context.Context, func(err error)) {
	return withTxPersist(ctx, false)
}

func withTxPersist(ctx context.Context, forceReplace bool) (context.Context, func(err error)) {
	tx := Tx(ctx)
	if tx != nil && !forceReplace {
		return ctx, func(err error) {
			logger := ctxlog.Extract(ctx)
			logger.Debug("commit or rollback will be done by parent", zap.Error(err))
		}
	}

	tx = query.Q.Begin()
	ctx = WithTx(ctx, tx)
	return ctx, func(ctx context.Context) func(err error) {
		return func(err error) {
			persist(ctx, err)
		}
	}(ctx)
}
