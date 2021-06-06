package dictionary

var (
	FailToQueryProductList   = "Fail to query product list"
	FailToMarshalProducts    = "Fail to marshal products"
	FailToQueryProductByUUID = "Fail to query product by uuid"
	FailToMarshalProduct     = "Fail to marshal product"
	FailToQuerySellerByUUID  = "Fail to query seller by UUID"
	SellerIsNotFound         = "Seller is not found"
	FailToInsertProduct      = "Fail to insert product"
	ProductIsNotFound        = "Product is not found"
	FailToDeleteProduct      = "Fail to delete product"

	FailToQuerySellerList     = "Fail to query seller list"
	FailToMarshalSellers      = "Fail to marshal sellers"
	FailToQueryTopNSellers    = "Fail to query %v sellers"
	FailToMarshalTopNSellers  = "Fail to marshal top %v sellers"
	TopSellerMustANumber      = "Top seller number must be a number"
	TopSellerMustGreaterThan0 = "Top seller number must greater than 0"

	SendNotificationError = "Send notification error"
)
