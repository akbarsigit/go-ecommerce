package database

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	// adding custom error for each fail use case
	ErrCantFindProduct = errors.New("cant find the product")
	ErrCantDecodeProduct = errors.New("cant findd the product")
	ErrUserIdIsNotValid = errors.New("this user is not valid")
	ErrCantUpdateUser = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCantGetItem = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem = errors.New("cannot update the purchase")
)

func AddProductToCart() gin.HandlerFunc {

}

func RemoveCartItem() gin.HandlerFunc {

}

func BuyItemFromCart() gin.HandlerFunc {

}

func InstantBuyer() gin.HandlerFunc {

}