package main

import "time"

type loopCtrl struct {
	fps fpsControls
	quit bool
}

type fpsControls struct{
	number int
	capped bool
}

func loop(init loopCtrl, toExec func (loopCtrl) (loopCtrl, error)) error {
	quit := init.quit
	for(!quit) {
		beforeLoop := time.Now()
		ctrl, err := toExec()
		if nil != err {
			return err
		}
		quit = ctrl.quit
		if (ctrl.fps.capped && !quit) {
			afterLoop := time.Now()
			timeToWait := 1000 / ctrl.fps.number
			time.Sleep(timeToWait - (afterLoop - beforeLoop))
		}
	}
	return nil
}
