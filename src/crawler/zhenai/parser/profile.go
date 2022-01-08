package parser

import (
	"crawler/engine"
	"crawler/model"
)

//var ageRe = regexp.MustCompile("")

func ParseProfile(contents []byte, name string, gender string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender
	profile.Age = 1
	profile.Height = 22
	profile.Weight = 33
	profile.Income = "111-222"
	profile.Marriage = "未婚"

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
