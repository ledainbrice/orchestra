package config

import (
	"github.com/ivpusic/grpool"
)

var pool *grpool.Pool = grpool.NewPool(100, 50)

type Task func()

func AddTask(foo Task) {
	pool.JobQueue <- func() {
		foo()
	}
}
