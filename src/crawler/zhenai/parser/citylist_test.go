package parser

import (
	"crawler/zhenai/parser"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {

	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parser.ParseCityList(contents)
	const resultListSize = 470
	expectedItems := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	if len(result.Items) != resultListSize {
		t.Errorf("result should have %d items; but had %d",
			resultListSize, len(result.Items))
	}
	for i, item := range expectedItems {
		if result.Items[i].(string) != item {
			t.Errorf("excepted item #%d: %s; but was %s", i, item, result.Items[i].(string))
		}
	}

	if len(result.Requests) != resultListSize {
		t.Errorf("result should have %d requests; but had %d",
			resultListSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("excepted url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}

}
