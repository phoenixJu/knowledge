package retry

import (
	"code.byted.org/gopkg/logs"
	"context"
	"time"
)

type stop struct {
	error
}

func NoRetryError(err error) error {
	return stop{err}
}
func Retry(retryTimes int, sleep time.Duration, fun func() error) error {
	if err := fun(); err != nil {
		if s, ok := err.(stop); ok {
			logs.CtxError(context.Background(), "Retry do func failed. retryTimes: %v, sleep: %v ", retryTimes, sleep)
			return s.error
		}
		if retryTimes--; retryTimes > 0 {
			logs.CtxWarn(context.Background(), "retry func failed. attemp after %v time.", sleep)
			time.Sleep(sleep)
			return Retry(retryTimes, sleep*2, fun)
		}
		return err
	}
	return nil
}
