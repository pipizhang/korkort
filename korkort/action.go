package korkort

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/urfave/cli"
	"log"
	"regexp"
	"time"
)

func Setup(c *cli.Context) {
	InitConfig()

	DB := GetDB()
	defer DB.Close()

	DB.AutoMigrate(&Question{}, &Choice{})
}

func Scrape(c *cli.Context) {
	InitConfig()

	log.Println("start ...")

	var useragent, email, password string
	useragent, _ = Cfg.GetValue("app", "useragent")
	email, _ = Cfg.GetValue("korkort", "account")
	password, _ = Cfg.GetValue("korkort", "password")

	b := colly.NewCollector()
	b.UserAgent = useragent
	b.Limit(&colly.LimitRule{DomainGlob: "korkortonline"})

	b.OnRequest(func(r *colly.Request) {
		log.Println("[REQUEST] ", r.URL)
	})

	// Login
	err := b.Post("https://korkortonline.se/en/theory-test/sign-in/", map[string]string{"email": email, "passw": password})
	if err != nil {
		log.Fatal(err)
	}

	b.OnHTML("div.provyta > div:not([class])", func(e *colly.HTMLElement) {
		log.Println("--------------------------------")

		// Convert image url to absolute url
		e.DOM.Find("img").Each(func(i int, s *goquery.Selection) {
			if src, _ := s.Attr("src"); src != "" {
				s.SetAttr("src", e.Request.AbsoluteURL(src))
			}
		})

		title := e.ChildText("p.fraga")
		log.Println(title)

		e.DOM.Find("label").Each(func(i int, s *goquery.Selection) {
			option := s.Text()
			log.Println(i, ": ", option)
		})

		explanation, _ := e.DOM.Find("div.forklaringsruta").Html()
		log.Printf("explanation: %s\n", explanation)

		info := e.ChildText("div.navlinks")
		log.Printf("info: %s\n", info)

		imgSrc := e.ChildAttr("div.img img", "src")
		var imgURL string = ""
		if imgSrc != "" {
			imgURL = e.Request.AbsoluteURL(imgSrc)
			log.Println("image: >>> " + imgURL)
		}

	})

	b.OnHTML("a[href]", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		matched, _ := regexp.MatchString(`\/inloggad\/statistik/\w+\/$`, url)
		if matched {
			log.Println("======> ", url)
			time.Sleep(time.Second * 5)
			e.Request.Visit(url)
		}
	})

	// Statistics page
	b.Visit("https://korkortonline.se/inloggad/statistik/")
}
