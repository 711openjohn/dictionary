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
