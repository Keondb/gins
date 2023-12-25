package config

type QiNiu struct {
	Enable    bool    `json:"enable" yaml:"enable"` // 是否启用
	AccessKey string  `json:"access_key" yaml:"access_key"`
	SecretKey string  `json:"secret_key" yaml:"secret_key"`
	Bucket    string  `json:"bucket" yaml:"bucket"` //存储桶的名字
	CDN       string  `json:"cdn" yaml:"cdn"`
	Zone      string  `json:"zone" yaml:"zone"`
	Size      float64 `json:"size" yaml:"size"`
}
