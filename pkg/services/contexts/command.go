package contexts

import (
	"encoding/json"
	"os/exec"

	"github.com/maxilian/panic-reloader/pkg/services/contexts/model"
)

type Config struct {
	*model.Config
}

func (c *Config) GetContext(name string) *model.Context {
	for _, ctx := range c.Contexts {
		if ctx.Name == name {
			return ctx
		}
	}

	return nil
}

func getConfig() *Config {
	cmd := exec.Command("kubectl", "config", "view", "-o", "json")
	res, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	config := &Config{}
	err = json.Unmarshal(res, config)
	if err != nil {
		panic(err)
	}
	return config
}

func useContext(c string) {
	cmd := exec.Command("kubectl", "config", "use-context", c)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
