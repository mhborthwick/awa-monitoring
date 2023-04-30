package colly

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Item struct {
	Provider string
	Name     string
	Status   string
}

type Selector struct {
	Container string
	Name      string
	Status    string
}

func GetSelectors(provider string) *Selector {
	if provider == "Klaviyo" {
		return &Selector{
			Container: ".components-container .component-inner-container",
			Name:      ".name",
			Status:    ".component-status",
		}
	}
	if provider == "Hover" {
		return &Selector{
			Container: "#statusio_components .component",
			Name:      ".component_name",
			Status:    ".component-status",
		}
	}
	return nil
}

func ScrapeData(
	url string,
	provider string,
) []Item {
	c := colly.NewCollector()
	selectors := GetSelectors(provider)
	data := []Item{}
	c.OnHTML(selectors.Container, func(h *colly.HTMLElement) {
		temp := Item{}
		temp.Provider = provider
		temp.Name = h.ChildText(selectors.Name)
		temp.Status = h.ChildText(selectors.Status)
		data = append(data, temp)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit(url)
	// fmt.Println(data, url) //TODO: Remove
	return data
}
