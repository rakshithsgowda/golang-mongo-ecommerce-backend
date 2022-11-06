package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rakshithsgowda/e-commerce-golang-practise/controllers"
	"github.com/rakshithsgowda/e-commerce-golang-practise/database"
	"github.com/rakshithsgowda/e-commerce-golang-practise/middlewares"
	"github.com/rakshithsgowda/e-commerce-golang-practise/routes"
)


func main()  {
	port := os.Getenv("PORT")
	if port= ""{
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client,"Products"), database.UserData(database.Client,"Users"))

	router := gin.New()
	router.Use(gin.Logger())
}