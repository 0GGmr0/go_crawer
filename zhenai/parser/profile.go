package parser

import (
	"crawler/engine"
	"regexp"
	"strconv"
	"crawler/model"
)

var (
	ageRe = regexp.MustCompile(`<td><span class="label">年龄: </span>([\d]+岁)</td>`) //年龄
	MarriageRe = regexp.MustCompile(`<td><span class="label">婚况: </span>([^<]+)</td>`) //婚况
)

//匹配用户页信息
func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents, MarriageRe)


	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

//一个抽象出来的给予content和re，匹配到符合的结果
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}