package model

type Voucher struct {
	ID       int64  `json:"id"`
	Code     string `json:"code"`
	Discount int64  `json:"discount"`
}
