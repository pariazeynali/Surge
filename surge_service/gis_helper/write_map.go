package gis_helper

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// UpdateOSMData Removes old .osm.pbf file and download Irans's osm data
func UpdateOSMData() {
	osmFileName := os.Getenv("OSM_PBF_FILE")
	osmURL := os.Getenv("OSM_PBF_URL")
	if err := os.Remove(osmFileName); err != nil {
		log.Panicf("ERROR: flailed to remove osm.pbf file %v\n")
	}
	osmFile, err := os.Create(osmFileName)
	if err != nil {
		log.Panicf("ERROR: flailed to create osm.pbf file %v\n")
	}
	defer osmFile.Close()
	osmData, err := http.Get(osmURL)
	if err != nil {
		log.Panicf("ERROR: flailed to download osm.pbf file %v\n")
	}
	defer osmData.Body.Close()
	size, err := io.Copy(osmFile, osmData.Body)
	if err != nil {
		log.Panicf("ERROR: flailed to write osm data osm.pbf file %v\n")
	}
	log.Printf("Successfully downloaded and saved %v of OSM data", size)
}

// SaveOSMToDB Save data into postgresql database using osm2pgsql
func SaveOSMToDB() {
	osmDb := os.Getenv("OSM_DB")
	osmUser := os.Getenv("OSM_USER")
	osmHost := os.Getenv("OSM_HOST")
	osm2pgsqlStyleFile := os.Getenv("STYLE_FILE")
	osmFileName := os.Getenv("OSM_PBF_FILE")
	cmd := exec.Command("osm2pgsql", "-c", "-d "+osmDb, "-U"+osmUser, "-H"+osmHost, "-S"+osm2pgsqlStyleFile, osmFileName)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Println("Could not save osm data: ", err)
	}
}
