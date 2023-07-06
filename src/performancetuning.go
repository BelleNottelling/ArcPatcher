package main

import (
	"os"
	"regexp"
)

func perfBoostSliderMhz() error {
	performanceJS, err := os.ReadFile(performanceTuningJS)
	if err != nil {
		return err
	}

	// First we enabled the positive / negative indicator
	pattern := `visible: activePerformanceAdapter\(\)\?\.supported_oc_features\?\.gpu_performance_boost\?\.bSupported,`
	replacement := `visible: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.bSupported,
	showPosNeg: true,`

	re := regexp.MustCompile(pattern)
	modified := re.ReplaceAllString(string(performanceJS), replacement)

	// Then we disable the "normalize" option so that it is no longer displayed as a percentage
	pattern = `normalize: true,\s*min: activePerformanceAdapter\(\)\?\.supported_oc_features\?\.gpu_performance_boost\?\.min,`
	replacement = `normalize: false,
	min: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.min,`

	re = regexp.MustCompile(pattern)
	modified = re.ReplaceAllString(modified, replacement)

	// Finally, set it to display the MHz unit
	pattern = `defaultValue: activePerformanceAdapter\(\)\?\.supported_oc_features\?\.gpu_performance_boost\?\.default,\s*units: null,`
	replacement = `defaultValue: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.default,
    units: 'units-mhz',`

	re = regexp.MustCompile(pattern)
	modified = re.ReplaceAllString(modified, replacement)

	err = os.WriteFile(performanceTuningJS, []byte(modified), 0644)
	if err != nil {
		return err
	}

	return nil
}
