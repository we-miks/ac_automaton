package ac_automaton

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"
)

func TestACAutomaton(t *testing.T) {
	ac := NewACAutomaton()
	ss := []string{"a", "aa", "aaa"}
	for _, s := range ss {
		ac.Insert(s)
	}
	ac.Build()
	s := `aaaaa`
	res := ac.FindMatches(s)
	t.Log(res)
}

type Sample struct {
	Content string `json:"content"`
	Keyword string `json:"keyword"`
}

func TestACAutomatonProfile(t *testing.T) {
	sampleFilename := "ac_sample.json"

	file, err := os.Open(sampleFilename)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() { _ = file.Close() }()
	reader := bufio.NewReader(file)
	var line string

	contents := make([]string, 0)
	keywords := make([]string, 0)
	ac := NewACAutomaton()
	sample := new(Sample)
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			t.Error(err)
			return
		}

		if err := json.Unmarshal([]byte(line), sample); err != nil {
			t.Error(err)
			return
		}
		ac.Insert(sample.Keyword)
		contents = append(contents, sample.Content)
		keywords = append(keywords, sample.Keyword)
	}

	ac.Build()
	for _, content := range contents {
		res := ac.FindMatches(content)
		for _, keyword := range keywords {
			cnt := strings.Count(content, keyword)
			if cnt != res[keyword] {
				t.Error("not equal")
				t.Log(content)
				t.Log(keyword)
				t.Log(cnt)
				t.Log(res[keyword])
				return
			}
		}
	}
}
