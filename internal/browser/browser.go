package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Browser struct {
	browser *rod.Browser
	page    *rod.Page
}

func New() *Browser {
	return &Browser{}
}

func (b *Browser) Launch() error {
	// launches a headless browser
	url := launcher.New().Headless(true).MustLaunch()
	b.browser = rod.New().ControlURL(url).MustConnect()
	return nil
}

func (b *Browser) Navigate(url string) error {
	// navigates to a specific url
	b.page = b.browser.MustPage(url)
	b.page.MustWaitStable()
	return nil
}

func (b *Browser) Close() {
	// closes browser if it is open
	if b.browser != nil {
		b.browser.MustClose()
	}
}

func (b *Browser) GetPage() *rod.Page {
	// get the current page opened in browser
	return b.page
}
