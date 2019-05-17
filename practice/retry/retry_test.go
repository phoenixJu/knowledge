package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	Retry(3, time.Duration(3 * time.Second), func() error {
		fmt.Println("download s3 / tos")
		return errors.New("fail")
	})
}
