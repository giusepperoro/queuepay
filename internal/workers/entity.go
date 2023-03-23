package workers

type worker struct {
	logger   logger
	database database
}

func NewWorker(logger logger, db database) worker {
	return worker{
		logger:   logger,
		database: db,
	}
}
