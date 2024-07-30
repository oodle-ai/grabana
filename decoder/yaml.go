package decoder

import (
	"io"

	"github.com/oodle-ai/grabana/dashboard"
	"gopkg.in/yaml.v3"
)

func UnmarshalYAML(input io.Reader) (dashboard.Builder, error) {
	decoder := yaml.NewDecoder(input)
	decoder.KnownFields(true)

	parsed := &DashboardModel{}
	if err := decoder.Decode(parsed); err != nil {
		return dashboard.Builder{}, err
	}

	return parsed.ToBuilder()
}
