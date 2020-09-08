package store

type AcmOptions struct {
	AccessKey   string `json:"access_key"`
	SecretKey   string `json:"secret_key"`
	Endpoint    string `json:"endpoint"`
	ServiceName string `json:"service_name"`
	NamespaceID string `json:"namespace_id"`
	LogDir      string `json:"log_dir"`
	CacheDir    string `json:"cache_dir"`
}


