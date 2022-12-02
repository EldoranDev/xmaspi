package controller

var (
	defaultBrightness = 128
	defaultDataPin    = 18
)

type Options struct {
	brightness int
	dataPin    int
}

func resolveOptions(options []Option) (*Options, error) {
	o := &Options{
		brightness: defaultBrightness,
		dataPin:    defaultDataPin,
	}

	for _, option := range options {
		err := option(o)
		if err != nil {
			return nil, err
		}
	}

	return o, nil
}

type Option func(*Options) error

func WithBrightness(brightness int) Option {
	return func(o *Options) error {
		o.brightness = brightness
		return nil
	}
}

func WithDataPin(pin int) Option {
	return func(o *Options) error {
		o.dataPin = pin
		return nil
	}
}
