package handlers

import (
	"encoding/json"
	"log"

	"github.com/abraaoneves/channel/infra"
	"github.com/abraaoneves/channel/models"
	"github.com/gofiber/fiber/v2"
)

func RegisterChannelRoutes(app *fiber.App) {
	app.Post("/channels", CreateChannel)
	app.Get("/channels/:id", GetChannel)
}

/**
* Create an channel
**/
//TODO: Create validation and specific methods to create channel
func CreateChannel(context *fiber.Ctx) error {
	channel := new(models.Channel)

	if err := context.BodyParser(channel); err != nil {
		log.Fatal("Invalid content")
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	log.Printf("persist new channel: %v \n", channel)

	redis := infra.CreateDataBaseConnection()
	channelJson, err := json.Marshal(channel)

	if err != nil {
		log.Fatal("Error at parse to JSON")
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	err = redis.Set(channel.Key, channelJson, 0).Err()

	if err != nil {
		log.Fatal("Error to persist")
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Channel successfully created",
		"channel": channel,
	})
}

/**
* Return Channel by Ids
 */
func GetChannel(context *fiber.Ctx) error {
	id := context.Params("id")
	db := infra.CreateDataBaseConnection()
	channel, err := db.Get(id).Result()
	log.Println(channel)
	if err != nil {
		return context.SendStatus(fiber.StatusNotFound)
	}

	return context.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Channel found",
		"channel": channel,
	})
}
