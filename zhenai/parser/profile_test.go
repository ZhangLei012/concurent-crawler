package parser

import (
	"crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	/* 获取内容写入文件中
	contents, err := fetcher.Fetch("https://album.zhenai.com/u/1476375288")
	file, err := os.Create(`profile_test_data.html`)
	if err != nil {
		fmt.Println("文件创建失败：", err)
	}
	defer file.Close()
	_, err = file.WriteString(string(contents))
	if err != nil {
		t.Errorf("File write %s", err)
	}
	*/

	content, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	profile := ParseProfile(content, "左岸定心").Items[0]

	rightProfile := model.Profile{
		Name:          "左岸定心",
		Gender:        "男士",
		Age:           20,
		Height:        175,
		MaritalStatus: "未婚",
		WorkingPlace:  "阿坝九寨沟",
		Income:        "3-5千",
		Education:     "高中及以下",
		Car:           "未买车",
		Constellation: "魔羯座(12.22-01.19)",
	}
	t.Logf("rightProfile = %v, profile = %v", rightProfile, profile)
	if profile != rightProfile {
		t.Errorf("expected profile %v, but had %v", rightProfile, profile)
	}
}
