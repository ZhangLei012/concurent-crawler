package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	/* 获取内容写入文件中
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	file, err := os.Create(`citylist_test_data.html`)
	if err != nil {
		fmt.Println("文件创建失败：", err)
	}
	defer file.Close()
	_, err = file.WriteString(string(contents))
	if err != nil {
		t.Errorf("File write %s", err)
	}
	*/

	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	//使用t.Log/t.Logf进行输出，go test时需要使用-v参数
	//t.Logf("%s", contents)

	result := ParseCityList(contents)

	const resultSize = 470
	expectedURLs := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Requests))
	}
	for i, URL := range expectedURLs {
		if URL != result.Requests[i].URL {
			t.Errorf("expected url #%d: %s, but was %s", i, URL, result.Requests[i].URL)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Items))
	}
	for i, city := range expectedCities {
		if city != result.Items[i] {
			t.Errorf("expected url #%d: %s, but was %s", i, city, result.Items[i])
		}
	}
}
