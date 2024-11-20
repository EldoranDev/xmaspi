package rendering

import "errors"

var renderers = map[string]func() Renderer{
	StaticRenderer: func() Renderer {
		return &static{}
	},
	BlueWhiteStatic: func() Renderer { return &blueWhiteStatic{} },
}

func GetRenderer(name string) (Renderer, error) {
	getRenderer, ok := renderers[name]

	if !ok {
		return nil, errors.New("renderer can not be found")
	}

	return getRenderer(), nil
}

func GetAllRendererNames() []string {
	list := make([]string, 0)

	for key, _ := range renderers {
		list = append(list, key)
	}

	return list
}
