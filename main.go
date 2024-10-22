package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
)


func main()  {
	
    database.StartDb()
	routes.HandleRequest()
	
}
