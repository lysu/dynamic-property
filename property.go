package dynamic

import "time"

type DynamicInt func() int

type DynamicInterface func() interface{}

type DynamicString func() string

type DynamicBool func() bool

type DynamicFloat64 func() float64

type DynamicTime func() time.Time

type DynamicDuration func() time.Duration

type DynamicStringSlice func() []string

type DynamicMap func() map[string]interface{}

type DynamicStringMap func() map[string]string

type DynamicMapStringSlice func() map[string][]string

type DynamicBytes func() uint
