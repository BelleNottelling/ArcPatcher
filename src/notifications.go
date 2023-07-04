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

	pattern := `showToast\({\s*type: notificationTypes.error,\s*toggleType: notificationToggleTypes.notification_driver_info,\s*mainMessageId: 'main-drivers',\s*secondaryMessageId: 'drivers-checking-for-updates-timeout',\s*}\);`

	// Perform the replacement using regular expressions
	re := regexp.MustCompile(pattern)
	modifiedUpdatesJS := re.ReplaceAllString(string(updatesJSContent), "")

	err = os.WriteFile(updatesJS, []byte(modifiedUpdatesJS), 0644)
	if err != nil {
		return err
	}

	return nil
}
