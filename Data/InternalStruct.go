package DATA

type Item struct {
	ItemID        int
	ItemName      string
	ItemPrice     int
	ItemStock     int
	ItemImagePath string
}

type Order struct {
	OrderItemID   int
	OrderItemName string
	OrderAmount   int
}
