package korkort

import (
	_ "fmt"
	"testing"
)

func TestGetFileExtension(t *testing.T) {
	var uri string

	uri = "https://www.example.com/images/M141.png?a=12#abc"
	if GetFileExtension(uri) != ".png" {
		t.Error("Expected ext == '.png'")
	}

	uri = "https://www.example.com/images/test"
	if GetFileExtension(uri) != "" {
		t.Error("Expected ext == ''")
	}

	uri = "images/test.jpg"
	if GetFileExtension(uri) != ".jpg" {
		t.Error("Expected ext == '.jpg'")
	}
}
