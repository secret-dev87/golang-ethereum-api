package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/secret-dev87/golang-ethereum-api/app/controllers"
)

type ApiRouter struct {
}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api/v1", limiter.New())

	api.Get("/", controllers.RenderHello)

	deploy := api.Group("/deploy")
	deploy.Post("/usdc", controllers.DeployUSDC)
	deploy.Post("/factory", controllers.DeployFactory)

	api.Get("/wallet-logic-address", controllers.GetCustodianWalletLogicAddress)
	api.Get("/escrow-address", controllers.GetEscrowAddress)
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}
