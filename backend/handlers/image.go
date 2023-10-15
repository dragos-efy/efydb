package handlers

import (
	"fmt"

	"github.com/efydb/config"
	"github.com/gofiber/fiber/v2"
	"github.com/h2non/bimg"
)

func GetImageThumb(c *fiber.Ctx) error {
	filename := c.Params("filename")
	buffer, err := bimg.Read(fmt.Sprintf("%s/files/%s", config.RootDir(), filename))
	if err != nil {
		return err
	}
	image, err := bimg.NewImage(buffer).Resize(320, 180)
	if err != nil {
		return err
	}
	return c.Send(image)
}
