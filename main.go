package main

import (
	"beHealthyBE/api"
	"beHealthyBE/db"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	db.InitDatabase()
}
func main() {
	router := gin.Default()
	f, fErr := os.Create("gin.log")
	if fErr != nil {
		log.Fatal(fErr)
		return
	}

	gin.DefaultWriter = io.MultiWriter(f)
	v1 := router.Group("/v1")
	v1.Use(cors.Default())
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	}

	recipe := v1.Group("/recipe")
	{
		recipe.GET("/", api.GetRecipes)
		recipe.POST("/", api.PostRecipe)
		recipe.GET("/random/:count", api.GetRandomRecipe)
		recipe.PUT("/:id", api.UpdateRecipe)
		recipe.GET("/:id", api.GetRecipeByID)
		recipe.DELETE("/", api.DeleteAllRecipes)
		recipe.DELETE("/:id", api.DeleteRecipeByID)
	}

	srv := &http.Server{
		Addr:    "0.0.0.0:8888",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	ctx.Done()
	log.Println("Server exiting")
}
