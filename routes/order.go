package routes

import (
	"errors"

	"github.com/atifali-pm/go-fiber-api/database"
	"github.com/atifali-pm/go-fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID      uint    `json:"id"`
	User    User    `json:"user"`
	Product Product `json:"product"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product}
}

func CreateOrder(c *fiber.Ctx) error {

	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product

	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)

}

func FindOrder(id int, order *models.Order) error {

	database.Database.DB.Find(&order, "id = ?", id)

	if order.ID == 0 {
		return errors.New("order doesnot exist")
	}

	return nil

}

func GetOrders(c *fiber.Ctx) error {

	orders := []models.Order{}

	database.Database.DB.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders {

		var user models.User
		var product models.Product

		database.Database.DB.Find(&user, "id = ?", order.UserRefer)
		database.Database.DB.Find(&product, "id = ?", order.ProductRefer)

		responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(200).JSON(responseOrders)

}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please ensure the :id is an integer")
	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product

	database.Database.DB.First(&user, order.UserRefer)
	database.Database.DB.First(&product, order.ProductRefer)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)

	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)

}
