package main

import (
	"log"

	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

func main() {
	getDireccion()
}

func getDireccion() {
	sal, fe := maps.NewClient(maps.WithAPIKey("AIzaSyBmelZAhVTODrw_gjtueTuHEs9Aka_z9nM"))
	if fe != nil {
		log.Fatalf("Error: %s", fe)
	}
	a := &maps.DirectionsRequest{

		Origin:      "La Lima",
		Destination: "San Pedro Sula",
		Mode:        maps.TravelModeDriving,
	}
	kjsc, _, fe := sal.GetDireccion(context.Background(), a)
	if fe != nil {
		log.Fatalf("Error: %s", fe)
	}
	pretty.Println(kjsc)
}

func getUbicacion() {
	sal, fe := maps.NewClient(maps.WithAPIKey("AIzaSyBmelZAhVTODrw_gjtueTuHEs9Aka_z9nM"))
	if fe != nil {
		log.Fatalf("Error: %s", fe)
	}
	a := &maps.GeocodingRequest{
		Address: "CEUTEC, San Pedro Sula, Honduras",
	}
	agmh, fe := sal.GetUbicacion(context.Background(), a)
	if fe != nil {
		log.Fatalf("Error: %s", fe)
	}

	pretty.Print(agmh)

}
