package dynamic

import (
	"github.com/spf13/viper"
	"time"
)

type DViper viper.Viper

// Viper is essentially repository for configurations
// Get can retrieve any value given the key to use
// Get has the behavior of returning the value associated with the first
// place from where it is set. Viper will check in the following order:
// override, flag, env, config file, key/value store, default
//
// Get returns an interface. For a specific value use one of the Get____ methods.
func Get(key string) DynamicInterface { return func() interface{} { return viper.Get(key) } }
func (v *DViper) Get(key string) DynamicInterface {
	return func() interface{} {
		vv := (*viper.Viper)(v)
		return vv.Get(key)
	}
}

// Returns the value associated with the key as a string
func GetString(key string) DynamicString { return func() string { return viper.GetString(key) } }
func (v *DViper) GetString(key string) DynamicString {
	return func() string {
		vv := (*viper.Viper)(v)
		return vv.GetString(key)
	}
}

// Returns the value associated with the key asa boolean
func GetBool(key string) DynamicBool { return func() bool { return viper.GetBool(key) } }
func (v *DViper) GetBool(key string) DynamicBool {
	return func() bool {
		vv := (*viper.Viper)(v)
		return vv.GetBool(key)
	}
}

// Returns the value associated with the key as an integer
func GetInt(key string) DynamicInt { return func() int { return viper.GetInt(key) } }
func (v *DViper) GetInt(key string) DynamicInt {
	return func() int {
		vv := (*viper.Viper)(v)
		return vv.GetInt(key)
	}
}

// Returns the value associated with the key as a float64
func GetFloat64(key string) DynamicFloat64 { return func() float64 { return viper.GetFloat64(key) } }
func (v *DViper) GetFloat64(key string) DynamicFloat64 {
	return func() float64 {
		vv := (*viper.Viper)(v)
		return vv.GetFloat64(key)
	}
}

// Returns the value associated with the key as time
func GetTime(key string) DynamicTime { return func() time.Time { return viper.GetTime(key) } }
func (v *DViper) GetTime(key string) DynamicTime {
	return func() time.Time {
		vv := (*viper.Viper)(v)
		return vv.GetTime(key)
	}
}

// Returns the value associated with the key as a duration
func GetDuration(key string) DynamicDuration {
	return func() time.Duration { return viper.GetDuration(key) }
}
func (v *DViper) GetDuration(key string) DynamicDuration {
	return func() time.Duration {
		vv := (*viper.Viper)(v)
		return vv.GetDuration(key)
	}
}

// Returns the value associated with the key as a slice of strings
func GetStringSlice(key string) DynamicStringSlice {
	return func() []string {
		return viper.GetStringSlice(key)
	}
}
func (v *DViper) GetStringSlice(key string) DynamicStringSlice {
	return func() []string {
		vv := (*viper.Viper)(v)
		return vv.GetStringSlice(key)
	}
}

// Returns the value associated with the key as a map of interfaces
func GetStringMap(key string) DynamicMap {
	return func() map[string]interface{} {
		return viper.GetStringMap(key)
	}
}
func (v *DViper) GetStringMap(key string) DynamicMap {
	return func() map[string]interface{} {
		vv := (*viper.Viper)(v)
		return vv.GetStringMap(key)
	}
}

// Returns the value associated with the key as a map of strings
func GetStringMapString(key string) DynamicStringMap {
	return func() map[string]string {
		return viper.GetStringMapString(key)
	}
}
func (v *DViper) GetStringMapString(key string) DynamicStringMap {
	return func() map[string]string {
		vv := (*viper.Viper)(v)
		return vv.GetStringMapString(key)
	}
}

// Returns the value associated with the key as a map to a slice of strings.
func GetStringMapStringSlice(key string) DynamicMapStringSlice {
	return func() map[string][]string {
		return viper.GetStringMapStringSlice(key)
	}
}
func (v *DViper) GetStringMapStringSlice(key string) DynamicMapStringSlice {
	return func() map[string][]string {
		vv := (*viper.Viper)(v)
		return vv.GetStringMapStringSlice(key)
	}
}

// Returns the size of the value associated with the given key
// in bytes.
func GetSizeInBytes(key string) DynamicBytes {
	return func() uint {
		return viper.GetSizeInBytes(key)
	}
}
func (v *DViper) GetSizeInBytes(key string) DynamicBytes {
	return func() uint {
		vv := (*viper.Viper)(v)
		return vv.GetSizeInBytes(key)
	}
}
