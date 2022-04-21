package configer

import (
	"io"
	"net/url"
	"time"
)

// Configurable configurable interface
type Configurable interface {
	// ContainsKey returns true if the given key is present in the configuration.
	ContainsKey(key string) bool
	// GetString returns the string value of the configuration
	GetString(key string) (string, error)
	// GetStringSlice returns the string slice value of the configuration
	GetStringSlice(key string) ([]string, error)
	// GetInt returns the int value of the configuration
	GetInt(key string) (int, error)
	// GetIntSlice returns the int slice value of the configuration
	GetIntSlice(key string) ([]int, error)
	// GetInt32 returns the int32 value of the configuration
	GetInt32(key string) (int32, error)
	// GetInt32Slice returns the int32 slice value of the configuration
	GetInt32Slice(key string) ([]int32, error)
	// GetInt64 returns the int64 value of the configuration
	GetInt64(key string) (int64, error)
	// GetInt64Slice returns the int64 slice value of the configuration
	GetInt64Slice(key string) ([]int64, error)
	// GetBool returns the bool value of the configuration
	GetBool(key string) (bool, error)
	// GetBoolSlice returns the bool slice value of the configuration
	GetBoolSlice(key string) ([]bool, error)
	// GetUint returns the uint value of the configuration
	GetUint(key string) (uint, error)
	// GetUintSlice returns the uint slice value of the configuration
	GetUintSlice(key string) ([]uint, error)
	// GetUint32 returns the uint32 value of the configuration
	GetUint32(key string) (uint32, error)
	// GetUint32Slice returns the uint32 slice value of the configuration
	GetUint32Slice(key string) ([]uint32, error)
	// GetUint64 returns the uint64 value of the configuration
	GetUint64(key string) (uint64, error)
	// GetUint64Slice returns the uint64 slice value of the configuration
	GetUint64Slice(key string) ([]uint64, error)
	// GetFloat32 returns the float32 value of the configuration
	GetFloat32(key string) (float32, error)
	// GetFloat32Slice returns the float32 slice value of the configuration
	GetFloat32Slice(key string) ([]float32, error)
	// GetFloat64 returns the float64 value of the configuration
	GetFloat64(key string) (float64, error)
	// GetFloat64Slice returns the float64 slice value of the configuration
	GetFloat64Slice(key string) ([]float64, error)
	// GetDuration returns the time.Duration value of the configuration
	GetDuration(key string) (time.Duration, error)
	// GetTime returns the time.Time value of the configuration
	GetTime(key string) (time.Time, error)
	// GetSection returns the string map string slice value of the configuration
	GetSection(key string) (Configurable, error)
	// Get is a helper function to access the configuration value
	Get(key string) (any, error)
	// Set is a helper function to set the configuration value
	Set(key string, value any) error
}

// FileConfiguration is the configuration interface for the application
type FileConfiguration interface {
	// Configurable returns the value of the configuration
	Configurable
	// Load loads the configuration from the given path
	Load(paths string) error
	// LoadStream loads the configuration from the given stream reader
	LoadStream(reader io.Reader) error
	// LoadRemote loads the configuration from the given URL
	LoadRemote(url *url.URL) error
	// Save save the configuration to the given path files
	Save(paths string) error
	// SaveStream save the configuration to the given stream writer
	SaveStream(writer io.Writer) error
	// SaveRemote save the configuration to the given URL
	SaveRemote(url *url.URL) error
	// GetFileName return the configuration file name
	GetFileName() string
	// GetFilePath return the configuration file path
	GetFilePath() string
	// GetURL return configuration URL
	GetURL() *url.URL
	// SetURL set configuration URL
	SetURL(url *url.URL)
	// Merge merges the given configuration into the current one
	Merge(config FileConfiguration) error
	// Sync syncs the configuration changes
	Sync() error
	// Watch watches the configuration for changes
	Watch(paths ...string) error
	// Reload reload file configuration
	Reload() error
	// GetReloadStrategy return reloading strategy
	GetReloadStrategy() ReloadingStrategy
	// SetReloadStrategy set reloading strategy
	SetReloadStrategy(strategy ReloadingStrategy)
	// GetEncoder returns the configuration encoder
	GetEncoder() ConfigEncoder
	// GetDecoder returns the configuration decoder
	GetDecoder() ConfigDecoder
	// SetEncoder sets the configuration encoder
	SetEncoder(encoder ConfigEncoder)
	// SetDecoder sets the configuration decoder
	SetDecoder(decoder ConfigDecoder)
}

// ReloadingStrategy file configuration reloading strategy
type ReloadingStrategy interface {
	// SetConfiguration set file configuration
	SetConfiguration(fileConfig FileConfiguration)
	// Init init strategy
	Init() error
	// NeedReloading return if the configuration need reload
	NeedReloading() (bool, error)
	// ReloadingPerformed notify the FileConfiguration has been reloaded
	ReloadingPerformed() error
}

// ConfigCodec is the configuration codec interface
type ConfigCodec interface {
	// ConfigEncoder encodes the configuration
	ConfigEncoder
	// ConfigDecoder decodes the configuration
	ConfigDecoder
}

// ConfigDecoder is the interface that wraps the Encode method.
type ConfigDecoder interface {
	// Decode decodes the configuration from the given bytes.
	Decode(bytes []byte) (map[string]ConfigField, error)
}

// ConfigEncoder is the interface that wraps the Decode method.
type ConfigEncoder interface {
	// Encode encodes the configuration to bytes.
	Encode(configMap map[string]ConfigField) ([]byte, error)
}
