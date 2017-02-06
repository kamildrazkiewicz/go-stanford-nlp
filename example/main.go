package main

import (
	"fmt"

	"github.com/kamildrazkiewicz/go-stanford-nlp"
)

func main() {
	tagger := pos.NewPOSTagger(
		"ext/english-left3words-distsim.tagger", // path to model
		"ext/stanford-postagger.jar")            // path to jar tagger file

	if res, err := tagger.Tag("What is your name?"); err == nil {
		for _, r := range res {
			fmt.Println(r.Word, r.TAG)
		}
	}

}
