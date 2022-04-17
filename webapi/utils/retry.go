package utils

import (
	"fmt"
	"time"
)

func Retry(count int, timeoutRaw int64, f func() error) error {
	timeout := time.Duration(timeoutRaw) * time.Second

	for ; count > 0; count-- {
		err := func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("Retry panic. %s", r)
				}
			}()

			err = f()
			if err != nil {
				return err
			}

			return nil
		}()
		if err != nil {
			fmt.Printf("Retry func exec fail, sleep. %s\n", err.Error())

			time.Sleep(timeout)
			continue
		}

		return nil
	}

	return fmt.Errorf("Retry func exec fail")
}
