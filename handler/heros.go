package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Hero struct
type Hero struct {
	Family    string
	UID       int
	HeroName  string `json:"hero_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

//GetHeros handler returning all heros
func GetHeros(ctx *gin.Context) {
	file, err := ioutil.ReadFile("data/heros.json")

	if err != nil {
		log.Fatal(err)
	}

	var data []Hero
	json.Unmarshal([]byte(file), &data)
	ctx.JSON(http.StatusOK, data)
}

//GetHeroByID handler return hero by uid
func GetHeroByID(ctx *gin.Context) {
	file, err := ioutil.ReadFile("data/heros.json")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	var data []Hero
	json.Unmarshal([]byte(file), &data)
	var selectedHero Hero
	for i := range data {
		if data[i].UID == id {
			selectedHero = data[i]
		}
	}

	ctx.JSON(http.StatusOK, selectedHero)
}
