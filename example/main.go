package main

import (
	"fmt"

	"github.com/kamildrazkiewicz/go-stanford-nlp"
)

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
