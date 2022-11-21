package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rakshithsgowda/e-commerce-golang-practise/database"
	"github.com/rakshithsgowda/e-commerce-golang-practise/models"
	generate "github.com/rakshithsgowda/e-commerce-golang-practise/tokens"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection= database.UserData(database.Client,"Users")
var validate= validator.New()


// basic controllers for user
// SignUp
// Login
// password hash and helpers

func HashPassword(password string) string {

}
func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

// take in the input values from the params/bind the json.
// validate USER-Model and database check 
// check if the user is already present
// also phone number to required check
// if not use the modelfrom db used to Create the user 
// initialize all the parameters requirted for the user object like the 
// hash the password 
// sign the tokens
// create the timestamps and retrurn a user signup succesfull message to response 
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx,cancel=context.WithTimeout(context.Background(),100*time.Second)
		defer cancel()

		var user models.User

		if err:=c.BindJSON(&user);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		// validate user data type from struct
		validationErr:= validate.Struct(user)
		if validationErr!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":validationErr})
			return
		}

	
		// cheack the USer EMail if it already exists
		count,err:=UserCollection.CountDocuments(ctx,bson.M{"email":user.Email})
		if err!=nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError,gin.H{"error":err})
			return
		}
		if count>0{
			c.JSON(http.StatusBadRequest,gin.H{"error":"User already exits"})
			return
		}
	// phone number check
		count,err=UserCollection.CountDocuments(ctx,bson.M{"phone":user.Phone})
		defer cancel()
		if err!=nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError,gin.H{"error":err})
			return
		}
		if count>0{
			c.JSON(http.StatusBadRequest,gin.H{"email":"Phone is already in use"})
			return
		}

		//  hashing password
		password:=HashPassword(*user.Password)
		user.Password=&password


		// What will the user object have as model parametrs
		user.Created_At,_=time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		user.Updated_At,_=time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID= user.ID.Hex()

		// tokens
		token,refreshtoken,_:=generate.TokenGenerator(*user.Email,)

		user.Token=&token
		user.Refresh_Token=&refreshtoken

		user.UserCart= make([]models.ProductUser, 0)
		user.Address_Details= make([]models.Address, 0)
		user.Order_Status= make([]models.Order,0)


		// insert to user collection
		_,inserterr:=UserCollection.InsertOne(ctx,user)
		if inserterr!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"user not created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated,"Successfully signed Up!")
}
}



func Login() gin.HandlerFunc {

}

func ProductViewerAdmin() gin.HandlerFunc    {}
func SearchProduct() gin.HandlerFunc         {}
func SearchProductsByQuery() gin.HandlerFunc {}
