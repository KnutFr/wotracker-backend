package service

import (
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"strconv"
	"time"
	health "wotracker-back/internal/app/api/models"
)

type HealthService struct {
	db *bun.DB
}

func NewHealthService(db *bun.DB) *HealthService {
	if db == nil {
		log.Fatalf("db not initialized")
	}
	return &HealthService{db: db}
}
func (s *HealthService) GetHealthService(ctx iris.Context) health.Health {
	begin := ctx.Values().GetString("begin")
	beginFullTime, _ := time.Parse(time.UnixDate, begin)
	elapsedTime := time.Since(beginFullTime)

	var rnd float64
	beginDb := time.Now()
	if err := s.db.NewSelect().ColumnExpr("random()").Scan(ctx, &rnd); err != nil {
		log.Fatalf("cannot run db request: %s", err)
	}
	elapsedTimeDb := time.Since(beginDb)
	myHealth := health.Health{
		Code:           200,
		ResponseTime:   strconv.FormatInt(int64(elapsedTime.Milliseconds()), 10) + "ms",
		DbResponseTime: strconv.FormatInt(int64(elapsedTimeDb.Milliseconds()), 10) + "ms",
	}
	return myHealth
}
