package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/secret-dev87/golang-ethereum-api/pkg/blockchain"
)

// Deployed to 0x1E98e069334d26Eb437913D146eB204FB0B543eA
func DeployUSDC(c *fiber.Ctx) error {

	conn := blockchain.CurrentConnection

	address, transaction := conn.DeployUSDC()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"address": address.String(),
		"hash":    transaction.Hash(),
	})
}

// Deployed to 0xc2Ce51479F3545d28c184400895A129d8c33c9cD
func DeployFactory(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection
	address, transaction := conn.DeployFactory()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"address": address.String(),
		"hash":    transaction.Hash(),
	})
}
