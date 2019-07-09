package transaction

import (
	"time"
)

type Transaction struct {
	Title string
	Amount float64
	DateFulfilled time.Time
	DateCreated time.Time
	DateUpdated time.Time
}
