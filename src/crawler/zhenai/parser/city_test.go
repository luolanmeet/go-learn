package parser

import (
	"crawler/zhenai/parser"
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parser.ParseCity(contents)

	expectedItems := []string{
		"User 希恩",
	}
	expectedUrls := []string{
		"http://album.zhenai.com/u/1858121739",
	}

	for i, item := range expectedItems {
		if result.Items[i].(string) != item {
			t.Errorf("excepted item #%d: %s; but was %s", i, item, result.Items[i].(string))
		}
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("excepted url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}
}
