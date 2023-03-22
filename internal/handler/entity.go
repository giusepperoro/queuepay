package handler

type WithdrawalRequest struct {
	Id     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

type WithdrawalResponse struct {
	Status string `json:"status"`
}

type HandleWithdrawal struct {
	ForWorkers map[int64]struct{}
}
