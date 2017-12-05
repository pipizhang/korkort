package korkort

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type (
	IQuestion struct {
		Content     string
		Explanation string
		OriginalID  int
		Category    string
		Image       IImage
		Choices     []IChoice
	}

	IChoice struct {
		Content string
		Status  int
	}

	IImage struct {
		URL       string
		LocalFile string
	}
)

// Parse title fo question
func (iq *IQuestion) ParseContent(raw string) {
	raw = strings.TrimSpace(raw)
	re := regexp.MustCompile(`^\d+\)\s+`)
	iq.Content = re.ReplaceAllString(raw, "")
}

func (iq *IQuestion) ParseExplanation(raw string) {
	raw = strings.TrimSpace(raw)
	iq.Explanation = raw
}

// Parse and extract orginal id of question
func (iq *IQuestion) ParseOrignalID(raw string) {
	re := regexp.MustCompile(`\(no\. (\d+)\)`)
	m := re.FindStringSubmatch(raw)
	if len(m) == 0 {
		log.Fatalln("Can't find OriginalID")
	}
	iq.OriginalID, _ = strconv.Atoi(m[1])
}

// Parse and extract category of question
func (iq *IQuestion) ParseCategory(raw string) {
	re := regexp.MustCompile(`Show explanation (.*) \(no`)
	m := re.FindStringSubmatch(raw)
	if len(m) == 0 {
		log.Fatalln("Can't find Category")
	}
	iq.Category = m[1]
}

func (ic *IChoice) Parse(raw string) {
	raw = strings.TrimSpace(raw)

	ic.Status = 0
	if strings.Contains(raw, "✓") {
		ic.Status = 1
	}

	re := regexp.MustCompile(`^(✓|✗)\s+`)
	ic.Content = re.ReplaceAllString(raw, "")
}
