package nelson

import (
	"io/ioutil"
	"os"

	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v2"
)

func (n *Nelson) readConfig() error {
	b, err := ioutil.ReadFile(n.configPath)
	if err != nil {
		return multierror.Prefix(err, "config reading")
	}
	if err := yaml.Unmarshal(b, n.NelsonConfig); err != nil {
		return err
	}

	return nil
}

func (n *Nelson) writeConfig() error {
	b, err := yaml.Marshal(n.NelsonConfig)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(n.configPath, b, os.FileMode(0666)); err != nil {
		return err
	}
	return nil
}
