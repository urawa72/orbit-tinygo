package data

//go:generate json-ice --type=Output
type Output struct {
	Lat			string	`json:"lat"`
	Lon			string	`json:"lon"`
	Bat			int64 	`json:"bat"`
	Rs			int64 	`json:"rs"`
	Temp		string	`json:"temp"`
	Humi		string	`json:"humi"`
	Timestamp	int64	`json:"timestamp"`
}

func (o *Output) SerializeJSON() ([]byte, error) {
	return MarshalOutputAsJSON(o)
}
