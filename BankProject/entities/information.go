package entities

type Information struct {
	AccountNumber   int64  `json:"accountnumber"`
	Name            string `json:"Name"`
	City            string `json:"city"`
	Pincode         int    `json:"pincode"`
	DateofBirth     string `json:"DateofBirth"`
	FatherName      string `json:"fathersname"`
	MotherName      string `json:"mothersname"`
	AadharNumber    uint64 `json:"aadharnumber"`
	PanNumber       uint64 `json:"pannumber"`
	AccountBalance  int64  `json:"accountbalance"`
	Credit          int32  `json:"credit"`
	Debit           int32  `json:"debit"`
	TransactionTime string `json:"transactiontime"`
}
