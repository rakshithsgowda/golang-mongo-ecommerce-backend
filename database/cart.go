package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/rakshithsgowda/golang-mongo-ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("cant find the product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("cant remove item from cart")
	ErrCantGetItem        = errors.New("cannot get item from cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddProductToCart(ctx context.Context, prodCollection *mongo.Collection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {

	searchfromdb, err := prodCollection.Find(ctx, bson.M{"_id": productID})
	if err != nil {
		log.Println(err)
		return ErrCantFindProduct
	}
	var productcart []models.ProductUser
	err = searchfromdb.All(ctx, &productcart)
	if err != nil {
		log.Println(err)
		return ErrCantDecodeProducts
	}

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "usercart", Value: bson.D{{Key: "$each", Value: productcart}}}}}}
	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ErrCantUpdateUser
	}
	return nil
}

// jsut deletes a particular product for the particular user
func RemoveCartItem(ctx context.Context, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}
	// come back
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.M{"$pull": bson.M{"usercart": bson.M{"_id": productID}}}

	_, err = userCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return ErrCantRemoveItemCart
	}
	return nil
}

// to buy get the orders from the usercart
// userid
// price calculation
// product id 
// edit product cart
// create an order
// after buy also empty the cart.
func BuyItemFromCart(ctx context.Context ,userCollection *mongo.Collection, userID string)error {
		id,err:=primitive.ObjectIDFromHex(userID)
		if err!=nil{
			fmt.Println(err)
			return ErrUserIdIsNotValid
		}
		// parameters are cartitems ordercart struct
		var getcartitems models.User
		var ordercart models.Order
			ordercart.Order_ID     =primitive.NewObjectID()
			ordercart.Order_Cart     = make([]models.ProductUser, 0)
			ordercart.Ordered_At     =time.Now()
			ordercart.Payment_Method.COD=true

		// agregate all the data for products in the cart for the user with product id.
		unwind:=bson.D{{Key: "$unwind",Value: bson.D{primitive.E{Key: "path",Value: "$usercart"}}}}
		grouping:=bson.D{{Key: "$group",Value: bson.D{primitive.E{Key: "_id",Value:"$_id"},{Key: "total",Value: bson.D{primitive.E{Key: "$sum",Value: "$usercart.price"}}}}}}

	currentresults,err:=userCollection.Aggregate(ctx,mongo.Pipeline{unwind,grouping})
	ctx.Done()
	if err!=nil{
		panic(err)
	}
	var getusercart []bson.M
	if err = currentresults.All(ctx,&getusercart);err!=nil{
		panic(err)
	}
// get all the price
	var total_price int32
	for _, user_item := range getusercart {
		price := user_item["total"]
		total_price = price.(int32)
	}
	ordercart.Price = int(total_price)
	filter:=bson.D{primitive.E{Key: "_id",Value: id}}
	update:=bson.D{{Key: "$push",Value: bson.D{primitive.E{Key: "orders",Value: ordercart}}}}
	_ , err= userCollection.UpdateMany(ctx, filter, update)
		if err!=nil{
			log.Println(err)
		}
// identify the useritems in the cart
		err = userCollection.FindOne(ctx,bson.D{primitive.E{Key:"_id",Value:id}}).Decode(&getcartitems)
		if err!=nil{
			log.Println(err)
		}

		// update 
		filter2:=bson.D{primitive.E{Key:"_id",Value: id}}
		update2:=bson.M{"$push":bson.M{"order.$[].order_list":bson.M{"$each":getcartitems.UserCart}}}
		_ , err = userCollection.UpdateOne(ctx, filter2, update2)
		if err!=nil{
			log.Println(err)
		}

		// after buy prcocess emptty the cart
		userempty_cart:= make([]models.ProductUser, 0)
		filteredcart:=bson.D{primitive.E{Key:"_id",Value:id}}
		updatedcart:=bson.D{{Key:"$set",Value:bson.D{primitive.E{Key:"usercart",Value:userempty_cart}}}}
		_, err = userCollection.UpdateOne(ctx,filteredcart,updatedcart);
		if err!=nil{
		return ErrCantBuyCartItem
				}

		return nil
}


// for instant buy check for address
func InstantBuyer(ctx context.Context,prodCollection *mongo.Collection, userCollection *mongo.Collection,productID primitive.ObjectID,UserID string)  error  {
	id,err := primitive.ObjectIDFromHex(UserID)
	if err!=nil{
		log.Println(err)
		return ErrUserIdIsNotValid
	}
	var product_details models.ProductUser
	var orders_details models.Order

	orders_details.Order_ID=primitive.NewObjectID()
	orders_details.Ordered_At=time.Now()
	orders_details.Order_Cart=make([]models.ProductUser, 0)
	orders_details.Payment_Method.COD=true

	err=prodCollection.FindOne(ctx,bson.D{primitive.E{Key: "_id",Value: productID}}).Decode(&product_details)
	if err!=nil{
		log.Println(err)
	}
	orders_details.Price=product_details.Price
	filter:=bson.D{primitive.E{Key: "_id",Value: id}}
	update:=bson.D{{Key: "$push",Value: bson.D{primitive.E{Key: "orders",Value: orders_details}}}}
	_,err= userCollection.UpdateOne(ctx,filter,update)
	if err!=nil{
		log.Println(err)
	}
	filter2:=bson.D{primitive.E{Key:"_id",Value: id}}
	update2:=bson.M{"$push":bson.M{"orders.$[].order_list":product_details}}
	_,err=userCollection.UpdateOne(ctx,filter2,update2)
	if err!=nil{
		log.Println(err)
	}
	return nil
}
