package led

var (
	defaultLedCount   = 50
	defaultBrightness = 128
	defaultDataPin    = 18
)

type Options struct {
	ledCount   int
	brightness int
	dataPin    int
}

func resolveOptions(options []Option) (*Options, error) {
	o := &Options{
		ledCount:   defaultLedCount,
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

// WithLedCount sets the number of leds that are connected to the controller
func WithLedCount(count int) Option {
	return func(o *Options) error {
		o.ledCount = count
		return nil
	}
}

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
