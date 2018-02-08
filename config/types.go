package config

type HeadConfig struct {
	CssIncludes []string `toml:"css_includes"`
}

type BodyConfig struct {
	AppId string `toml:"app_id"`
}

type Type struct {
	AppName string `toml:"app_name"`
	Head    HeadConfig
	Body    BodyConfig
}
