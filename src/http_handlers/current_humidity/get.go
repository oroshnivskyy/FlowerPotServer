package current_humidity

import (
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/dancannon/gorethink"
	_ "models"
	"models"
	"strconv"
)

type HumidityPage struct{
	Humidities []*models.HumidityRecord
	CurrentPage int
	NextPage int
	PreviousPage int
	HasNext bool
}

func Get(session *gorethink.Session, w http.ResponseWriter, r *http.Request, render render.Render) {
	humidityPage := HumidityPage{}
	var offset int
	{
		strPage := r.URL.Query().Get("page")
		offset, _ = strconv.Atoi(strPage)
		if offset <= 0 {
			offset = 0
		}
	}
	limit := 10
	humidityPage.CurrentPage = offset*limit
	res, err := gorethink.Table("humidity").OrderBy(gorethink.Desc("date_created")).Skip(humidityPage.CurrentPage).Limit(limit).Run(session)
	if err != nil {
		render.HTML(404, "current_humidity", nil)
		return
	}

	humidityPage.Humidities = make([]*models.HumidityRecord, limit)
	for i := 0; i < limit; i++ {
		humidityPage.Humidities[i] = models.NewCleanRecord()
	}
	err = res.All(&humidityPage.Humidities)
	humidityPage.HasNext = false
	if len(humidityPage.Humidities) == limit {
		humidityPage.HasNext = true
	}

	if err != nil {
		render.HTML(404, "current_humidity", nil)
		return
	}

	humidityPage.NextPage = humidityPage.CurrentPage + 1
	humidityPage.PreviousPage = humidityPage.CurrentPage - 1

	render.HTML(200, "current_humidity", humidityPage)
}