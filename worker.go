package main

type Work struct {
	db     DbManager
	idChan chan int
}

func NewWorker(db DbManager, idCh chan int) *Work {
	return &Work{
		db:     db,
		idChan: idCh,
	}
}
