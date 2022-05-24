package mbase

import (
	"fmt"
	"time"
)

func Retry(maxRetries int, maxRetryTime int64, f func() error) error {
	timeout := time.Duration(maxRetryTime) * time.Second

	for ; maxRetries > 0; maxRetries-- {
		err := func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("retry panic. %s", r)
				}
			}()

			if err = f(); err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
			fmt.Printf("Retry func exec fail, sleep. %v\n", err.Error())

			time.Sleep(timeout)
			continue
		}
		return nil
	}

	return fmt.Errorf("retry func exec fail")
}
