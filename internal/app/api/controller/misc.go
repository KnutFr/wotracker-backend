package controller

import (
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	misc "wotracker-back/internal/app/api/service"
)

type MiscController struct {
	healthService *misc.HealthService
}

func (c *MiscController) NewHealthController(healthService *misc.HealthService) {
	if healthService == nil {
		log.Fatalf("cannot instanciate service")
	}
}

// GetHealth     godoc
// @Summary      Health
// @Description  get health info
// @Tags         misc
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Health
// @Security     None
// @Router       /misc/health [get]
func (c *MiscController) GetHealth(ctx iris.Context) {
	myHealth := c.healthService.GetHealthService(ctx)
	_, err := ctx.JSON(myHealth)
	if err != nil {
		log.Errorf("cannot encode health: %s", err)
	}

}
