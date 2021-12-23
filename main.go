package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
)

var PORT, set = os.LookupEnv("PORT")

func init() {
	if !set {
		PORT = "8080"
	}
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Post("/ocr", func(c *fiber.Ctx) error {
		multipartForm, err := c.MultipartForm()
		if err != nil {
			return err
		}

		files := multipartForm.File["file"]

		if len(files) == 0 {

			log.Println("no files uploaded by the user")
			return fiber.NewError(fiber.ErrBadRequest.Code, "no files uploaded")
		} else {

			//TODO: do something with the files
			file := files[0]

			//check if filename is not pdf
			if !strings.HasSuffix(file.Filename, ".pdf") {
				return fiber.NewError(fiber.ErrBadRequest.Code, "file is not pdf")
			}

			c.SaveFile(file, os.TempDir()+"/"+file.Filename)

			id := uuid.New()
			performOCR(os.TempDir()+"/"+file.Filename, file.Filename, id)

			c.Redirect("/ocr/" + id.String() + ".pdf")

		}

		return nil
	})

	app.Static("/ocr", "./output")

	log.Println("App listening on port :" + PORT)
	app.Listen(fmt.Sprintf(":%s", PORT))
}
