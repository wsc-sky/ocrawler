package model

import 	"encoding/json"


type Profile struct {
	Id 			string
	Url 		string
	Name	    string
	Gender      string
	Age 	    int
	Height      int
	Weight      int
	Income      string
	Marriage    string
	House	string
	Location	string
	Education   string
	Occupation  string
	Xinzuo 		string
	Car 		string
}

func FromJsonObj(o interface{}) (Profile, error){
	var profile Profile

	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err

}