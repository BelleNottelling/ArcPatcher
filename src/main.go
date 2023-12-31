package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/bi-zone/go-fileversion"
	"github.com/fatih/color"
	"golang.org/x/exp/slices"
	"golang.org/x/sys/windows"
)

var (
	indexHtml           = `C:\Program Files\Intel\Intel Arc Control\resource\index.html`
	overlayHtml         = `C:\Program Files\Intel\Intel Arc Control\resource\overlay.html`
	fanChartConfig      = `C:\Program Files\Intel\\Intel Arc Control\resource\js\chart_configs\charts.js`
	performanceTuningJS = `C:\Program Files\Intel\\Intel Arc Control\resource\js\pages\performance\performance_tuning.js`
	updatesJS           = `C:\Program Files\Intel\\Intel Arc Control\resource\js\pages\drivers\updates.js`
	overlayJS           = `C:\Program Files\Intel\\Intel Arc Control\resource\js\overlay.js`
)

var testedArcControlVersions = []string{
	`1.69.5033.3`,
}

var availablePatches = []string{
	"Restore backup",
	"Improved fan control",
	"Remove driver update timeout notification",
	"Minimal overlay",
	"Show MHz on the performance boost slider",
	"Fix Arc Control Bugs",
}

func main() {
	escalateIfNotAdmin()

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

	checkArControlVersion()

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
	case 6:
		err = fixArcControlBugs()
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
	fmt.Println("Available tasks: ")
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

func checkArControlVersion() {
	f, err := fileversion.New(`C:\Program Files\Intel\Intel Arc Control\ArcControl.exe`)
	if err != nil {
		coloredOutput("error", "Unable to read your Arc Control version", err)
		return
	}

	if slices.Contains(testedArcControlVersions, f.FileVersion()) {
		coloredOutput("good", "Your version of Arc Control has been tested with ArcPatcher", nil)
	} else {
		coloredOutput("warning", "Your version of Arc Control has not been tested with ArcPatcher", nil)
	}
}

// Based on the gist here: https://gist.github.com/jerblack/d0eb182cc5a1c1d92d92a4c4fcc416c6
func escalateIfNotAdmin() {
	if !windows.GetCurrentProcessToken().IsElevated() {
		exe, _ := os.Executable()
		cwd, _ := os.Getwd()
		args := strings.Join(os.Args[1:], " ")

		verbPtr, _ := syscall.UTF16PtrFromString("runas")
		exePtr, _ := syscall.UTF16PtrFromString(exe)
		cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
		argPtr, _ := syscall.UTF16PtrFromString(args)

		var showCmd int32 = 1 //SW_NORMAL

		err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
		if err != nil {
			coloredOutput("error", "ArcPatcher requires administrator rights in order to function correctly", nil)
			coloredOutput("warning", "This error occured while attempting to launch ArcPatcher with administrator rights: ", err)
		} else {
			os.Exit(0)
		}
	}
}
