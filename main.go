package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

var PORT, set = os.LookupEnv("PORT")

func init() {
	if !set {
		PORT = "8080"
	}
}

func main() {
	app := fiber.New(fiber.Config{BodyLimit: 1024 * 1024 * 1024})

	app.Use(cors.New())

	app.Use("/ocr/ws*", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ocr/ws/:id", websocket.New(func(c *websocket.Conn) {
		id := c.Params("id")
		if db[id] == nil {
			log.Printf("Channel does not exist for id %s\n", id)
			db[id] = make(chan string, 10)
		}
		db[id] <- "Uploading file"

		for message := range db[id] {
			c.WriteMessage(websocket.TextMessage, []byte(message))
		}
		c.WriteMessage(websocket.TextMessage, []byte("done"))
		c.Close()
	}))

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

			id := c.FormValue("id", "")
			if id == "" {
				return fiber.NewError(fiber.ErrBadGateway.Code, "no id sent")
			}

			if db[id] == nil {
				db[id] = make(chan string, 10)
			}
			db[id] <- "initializing...."

			c.SaveFile(file, os.TempDir()+"/"+file.Filename)

			performOCR(os.TempDir()+"/"+file.Filename, file.Filename, id)

			c.Redirect("/ocr/" + id + ".pdf")

		}

		return nil
	})

	app.Static("/ocr", "./output")

	log.Println("App listening on port :" + PORT)
	app.Listen(fmt.Sprintf(":%s", PORT))
}
