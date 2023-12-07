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

func NewWallet(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection
	uuid := c.Params("uuid")

	trx, err := conn.NewWallet(uuid)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"hash":   trx.Hash(),
	})
}

func GetWallet(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection
	uuid := c.Params("uuid")

	address, err := conn.GetAccountByUUID(uuid)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"hash":   address.String(),
	})
}

//UUID									| Account Address
//b93e42b0-33a2-11ed-a261-0242ac120002 	| 0xAaDCa4E6f0D96CddCE311Cb5b3027a5854e1C466
//cec3df26-339a-11ed-a261-0242ac120002 	| 0x77aB3945978b676B863aA5eBf2A45193b65F6dB8
//cec3dd14-339a-11ed-a261-0242ac120002 	| 0x254eDb0820191967EC1E439675f65F52505D725b
