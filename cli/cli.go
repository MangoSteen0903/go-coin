package cli

import (
	"flag"
	"fmt"
	"os"

	explorer "github.com/mangosteen0903/go-coin/explorer/templates"
	"github.com/mangosteen0903/go-coin/rest"
)

func usage() {
	fmt.Printf("Welcome to Go-Coin! \n\n")
	fmt.Printf("Please use the following flags: \n\n")
	fmt.Printf("-port:		Set port of the server \n")
	fmt.Printf("-mode:		Choose 'rest' or 'html' or 'both' \n\t\t if you choose 'both', Rest API and HTML Explorer will run together. \n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}
	port := flag.Int("port", 5000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose 'rest' or 'html' or 'both' \n if you choose 'both', Rest API and HTML Explorer will run together.")

	flag.Parse()
	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "both":
		go rest.Start(*port)
		explorer.Start(*port + 50)
	default:
		usage()
	}
	fmt.Println(*port, *mode)
}
