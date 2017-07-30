package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := []string{`div|id="one"|class="two"`, `div|id="A"|class="B"`, `h1`, `p`}
	s, err := parseInput(input)
	if err != nil {
		t.Fail()
	}
	if strings.Join(input, " ") != s.String() {
		t.Fail()
	}
}

func TestParseInputError(t *testing.T) {
	input := []string{"div|"}
	_, err := parseInput(input)
	if err == nil {
		t.Fail()
	}
	if !strings.Contains(err.Error(), "wrong input") {
		t.Fail()
	}
}

func TestXmlSelect(t *testing.T) {
	tests := []struct {
		selector, xml string
	}{
		{
			"p",
			"<p>ok</p>",
		},
		{
			`div|id="one"|class="two" p`,
			`<div id="one" class="two">
				<p>ok</p>
			</div>
			<div id="one">
				<p>fail</p>
			</div>`,
		},
		{
			`div|id="one"|class="two" div|id="A"|class="B" p`,
			`<div id="one" class="two">
				<div id="A" class="B">
					<p>ok</p>
					<div id="A">
						<p>ok</p>
					</div>
				</div>
			</div>
			<div id="one">
				<p>fail</p>
			</div>`,
		},
	}
	for _, test := range tests {
		sel, err := parseInput(strings.Split(test.selector, " "))
		if err != nil {
			t.Fatalf("parse failed: %s", err.Error())
		}
		r := strings.NewReader(test.xml)
		w := &bytes.Buffer{}
		xmlSelect(w, r, sel)
		res := w.String()
		fmt.Print(res)
		if !strings.HasSuffix(res, "ok\n") {
			t.Errorf("%s", w.String())
		}
	}
}
