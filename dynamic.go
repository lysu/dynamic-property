// Dynamic-Property enhance viper with API return dynamic(or lazy) value.
//
// Viper is essentially repository for configurations
// Get can retrieve any value given the key to use
// Get has the behavior of returning the value associated with the first
// place from where it is set. Viper will check in the following order:
// override, flag, env, config file, key/value store, default
// This package is base on viper, so we recommend to read viper document first at https://github.com/spf13/viper
//
// It's common for us to use viper with Etcd/consul or company's configuration management system.
// And let properties have ability to modify dynamically during running...
// Viper's origin API is get a `Value`, it's easy to lead us to save value into field or some cached-like hold
// which are not easy to modify. So we extend them and return a function/a lazy value/ dynamic value.
//
// because Viper do memory cached internal, so there is no performance degrade to use dynamic value
// Just replace viper.* with dynamic.* and take free to hold return value~
//
// At last, we recommend to see this example:
package dynamic

import "time"

// DynamicInt present lazy eval integer
type DynamicInt func() int

// DynamicInterface present lazy eval interface
type DynamicInterface func() interface{}

// DynamicString present lazy eval string
type DynamicString func() string

// DynamicBool present lazy eval boolean
type DynamicBool func() bool

// DynamicFloat64 present lazy eval float64
type DynamicFloat64 func() float64

// DynamicTime present lazy eval time.Time
type DynamicTime func() time.Time

// DynamicDuration present lazy eval time.Duration
type DynamicDuration func() time.Duration

// DynamicStringSlice present lazy eval []string
type DynamicStringSlice func() []string

// DynamicMap present lazy eval map[string]interface{}
type DynamicMap func() map[string]interface{}

// DynamicStringMap present lazy eval map[string]string
type DynamicStringMap func() map[string]string

// DynamicMapStringSlice present lazy eval map[string][]string
type DynamicMapStringSlice func() map[string][]string

// DynamicBytes present lazy eval uint
type DynamicBytes func() uint
