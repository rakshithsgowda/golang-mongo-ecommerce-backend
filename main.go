package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rakshithsgowda/golang-mongo-ecommerce-backend/controllers"
	"github.com/rakshithsgowda/golang-mongo-ecommerce-backend/database"
	"github.com/rakshithsgowda/golang-mongo-ecommerce-backend/middlewares"
	"github.com/rakshithsgowda/golang-mongo-ecommerce-backend/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middlewares.Authentication())

	// direct homepage router actions
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	// after buy non authenticated user address input(CRUD)

	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddress", controllers.DeleteAddress())
	router.GET("/listcart", controllers.GetItemFromCart())
	log.Fatal(router.Run(":" + port))
}
