package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

//var occupationRe = regexp.MustCompile(``)
var genderRe = regexp.MustCompile(`"genderString":"([^"]*)"`)
var ageRe = regexp.MustCompile(`<div class="[^"]*" data-v-8b1eac0c>([0-9]+)岁</div>`)
var heightRe = regexp.MustCompile(`"heightString":"([^"]*)cm"`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([0-9]+)kg</div>`)
var maritalStatusRe = regexp.MustCompile(`"marriageString":"([^"]+)"`)
var workPlaceRe = regexp.MustCompile(`<div class="[^"]*" data-v-8b1eac0c>工作地:([^<]+)</div>`)
var inComeRe = regexp.MustCompile(`<div class="[^"]*" data-v-8b1eac0c>[^:]收入:([^<]+)</div>`)
var educationRe = regexp.MustCompile(`"educationString":"([^"]+)"`)
var nativePlaceRe = regexp.MustCompile(`<div class="[^"]*" data-v-8b1eac0c>籍贯:([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="[^"]*" data-v-8b1eac0c>([^<]*房[^<]*)</div>`)
var carRe = regexp.MustCompile(`<div class="[^"]*" data-v-8b1eac0c>([^<]*车[^<]*)</div>`)
var constellationRe = regexp.MustCompile(`<div class="[^"]*" data-v-8b1eac0c>([^<]*座[^<]*)</div>`)

//ParseProfile is to parse content to user profile info
func ParseProfile(content []byte, name string) engine.ParserResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(content, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(content, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(content, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Name = name
	profile.Gender = extractString(content, genderRe)
	profile.MaritalStatus = extractString(content, maritalStatusRe)
	profile.WorkingPlace = extractString(content, workPlaceRe)
	profile.Income = extractString(content, inComeRe)
	profile.Education = extractString(content, educationRe)
	profile.NativePlace = extractString(content, nativePlaceRe)
	profile.House = extractString(content, houseRe)
	profile.Car = extractString(content, carRe)
	profile.Constellation = extractString(content, constellationRe)

	result := engine.ParserResult{}
	result.Items = append(result.Items, profile)

	return result
}

func extractString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
