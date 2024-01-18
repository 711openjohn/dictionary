package client

import (
	"dictionary/domain"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type Cambridge struct {
	Host string
}

func NewCambridge() Cambridge {
	return Cambridge{
		Host: "https://dictionary.cambridge.org/",
	}
}

func (c *Cambridge) Lookup(word string, language string) (domain.DictionaryPage, error) {
	// https://dictionary.cambridge.org/dictionary/english-chinese-traditional/equal?q=equals
	var dp domain.DictionaryPage
	definitions := []domain.Definition{}
	col := colly.NewCollector()
	col.OnRequest(func(req *colly.Request) {
		req.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Headers.Set("Accept-Language", "en-US,en;q=0.9,ja;q=0.8")
		req.Headers.Set("Cache-Control", "no-cache")
		req.Headers.Set("Connection", "keep-alive")
		req.Headers.Set("Pragma", "no-cache")
		req.Headers.Set("Referer", "https://dictionary.cambridge.org/dictionary/english-chinese-traditional/")
		req.Headers.Set("Sec-Fetch-Dest", "document")
		req.Headers.Set("Sec-Fetch-Mode", "navigate")
		req.Headers.Set("Sec-Fetch-Site", "same-origin")
		req.Headers.Set("Upgrade-Insecure-Requests", "1")
		req.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
		req.Headers.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
		req.Headers.Set("sec-ch-ua-mobile", "?0")
		req.Headers.Set("sec-ch-ua-platform", `"macOS"`)
	})
	col.OnHTML(".pos-header .di-title", func(e *colly.HTMLElement) {
		dp.Symbol = e.Text
	})
	col.OnHTML(".pr.entry-body__el", func(e *colly.HTMLElement) {
		pronunciationLink := ""
		if source, exists := e.DOM.Find(".pos-header.dpos-h .us.dpron-i source").First().Attr("src"); exists {
			pronunciationLink = fmt.Sprintf("%s%s", c.Host, source)
		}
		pos := e.DOM.Find(".posgram.dpos-g").First().Text()
		e.DOM.Find(".def-block.ddef_block").Each(func(_ int, s *goquery.Selection) {
			if len(s.Find(".x.dx").First().Nodes) != 0 {
				return
			}
			definition := domain.Definition{}
			definition.PronunciationLink = pronunciationLink
			definition.POS = pos
			definition.Meaning = domain.Meaning{}
			text := strings.TrimSpace(s.Find(".def.ddef_d.db").First().Text())
			if len(text) == 0 {
				return
			}
			translation := strings.TrimSpace(s.Find(".trans.dtrans").First().Text())
			definition.Meaning.Text = text
			definition.Meaning.Translation = translation

			examples := []domain.Example{}
			s.Find(".examp.dexamp").Each(func(_ int, s1 *goquery.Selection) {
				example := domain.Example{}
				example.Text = s1.Find(".eg").Text()
				example.Translation = s1.Find(".trans").Text()
				examples = append(examples, example)
			})
			definition.Meaning.Examples = examples
			definitions = append(definitions, definition)
		})
	})
	col.Visit(fmt.Sprintf("https://dictionary.cambridge.org/dictionary/english-%v/%v", language, word))
	dp.Definitions = definitions

	return dp, nil
}
