package config

func (c *Config) defaultConfig() {
	c.Debug = false
	c.MaxTokens = 50
	c.Temperature = 0.6
	c.SystemMessage = "You are a {shell} shell command generator on the {os} operating system. Try to get paths from envs. You can use pipes as well. Provide the desired command without any explanations."
	c.SystemMessageDebug = "You are a {shell} shell command generator on the {os} operating system. Try to get paths from envs. You can use pipes as well. Explain your reasoning in plain english"
	c.PromptTemplate = "Current directory: '{path}'; Files: {files}; Task: '{input}'; Don't explanations."
	c.Model = "gpt-3.5-turbo"
	c.Color = true
}
