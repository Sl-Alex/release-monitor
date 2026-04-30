package transform

import (
	"fmt"

	"release-monitor/model"
)

type Transformer interface {
	Apply(input string, params []string) (string, error)
}

var registry = map[string]Transformer{
	"split": SplitTransformer{},
	"regex": RegexTransformer{},
}

func Apply(input string, transforms []model.Transform) (string, error) {
	current := input

	for _, t := range transforms {
		transformer, ok := registry[t.Type]
		if !ok {
			return "", fmt.Errorf("unknown transform: %s", t.Type)
		}

		var err error
		current, err = transformer.Apply(current, t.Params)
		if err != nil {
			return "", err
		}
	}

	return current, nil
}
