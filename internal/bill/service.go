package bill

import (
	"github.com/jinzhu/gorm"
)

type BillService interface {
	GetBill(ID uint) (Bill, error)
	GetBillByPayer(payer string) (Bill, error)
	AddBill(bill Bill) (Bill, error)
	UpdateBill(ID uint, newBill Bill) (Bill, error)
	DeleteBill(ID uint) error
	GetAllBills() ([]Bill, error)
}

type Service struct {
	DB *gorm.DB
}

// Returns a new Bill service
func New(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetBill(ID uint) (Bill, error) {
	var bill Bill
	result := s.DB.First(&bill, ID)
	return bill, result.Error
}

func (s *Service) GetBillByPayer(payer string) ([]Bill, error) {
	var bills []Bill
	result := s.DB.Find(&bills).Where("payer = ?", payer)
	return bills, result.Error
}

func (s *Service) AddBill(bill Bill) (Bill, error) {
	result := s.DB.Save(&bill)
	return bill, result.Error
}

func (s *Service) UpdateBill(ID uint, newBill Bill) (Bill, error) {
	bill, err := s.GetBill(ID)
	if err != nil {
		return Bill{}, err
	}

	result := s.DB.Model(&bill).Updates(newBill)
	return bill, result.Error
}

func (s *Service) DeleteBill(ID uint) error {
	result := s.DB.Delete(&Bill{}, ID)
	return result.Error
}

func (s *Service) GetAllBills() ([]Bill, error) {
	var bills []Bill
	result := s.DB.Find(&bills)
	return bills, result.Error
}
