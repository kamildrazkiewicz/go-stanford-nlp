package pos

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type POSTagger struct {
	model  string
	tagger string
	java   string
	opts   []string
}

func NewPOSTagger(m, t string) *POSTagger {
	return &POSTagger{
		model:  m,
		tagger: t,
		java:   "java",
		opts:   []string{"-mx300m"},
	}
}

func (p *POSTagger) SetModel(m string) {
	p.model = m
}

func (p *POSTagger) SetTagger(t string) {
	p.tagger = t
}

func (p *POSTagger) SetJavaPath(j string) {
	p.java = j
}

func (p *POSTagger) parse(out string) map[string]string {
	words := strings.Split(out, " ")

	res := make(map[string]string, len(words))
	for _, word := range words {
		split := strings.Split(word, "_")
		res[split[0]] = split[1]
	}

	return res
}

func (p *POSTagger) Tag(input string) (map[string]string, error) {
	var (
		tmp *os.File
		err error
	)

	if tmp, err = ioutil.TempFile("", "nlptemp"); err != nil {
		return nil, err
	}
	tmp.WriteString(input)

	args := []string{
		"-cp",
		p.tagger + ":",
		"edu.stanford.nlp.tagger.maxent.MaxentTagger",
		"-model",
		p.model,
		"-textFile",
		tmp.Name(),
		"-encoding",
		"utf8",
	}

	cmd := exec.Command(p.java, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		return nil, err
	}

	return p.parse(out.String()), nil
}
