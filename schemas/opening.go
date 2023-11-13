package schemas

import (
	"time"
)

type Openings struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Role      string
	Company   string
	Location  string
	Remote    bool
	Link      string
	Salary    int64
}

// type OpeningResponse struct {
// 	ID        uint      `json:"id"`
// 	CreatedAt time.Time `json:"createdAt"`
// 	UpdatedAt time.Time `json:"updatedAt"`
// 	DeletedAt time.Time `json:"deteledAt,omitempty"`
// 	Role      string    `json:"role"`
// 	Company   string    `json:"company"`
// 	Location  string    `json:"location"`
// 	Remote    bool      `json:"remote"`
// 	Link      string    `json:"link"`
// 	Salary    int64     `json:"salary"`
// }
