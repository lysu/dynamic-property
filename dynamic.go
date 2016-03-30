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
