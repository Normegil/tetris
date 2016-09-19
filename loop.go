package main

import (
	"time"

	"github.com/Sirupsen/logrus"
)

type loopCtrl struct {
	fps  fpsControls
	quit bool
}

type fpsControls struct {
	number int
	capped bool
}

func loop(init loopCtrl, toExec func(loopCtrl) (loopCtrl, error)) error {
	ctrl := init
	quit := ctrl.quit
	logrus.WithField("Controls", ctrl).Debug("Launching main loop")
	for !quit {
		beforeLoop := time.Now()
		ctrl, err := toExec(ctrl)
		if nil != err {
			return err
		}
		quit = ctrl.quit
		if ctrl.fps.capped && !quit {
			afterLoop := time.Now()
			time.Sleep(timeToSleep(ctrl.fps.number, beforeLoop, afterLoop))
		}
	}
	return nil
}

func timeToSleep(framePerSeconds int, beforeLoop, afterLoop time.Time) time.Duration {
	spentTime := time.Duration(toMilliseconds(afterLoop) - toMilliseconds(beforeLoop))
	theoriticalTimeToWait := time.Duration(1000 / framePerSeconds)
	timeToWait := theoriticalTimeToWait - spentTime
	return timeToWait * time.Millisecond
}

func toMilliseconds(t time.Time) int64 {
	return int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
}
