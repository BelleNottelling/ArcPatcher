package main

import (
	"os"
	"regexp"
)

func fixArcControlBugs() error {
	err := fixPerfNotUpdateSlider()
	if err != nil {
		return err
	}

	return nil
}

func fixPerfNotUpdateSlider() error {
	performanceJS, err := os.ReadFile(performanceTuningJS)
	if err != nil {
		return err
	}

	// Replace the chart initialization code with the custom one that updates it to have more fan curve points
	pattern := `newTuningSettings = null; \/\/ reset the state so we can make more changes\s*warningRequired = false; \/\/ clear warning flag`
	replacement := `updateOverclockingSettings(ret?.data);
				updatePerformanceTuningSettings();
				warningRequired = false; // clear warning flag`

	re := regexp.MustCompile(pattern)
	modified := re.ReplaceAllString(string(performanceJS), replacement)

	err = os.WriteFile(performanceTuningJS, []byte(modified), 0644)
	return err
}
