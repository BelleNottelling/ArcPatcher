package main

import (
	"fmt"
)

var (
	indexHtml            = `C:\Program Files\Intel\Intel Arc Control\resource\index.html`
	fanChartConfig       = `C:\\Program Files\\Intel\\Intel Arc Control\\resource\\js\\chart_configs\\charts.js`
	performanceTurningJS = `C:\\Program Files\\Intel\\Intel Arc Control\\resource\\js\\pages\\performance\\performance_tuning.js`
	updatesJS            = `C:\\Program Files\\Intel\\Intel Arc Control\\resource\\js\\pages\\drivers\\updates.js`
	overlayJS            = `C:\\Program Files\\Intel\\Intel Arc Control\\resource\\js\\overlay.js`
)

var availablePatches = []string{
	"Improved fan control",
	"Remove driver update timeout notification",
	"Minimal overlay",
	"Option 4",
}

func main() {
promt:
	selection, err := promtOptions()
	if err != nil {
		goto promt
	}

	// Just a blank like to make the screen more readable.
	fmt.Println("")

	switch selection {
	case 1:
		err = betterFanControl()
		if err != nil {
			fmt.Println("Error:", err)
		}
	case 2:
		removeDriverTimeoutNotification()
	case 3:
		minimalOverlayJS()
	default:
		fmt.Println("Unknown option selected.")
	}
	goto promt
}

func promtOptions() (int, error) {
	fmt.Println("Available patches: ")
	// Display numbered options
	for i, option := range availablePatches {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	// Prompt for user input
	fmt.Print("Enter the number of your choice: ")
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return 0, err
	}

	return choice, nil
}
