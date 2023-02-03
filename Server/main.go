// define and config my go server

package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// get the env port config :
	port := os.Getenv("PORT")
	// check if port is a valable port id :
	if port == "" {
		// define a custom port id
		port = "8080"
	}
	// define my router engine :
	router := gin.New()
	// use middelware :
	router.Use(gin.Logger())
	// config Cors policy :
	router.Use(cors.Default())
	// define server endpoint :
	/*
		// Create New Order:
		router.POST("/order/create", routers.AddOrder)
		// get List of orders by waiter:
		router.GET("/waiter/:waiter", routers.GetOrdersByWaitter)
		// get All orders :
		router.GET("/orders", routers.GetOrders)
		// get order By id :
		router.GET("/order/:id", routers.GetOrderById)
		// update :
		router.PUT("/waiter/update/:id", routers.UpdateWaiter)
		router.PUT("/order/update/:id", routers.UpdateOrder)
		// delete:
		router.DELETE("/order/delete/:id", routers.DeleteOrder)
	*/
	// run the server :
	router.Run(":" + port)
}
