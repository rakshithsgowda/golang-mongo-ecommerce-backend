package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCantFindProduct     = errors.New("can't find the product")
	ErrCantDecodeProdducts = errors.New("cant find the product")
	ErrUserIdIsNotValid    = errors.New("this user is not valid")
	ErrCantUpdateUser      = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart  = errors.New("cant remove item from cart")
	ErrCantGetItem         = errors.New("cannot get item from cart")
	ErrCantBuyCartItem     = errors.New("cannot update the purchase")
)

func AddProductToCart(ctx context.Context, prodCollection *mongo.Collection, userCollection *mongo.Collection, productID primitive.ObjectID, userQueryID string) {
}
func RemoveCartItem(ctx context.Context, prodCollection *mongo.Collection, userCollection *mongo.Collection, productID primitive.ObjectID, userQueryID string) {
}
func BuyItemFromCart() {}
func InstantBuyer()    {}
