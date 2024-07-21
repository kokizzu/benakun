package model

const (
	COA_LABEL_PRODUCT = `product`
	COA_LABEL_PRODUCTS = `products`
)

var COA_LABEL_LiST []string = []string{
	COA_LABEL_PRODUCT, COA_LABEL_PRODUCTS,
}

func IsValidCoaLabel(label string) bool {
	switch label {
	case COA_LABEL_PRODUCT:
		return true
	case COA_LABEL_PRODUCTS:
		return true
	default:
		return false
	}
}