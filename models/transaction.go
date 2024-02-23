package models

import "time"

// Transaction represents the data structure for a transaction
type Transaction struct {
	TransactionID   string    `json:"transactionID"`
	SenderID        string    `json:"senderID"`
	ReceiverID      string    `json:"receiverID"`
	CampaignID      string    `json:"campaignID"`
	Amount          int64     `json:"amount"`
	TransactionTime time.Time `json:"transactionTime"`
}
