package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/secret-dev87/golang-ethereum-api/pkg/blockchain"
)

func GetCustodianWalletLogicAddress(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection

	address := conn.GetLogicAddress()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"address": address.String(),
	})
}

func GetEscrowAddress(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection

	address := conn.GetEscrowAddress()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"address": address.String(),
	})
}
