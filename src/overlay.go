package main

import (
	"bytes"
	"os"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func patchMinimalOverlay() error {
	err := minimalOverlayJS()
	if err != nil {
		return err
	}

	err = minimalOverlayHTML()
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

	// Find the Overlay's li template and replace it with one that has reduced padding
	pattern := `<li\s*id="\${setting\?\.settingId}-wrapper"\s*class="\${hidden\s*\?\s*'is-hidden'\s*:\s*''}">`

	re := regexp.MustCompile(pattern)
	modifiedOverlayJS := re.ReplaceAllString(string(content), `<li id="${setting?.settingId}-wrapper" class="${hidden ? 'is-hidden' : ''}" style="margin: 0; padding: 0.25em;">`)

	err = os.WriteFile(overlayJS, []byte(modifiedOverlayJS), 0644)
	if err != nil {
		return err
	}

	return nil
}

func minimalOverlayHTML() error {
	overlayHtmlBytes, err := os.ReadFile(overlayHtml)
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(overlayHtmlBytes))
	if err != nil {
		return err
	}

	// Removes the "Intel® Performance Telemetry" header
	doc.Find(".header").Remove()

	updatedHTML, err := doc.Html()
	if err != nil {
		return err
	}
	err = os.WriteFile(overlayHtml, []byte(updatedHTML), 0644)
	if err != nil {
		return err
	}

	return nil
}
