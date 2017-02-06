# Go-Stanford-NLP

Go interface for Stanford NLP POS Tagger


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
		"ext/english-left3words-distsim.tagger",
		"ext/stanford-postagger.jar")

	if res, err := tagger.Tag("What is your name?"); err == nil {
		fmt.Println(res)
	}
}
```

Output will be:
```
map[is:VBZ your:PRP$ name:NN ?:. What:WP]
```