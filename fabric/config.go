package fabric

type SdkConfig struct {
	Channel     string
	User        string
	Org         string
	ConfigFile  string
	ChainCodeId string
}

func NewSdkConfig(channel, user, org, configFile, chainCodeId string) *SdkConfig{
	s := &SdkConfig {
		Channel     : channel,
		User        : user,
		Org         : org,
		ConfigFile  : configFile,
		ChainCodeId : chainCodeId,
	}
	return s
}
