package config

type HeadConfig struct {
	JavascriptIncludes []string `toml:"javascript_includes"`
	CssIncludes        []string `toml:"css_includes"`
	RawTags            []string `toml:"raw_tags"`
}

type BodyConfig struct {
	AppId string   `toml:"app_id"`
	Nodes []string `toml:"nodes"`
}

type Type struct {
	AppName string `toml:"app_name"`
	Head    HeadConfig
	Body    BodyConfig
}
