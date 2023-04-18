package config

func (c *Config) applyDefaultValues() {
	c.Color = true
	c.Version = "1.2.1"
	if c.Debug == false {
		c.Debug = false
	}
	if c.MaxTokens == 0 {
		c.MaxTokens = 500
	}
	if c.Temperature == 0 {
		c.Temperature = 0.8
	}
	if c.PromptTemplate == "" {
		c.PromptTemplate = "Write a single line shell script for the task: {input}"
	}
	if c.Model == "" {
		c.Model = "gpt-3.5-turbo"
	}
	if c.SystemMessage == "" {
		c.SystemMessage = "You are an administrator who works in a {shell} on the {os} operating system."
	}
	if c.SystemMessageDebug == "" {
		c.SystemMessageDebug = "You are a {shell} shell command generator on the {os} operating system. Try to get paths from envs. You can use pipes as well. Explain your reasoning in plain english"
	}
}