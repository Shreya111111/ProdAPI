package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var products = []Product{
	{ID: "1", Name: "Laptop", Price: 10.99, Quantity: 50},
	{ID: "2", Name: "Computer", Price: 24.99, Quantity: 25},
	{ID: "3", Name: "Charger", Price: 45.99, Quantity: 100},
	{ID: "4", Name: "Smartphone", Price: 800.0, Quantity: 1032},
	{ID: "5", Name: "Headphones", Price: 50.0, Quantity: 203},
	{ID: "6", Name: "Speaker", Price: 140.99, Quantity: 503},
	{ID: "7", Name: "Wiredphones", Price: 243.99, Quantity: 255},
	{ID: "8", Name: "DVD player", Price: 5.99, Quantity: 1005},
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func productById(c *gin.Context) {
	id := c.Param("id")
	product, err := getProductById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}

func purchaseProduct(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	product, err := getProductById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found."})
		return
	}

	if product.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Product not available."})
		return
	}

	product.Quantity -= 1
	c.IndentedJSON(http.StatusOK, product)
}

func restockProduct(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	product, err := getProductById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found."})
		return
	}

	product.Quantity += 10 // Restock by adding 10 units
	c.IndentedJSON(http.StatusOK, product)
}

func getProductById(id string) (*Product, error) {
	for i, p := range products {
		if p.ID == id {
			return &products[i], nil
		}
	}

	return nil, errors.New("product not found")
}

func createProduct(c *gin.Context) {
	var newProduct Product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", productById)
	router.POST("/products", createProduct)
	router.PATCH("/purchase", purchaseProduct)
	router.PATCH("/restock", restockProduct)
	router.Run("localhost:8080")
}
