package sensors

// map of functions to call via string
func SensorMap(s string) string {
	m := map[string]func() string {
	"GroveDigitalLightI2C1": groveDigitalLightI2C1,
	}

	r := m[s]()
	return r
}

func groveDigitalLightI2C1() string {
	return "/lux"
}