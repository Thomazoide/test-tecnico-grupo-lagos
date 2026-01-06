package entities

import (
	"encoding/json"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID   uint           `gorm:"uniqueIndex"`
	Barcodes datatypes.JSON `gorm:"type:json"`
}

func (c *Cart) GetBarcodes() ([]string, error) {
	if len(c.Barcodes) == 0 {
		return []string{}, nil
	}
	var out []string
	if err := json.Unmarshal(c.Barcodes, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Cart) SetBarcodes(barcodes []string) error {
	b, err := json.Marshal(barcodes)
	if err != nil {
		return err
	}
	c.Barcodes = datatypes.JSON(b)
	return nil
}
