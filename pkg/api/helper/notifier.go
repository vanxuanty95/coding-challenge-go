package helper

type NotificationsInfo struct {
	SellerUUID  string
	SellerPhone string
	SellerName  string
	SellerEmail string
	ProductName string
	OldStock    int
	NewStock    int
}

type Notifier interface {
	StockChanged(info NotificationsInfo) error
}
