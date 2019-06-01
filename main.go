package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Unit data type, a struct containing all
// data of an Age of Empires unit
type Unit struct {
	ID                                int `json:"id"`
	Name, Description, Expansion, Age string
	CreatedIn                         string `json:"created_in"`
	Cost                              struct {
		Wood int
		Gold int
	} `json:"cost"`
	BuildTime       int     `json:"build_time"`
	ReloadTime      float32 `json:"reload_time"`
	AttackDelay     float32 `json:"attack_delay"`
	MovementRate    float32 `json:"movement_rate"`
	LineOfSight     int     `json:"line_of_sight"`
	HitPoints       int     `json:"hit_points"`
	MyRange         int     `json:"range"`
	Attack          int
	Armor, Accuracy string
}

func get(url string) *http.Response {
	resp, err := http.Get("https://age-of-empires-2-api.herokuapp.com/api/v1/unit/1")
	if err != nil {
		panic("oops")
	}
	return resp
}

func main() {
	resp := get("https://age-of-empires-2-api.herokuapp.com/api/v1/unit/1")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var myUnit Unit
	json.Unmarshal(body, &myUnit)
	fmt.Println(myUnit)

}
