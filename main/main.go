package main

import (
	"encoding/json"
	"fatality14/godp"
	"flag"
	"fmt"
	"log"
	"os"
)

// Main function for the CLI tool
func main() {
	//logger
	var err error
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
		os.Exit(1)
	}
	log.SetOutput(file)

	var (
		route     string
		fromFiles bool
		saveFiles bool
		logstats  bool
	)

	//cli flags
	flag.StringVar(&route, "api-route", "https://rdb.altlinux.org/api/export/branch_binary_packages", "route url")
	flag.BoolVar(&fromFiles, "from-files", false, "load packages from files if any")
	flag.BoolVar(&saveFiles, "save-files", false, "save loaded packages to files")
	flag.BoolVar(&logstats, "log-stats", false, "log branch comparsion stats")

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Println("example usage: ./godp --from-files sisyphus p10")
	}

	//cli comand execution
	var firstResp godp.APIResponse
	var secondResp godp.APIResponse

	firstBranch := flag.Args()[0]
	secondBranch := flag.Args()[1]

	if !fromFiles {
		//load from server
		log.Println("Loading data from server...")
		firstResp, err = godp.FetchPackagesFromAPI[godp.APIResponse](route, firstBranch)
		if err != nil {
			log.Println("Error fetching first branch data: ", err)
			os.Exit(1)
		}
		secondResp, err = godp.FetchPackagesFromAPI[godp.APIResponse](route, secondBranch)
		if err != nil {
			log.Println("Error fetching second branch data: ", err)
			os.Exit(1)
		}
		log.Println("Data loaded successfully from server.")
	} else {
		//load from files
		log.Println("Loading data from filesystem...")
		firstResp, err = godp.DeserializeData[godp.APIResponse](firstBranch)
		if err != nil {
			log.Println("Error loading first branch data:", err)
			os.Exit(1)
		}
		secondResp, err = godp.DeserializeData[godp.APIResponse](secondBranch)
		if err != nil {
			log.Println("Error loading second branch data:", err)
			os.Exit(1)
		}
		log.Println("Data loaded successfully from local files.")
	}

	if saveFiles {
		//save data from server
		log.Println("Saving data to filesystem...")
		err := godp.SerializeData(firstResp, firstBranch)
		if err != nil {
			log.Println("Error saving first branch data: ", err)
			os.Exit(1)
		}

		err = godp.SerializeData(secondResp, secondBranch)
		if err != nil {
			log.Println("Error saving second branch data:", err)
			os.Exit(1)
		}
		log.Println("Data saved successfully to local files.")
	}

	//check if there is anything to compare
	if len(firstResp.Packages) == 0 || len(secondResp.Packages) == 0 {
		log.Println("Data for both branches must be loaded first.")
	}

	log.Println("Comparing packages between", firstBranch, " and ", secondBranch)
	results := godp.ComparePackages(firstResp, secondResp)

	//comparsion stats
	if logstats {
		for _, result := range results {
			log.Printf("Architecture: %s\n", result.Arch)
			log.Printf("Packages in %s but not in %s: %d\n", firstBranch, secondBranch, len(result.InSecondNotInFirst))
			log.Printf("Packages in %s but not in %s: %d\n", secondBranch, firstBranch, len(result.InFirstNotInSecond))
			log.Printf("Packages with higher versions in %s: %d\n\n", firstBranch, len(result.HigherInFirst))
		}
	}

	res, err := json.Marshal(results)
	if err != nil {
		log.Println("Error marshaling results to JSON:", err)
		os.Exit(1)
	}
	fmt.Println(string(res))
}
