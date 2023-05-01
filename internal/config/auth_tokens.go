package config

type FlashBlade struct {
        Address       string    `yaml:"address"`
        ApiToken      string    `yaml:"api_token"`
}

type FlashBladeList map[string]FlashBlade

func (f *FlashBladeList) GetArrayParams(fb string) (string, string) {
	for a_name, a := range *f {
                if a_name == fb {
			return a.Address, a.ApiToken
		}
	}
	return "", ""
}
