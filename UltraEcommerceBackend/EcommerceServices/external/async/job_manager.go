package async

import (
	"context"
	"log"
	"sync"
)

type group struct {
	jobs         []Job
	isConcurrent bool
	wg           *sync.WaitGroup
}

func NewGroup(isConcurrent bool, jobs ...Job) *group {
	g := &group{
		isConcurrent: isConcurrent,
		jobs:         jobs,
		wg:           new(sync.WaitGroup),
	}

	return g
}

func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))

	errChan := make(chan error, len(g.jobs))

	var err error
	for i, _ := range g.jobs {

		//if g.isConcurrent {
		//	go func(aj Job) {
		//		errChan <- g.runJob(ctx, aj)
		//	}(g.jobs[i])
		//
		//	continue
		//}

		j := g.jobs[i]
		errChan <- g.runJob(ctx, j)
		//g.wg.Done()
	}

	g.wg.Wait()
	return err
}
func (g *group) runJob(ctx context.Context, j Job) error {
	if err := j.Execute(ctx); err != nil {
		for {
			log.Println(err)
			if j.State() == StateRetryFailed {
				return err
			}

			if j.Retry(ctx) != nil {
				return nil
			}
		}
	}

	return nil
}
