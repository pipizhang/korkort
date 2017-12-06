package korkort

import (
	"bytes"
	"fmt"
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
		Image       *IImage
		Choices     []*IChoice
	}

	IChoice struct {
		Content string
		Status  int
	}

	IImage struct {
		URL           string
		LocalFileName string
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

func (iq *IQuestion) AddChoice(raw string) {
	ic := &IChoice{}
	ic.Parse(raw)
	iq.Choices = append(iq.Choices, ic)
}

func (iq *IQuestion) AddImage(url string) {
	iq.Image = &IImage{
		URL:           url,
		LocalFileName: fmt.Sprintf("%s%s", RandStringBytes(10), GetFileExtension(url)),
	}
}

func (iq *IQuestion) GetImageFileName() (fileName string) {
	fileName = ""
	if iq.Image != nil {
		fileName = iq.Image.LocalFileName
	}
	return fileName
}

func (iq *IQuestion) Save() {
	db := GetDB()
	defer db.Close()

	var (
		n int = 0
		q Question
		c Choice
	)

	db.Model(&Question{}).Where("original_id = ?", iq.OriginalID).Count(&n)

	if n == 0 {
		if iq.Image != nil {
			iq.Image.Download()
		}

		q = Question{
			Content:     iq.Content,
			Explanation: iq.Explanation,
			OriginalID:  iq.OriginalID,
			Category:    iq.Category,
			Image:       iq.GetImageFileName(),
		}

		db.NewRecord(q)
		db.Create(&q)

		for _, v := range iq.Choices {
			c = Choice{
				Content:    v.Content,
				Status:     v.Status,
				QuestionID: q.ID,
			}
			db.NewRecord(c)
			db.Create(&c)
		}
	}

}

func (iq IQuestion) String() string {
	var _buffer bytes.Buffer

	_buffer.WriteString(fmt.Sprintf("Quesiton: %s\n", iq.Content))
	_buffer.WriteString("Options:\n")
	for k, v := range iq.Choices {
		_buffer.WriteString(fmt.Sprintf("%d [%d] %s\n", k+1, v.Status, v.Content))
	}
	_buffer.WriteString(fmt.Sprintf("Explanation: \n%s\n", iq.Explanation))
	_buffer.WriteString(fmt.Sprintf("OriginalID: %d\n", iq.OriginalID))
	_buffer.WriteString(fmt.Sprintf("Category: %s\n", iq.Category))
	if iq.Image != nil {
		_buffer.WriteString(fmt.Sprintf("Image: %s\n", iq.Image.URL))
	} else {
		_buffer.WriteString("Image: nil")
	}

	return _buffer.String()
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

func (ii *IImage) GetFilePath() string {
	imagePath, _ := Cfg.GetValue("app", "image")
	return imagePath + "/" + ii.LocalFileName
}

func (ii *IImage) Download() {
	DownloadImage(ii.URL, ii.GetFilePath())
}
