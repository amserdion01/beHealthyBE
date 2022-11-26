package api

import (
	"beHealthyBE/db"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateRecipe(ctx *gin.Context) {
	var recipe db.Recipe
	if err := ctx.BindJSON(&recipe); err != nil {
		log.Printf("updateRecipe: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if !(recipe.ID == ctx.Param("id")) && (len(recipe.ID) > 0) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot update ID",
		})
		return
	}
	db.GetDB().Model(&db.Recipe{}).Where("id=?", ctx.Param("id")).Updates(&recipe)
	var updatedRecipe db.Recipe
	db.GetDB().Where("id=?", ctx.Param("id")).Find(&updatedRecipe)
	log.Println(updatedRecipe.AuthorID)
	ctx.JSON(http.StatusOK, updatedRecipe)

}
func GetRandomRecipe(ctx *gin.Context) {
	var recipes []db.Recipe
	result := db.GetDB().Find(&recipes)
	if result.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	rand.Seed(time.Now().Unix())
	ctx.JSON(http.StatusOK, recipes[rand.Intn(len(recipes)-1)])

}
func GetRecipeByID(ctx *gin.Context) {
	var recipe db.Recipe
	result := db.GetDB().Where("id=?", ctx.Param("id")).Find(&recipe)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, recipe)
}
func GetRecipes(ctx *gin.Context) {
	var recipes []db.Recipe
	result := db.GetDB().Find(&recipes)
	if result.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, recipes)
}
func DeleteAllRecipes(ctx *gin.Context) {
	var recipes []db.Recipe
	result := db.GetDB().Find(&recipes)
	if result.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	db.GetDB().Delete(recipes)
	ctx.JSON(http.StatusOK, recipes)
}
func DeleteRecipeByID(ctx *gin.Context) {
	var recipe db.Recipe
	result := db.GetDB().Where("id=?", ctx.Param("id")).Find(&recipe)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	db.GetDB().Delete(recipe)
	ctx.JSON(http.StatusOK, recipe)
}
func PostRecipe(ctx *gin.Context) {
	var newRecipe db.Recipe

	if err := ctx.BindJSON(&newRecipe); err != nil {
		return
	}
	var checkRecipe db.Recipe
	db.GetDB().Where(&db.Recipe{Name: newRecipe.Name, Author: newRecipe.Author}).First(&checkRecipe)
	if len(checkRecipe.ID) != 0 {
		ctx.JSON(http.StatusConflict, gin.H{"message": "Recipe already exists"})
		return
	}
	log.Println(newRecipe.Cooking.String())

	db.GetDB().Create(&newRecipe)

	ctx.JSON(http.StatusCreated, newRecipe)
}
