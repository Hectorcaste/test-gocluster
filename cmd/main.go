package main

import (
	"fmt"

	"github.com/Hectorcaste/test-gocluster/database"
	"github.com/Hectorcaste/test-gocluster/database/models"
	cluster "github.com/MadAppGang/gocluster"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Print("error reading")
	}
}

func main() {
	fmt.Printf("si inicia mi pez %s", "hola")
	c := cluster.NewCluster()
	c.PointSize = 60
	c.MaxZoom = 3
	c.TileSize = 256
	adapter := database.NewConnection()
	response := adapter.GetSucursales()
	fmt.Printf("the response %+v \n", response)
	geoPoints := make([]cluster.GeoPoint, len(*response))
	for index := range *response {
		current := (*response)[index]
		currentTestpoint := TestPoint{
			Branch: current,
		}
		geoPoints[index] = &currentTestpoint
	}
	err := c.ClusterPoints(geoPoints)
	fmt.Printf("the error %v ", err)

	northWest := SimplePoint{-74.120651, 4.696991}
	southEast := SimplePoint{-74.051900, 4.646335}

	result := c.GetClusters(&northWest, &southEast, 0)
	fmt.Printf("result %+v \n", result)
}

// /////////////////////////////////////////////////////////
type SimplePoint struct {
	Lon, Lat float64
}

func (tp *SimplePoint) GetCoordinates() cluster.GeoCoordinates {
	return cluster.GeoCoordinates{
		Lon: tp.Lon,
		Lat: tp.Lat,
	}
}

type TestPoint struct {
	Branch models.BranchModel
}

func (tp *TestPoint) GetCoordinates() cluster.GeoCoordinates {
	return cluster.GeoCoordinates{
		Lon: tp.Branch.BranchLongitude,
		Lat: tp.Branch.BranchLatitude,
	}
}
