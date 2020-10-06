package parser

import (
	"crawler-go/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "真的爱你")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1" +
			"element; but wa %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name: "真的爱你",
		Age: 33,
		Education: "中专2",
		Marriage: "未婚",
		Height: 175,
		Income: "12001-20000元",
		City: "阿坝",
	}

	if profile != expected {
		t.Errorf("excepted %v; but was %v", expected, profile)
	}
}
