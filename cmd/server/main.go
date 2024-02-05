package main

import (
	"log"

	"github.com/chista.dev/internal"
	"github.com/chista.dev/pkg"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){

	db := pkg.GetRepo().Db

	// migrate the schema
	db.AutoMigrate(&pkg.User{},&pkg.ApiKey{},&pkg.Service{},&pkg.UserService{})

	//seed
	router := gin.Default()
	router.Use(func(c *gin.Context){
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file")
		}

		c.Next()
	})


	router.GET("/seed", func(c *gin.Context) {
		// seed on dev
		pkg.GetRepo().Seed()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	externalV1 := router.Group("/api/v1").Use(authenticate)
	{
		externalV1.POST("/chistadotdev",internal.Chista)
	}

	router.Run()

}



func authenticate(c *gin.Context){

	clientApikey :=  c.Query("apikey")

	if clientApikey == "" {
		c.JSON(401,gin.H{"error":"Unauthorized"})
		c.Abort()
		return
	}else{

	
		apikey,err := pkg.GetRepo().GetApiKey(clientApikey)

		if err != nil {
			c.JSON(401,gin.H{"error" : "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userId",apikey.UserID)
		c.Set("apiKey",apikey.ApiKey)

		c.Next()


	}

}


// func extractBearerToken(header string) (string, error) {
// 	if header == "" {
// 		return "", errors.New("bad header value given")
// 	}

// 	jwtToken := strings.Split(header, " ")
// 	if len(jwtToken) != 2 {
// 		return "", errors.New("incorrectly formatted authorization header")
// 	}

// 	return jwtToken[1], nil
// }