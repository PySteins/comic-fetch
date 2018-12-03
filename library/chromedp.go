package library

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
)

func ChromedpText(site, sel, referer string) (res string, err error) {
	log.Println("chromedp start: " + site)

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt)
	if err != nil {
		log.Println("error1: " + err.Error())
		return
	}

	// run task list
	err = c.Run(ctxt, chromedpSetheaders(
		site,
		sel,
		map[string]interface{}{
			"Referer": referer,
		},
		&res,
	))
	if err != nil {
		log.Println("error2: " + err.Error())
		return
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Println("error3: " + err.Error())
		return
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Println("error4: " + err.Error())
		return
	}

	return
}

func chromedpSetheaders(host, sel string, headers map[string]interface{}, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(headers)),
		chromedp.Navigate(host),
		chromedp.OuterHTML(sel, res, chromedp.ByQuery),
	}
}
