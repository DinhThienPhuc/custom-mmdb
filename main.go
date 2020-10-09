package main

import (
	"fmt"
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

	// List Facebook IPs
	facebookCIDRs := []string{
		"31.13.24.0/21",
		"31.13.64.0/19",
		"31.13.64.0/24",
		"31.13.69.0/24",
		"31.13.70.0/24",
		"31.13.71.0/24",
		"31.13.72.0/24",
		"31.13.73.0/24",
		"31.13.75.0/24",
		"31.13.76.0/24",
		"31.13.77.0/24",
		"31.13.78.0/24",
		"31.13.79.0/24",
		"31.13.80.0/24",
		"66.220.144.0/20",
		"66.220.144.0/21",
		"66.220.149.11/16",
		"66.220.152.0/21",
		"66.220.158.11/16",
		"66.220.159.0/24",
		"69.63.176.0/21",
		"69.63.176.0/24",
		"69.63.184.0/21",
		"69.171.224.0/19",
		"69.171.224.0/20",
		"69.171.224.37/16",
		"69.171.229.11/16",
		"69.171.239.0/24",
		"69.171.240.0/20",
		"69.171.242.11/16",
		"69.171.255.0/24",
		"74.119.76.0/22",
		"173.252.64.0/19",
		"173.252.70.0/24",
		"173.252.96.0/19",
		"204.15.20.0/22",
	}

	sreData := mmdbtype.Map{
		"platform": mmdbtype.Map{
			"geoname_id": mmdbtype.Uint32(7777777),
			"iso_code":   mmdbtype.String("FC"),
			"names": mmdbtype.Map{
				"de":    mmdbtype.String("Fa'c'ebook"),
				"en":    mmdbtype.String("Facebook"),
				"es":    mmdbtype.String("Facebookee"),
				"fr":    mmdbtype.String("Faceletbokuu"),
				"ja":    mmdbtype.String("ベトナム"),
				"ru":    mmdbtype.String("Вьетнам"),
				"zh-CN": mmdbtype.String("越南"),
			},
		},
	}

	for _, cidr := range facebookCIDRs {
		// Define and insert the new data.
		_, sreNet, err := net.ParseCIDR(cidr)
		if err != nil {
			log.Fatal(err)
		}

		if err := writer.InsertFunc(sreNet, inserter.TopLevelMergeWith(sreData)); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Success: ", cidr)
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
