package pos

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Descriptions - word tags description
// https://www.ling.upenn.edu/courses/Fall_2003/ling001/penn_treebank_pos.html
var Descriptions = map[string]string{
	"CC":   "Coordinating conjunction",
	"CD":   "Cardinal number",
	"DT":   "Determiner",
	"EX":   "Existential there",
	"FW":   "Foreign word",
	"IN":   "Preposition or subordinating conjunction",
	"JJ":   "Adjective",
	"JJR":  "Adjective, comparative",
	"JJS":  "Adjective, superlative",
	"LS":   "List item marker",
	"MD":   "Modal",
	"NN":   "Noun, singular or mass",
	"NNS":  "Noun, plural",
	"NNP":  "Proper noun, singular",
	"NNPS": "Proper noun, plural",
	"PDT":  "Predeterminer",
	"POS":  "Possessive ending",
	"PRP":  "Personal pronoun",
	"PRP$": "Possessive pronoun",
	"RB":   "Adverb",
	"RBR":  "Adverb, comparative",
	"RBS":  "Adverb, superlative",
	"RP":   "Particle",
	"SYM":  "Symbol",
	"TO":   "to",
	"UH":   "Interjection",
	"VB":   "Verb, base form",
	"VBD":  "Verb, past tense",
	"VBG":  "Verb, gerund or present participle",
	"VBN":  "Verb, past participle",
	"VBP":  "Verb, non-3rd person singular present",
	"VBZ":  "Verb, 3rd person singular present",
	"WDT":  "Wh-determiner",
	"WP":   "Wh-pronoun",
	"WP$":  "Possessive wh-pronoun",
	"WRB":  "Wh-adverb",
}

// Tagger struct
type Tagger struct {
	model     string
	tagger    string
	java      string
	opts      []string
	separator string
	encoding  string
}

// Result struct
type Result struct {
	Word string
	TAG  string
}

// Description - returns tag description
func (r *Result) Description() string {
	if _, exists := Descriptions[r.TAG]; !exists {
		return ""
	}
	return Descriptions[r.TAG]
}

// NewTagger - returns Tagger pointer
func NewTagger(m, t string) (*Tagger, error) {
	separator := ":"
	if runtime.GOOS == "windows" {
		separator = ";"
	}

	pos := &Tagger{
		java:      "java",
		encoding:  "utf8",
		opts:      []string{"-mx300m"},
		separator: separator,
	}

	if err := pos.SetModel(m); err != nil {
		return nil, err
	}
	if err := pos.SetTagger(t); err != nil {
		return nil, err
	}

	return pos, nil
}

// SetModel - set stanford pos tagger model
func (p *Tagger) SetModel(m string) error {
	if _, err := os.Stat(m); err != nil {
		return errors.New("Model not exists!")
	}
	p.model = m

	return nil
}

// SetTagger - set stanford pos tagger jar file
func (p *Tagger) SetTagger(t string) error {
	if _, err := os.Stat(t); err != nil {
		return errors.New("Tagger not exists!")
	}
	p.tagger = t

	return nil
}

// SetJavaPath - set path to java executable file
func (p *Tagger) SetJavaPath(j string) {
	p.java = j
}

// SetJavaOpts - set java options (default: [mx300m])
func (p *Tagger) SetJavaOpts(opts []string) {
	p.opts = opts
}

// SetEncoding - set outupt encoding (default: utf8)
func (p *Tagger) SetEncoding(e string) {
	p.encoding = e
}

func (p *Tagger) parse(out string) []*Result {
	words := strings.Split(out, " ")

	res := make([]*Result, len(words))
	for i, word := range words {
		split := strings.Split(word, "_")
		res[i] = &Result{
			Word: split[0],
			TAG:  split[1],
		}
	}

	return res
}

// Tag - use stanford pos tagger to tag input sentence
func (p *Tagger) Tag(input string) ([]*Result, error) {
	var (
		tmp  *os.File
		err  error
		args []string
	)

	if tmp, err = ioutil.TempFile("", "nlptemp"); err != nil {
		return nil, err
	}
	defer os.Remove(tmp.Name())
	if _, err = tmp.WriteString(input); err != nil {
		return nil, err
	}

	args = append(p.opts, []string{
		"-cp",
		p.tagger + p.separator,
		"edu.stanford.nlp.tagger.maxent.MaxentTagger",
		"-model",
		p.model,
		"-textFile",
		tmp.Name(),
		"-encoding",
		p.encoding,
	}...)

	cmd := exec.Command(p.java, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s: %s", err, stderr.String())
	}

	return p.parse(out.String()), nil
}
