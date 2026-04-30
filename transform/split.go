package transform

import (
	"fmt"
	"strconv"
	"strings"
)

type SplitTransformer struct{}

func (s SplitTransformer) Apply(input string, params []string) (string, error) {
	if len(params) != 2 {
		return "", fmt.Errorf("split requires delimiter and index")
	}

	delimiter := params[0]

	index, err := strconv.Atoi(params[1])
	if err != nil {
		return "", err
	}

	parts := strings.Split(input, delimiter)

	if index < 1 || index > len(parts) {
		return "", fmt.Errorf("split index out of range")
	}

	return strings.TrimSpace(parts[index-1]), nil
}
