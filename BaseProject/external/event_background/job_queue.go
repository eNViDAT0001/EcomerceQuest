package event_background

import "context"

type backGroundJobs struct {
	Group chan *group
}

var bgJobs *backGroundJobs

func GetBackGroundJobs() *backGroundJobs {
	if bgJobs == nil {
		bgJobs = &backGroundJobs{
			Group: make(chan *group),
		}
	}
	return bgJobs
}
func (b *backGroundJobs) Run() {
	for {
		select {
		case g := <-b.Group:
			go g.Run(context.Background())
		}
	}
}
