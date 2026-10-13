package main

import "fmt"

func SafeExecutor(task func()) (err error) {
	// TODO: Execute task safely using defer and recover; convert any panic to error using fmt.Errorf
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	task()
	return err
}
