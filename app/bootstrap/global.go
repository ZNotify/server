//go:build !test

package bootstrap

import (
	"context"

	"notify-api/app/db/dao"
)

func initializeGlobalVar() {
	dao.SequenceID.Store(int64(dao.GetLatestMessageID(context.Background())))
}
