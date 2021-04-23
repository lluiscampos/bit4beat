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

func (api *Api) createRecordHandler(context *gin.Context) {
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
}

func (api *Api) showRecordHandler(context *gin.Context) {
	id := context.Param("id")
	record, err := api.store.ReadRecord(id)
	if err != nil {
		context.JSON(500, gin.H{"error": string(err.Error())})
	} else {
		context.JSON(200, record)
	}
}

func (api *Api) listRecordsHandler(context *gin.Context) {
	records, err := api.store.ListRecords()
	if err != nil {
		context.JSON(500, gin.H{"error": string(err.Error())})
	} else {
		context.JSON(200, records)
	}
}

func (api *Api) Serve() {
	router := gin.Default()

	router.POST("/record", api.createRecordHandler)

	router.GET("/record/:id", api.showRecordHandler)

	router.GET("/records", api.listRecordsHandler)

	router.Run(":8080")
}
