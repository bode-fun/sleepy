package sleep

import (
	"os/exec"
)

func Sleep() error {
	// Kinda sucks because the App has to be allowed to send messages to System Events
	// sleepCmd := exec.Command("osascript", "-e", `tell app "System Events" to sleep`)
	// Therefore pmset is used
	sleepCmd := exec.Command("pmset", "sleepnow")

	err := sleepCmd.Start()
	if err != nil {
		return err
	}

	err = sleepCmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
