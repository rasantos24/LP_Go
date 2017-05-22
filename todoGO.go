package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kr/pretty"
	"golang.org/x/image/bmp"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

const GOOGLE_GEOLOCATION_API_KEY = "AIzaSyDHjrRQfVX98LLmvEMTsjLbrQvspG4QbOc"
const GOOGLE_PLACES_API_KEY = "AIzaSyDSOkL3swVzkqHc6A99qa791zD2h4qSQpY"

type Route struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

type Coordinates struct {
	Longitude string `json:"long"`
	Latitude  string `json:"lat"`
}

type Place struct {
	Origin string `json:"origin"`
}

type MiniBMP struct {
	Nombre string `json:"nombre"`
	Size   Size   `json:size`
}

type AntiBMP struct {
	Nombre string `json:"nombre"`
}

type Size struct {
	Alto  int `json:alto`
	Ancho int `json:ancho`
}

func main() {
	agmh := mux.NewRouter()
	agmh.HandleFunc("/ejercicio1", Create).Methods("POST")
	agmh.HandleFunc("/ejercicio2", GetDetalles).Methods("POST")
	agmh.HandleFunc("/ejercicio3", Mini).Methods("POST")
	agmh.HandleFunc("/ejercicio4", Anti).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", agmh))
}

func GetDetalles(w http.ResponseWriter, req *http.Request) {

	var route Route
	_ = json.NewDecoder(req.Body).Decode(&route)

	client, fe := maps.NewClient(maps.WithAPIKey(GOOGLE_GEOLOCATION_API_KEY))
	if fe != nil {
		log.Fatalf("fatal error: %s", fe)
	}

	r := &maps.GeocodingRequest{
		Address: route.Origin,
	}

	r2 := &maps.NearbySearchRequest{
		Radius: 500,
		Type:   maps.PlaceTypeRestaurant,
	}

	ll, fe := maps.ParseLatLng(route.Origin)
	pretty.Print(ll.Lat)

	resp, fe := client.Geocode(context.Background(), r)
	if fe != nil {
		log.Fatalf("fatal error: %s", fe)
	}

	resp2, fe := client.NearbySearch(context.Background(), r2)
	pretty.Print(resp)
	pretty.Println(resp2)
}

func Create(w http.ResponseWriter, req *http.Request) {
	var route Route
	_ = json.NewDecoder(req.Body).Decode(&route)

	client, fe := maps.NewClient(maps.WithAPIKey(GOOGLE_GEOLOCATION_API_KEY))
	if fe != nil {
		log.Fatalf("Fatal Error: %s", fe)
	}

	r := &maps.DirectionsRequest{

		Origin:      route.Origin,
		Destination: route.Destination,
		Mode:        maps.TravelModeDriving,
	}

	routes, _, fe := client.Directions(context.Background(), r)
	if fe != nil {
		log.Fatalf("Fatal Error: %s", fe)
	}

	json.NewEncoder(w).Encode(routes)
}

func Mini(w http.ResponseWriter, req *http.Request) {

	var img MiniBMP
	_ = json.NewDecoder(req.Body).Decode(&img)

	bitmap, fe := openImage(img.Nombre)
	if fe != nil {
		fmt.Println(fe)
	}

	pretty.Println("AGMH")
	bounds := bitmap.Bounds()
	ancho, altura := bounds.Max.X/img.Size.Alto, bounds.Max.Y/img.Size.Ancho
	myBMP := image.NewRGBA(image.Rect(0, 0, ancho, altura))

	for y := 0; y < img.Size.Alto; y++ {
		for x := 0; x < img.Size.Ancho; x++ {
			pixel := bitmap.At(x*ancho, y*altura)
			myBMP.Set(x, y, pixel)
		}
	}

	archiOUT, fe := os.Create("Mini_Kaiba.bmp")
	if fe != nil {
		fmt.Println(fe)
	}

	defer archiOUT.Close()

	pretty.Println(myBMP)
	bmp.Encode(archiOUT, myBMP)

	res, fe := openImage("Mini_Kaiba.bmp")
	if fe != nil {
		fmt.Println(fe)
	}

	json.NewEncoder(w).Encode(res.ColorModel())
	pretty.Println(res)
}

func Anti(w http.ResponseWriter, req *http.Request) {

	var img AntiBMP
	_ = json.NewDecoder(req.Body).Decode(&img)

	bitmap, fe := openImage(img.Nombre)
	if fe != nil {
		fmt.Println(fe)
	}

	limi := bitmap.Bounds()
	ancho, altura := limi.Max.X, limi.Max.Y

	myBMP := image.NewRGBA(limi)
	for y := 0; y < altura; y++ {
		for x := 0; x < ancho; x++ {
			pu := bitmap.At(x, y)
			red, green, blue, _ := pu.RGBA()
			rgb := (red + green + blue) / 3
			pi := color.Gray{uint8(rgb / 256)}
			myBMP.Set(x, y, pi)
		}
	}

	archiOUT, fe := os.Create("Kaiba_2.bmp")
	if fe != nil {
		fmt.Println(fe)
	}

	defer archiOUT.Close()

	pretty.Println(myBMP)
	bmp.Encode(archiOUT, myBMP)
}

func openImage(rafaam string) (image.Image, error) {
	file, fe := os.Open(rafaam)
	if fe != nil {
		return nil, fe
	}
	defer file.Close()
	return bmp.Decode(file)
}
