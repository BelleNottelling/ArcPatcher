package main

import (
	"os"
	"regexp"
)

func removeDriverTimeoutNotification() error {
	updatesJSContent, err := os.ReadFile(updatesJS)
	if err != nil {
		return err
	}

	// Find the code that throws the driver timeout notification and then completely remove it.
	pattern := `showToast\({\s*type: notificationTypes.error,\s*toggleType: notificationToggleTypes.notification_driver_info,\s*mainMessageId: 'main-drivers',\s*secondaryMessageId: 'drivers-checking-for-updates-timeout',\s*}\);`

	re := regexp.MustCompile(pattern)
	modifiedUpdatesJS := re.ReplaceAllString(string(updatesJSContent), "")

	err = os.WriteFile(updatesJS, []byte(modifiedUpdatesJS), 0644)
	if err != nil {
		return err
	}

	return nil
}
