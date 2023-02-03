package wrap_gorm

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

const ctLayout = "02-01-2006 15:04:05 Z07:00"

type DateOnly time.Time

type SoftDeleteModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt DateOnly       `json:"created_at"`
	UpdatedAt DateOnly       `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
type HardDeleteModel struct {
	ID        uint     `gorm:"primarykey" json:"id"`
	CreatedAt DateOnly `json:"created_at"`
	UpdatedAt DateOnly `json:"updated_at"`
}

// UnmarshalJSON Parses the json string in the custom format
func (ct *DateOnly) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(ctLayout, s)
	*ct = DateOnly(nt)
	return
}

// MarshalJSON writes a quoted string in the custom format
func (ct DateOnly) MarshalJSON() ([]byte, error) {
	date := ct.String()
	if len(date) > 0 {
		date = strings.Split(strings.Trim(ct.String(), `"`), " ")[0]
	}
	return json.Marshal(date)
}

// String returns the time in the custom format
func (ct *DateOnly) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(ctLayout))
}

func StringToDate(s string) (*time.Time, error) {
	date := strings.Split(s, "-")
	year, err := strconv.Atoi(date[0])
	if err != nil {
		return nil, err
	}
	month, err := strconv.Atoi(date[1])
	if err != nil {
		return nil, err
	}
	day, err := strconv.Atoi(date[2])
	if err != nil {
		return nil, err
	}
	dateTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return &dateTime, nil
}
