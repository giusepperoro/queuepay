package main

import "net/http"

type changeBalanceHandler interface {
	HandleBalanceChange(w http.ResponseWriter, r *http.Request)
}
