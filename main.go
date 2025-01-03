package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/inserter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
)

func main() {
	// Load the database we wish to enrich.
	writer, err := mmdbwriter.Load("maxmind-country-origin.mmdb", mmdbwriter.Options{})
	if err != nil {
		log.Fatal(err)
	}

	// Define json file path
	path := os.Args[1]

	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully opened json file!")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// Names struct
	type Names struct {
		De string `json:"de"`
		En string `json:"en"`
		Es string `json:"es"`
		Fr string `json:"fr"`
		Ja string `json:"ja"`
		Ru string `json:"ru"`
		Cn string `json:"zh-CN"`
	}

	// Platform struct
	type Platform struct {
		Name      string   `json:"name"`
		GeoNameID int      `json:"geoNameId"`
		IsoCode   string   `json:"isoCode"`
		Names     Names    `json:"names"`
		Cidrs     []string `json:"cidrs"`
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// initialize array
	var jsonData = []Platform{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'jsonData' which we defined above
	json.Unmarshal(byteValue, &jsonData)

	// we iterate through every platform within our platform array
	for i := 0; i < len(jsonData); i++ {
		fmt.Println("User Type: " + jsonData[i].Name)

		sreData := mmdbtype.Map{
			"platform": mmdbtype.Map{
				"geoname_id": mmdbtype.Uint32(jsonData[i].GeoNameID),
				"iso_code":   mmdbtype.String(jsonData[i].IsoCode),
				"names": mmdbtype.Map{
					"de":    mmdbtype.String(jsonData[i].Names.De),
					"en":    mmdbtype.String(jsonData[i].Names.En),
					"es":    mmdbtype.String(jsonData[i].Names.Es),
					"fr":    mmdbtype.String(jsonData[i].Names.Fr),
					"ja":    mmdbtype.String(jsonData[i].Names.Ja),
					"ru":    mmdbtype.String(jsonData[i].Names.Ru),
					"zh-CN": mmdbtype.String(jsonData[i].Names.Cn),
				},
			},
		}

		for _, cidr := range jsonData[i].Cidrs {
			// Define and insert the new data.
			_, sreNet, err := net.ParseCIDR(cidr)
			if err != nil {
				log.Fatal(err)
			}

			if err := writer.InsertFunc(sreNet, inserter.TopLevelMergeWith(sreData)); err != nil {
				log.Fatal(err)
			}

			fmt.Println(jsonData[i].Name, ": ", cidr)
		}
	}

	// Write the newly enriched DB to the filesystem.
	fh, err := os.Create("maxmind-country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	_, err = writer.WriteTo(fh)
	if err != nil {
		log.Fatal(err)
	}
}
