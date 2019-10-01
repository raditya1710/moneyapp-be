package transaction

import (
	"time"
)

type Transaction struct {
	title string
	amount float64
	dateFulfilled time.Time
	dateCreated time.Time
	dateUpdated time.Time
}
