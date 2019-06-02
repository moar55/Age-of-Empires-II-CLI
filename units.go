package main

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
