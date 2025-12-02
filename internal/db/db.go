package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/oita-u/campus-api/internal/config"
	"github.com/oita-u/campus-api/internal/logger"
	"go.uber.org/zap"
)

var Conn *sqlx.DB

func Init() {
	var err error

	Conn, err = sqlx.Connect("postgres", config.C.DBUrl)
	if err != nil {
		logger.Get().Fatal("Failed to connect to database", zap.Error(err))
	}
	logger.Get().Info("Connected to database")
}
