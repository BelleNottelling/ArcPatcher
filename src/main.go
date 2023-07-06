package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	indexHtml           = `C:\Program Files\Intel\Intel Arc Control\resource\index.html`
	overlayHtml         = `C:\Program Files\Intel\Intel Arc Control\resource\overlay.html`
	fanChartConfig      = `C:\\Program Files\\Intel\\Intel Arc Control\\resource\\js\\chart_configs\\charts.js`
	performanceTuningJS = `C:\\Program Files\\Intel\\Intel Arc Control\\resource\\js\\pages\\performance\\performance_tuning.js`
	updatesJS           = `C:\\Program Files\\Intel\\Intel Arc Control\\resource\\js\\pages\\drivers\\updates.js`
	overlayJS           = `C:\\Program Files\\Intel\\Intel Arc Control\\resource\\js\\overlay.js`
)

var availablePatches = []string{
	"Restore backup",
	"Improved fan control",
	"Remove driver update timeout notification",
	"Minimal overlay",
	"Show Mhz on the performance boost slider",
}

func main() {
	if _, err := os.Stat("backup.zip"); err != nil {
		fmt.Println("Backup doesn't exist yet, creating now.")
		err = createBackup()

		if err != nil {
			coloredOutput(`error`, `An error occured when creating a backup:`, nil)
		} else {
			coloredOutput(`good`, `Successfully created a backup.`, nil)
		}
	} else {
		coloredOutput(`warning`, `Backup already exists, skipping backup step.`, nil)
	}

promt:
	selection, err := promtOptions()
	if err != nil {
		goto promt
	}

	// Just a blank like to make the screen more readable.
	fmt.Println("")

	switch selection {
	case 1:
		err = restoreBackup()
	case 2:
		err = betterFanControl()
	case 3:
		err = removeDriverTimeoutNotification()
	case 4:
		err = patchMinimalOverlay()
	case 5:
		err = perfBoostSliderMhz()
	default:
		coloredOutput(`warning`, `Unknown option selected.`, nil)
		goto promt
	}

	if err != nil {
		coloredOutput(`error`, `Error:`, err)
	} else {
		coloredOutput(`good`, `Patch applied without any errors.`, nil)
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
		coloredOutput(`warning`, `Invalid input. Please enter a valid number.`, nil)
		return 0, err
	}

	return choice, nil
}

func coloredOutput(outputType string, message string, err error) {
	var d *color.Color
	switch outputType {
	case `error`:
		d = color.New(color.FgRed, color.Bold)
	case `warning`:
		d = color.New(color.FgYellow, color.Bold)
	case `good`:
		d = color.New(color.FgGreen, color.Bold)
	default:
		d = color.New(color.FgYellow, color.Bold)
	}

	if err == nil {
		d.Println(message)
	} else {
		d.Println(message, err)
	}
}
