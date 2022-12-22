package serde

type MchConfig struct {
	Data struct {
		ConfigurationID string                            `json:"configurationId"`
		ComponentMap    map[string]map[string]interface{} `json:"componentMap"`
	} `json:"data"`
}

func (c *MchConfig) GetString(section, config string) string {
	return c.Data.ComponentMap[section][config].(string)
}
