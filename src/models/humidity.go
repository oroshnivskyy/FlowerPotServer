package models
import (
	"time"
	"strconv"
)

type HumidityRecord struct{
	Id    string   `gorethink:"id,omitempty"`
	Name  string   `gorethink:"name"`
	Data []*Humidity  `gorethink:"data"`
	DateCreated time.Time `gorethink:"date_created"`
	ClientIp string `gorethink:"client_ip"`
}

type Humidity struct{
	Id uint16 `gorethink:"id"`
	Humidity uint16 `gorethink:"humidity"`
}

func NewCleanRecord() (*HumidityRecord){
	humidityRecord := new(HumidityRecord)
	humidityRecord.Data = make([]*Humidity, 20)
	return humidityRecord
}

func NewRecord(data []string, name string, clientIP string) (*HumidityRecord, error){
	humidityRecord := new(HumidityRecord)
	humidityRecord.Name = name
	humidityRecord.DateCreated = time.Now()
	humidityRecord.ClientIp = clientIP

	allHumidities := make([]*Humidity, len(data))
	for index, humidity := range data{
		uint32humidity, err := strconv.ParseUint(humidity, 16, 0)
		if err!=nil{
			continue
		}
		allHumidities[index] = new(Humidity)
		allHumidities[index].Id = uint16(index)
		allHumidities[index].Humidity = uint16(uint32humidity)
	}
	humidityRecord.Data = allHumidities
	return humidityRecord, nil
}