package transform

import (
	"fmt"
	"regexp"
)

type RegexTransformer struct{}

func (r RegexTransformer) Apply(input string, params []string) (string, error) {
	if len(params) != 1 {
		return "", fmt.Errorf("regex requires 1 param")
	}

	re, err := regexp.Compile(params[0])
	if err != nil {
		return "", err
	}

	match := re.FindString(input)
	if match == "" {
		return "", fmt.Errorf("regex no match")
	}

	return match, nil
}
