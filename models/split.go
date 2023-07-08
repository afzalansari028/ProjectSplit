package models

import "time"

type SplitBill struct {
	Amount    float32
	SpentOn   string
	SpentDate time.Time
}
