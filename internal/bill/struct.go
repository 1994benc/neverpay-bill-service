package bill

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Bill struct {
	gorm.Model
	Price float64
	Payer string
}

func (b *Bill) FromJSON(body io.ReadCloser) error {
	err := json.NewDecoder(body).Decode(b)
	return err
}

func (b *Bill) ToJSON(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(b)
}
