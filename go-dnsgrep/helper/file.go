package helper

import (
	"os"
	"log"
)
// WriteDomainsToFile writes the domain to output file
func WriteDomainsToFile(outputFile string , domains []string){
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for _,domain:=range domains{
	 if _, err := f.Write([]byte(domain+"\n")); err != nil {
		 log.Fatal(err)
	 }
	}
  
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("Domains written to",outputFile)
}