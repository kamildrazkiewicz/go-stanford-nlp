# Go-Stanford-NLP

[![License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/kamildrazkiewicz/go-stanford-nlp/blob/master/LICENSE)  [![GoReport](https://goreportcard.com/badge/github.com/kamildrazkiewicz/go-stanford-nlp)](https://goreportcard.com/report/github.com/kamildrazkiewicz/go-stanford-nlp) 

Go interface for Stanford NLP POS Tagger

More info: http://nlp.stanford.edu/software/tagger.shtml


## Install

Install the package with:

```bash
go get github.com/kamildrazkiewicz/go-stanford-nlp
```

Import it with:

```go
import "github.com/kamildrazkiewicz/go-stanford-nlp"
```

and use `pos` as the package name inside the code.

## Example

```go
func main() {
	var (
		tagger *pos.POSTagger
		err    error
		res    []*pos.Result
	)

	if tagger, err = pos.NewPOSTagger(
		"ext/english-left3words-distsim.tagger",    // path to model
		"ext/stanford-postagger.jar"); err != nil { // path to jar tagger file
		fmt.Print(err)
		return
	}
	if res, err = tagger.Tag("What is your name?"); err != nil {
		fmt.Print(err)
		return
	}
	for _, r := range res {
		fmt.Println(r.Word, r.TAG, r.TAGDescription())
	}

}
```

Output will be:
```
What WP Wh-pronoun
is VBZ Verb, 3rd person singular present
your PRP$ Possessive pronoun
name NN Noun, singular or mass
? .
```
