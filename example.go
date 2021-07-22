package selex

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium"
)

type WebRunner struct {
    wd selenium.WebDriver
    srv *selenium.Service
}

func (wr *WebRunner) Driver() selenium.WebDriver {
    return wr.wd
}

func (wr *WebRunner) Teardown() {
    wr.wd.Quit()
    wr.srv.Stop()
}

func NewWebRunner() (*WebRunner, error) {
	const port int = 8083
    wr := &WebRunner { }

    var err error

	selenium.SetDebug(true)
	wr.srv, err = selenium.NewSeleniumService(
        "/usr/local/Cellar/selenium-server-standalone/3.141.59_2/libexec/selenium-server-standalone-3.141.59.jar",
        port,
		selenium.GeckoDriver("/usr/local/bin/geckodriver"),
		selenium.Output(os.Stderr),
    )
	if err != nil {
       return nil, fmt.Errorf("Error starting selenium: %v", err)
	}

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wr.wd, err = selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
       return nil, fmt.Errorf("Error connecting to firefox: %v", err)
	}

    return wr, nil
}
