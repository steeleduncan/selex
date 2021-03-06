package selex

import (
    "fmt"
	"strings"
	"time"
    "testing"
	"github.com/tebeka/selenium"
)

func TestGoPlayground(t *testing.T) {
    wr, err := NewWebRunner()
    if err != nil {
        t.Errorf("Error booting web runner")
    }
    defer wr.Teardown()

	if err := wr.Driver().Get("http://play.golang.org/?simple=1"); err != nil {
        t.Errorf("Error navigating to website: %v", err)
	}
	elem, err := wr.Driver().FindElement(selenium.ByCSSSelector, "#code")
	if err != nil {
        t.Errorf("Error finding element: %v", err)
	}
	if err := elem.Clear(); err != nil {
        t.Errorf("Error clearing the box")
	}

	// Enter some new code in text box.
	err = elem.SendKeys(`
		package main
		import "fmt"
		func main() {
			fmt.Println("Hello WebDriver!")
		}
	`)
	if err != nil {
        t.Errorf("Error writing in text")
	}

	btn, err := wr.Driver().FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
        t.Errorf("Error finding run")
	}
	if err := btn.Click(); err != nil {
        t.Errorf("Error clicking run")
	}

	// Wait for the program to finish running and get the output.
	outputDiv, err := wr.Driver().FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
        t.Errorf("Error finding output box")
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
            t.Errorf("Error reading text")
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))
}

func TestTexpad(t *testing.T) {
    wr, err := NewWebRunner()
    if err != nil {
        t.Errorf("Error booting web runner")
    }
    defer wr.Teardown()

	if err := wr.Driver().Get("https://texpad.com"); err != nil {
        t.Errorf("Error navigating to website: %v", err)
	}

	loginButton, err := wr.Driver().FindElement(selenium.ByLinkText, "Log in")
	if err != nil {
        t.Errorf("Error finding login button")
	}

	if err := loginButton.Click(); err != nil {
        t.Errorf("Error clicking login button")
	}

	usernameField, err := wr.Driver().FindElement(selenium.ByID, "username")
	if err != nil {
        t.Errorf("Error finding username field")
	}
	if err := usernameField.Clear(); err != nil {
        t.Errorf("Error clearing the username field")
	}
	err = usernameField.SendKeys("chuffy")
	if err != nil {
        t.Errorf("Error writing in text")
	}

    time.Sleep(1 * time.Second)

	passwordField, err := wr.Driver().FindElement(selenium.ByID, "password")
	if err != nil {
        t.Errorf("Error finding password field")
	}
	if err := passwordField.Clear(); err != nil {
        t.Errorf("Error clearing the password field")
	}
	err = passwordField.SendKeys("not-my-password")
	if err != nil {
        t.Errorf("Error writing in text to password")
	}

    time.Sleep(1 * time.Second)

	loginButton2, err := wr.Driver().FindElement(selenium.ByLinkText, "Log in")
	if err != nil {
        t.Errorf("Error finding login button")
	}
	if err := loginButton2.Click(); err != nil {
        t.Errorf("Error clicking login button")
	}

    time.Sleep(10 * time.Second)
}
