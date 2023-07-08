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

	err := patchFanControlHTML()
	if err != nil {
		return err
	}

	err = patchfanControlConfig()
	if err != nil {
		return err
	}

	err = patchfanControlPerformanceJS()
	if err != nil {
		return err
	}

	return nil
}

func patchFanControlHTML() error {
	indexHtmlBytes, err := os.ReadFile(indexHtml)
	if err != nil {
		return err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(indexHtmlBytes))
	if err != nil {
		return err
	}

	// Remove the "blockout" as it doesn't really provide a purpose and the positioning get's messed up
	doc.Find("#fan-speed-graph-blockout").Remove()

	// Update the height of the canvas to give it more space
	canvas := doc.Find("#fan-speed-graph-dragable")
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

	return nil
}

func patchfanControlConfig() error {
	chartJSConf, err := os.ReadFile(fanChartConfig)
	if err != nil {
		return err
	}

	// Find the old 'x' axis configuration and replace it with the new one that includes the ticks & that calculates the correct tempature
	pattern := `x: {\s*display: false,\s*grid: {\s*display: false,\s*},\s*ticks: {\s*color: energyBlue,\s*},\s*},`
	replacement := `x: {
        display: true,
        grid: {
          display: false,
        },
        ticks: {
          color: energyBlue,
          callback: function (value, index, values) {
            if (!Number.isInteger(value)) {
              return "";
            }
            length = isEmpty(activeOverclockingSettings()?.fan_speed_table) ? 12 : activeOverclockingSettings()?.fan_speed_table?.length;
            interval = 75 / (length - 1);
            return Math.round((value * interval + 25) * 2) / 2 + "°C";
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

	re := regexp.MustCompile(pattern)
	modifiedchartJSConf := re.ReplaceAllString(string(chartJSConf), replacement)

	err = os.WriteFile(fanChartConfig, []byte(modifiedchartJSConf), 0644)
	if err != nil {
		return err
	}

	return nil
}

func patchfanControlPerformanceJS() error {
	performanceJS, err := os.ReadFile(performanceTuningJS)
	if err != nil {
		return err
	}

	// Replace the chart initialization code with the custom one that updates it to have more fan curve points
	pattern := `/if \(isEmpty\(activeOverclockingSettings\(\)\?\.fan_speed_table\)\) {\s*activeOverclockingSettings\(\)\.fan_speed_table = \[30, 30, 40, 55, 75, 90];\s*}`
	replacement := `if (isEmpty(activeOverclockingSettings()?.fan_speed_table) || activeOverclockingSettings()?.fan_speed_table?.length <= 6) {
		activeOverclockingSettings().fan_speed_table = [30, 30, 30, 30, 30, 55, 65, 75, 82.5, 90]; //[30, 30, 40, 55, 75, 90]
		}`

	re := regexp.MustCompile(pattern)
	modified := re.ReplaceAllString(string(performanceJS), replacement)

	// Remove the now redundant tempature lables.
	modified = strings.Replace(modified, "document.getElementById('fan-graph-x-max').innerHTML = 100 + getTranslationFromId('units-celsius');", "", 1)
	modified = strings.Replace(modified, "document.getElementById('fan-graph-x-min').innerHTML = 25 + getTranslationFromId('units-celsius');", "", 1)

	// Update the reset option so that it also restores the custom fan curve with additional points.
	pattern = `newTuningSettings = null; \/\/ reset the state so we can make new changes\s*updateOverclockingSettings\(ret\?\.data\);`
	replacement = `newTuningSettings = null; // reset the state so we can make new changes
	ret.data.fan_speed_table = [30, 30, 30, 30, 30, 55, 65, 75, 82.5, 90];
	updateOverclockingSettings(ret?.data);`

	re = regexp.MustCompile(pattern)
	modified = re.ReplaceAllString(modified, replacement)

	err = os.WriteFile(performanceTuningJS, []byte(modified), 0644)
	if err != nil {
		return err
	}

	return nil
}
