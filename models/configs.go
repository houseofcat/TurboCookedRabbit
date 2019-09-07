package models

// RabbitSeasoning represents the configuration values.
type RabbitSeasoning struct {
	PoolConfig      *PoolConfig                `json:"PoolConfig"`
	ConsumerConfigs map[string]*ConsumerConfig `json:"ConsumerConfigs"`
	PublisherConfig *PublisherConfig           `json:"PublisherConfig"`
}

// PoolConfig represents settings for creating/configuring pools.
type PoolConfig struct {
	ChannelPoolConfig    *ChannelPoolConfig    `json:"ChannelPoolConfig"`
	ConnectionPoolConfig *ConnectionPoolConfig `json:"ConnectionPoolConfig"`
}

// ChannelPoolConfig represents settings for creating channel pools.
type ChannelPoolConfig struct {
	ErrorBuffer             uint16 `json:"ErrorBuffer"`
	BreakOnInitializeError  bool   `json:"BreakOnInitializeError"`  // error out on first error during Initialize.
	MaxInitializeErrorCount uint16 `json:"MaxInitializeErrorCount"` // error out after n errors during Initialize.
	SleepOnErrorInterval    uint32 `json:"SleepOnErrorInterval"`    // sleep length on errors
	CreateChannelRetryCount uint16 `json:"CreateChannelRetryCount"`
	ChannelCount            uint64 `json:"ChannelCount"`
	AckChannelCount         uint64 `json:"AckChannelCount"`
	GlobalQosCount          int    `json:"GlobalQosCount"` // Leave at 0 if you want to ignore them.
}

// ConnectionPoolConfig represents settings for creating connection pools.
type ConnectionPoolConfig struct {
	URI                        string     `json:"URI"`
	ErrorBuffer                uint16     `json:"ErrorBuffer"`
	BreakOnInitializeError     bool       `json:"BreakOnInitializeError"`     // error out on first error during Initialize.
	MaxInitializeErrorCount    uint16     `json:"MaxInitializeErrorCount"`    // error out after n errors during Initialize.
	SleepOnErrorInterval       uint32     `json:"SleepOnErrorInterval"`       // sleep length on errors
	EnableTLS                  bool       `json:"EnableTLS"`                  // Use TLSConfig to create connections with AMQPS uri.
	CreateConnectionRetryCount uint16     `json:"CreateConnectionRetryCount"` // when creating connections, add a retry, ignored when 0
	ConnectionCount            uint64     `json:"ConnectionCount"`            // number of connections to create in the pool
	TLSConfig                  *TLSConfig `json:"TLSConfig"`                  // TLS settings for connection with AMQPS.
}

// TLSConfig represents settings for configuring TLS.
type TLSConfig struct {
	PEMCertLocation   string `json:"PEMCertLocation"`
	LocalCertLocation string `json:"LocalCertLocation"`
	CertServerName    string `json:"CertServerName"`
}

// ConsumerConfig represents settings for configuring a consumer with ease.
type ConsumerConfig struct {
	QueueName            string                 `json:"QueueName"`
	ConsumerName         string                 `json:"ConsumerName"`
	AutoAck              bool                   `json:"AutoAck"`
	Exclusive            bool                   `json:"Exclusive"`
	NoWait               bool                   `json:"NoWait"`
	Args                 map[string]interface{} `json:"Args"`
	QosCountOverride     int                    `json:"QosCountOverride"` // if zero ignored
	MessageBuffer        uint32                 `json:"MessageBuffer"`
	ErrorBuffer          uint32                 `json:"ErrorBuffer"`
	SleepOnErrorInterval uint32                 `json:"SleepOnErrorInterval"`
}

// PublisherConfig represents settings for configuring global settings for all Publishers with ease.
type PublisherConfig struct {
	SleepOnIdleInterval uint32 `json:"SleepOnIdleInterval"`
	LetterBuffer        uint32 `json:"LetterBuffer"`
	NotificationBuffer  uint32 `json:"NotificationBuffer"`
}

// TopologyConfig allows you to build a simple toplogy from Json.
type TopologyConfig struct {
}
