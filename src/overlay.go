package main

import (
	"os"
	"regexp"
)

func minimalOverlayJS() error {
	content, err := os.ReadFile(overlayJS)
	if err != nil {
		return err
	}

	pattern := `<li\s*id="\${setting\?\.settingId}-wrapper"\s*class="\${hidden\s*\?\s*'is-hidden'\s*:\s*''}">`

	// Perform the replacement using regular expressions
	re := regexp.MustCompile(pattern)
	modifiedOverlayJS := re.ReplaceAllString(string(content), `<li id="${setting?.settingId}-wrapper" class="${hidden ? 'is-hidden' : ''}" style="margin: 0; padding: 0.25em;">`)

	err = os.WriteFile(overlayJS, []byte(modifiedOverlayJS), 0644)
	if err != nil {
		return err
	}

	return nil
}
