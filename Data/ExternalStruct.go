package DATA

import "time"

// Customer struct to match the customer structure in the response
type Customer struct {
	ID             uint       `json:"ID"`
	CreatedAt      time.Time  `json:"CreatedAt"`
	UpdatedAt      time.Time  `json:"UpdatedAt"`
	DeletedAt      *time.Time `json:"DeletedAt,omitempty"`
	CustomerName   string     `json:"customer_name"`
	Phone          string     `json:"phone"`
	Email          string     `json:"email"`
	CardUID        string     `json:"card_uid"`
	CustomerPoints uint       `json:"customer_points"`
	Bill           []Bill     `json:"Bill"`
}

type Bill_Details struct {
	BillID    uint    `gorm:"index"`
	ProductID uint    `json:"product_id" gorm:"index"`
	Quantity  uint    `json:"quantity"`
	Total     uint    `json:"total"`
	Product   Product `gorm:"foreignKey:ProductID;"`
}

// Bill struct to match the bill details in the response
type Bill struct {
	ID          uint           `json:"ID"`
	CreatedAt   time.Time      `json:"CreatedAt"`
	UpdatedAt   time.Time      `json:"UpdatedAt"`
	DeletedAt   *time.Time     `json:"DeletedAt,omitempty"`
	CustomerID  uint           `json:"customer_id"`
	TotalAmount uint           `json:"total_amount"`
	BillDetails []Bill_Details `json:"Bill_Details" gorm:"foreignKey:BillID"`
}

type Product struct {
	ID             uint       `json:"ID"`
	CreatedAt      time.Time  `json:"CreatedAt"`
	UpdatedAt      time.Time  `json:"UpdatedAt"`
	DeletedAt      *time.Time `json:"DeletedAt,omitempty"`
	ProductBarcode string     `json:"product_barcode" gorm:"unique"`
	ProductName    string     `json:"product_name"`
	ImagePath      string     `json:"image_path"`
	Price          uint       `json:"price"`
	Stock          Stock      `gorm:"constraint:OnDelete:CASCADE;"`
}

type Stock struct {
	ProductID uint `gorm:"uniqueIndex;not null"`
	Quantity  uint `json:"quantity"`
}

// Response struct to match the outermost JSON structure
type CustomerResponse struct {
	Data   []Customer `json:"data"`
	Status string     `json:"status"`
}

type ProductResponse struct {
	Data   []Product `json:"data"`
	Status string    `json:"status"`
}
