package product

import "github.com/intersight/intersight-go"

const (
	//StatusInterest - probable interest in product
	StatusInterest intersight.ProductStatus = "interest"
	//StatusInterestRegistered - registered an interest
	StatusInterestRegistered intersight.ProductStatus = "interest.reg"
	//StatusPreOrder - pre-order
	StatusPreOrder intersight.ProductStatus = "preorder"
	//StatusCheckout - In checkout process
	StatusCheckout intersight.ProductStatus = "checkout"
	//StatusOrdered - Ordered Product
	StatusOrdered intersight.ProductStatus = "ordered"
	//StatusProvisioning - Provisioning a product
	StatusProvisioning intersight.ProductStatus = "provisioning"
	//StatusShipping - Shipping product
	StatusShipping intersight.ProductStatus = "shipping"
	//StatusDelivered - Product delivered to customer
	StatusDelivered intersight.ProductStatus = "delivered"
	//StatusReturned - Product returned by customer
	StatusReturned intersight.ProductStatus = "returned"
	//StatusActive - actively using the product
	StatusActive intersight.ProductStatus = "active"
	//StatusSuspended - use of product suspended
	StatusSuspended intersight.ProductStatus = "suspended"
	//StatusCancelled - Cancelled, may still be active
	StatusCancelled intersight.ProductStatus = "cancelled"
	//StatusRefunded
	StatusRefunded intersight.ProductStatus = "refunded"
	//StatusEnded - No longer in use
	StatusEnded intersight.ProductStatus = "ended"

	TypeService intersight.ProductType = "service"
	TypeProduct intersight.ProductType = "product"
)
