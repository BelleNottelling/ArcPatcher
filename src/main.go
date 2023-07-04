package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var indexHtml = "C:\\Program Files\\Intel\\Intel Arc Control\\resource\\index.html"
var fanChartConfig = "C:\\Program Files\\Intel\\Intel Arc Control\\resource\\js\\chart_configs\\charts.js"
var performanceTurningJS = "C:\\Program Files\\Intel\\Intel Arc Control\\resource\\js\\pages\\performance\\performance_tuning.js"
var updatesJS = "C:\\Program Files\\Intel\\Intel Arc Control\\resource\\js\\pages\\drivers\\updates.js"
var overlayJS = "C:\\Program Files\\Intel\\Intel Arc Control\\resource\\js\\overlay.js"

func main() {
	err := betterFanControl()
	removeDriverTimeoutNotification()
	minimalOverlayJS()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("HTML manipulation and saving complete.")
}

func betterFanControl() error {
	// HTML Based modifications
	indexHtmlBytes, err := os.ReadFile(indexHtml)
	if err != nil {
		return err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(indexHtmlBytes))
	if err != nil {
		return err
	}

	doc.Find("#fan-speed-graph-blockout").Remove()
	canvas := doc.Find("#fan-speed-graph-dragable")

	// Update the height of the canvas to give it more space
	if canvas.Length() == 0 {
		return fmt.Errorf("canvas element not found")
	}
	canvas.SetAttr("height", "125")
	canvas.SetAttr("style", "display: block; box-sizing: border-box; height: 125px; width: 527px; touch-action: none; -webkit-tap-highlight-color: rgba(0, 0, 0, 0); transform: translateZ(10px);")

	// Save the modified HTML back to the file
	updatedHTML, err := doc.Html()
	if err != nil {
		return err
	}
	err = os.WriteFile(indexHtml, []byte(updatedHTML), 0644)
	if err != nil {
		return err
	}

	// Now remove the 25C and 100C labels
	performanceJS, err := os.ReadFile(performanceTurningJS)
	if err != nil {
		return err
	}

	modifiedperformanceJS := strings.Replace(string(performanceJS), "document.getElementById('fan-graph-x-max').innerHTML = 100 + getTranslationFromId('units-celsius');", "", 1)
	modifiedperformanceJS = strings.Replace(modifiedperformanceJS, "document.getElementById('fan-graph-x-min').innerHTML = 25 + getTranslationFromId('units-celsius');", "", 1)
	err = os.WriteFile(performanceTurningJS, []byte(modifiedperformanceJS), 0644)
	if err != nil {
		return err
	}

	// Finally update the x-axis config for Chart.JS to includ the ticks with the tempatures
	chartJSConf, err := os.ReadFile(fanChartConfig)
	if err != nil {
		return err
	}

	pattern := `x: {\s*display: false,\s*grid: {\s*display: false,\s*},\s*ticks: {\s*color: energyBlue,\s*},\s*},`
	replacement := `x: {
		display: true,
		grid: {
			display: false,
		},
		ticks: {
			color: energyBlue,
			callback: function(value, index, values) {
				if (!Number.isInteger(value)) {
					return "";
				}
				return (value * 15) + 25 + "°C";
			},
			padding: 0, 
			font: {
				size: 14,
			},
			color: "#FFFFFF",
		},
		scaleLabel: {
			display: true,
		},
	},`

	// Perform the replacement using regular expressions
	re := regexp.MustCompile(pattern)
	modifiedchartJSConf := re.ReplaceAllString(string(chartJSConf), replacement)

	err = os.WriteFile(fanChartConfig, []byte(modifiedchartJSConf), 0644)
	if err != nil {
		return err
	}

	return nil
}

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

func minimalOverlayJS() error {
	content, err := os.ReadFile(overlayJS)
	if err != nil {
		return err
	}

	pattern := `<li\s*id="\${setting\?\.settingId}-wrapper"\s*class="\${hidden\s*\?\s*'is-hidden'\s*:\s*''}">`

	// Perform the replacement using regular expressions
	re := regexp.MustCompile(pattern)
	modifiedUpdatesJS := re.ReplaceAllString(string(content), `<li id="${setting?.settingId}-wrapper" class="${hidden ? 'is-hidden' : ''}" style="margin: 0; padding: 0.25em;">`)

	err = os.WriteFile(overlayJS, []byte(modifiedUpdatesJS), 0644)
	if err != nil {
		return err
	}

	return nil
}
