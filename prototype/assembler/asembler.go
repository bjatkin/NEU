package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type tokenMatcher interface {
	match(string) bool
	value() []byte
}

type asembler struct {
	tokenMatchers []tokenMatcher
}

func (a *asembler) addMatcher(matcher tokenMatcher) {
	a.tokenMatchers = append(a.tokenMatchers, matcher)
}

func (a *asembler) execute(filepath string) ([]byte, error) {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	code := string(raw)
	fmt.Println("code: ", code)

	var newByteCode []byte
	lines := strings.Split(code, "\n")
	for _, line := range lines {
		for _, token := range strings.Split(line, " ") {
			var matched bool
			for _, matcher := range a.tokenMatchers {
				if matcher.match(token) {
					newByteCode = append(newByteCode, matcher.value()...)
					matched = true
					break
				}
			}
			if !matched {
				log.Fatalln("could not find a valid match for " + token)
			}
		}
	}

	return newByteCode, nil
}
