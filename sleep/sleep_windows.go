package sleep

import "os/exec"

func Sleep() error {
	sleep := exec.Command("shutdown", "/h")

	err := sleep.Start()
	if err != nil{
		return err
	}

	err = sleep.Wait()
	if err != nil{
		return err
	}

	return nil
}