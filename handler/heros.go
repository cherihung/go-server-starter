package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	data := readHerosToFile()

	if len(ctx.Request.URL.Query()) > 0 {
		searchHerosByAnyName(ctx, data)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

//GetHeroByID handler return hero by uid: /hero/id/1
func GetHeroByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Fatal(err)
	}

	data := readHerosToFile()
	var selectedHero Hero
	for i := range data {
		if data[i].UID == id {
			selectedHero = data[i]
		}
	}

	ctx.JSON(http.StatusOK, selectedHero)
}

//GetHerosByFamily handler returns all heros by family name: /heros/family/marvel
func GetHerosByFamily(ctx *gin.Context) {
	term := strings.ToLower(ctx.Param("name"))

	data := readHerosToFile()
	var selectedHeros []Hero

	for i := range data {
		if strings.ToLower(data[i].Family) == term {
			selectedHeros = append(selectedHeros, data[i])
		}
	}

	ctx.JSON(http.StatusOK, selectedHeros)
}

/* PRIVATE */

func searchHerosByAnyName(ctx *gin.Context, data []Hero) {

	var selectedHeros []Hero
	var terms []string

	// ?name=anyname
	rawTerms := strings.ToLower(ctx.Query("name"))
	// ?name=firstname,lastname
	byRealName := strings.ContainsAny(",", rawTerms)
	// ["firstname", "lastname"]
	terms = strings.Split(rawTerms, ",")

	for i := range data {
		heroName := strings.ToLower(data[i].HeroName)
		lastName := strings.ToLower(data[i].LastName)
		firstName := strings.ToLower(data[i].FirstName)

		if byRealName {
			if firstName == terms[0] && lastName == terms[1] {
				selectedHeros = append(selectedHeros, data[i])
			}
		} else {
			if heroName == terms[0] || lastName == terms[0] || firstName == terms[0] {
				selectedHeros = append(selectedHeros, data[i])
			}
		}
	}

	ctx.JSON(http.StatusOK, selectedHeros)
}

func readHerosToFile() []Hero {
	file, err := ioutil.ReadFile("data/heros.json")

	if err != nil {
		log.Fatal("reading json file error: ", err)
	}

	var data []Hero
	json.Unmarshal([]byte(file), &data)
	return data
}
