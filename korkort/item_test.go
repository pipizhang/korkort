package korkort

import (
	//"fmt"
	"testing"
)

func TestIQuestion(t *testing.T) {

	iq := IQuestion{}

	iq.ParseContent(" 58) What does it mean if this lamp lights up on the dashboard? ")
	if iq.Content != "What does it mean if this lamp lights up on the dashboard?" {
		t.Error("Failed to parse IQuestion.Content")
	}

	_info := "+ Show explanation Vehicle knowledge, manoeuvring (no. 925)"
	iq.ParseOrignalID(_info)
	if iq.OriginalID != 925 {
		t.Error("Failed to parse IQuestion.OriginalID")
	}

	iq.ParseCategory(_info)
	if iq.Category != "Vehicle knowledge, manoeuvring" {
		t.Error("Failed to parse IQuestion.Category")
	}
}

func TestIChoice(t *testing.T) {
	var raw string
	var ic = IChoice{}

	raw = "✓ It will be 9 times greater."
	ic.Parse(raw)
	if ic.Content != "It will be 9 times greater." {
		t.Error("Failed to parse IChoice.Content")
	}
	if ic.Status != 1 {
		t.Error("Failed to parse IChoice.Status")
	}

	raw = " ✗ Only third-party insurance."
	ic.Parse(raw)
	if ic.Content != "Only third-party insurance." {
		t.Error("Failed to parse IChoice.Content")
	}
	if ic.Status != 0 {
		t.Error("Failed to parse IChoice.Status")
	}
}
