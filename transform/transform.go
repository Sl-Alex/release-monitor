package transform

import (
	"fmt"

	"release-monitor/model"
	"release-monitor/app_context"
)

type Transformer interface {
	Apply(input string, params []string) (string, error)
}

var registry = map[string]Transformer{
	"split": SplitTransformer{},
	"regex": RegexTransformer{},
}

func Apply(ctx app_context.Context, input string, transforms []model.Transform) (string, error) {
	current := input

	for _, t := range transforms {
		transformer, ok := registry[t.Type]
		if !ok {
			return "", fmt.Errorf("unknown transform: %s", t.Type)
		}

		var err error

        app_context.Debug(ctx, "before transform: %s", current)
        current, err = transformer.Apply(current, t.Params)
        app_context.Debug(ctx, "after %s transform: %s", t.Type, current)
		if err != nil {
			return "", err
		}
	}

	return current, nil
}
