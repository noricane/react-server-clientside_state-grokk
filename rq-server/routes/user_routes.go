package routes

import (
	"fmt"
	"rq-server/services"
	//"time"
	"github.com/gofiber/fiber/v2"
)

// Hold state within this variable, I'm grokking react-query. No need for persistence.
// Json data starts at id=1 so it is alright that this initializes to 0
var lastCarClickedId int

func GetLastCarClickedId(c *fiber.Ctx) error {
	fmt.Printf("Get Last Clicked\n")
	//Timer just to test loading state of component etc.
	//time.Sleep(5 * time.Second)
	if lastCarClickedId == 0 {
		/* 
		425 Too Early is more appropriate here, but I
		will keep teapot for the lols */
		return c.SendStatus(fiber.StatusTeapot)
	}
	res, status, err := services.GetCar(lastCarClickedId)
	if err != nil {
		fmt.Printf("ERROR HERE\n")
		return c.Status(status).JSON(err.Error())
	}
	return c.Status(status).JSON(res)
}

func SetLastCarClickedId(c *fiber.Ctx) error {
	fmt.Printf("Set Last Clicked\n")
	payload := struct {
		CarId int `json:"car_id"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if !services.ValidCarId(payload.CarId) {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	lastCarClickedId = payload.CarId
	return c.Status(fiber.StatusOK).JSON(true)
}
