package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Hero struct
type Hero struct {
	family    string
	uid       int
	heroName  string `mapstructure:"hero_name"`
	firstName string `mapstructure:"first_name"`
	lastName  string `mapstructure:"last_name"`
}

//GetHeros handler returning all heros
func GetHeros(ctx *gin.Context) {
	file, err := ioutil.ReadFile("data/heros.json")

	if err != nil {
		log.Fatal(err)
	}
	var data interface{}
	_ = json.Unmarshal([]byte(file), &data)
	ctx.JSON(http.StatusOK, data)
}

//GetHeroByID handler return hero by uid
func GetHeroByID(ctx *gin.Context) {

}
