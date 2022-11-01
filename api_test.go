package main_test

import (
	"beHealthyBE/api"
	"beHealthyBE/db"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var router *gin.Engine
var testRecipes []db.Recipe

func deleteAllRecipes() {
	db.GetDB().Delete(&db.Recipe{})
}

var recipe db.Recipe
var recipeMsg string

func TestMain(m *testing.M) {
	router = gin.Default()
	recipe = db.Recipe{
		Name:        "Test",
		Author:      "Go_Test",
		AuthorID:    "testing123",
		Ingredients: "milk,sugar,eggs,vanilla",
		Details:     "inghetata",
		Portions:    10,
		Preparation: 10,
		Cooking:     10,
		Tools:       "go",
	}
	recipeMsg = `{"ID":"%s","Name":"Test","Author":"Go_Test","AuthorID":"testing123","Ingredients":"milk,sugar,eggs,vanilla","Details":"inghetata","Portions":10,"Preparation":10,"Cooking":10,"Tools":"go"}`
	code := m.Run()
	deleteAllRecipes()
	os.Exit(code)

}
func TestGetRecipeByID(t *testing.T) {

	db.GetDB().Create(&recipe)
	testRecipes = append(testRecipes, recipe)
	router.GET("/v1/recipe/:id", api.GetRecipeByID)

	req, _ := http.NewRequest("GET", "/v1/recipe/"+recipe.ID, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	responseData, _ := ioutil.ReadAll(w.Body)
	response := string(responseData)
	recipeResponse := fmt.Sprintf(recipeMsg, recipe.ID)
	assert.Equal(t, response, recipeResponse)

}

func TestGetAllRecipes(t *testing.T) {
	deleteAllRecipes()
	db.GetDB().Create(&recipe)

	router.GET("/v1/recipe/", api.GetRecipes)
	req, _ := http.NewRequest("GET", "/v1/recipe/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	responseData, _ := ioutil.ReadAll(w.Body)
	response := string(responseData)
	preExpectedResponse := fmt.Sprintf("[%s]", recipeMsg)
	expectedResponse := fmt.Sprintf(preExpectedResponse, recipe.ID)
	assert.Equal(t, response, expectedResponse)
}

// [{"ID":"b7eea277-efe6-459c-8101-d53a3e5b328d","Name":"Test","Author":"Go_Test","AuthorID":"testing123","Ingredients":"milk,sugar,eggs,vanilla","Details":"inghetata","Portions":10,"Preparation":10,"Cooking":10,"Tools":"go"},{"ID":"0da9e6f3-9f89-43db-91aa-a22ac9337656","Name":"Test","Author":"Go_Test","AuthorID":"testing123","Ingredients":"milk,sugar,eggs,vanilla","Details":"inghetata","Portions":10,"Preparation":10,"Cooking":10,"Tools":"go"}]

// [{"ID":"0da9e6f3-9f89-43db-91aa-a22ac9337656","Name":"Test","Author":"Go_Test","AuthorID":"testing123","Ingredients":"milk,sugar,eggs,vanilla","Details":"inghetata","Portions":10,"Preparation":10,"Cooking":10,"Tools":"go"}]
