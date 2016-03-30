package dynamic_test

import (
	"github.com/lysu/dynamic-property"
	"gopkg.in/olivere/elastic.v3"
	"time"
	"net/http"
)

// This example show how to get config value and hold in conf value, and use it.
// during conf can freely cache or hold a long time, it will always get newest changed value.
func Example_dynamicProperty() {

	conf := struct {
		ESHost dynamic.DynamicStringSlice
		SoTimeout dynamic.DynamicDuration
	}{
		ESHost: dynamic.GetStringSlice("es.hosts"), // get config value use dynamic.*, and hold it in conf
		SoTimeout: dynamic.GetDuration("es.SoTimeout"),
	}

	client, _ := elastic.NewClient(
		elastic.SetURL(conf.ESHost()), // use hold value with xx(), just eval it at call time~
		createHttpClient(conf.SoTimeout()),
	)

	client.Get()

}

func createHttpClient(soTimeout time.Duration) *http.Client {
	esHttpClient := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:  soTimeout,
	}
	return esHttpClient
}