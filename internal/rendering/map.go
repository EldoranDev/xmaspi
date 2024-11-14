package rendering

import "errors"

var renderers = map[string]func() Renderer{
	"static": func() Renderer {
		return &static{}
	},
}

func GetRenderer(name string) (Renderer, error) {
	getRenderer, ok := renderers[name]

	if !ok {
		return nil, errors.New("renderer can not be found")
	}

	return getRenderer(), nil
}
