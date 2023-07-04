package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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

	// Finally update the x-axis config for Chart.JS to include the ticks with the temperatures
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
