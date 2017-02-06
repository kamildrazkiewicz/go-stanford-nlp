# Go-Stanford-NLP

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
	tagger := pos.NewPOSTagger(
		"ext/english-left3words-distsim.tagger", // path to model
		"ext/stanford-postagger.jar")            // path to jar tagger file

	if res, err := tagger.Tag("What is your name?"); err == nil {
		for _, r := range res {
			fmt.Println(r.Word, r.TAG, r.TAGDescription())
		}
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