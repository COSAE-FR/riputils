package packages

import "github.com/COSAE-FR/riputils/pfsense/configuration/helpers"

type SquidConfig struct {
	Enable     helpers.OnOffBool          `xml:"enable_squid"`
	Interfaces helpers.CommaSeparatedList `xml:"active_interface"`
	ProxyPort  helpers.StringPort         `xml:"proxy_port"`
	ICPPort    helpers.StringPort         `xml:"icp_port"`
	KeepData   helpers.OnOffBool          `xml:"keep_squid_data"`
}

type SquidCacheConfig struct {
	CacheReplacementPolicy string                    `xml:"cache_replacement_policy"`
	CacheSwapLow           uint64                    `xml:"cache_swap_low"`
	CacheSwapHigh          uint64                    `xml:"cache_swap_high"`
	DoNotCache             helpers.TrueIfPresentBool `xml:"do_not_cache"`
	EnableOffline          helpers.TrueIfPresentBool `xml:"enable_offline"`
	//ext_cache_manager: <ext_cachemanager></ext_cachemanager>
	HarddiskCacheSize       uint64                    `xml:"harddisk_cache_size"`
	HarddiskCacheSystem     string                    `xml:"harddisk_cache_system"`
	Level1Subdirs           uint64                    `xml:"level1_subdirs"`
	HarddiskCacheLocation   string                    `xml:"harddisk_cache_location"`
	MinimumObjectSize       uint64                    `xml:"minimum_object_size"`
	MaximumObjectSize       uint64                    `xml:"maximum_object_size"`
	MemoryCacheSize         uint64                    `xml:"memory_cache_size"`
	MaximumObjSizeInMem     uint64                    `xml:"maximum_objsize_in_mem"`
	MemoryReplacementPolicy string                    `xml:"memory_replacement_policy"`
	CacheDynamicContent     helpers.TrueIfPresentBool `xml:"cache_dynamic_content"`
	CustomRefreshPatterns   string                    `xml:"custom_refresh_patterns"`
}

type SquidReverseGeneralConfig struct {
	// General settings
	Interfaces           helpers.CommaSeparatedList     `xml:"reverse_interface"`
	Ips                  helpers.SemiColonSeparatedList `xml:"reverse_ip"`
	ExternalFqdn         string                         `xml:"reverse_external_fqdn"`
	ResetUnauthorizedTcp helpers.OnOffBool              `xml:"deny_info_tcp_reset"`

	// Http settings
	EnableHttp      helpers.OnOffBool `xml:"reverse_http"`
	HttpPort        uint16            `xml:"reverse_http_port"`
	HttpDefaultSite string            `xml:"reverse_http_defsite"`

	// Https settings
	EnableHttps                  helpers.OnOffBool `xml:"reverse_https"`
	HttpsPort                    uint16            `xml:"reverse_https_port"`
	HttpsDefaultSite             string            `xml:"reverse_https_defsite"`
	TlsCert                      string            `xml:"reverse_ssl_cert"`
	IntermediateCa               string            `xml:"reverse_int_ca"`
	IgnoreInternalCertValidation helpers.OnOffBool `xml:"reverse_ignore_ssl_valid"`
	CheckClientCert              helpers.OnOffBool `xml:"reverse_check_clientca"`
	ClientCa                     string            `xml:"reverse_ssl_clientca"`
	ClientCrl                    string            `xml:"reverse_ssl_clientcrl"`

	// Security settings
	TlsCompatibilityMode string            `xml:"reverse_compatibility_mode"`
	TlsDhParamSize       uint16            `xml:"dhparams_size"`
	TlsDisableResumption helpers.OnOffBool `xml:"disable_session_reuse"`

	// OWA settings
	EnableOwa             helpers.OnOffBool              `xml:"reverse_owa"`
	OwaFrontendIps        helpers.SemiColonSeparatedList `xml:"reverse_owa_ip"`
	OwaEnableActiveSync   helpers.OnOffBool              `xml:"reverse_owa_activesync"`
	OwaEnableRpc          helpers.OnOffBool              `xml:"reverse_owa_rpchttp"`
	OwaEnableMapi         helpers.OnOffBool              `xml:"reverse_owa_mapihttp"`
	OwaEnableWebservice   helpers.OnOffBool              `xml:"reverse_owa_webservice"`
	OwaEnableAutodiscover helpers.OnOffBool              `xml:"reverse_owa_autodiscover"`
}

type SquidReversePeerConfig struct {
	Enable      helpers.OnOffBool `xml:"enable"`
	Name        string            `xml:"name"`
	Ip          string            `xml:"ip"`
	Port        uint16            `xml:"port"`
	Protocol    string            `xml:"prorocol"`
	Description string            `xml:"description"`
}

type SquidReverseUriConfig struct {
	Enable      helpers.OnOffBool          `xml:"enable"`
	Name        string                     `xml:"name"`
	Description string                     `xml:"description"`
	Peers       helpers.CommaSeparatedList `xml:"peers"`
	Uris        []string                   `xml:"row>uri"`
}
