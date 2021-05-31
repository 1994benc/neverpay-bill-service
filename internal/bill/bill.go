package bill

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

type Bill struct {
	gorm.Model
	Price float64
	Payer string
}

type BillService interface {
	GetBill(ID uint) (Bill, error)
	AddBill(bill Bill) (Bill, error)
	UpdateBill(ID uint, newBill Bill) (Bill, error)
	DeleteBill(ID uint) error
	GetAllBills() ([]Bill, error)
}

// Returns a new Bill service
func New(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
