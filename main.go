package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", getHome)
	app.Get("/about", getAbout)
	app.Get("/services", getServices)
	app.Get("/gallery", getGallery)
	app.Get("/team", getTeam)
	app.Get("/careers", getCareers)
	app.Get("/contact-us", getContactUs)

	log.Fatal(app.Listen("127.0.0.1:7777"))
}

// FUNCTION FOR GET HOME
func getHome(c *fiber.Ctx) error {
	htmlContent, err := os.ReadFile("./html_text/home.html")
	if err != nil {
		log.Fatalf("Error reading HTML file: %v", err)
	}

	return c.Type("html").SendString(string(htmlContent))
}

// FUNCTION FOR GET ABOUT
func getAbout(c *fiber.Ctx) error {
	htmlContent, err := os.ReadFile("./html_text/about.html")
	if err != nil {
		log.Fatalf("Error reading HTML file: %v", err)
	}

	return c.Type("html").SendString(string(htmlContent))
}

// FUNCTION FOR GET SERVICES
func getServices(c *fiber.Ctx) error {
	htmlContent, err := os.ReadFile("./html_text/services.html")
	if err != nil {
		log.Fatalf("Error reading HTML file: %v", err)
	}

	return c.Type("html").SendString(string(htmlContent))
}

// FUNCTION FOR GET GALLERY
func getGallery(c *fiber.Ctx) error {
	htmlContent, err := os.ReadFile("./html_text/gallery.html")
	if err != nil {
		log.Fatalf("Error reading HTML file: %v", err)
	}

	return c.Type("html").SendString(string(htmlContent))
}

// FUNCTION FOR GET TEAM
func getTeam(c *fiber.Ctx) error {
	htmlContent, err := os.ReadFile("./html_text/team.html")
	if err != nil {
		log.Fatalf("Error reading HTML file: %v", err)
	}

	return c.Type("html").SendString(string(htmlContent))
}

// FUNCTION FOR GET CAREERS
func getCareers(c *fiber.Ctx) error {
	htmlContent, err := os.ReadFile("./html_text/careers.html")
	if err != nil {
		log.Fatalf("Error reading HTML file: %v", err)
	}

	return c.Type("html").SendString(string(htmlContent))
}

// FUNCTION FOR GET CONTACT US
func getContactUs(c *fiber.Ctx) error {
	htmlContent, err := os.ReadFile("./html_text/contact_us.html")
	if err != nil {
		log.Fatalf("Error reading HTML file: %v", err)
	}

	return c.Type("html").SendString(string(htmlContent))
}
