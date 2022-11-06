package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rakshithsgowda/e-commerce-golang-practise/controllers"
	"github.com/rakshithsgowda/e-commerce-golang-practise/database"
	"github.com/rakshithsgowda/e-commerce-golang-practise/routes"
)


func main()  {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client,"Products"), database.UserData(database.Client,"Users"))

	router := gin.New()
	router.Use(gin.Logger())

			routes.UserRoutes(router)
			router.Use(middleware.Authentication())

	router.GET("/addtocart",app.Addtocart())
router.GET("/removeitem",app.RemoveItem())
router.GET("/cartcheckout",app.BuyFromCart())
router.GET("/instantbuy",app.InstantBuy())

log.Fatal(router.Run(":"+port))
}