package api

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"github.com/lluiscampos/bit4beat.back/model"
	"github.com/lluiscampos/bit4beat.back/store"
)

type Api struct {
	store store.Store
}

func NewApi(store store.Store) *Api {
	return &Api{
		store: store,
	}
}

func (api *Api) Serve() {
	router := gin.Default()

	router.POST("/record", func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, gin.H{"error": string(err.Error())})
		}

		var record model.Record
		err = json.Unmarshal(jsonData, &record)
		if err != nil {
			context.JSON(500, gin.H{"error": string(err.Error())})
		}

		err = api.store.CreateRecord(&record)
		if err != nil {
			context.JSON(500, gin.H{"error": string(err.Error())})
		}
	})

	router.GET("/record/:id", func(context *gin.Context) {
		id := context.Param("id")
		record, err := api.store.ReadRecord(id)
		if err != nil {
			context.JSON(500, gin.H{"error": string(err.Error())})
		} else {
			context.JSON(200, record)
		}
	})

	router.GET("/records", func(context *gin.Context) {
		records, err := api.store.ListRecords()
		if err != nil {
			context.JSON(500, nil)
		} else {
			context.JSON(200, records)
		}
	})

	router.Run(":8080")
}
