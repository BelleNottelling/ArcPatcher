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

	pattern := `range: {\s*normalize: true,\s*min: activePerformanceAdapter\(\)\?\.supported_oc_features\?\.gpu_performance_boost\?\.min,\s*max: activePerformanceAdapter\(\)\?\.supported_oc_features\?\.gpu_performance_boost\?\.max,\s*step: activePerformanceAdapter\(\)\?\.supported_oc_features\?\.gpu_performance_boost\?\.step,\s*defaultValue: activePerformanceAdapter\(\)\?\.supported_oc_features\?\.gpu_performance_boost\?\.default,\s*units: null,\s*}`
	replacement := `range: {
		normalize: false,
		min: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.min,
		max: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.max,
		step: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.step,
		defaultValue: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.default,
		units: 'units-mhz',
	}`

	// Perform the replacement using regular expressions
	re := regexp.MustCompile(pattern)
	modified := re.ReplaceAllString(string(performanceJS), replacement)

	err = os.WriteFile(performanceTuningJS, []byte(modified), 0644)
	if err != nil {
		return err
	}

	return nil
}
