package product

import (
	"github.com/intersights/intersights-go"
)

const (
	//StatusInterest - probable interest in product
	StatusInterest intersights.ProductStatus = "interest"
	//StatusInterestRegistered - registered an interest
	StatusInterestRegistered intersights.ProductStatus = "interest.reg"
	//StatusPreOrder - pre-order
	StatusPreOrder intersights.ProductStatus = "preorder"
	//StatusCheckout - In checkout process
	StatusCheckout intersights.ProductStatus = "checkout"
	//StatusOrdered - Ordered Product
	StatusOrdered intersights.ProductStatus = "ordered"
	//StatusProvisioning - Provisioning a product
	StatusProvisioning intersights.ProductStatus = "provisioning"
	//StatusShipping - Shipping product
	StatusShipping intersights.ProductStatus = "shipping"
	//StatusDelivered - Product delivered to customer
	StatusDelivered intersights.ProductStatus = "delivered"
	//StatusReturned - Product returned by customer
	StatusReturned intersights.ProductStatus = "returned"
	//StatusActive - actively using the product
	StatusActive intersights.ProductStatus = "active"
	//StatusSuspended - use of product suspended
	StatusSuspended intersights.ProductStatus = "suspended"
	//StatusCancelled - Cancelled, may still be active
	StatusCancelled intersights.ProductStatus = "cancelled"
	//StatusRefunded
	StatusRefunded intersights.ProductStatus = "refunded"
	//StatusEnded - No longer in use
	StatusEnded intersights.ProductStatus = "ended"

	TypeService intersights.ProductType = "service"
	TypeProduct intersights.ProductType = "product"
)
