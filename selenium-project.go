package main

import (
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const (
	seleniumURL = "http://localhost:4444/wd/hub"
)

func main() {
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	driver, err := selenium.NewRemote(caps, seleniumURL)
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}
	defer driver.Quit()

	url := "https://demoqa.com/automation-practice-form"
	if err := driver.Get(url); err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	inputFirstName, err := driver.FindElement(selenium.ByID, "firstName")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	err = inputFirstName.SendKeys("Daniel")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	inputLastName, err := driver.FindElement(selenium.ByID, "lastName")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	err = inputLastName.SendKeys("Certan")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	inputUserEmail, err := driver.FindElement(selenium.ByID, "userEmail")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	err = inputUserEmail.SendKeys("oviovovivoi7b@gmail.com")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	inputUserNumber, err := driver.FindElement(selenium.ByID, "userNumber")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	err = inputUserNumber.SendKeys("0637222710")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	inputBirth, err := driver.FindElement(selenium.ByID, "dateOfBirthInput")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	inputBirth.Click()
	time.Sleep(500 * time.Millisecond)

	monthSelect, err := driver.FindElement(selenium.ByCSSSelector, "select.react-datepicker__month-select")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}
	err = monthSelect.SendKeys("January")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	yearSelect, err := driver.FindElement(selenium.ByCSSSelector, "select.react-datepicker__year-select")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}
	err = yearSelect.SendKeys("2025")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	day5, err := driver.FindElement(selenium.ByCSSSelector, "div.react-datepicker__day--005")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	err = day5.Click()
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	subjectsInput, err := driver.FindElement(selenium.ByID, "subjectsInput")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}
	subjectsInput.SendKeys("Maths")

	time.Sleep(500 * time.Millisecond)

	mathOption, err := driver.FindElement(selenium.ByXPATH, "//div[contains(@class,'subjects-auto-complete__option') and text()='Maths']")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	time.Sleep(500 * time.Millisecond)

	driver.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{mathOption})
	time.Sleep(200 * time.Millisecond)
	driver.ExecuteScript("arguments[0].click();", []interface{}{mathOption})

	musicLabel, err := driver.FindElement(selenium.ByCSSSelector, "label[for='hobbies-checkbox-3']")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	err = musicLabel.Click()
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	sportLabel, err := driver.FindElement(selenium.ByCSSSelector, "label[for='hobbies-checkbox-1']")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	err = sportLabel.Click()
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	fileInput, err := driver.FindElement(selenium.ByCSSSelector, "input.form-control-file")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	err = fileInput.SendKeys("/home/debian/projects/selenium_project/dog.jpeg")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	inputAddress, err := driver.FindElement(selenium.ByID, "currentAddress")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	err = inputAddress.SendKeys("str.Igor Vieru 2/1")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	stateDropdown, err := driver.FindElement(selenium.ByXPATH, "//div[contains(@class,'css-1wa3eu0-placeholder') and text()='Select State']")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}
	stateDropdown.Click()
	time.Sleep(500 * time.Millisecond)

	haryanaOption, err := driver.FindElement(selenium.ByXPATH, "//div[contains(@class,'css-yt9ioa-option') and text()='Haryana']")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	err = haryanaOption.Click()
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	cityDropdown, err := driver.FindElement(selenium.ByXPATH, "//div[contains(@class,'css-1wa3eu0-placeholder') and text()='Select City']")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}
	cityDropdown.Click()
	time.Sleep(500 * time.Millisecond)

	panipatOption, err := driver.FindElement(selenium.ByXPATH, "//div[contains(@class,'css-yt9ioa-option') and text()='Panipat']")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	err = panipatOption.Click()
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	maleLabel, err := driver.FindElement(selenium.ByCSSSelector, "label[for='gender-radio-1']")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	err = maleLabel.Click()
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}

	submit, err := driver.FindElement(selenium.ByID, "submit")
	if err != nil {
		log.Fatalf("Eroare: %v", err)
	}
	submit.Click()

	time.Sleep(10 * time.Second)
}
