package parser

import (
	"crawler-go/engine"
	"crawler-go/model"
	"regexp"
	"strconv"
)

//var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>未婚</div>`)
//var heightRe = regexp.MustCompile( `<div class="m-btn purple" data-v-8b1eac0c>未婚</div>`)
//var weightRe = regexp.MustCompile( `<div class="m-btn purple" data-v-8b1eac0c>未婚</div>`)
//var incomeRe = regexp.MustCompile( `<div class="m-btn purple" data-v-8b1eac0c>未婚</div>`)
//var genderRe = regexp.MustCompile( `<div class="m-btn purple" data-v-8b1eac0c>未婚</div>`)
//var carRe = regexp.MustCompile( `<div class="m-btn purple" data-v-8b1eac0c>未婚</div>`)
//var deucationRe = regexp.MustCompile( `<div class="m-btn purple" data-v-8b1eac0c>未婚</div>`)
//var hokouRe = regexp.MustCompile( `<div class="m-btn purple" data-v-8b1eac0c>未婚</div>`)
//var houseRe = regexp.MustCompile( `<div class="m-btn purple" data-v-8b1eac0c>未婚</div>`)
//var marriageRe = regexp.MustCompile( `<div class="m-btn purple" data-v-8b1eac0c>未婚</div>`)
//var occupationRe = regexp.MustCompile( `<div class="m-btn purple" data-v-8b1eac0c>未婚</div>`)
//var xinzuoRe = regexp.MustCompile( `<div class="m-btn purple" data-v-8b1eac0c>未婚</div>`)

var ageRe = regexp.MustCompile(`<div[^>]*>(\d+)岁</div>`)
var heightRe = regexp.MustCompile(`<div[^>]*>(\d+)cm</div>`)
var incomeRe = regexp.MustCompile(`]*>月收入:([^>]+)`)
var weightRe = regexp.MustCompile(`]*>([\d]+)kg`)
var genderRe = regexp.MustCompile(`"genderString":"([^"]*)"`)
var xinzuoRe = regexp.MustCompile(`]*>([^>]+座)[^>]*`)
var marriageRe = regexp.MustCompile(`]*>([离异|丧偶|未婚]+)`)
var educationRe = regexp.MustCompile(`"educationString":"([^"]*)"`)
var occupationRe = regexp.MustCompile(``)
var hokouRe = regexp.MustCompile(`]*>籍贯:([^>]+)`)
var houseRe = regexp.MustCompile(`]*>([^>]+房)`)
var carRe = regexp.MustCompile(`]*>([^>]+车)`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

// 阿坝 | 33岁 | 中专 | 未婚 | 175cm | 12001-20000元
var profileRe = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^>]*[^<])</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		profile.Weight = weight
	}

	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Car = extractString(contents, carRe)
	profile.Education = extractString(contents, educationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.House = extractString(contents, houseRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func ParseProfile3(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		profile.Weight = weight
	}

	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Car = extractString(contents, carRe)
	profile.Education = extractString(contents, educationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.House = extractString(contents, houseRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)

	//profileSplit := strings.Split(extractString(contents, profileRe), "|")
	////阿坝 | 33岁 | 中专 | 未婚 | 175cm | 12001-20000元
	//profile.City = profileSplit[0]
	//re := regexp.MustCompile(`^((1[0-5])|[1-9])?\d$`)
	//age, err := strconv.Atoi(re.FindString(profileSplit[1]))
	//if err != nil {
	//	profile.Age = age
	//}
	//profile.Education = profileSplit[2]
	//profile.Marriage = profileSplit[3]
	//height, err := strconv.Atoi(re.FindString(profileSplit[4]))
	//if err != nil {
	//	profile.Height = height
	//}
	//profile.Income = profileSplit[5]

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func ParseProfile2(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	//找第一个
	match := ageRe.FindSubmatch(contents)

	if match != nil {
		age, err := strconv.Atoi(string(match[1]))
		if err != nil {
			// user age is age
			profile.Age = age
		}
	}

	////下面的内容提取出去
	//match = marriageRe.FindSubmatch(contents)
	//
	//if match != nil {
	//	profile.City = string(match[0])
	//}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	//fmt.Printf("match: %s\n", match[1])

	if len(match) >= 2 {
		return string(match[0])
	} else {
		return ""
	}
}