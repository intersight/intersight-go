package intersights

type ProductStatus string
type ProductType string

type Product struct {
	ID          string
	Name        string
	Description string
	Code        string
	Type        ProductType
	//Recurring - can this product be purchased on a recurring cycle
	Recurring bool
	GroupID   string
}

type ProductFeature struct {
	ProductID   string
	ID          string
	Name        string
	Code        string
	Description string
}
