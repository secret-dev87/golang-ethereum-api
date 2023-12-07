package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/secret-dev87/golang-ethereum-api/pkg/blockchain"
)

func NewBuyOrder(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection
	request := new(blockchain.Order)
	wallet := c.Params("address")

	blockchain.WalletAddress = wallet

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Malformed data",
			"data":    err,
		})
	}

	trx, err := conn.OpenBuyOrder(*request)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   trx.Hash(),
	})
}

func GetWalletUSDCBalance(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection
	walletAddress := c.Params("address")

	blockchain.WalletAddress = walletAddress

	balance, err := conn.GetUSDCBalance()

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   balance,
	})
}
