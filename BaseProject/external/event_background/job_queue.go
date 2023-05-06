package event_background

import (
	"context"
	"log"
)

type backGroundJobs struct {
	Group chan *group
}

var bgJobs *backGroundJobs

func GetBackGroundJobs() *backGroundJobs {
	if bgJobs == nil {
		bgJobs = &backGroundJobs{
			Group: make(chan *group),
		}
		go bgJobs.Run()
	}
	return bgJobs
}
func (b *backGroundJobs) Run() {
	for {
		select {
		case g := <-b.Group:
			go func() {
				err := g.Run(context.Background())
				if err != nil {
					log.Println("--------------------------")
					log.Println("Can not do job due to: ", err)
					log.Println("--------------------------")
				}
			}()
		}
	}
}
