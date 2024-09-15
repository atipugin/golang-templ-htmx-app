package domain

import "fmt"

type ValidationErrors map[string][]string

func (e ValidationErrors) Error() string {
	return fmt.Sprintf("found %d validation errors", len(e))
}

func (e ValidationErrors) Add(field string, msg string) {
	if _, ok := e[field]; !ok {
		e[field] = []string{}
	}

	e[field] = append(e[field], msg)
}

func (e ValidationErrors) Get(field string) []string {
	return e[field]
}

func (e ValidationErrors) Has(field string) bool {
	return len(e[field]) > 0
}

func (e ValidationErrors) Any() bool {
	return len(e) > 0
}
