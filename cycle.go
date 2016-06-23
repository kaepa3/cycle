package main

import "time"

// CycleProc is cycleprocess task
type CycleProc struct {
	time   int
	flg    bool
	action func()
}

func doProcess(obj CycleProc) {

	go func() {
		sleepTime := time.Duration(obj.time) * time.Millisecond
		t := time.NewTicker(sleepTime)
		for obj.flg == true {
			for {
				select {
				case <-t.C:
					obj.action()
				}
			}
		}
		t.Stop()
	}()

}
