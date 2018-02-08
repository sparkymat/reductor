package config

type HeadConfig struct {
	CssIncludes []string `toml:"css_includes"`
}

type BodyConfig struct {
	AppId string `toml:"app_id"`
}

type Type struct {
	PageTitle string `toml:"page_title"`
	Port      int
	Head      HeadConfig
	Body      BodyConfig
}
