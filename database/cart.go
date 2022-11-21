package database

import "errors"

 var(
	ErrCantFindProduct = errors.New("can't find the product")
	ErrCantDecodeProdducts= errors.New("cant find the product")
	ErrUserIdIsNotValid= errors.New("this user is not valid")
	ErrCantUpdateUser= errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart= errors.New("cant remove item from cart")
	ErrCantGetItem= errors.New("cannot get item from cart")
	ErrCantBuyCartItem= errors.New("cannot update the purchase")
 )


 func AddProductToCart(){}
 func RemoveCartItem(){}
 func BuyItemFromCart(){}
 func InstantBuyer(){}