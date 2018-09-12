package main

type AccOpArgs struct {
	AccID  string
	Amount float32
}

type TransferArgs struct {
	PayerID string
	PayeeID string
	Amount  float32
}
