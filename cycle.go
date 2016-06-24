package cycle

import "time"

// CycleProc is cycleprocess task
type CycleProc struct {
	Time   int
	Flg    bool
	Action func()
}

func DoProcess(obj CycleProc) {
	go func() {
		sleepTime := time.Duration(obj.Time) * time.Millisecond
		t := time.NewTicker(sleepTime)
		for obj.Flg == true {
			for {
				select {
				case <-t.C:
					obj.Action()
				}
			}
		}
		t.Stop()
	}()
}
