package google

import "github.com/mxschmitt/playwright-go"

type GStore struct {
	Url     string
	Pw      *playwright.Playwright
	Browser playwright.Browser
	Page    playwright.Page
}
