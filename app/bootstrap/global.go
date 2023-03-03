//go:build !test

package bootstrap

import (
	"context"

	"github.com/ZNotify/server/app/db/dao"
)

func initializeGlobalVar() {
	dao.SequenceID.Store(int64(dao.GetLatestMessageID(context.Background())))
}
