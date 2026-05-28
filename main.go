package main

import (
	"flag"
	"fmt"
)

func main() {
	helpS3 := flag.Bool("help", false, "Display help information")
	lsS3 := flag.Bool("list", false, "List all buckets")
	cbS3 := flag.Bool("create", false, "Create new S3 bucket(s)")
	dbS3 := flag.Bool("bktd", false, "Delete S3 bucket(s)")
	doS3 := flag.Bool("objd", false, "Delete S3 object(s)")
	getS3 := flag.Bool("get", false, "Download S3 object(s)")
	putS3 := flag.Bool("put", false, "Upload S3 object(s)")
	lgS3 := flag.Bool("large", false, "Upload large S3 object(s)")
	flag.Parse()

	if flag.NFlag() > 0 {
		client = createClient()

		switch {
		case *helpS3:
			logo()
			credits()
			help()
		case *cbS3:
			createBucket(client, flag.Arg(0))
		case *dbS3:
			deleteBucket(client, flag.Arg(0))
		case *doS3:
			deleteObject(client, flag.Arg(0), flag.Arg(1))
		case *getS3:
			downloadObject(client, flag.Arg(0), flag.Arg(1), flag.Arg(2))
		case *lsS3:
			listBuckets(client)
		case *lgS3:
			uploadLargeObject(client, flag.Arg(0), flag.Arg(1), flag.Arg(2))
		case *putS3:
			uploadObject(client, flag.Arg(0), flag.Arg(1), flag.Arg(2))
		default:
			inform("Arguments not recognized - ")
		}
	} else {
		alert("No arguments found - ")
	}

}

// Print help information for using the program
func help() {
	Yellow.Println("\nUsage:")
	fmt.Println("  cerberus [COMMAND] [ARGUMENTS...]")
	Yellow.Println("\nCommands:")
	Green.Printf("%s", "  -help")
	fmt.Println("		Display help information")
	Green.Printf("%s", "  -list")
	fmt.Println("		List buckets")
	Green.Printf("%s", "  -create")
	fmt.Println("	Create new S3 bucket(s)")
	Green.Printf("%s", "  -bktd")
	fmt.Println("		Delete S3 bucket(s)")
	Green.Printf("%s", "  -objd")
	fmt.Println("		Delete S3 object(s)")
	Green.Printf("%s", "  -get")
	fmt.Println("		Download S3 object(s)")
	Green.Printf("%s", "  -put")
	fmt.Println("		Upload S3 object(s)")
	Green.Printf("%s", "  -large")
	fmt.Println("	Upload large S3 object(s)")
	Yellow.Println("\nExample:")
	fmt.Println("  Adding your path to file if necessary, run:")
	Green.Printf("%s", "    cerberus ")
	Yellow.Printf("%s", "-put ")
	fmt.Println("[BUCKET-NAME] [SOURCE] [DESTINATION]")
	Green.Printf("%s", "    cerberus ")
	Yellow.Printf("%s", "-put ")
	fmt.Println("my-bucket ~/Pictures/image.jpg photos/2024/image.jpg")
	Yellow.Println("\nHelp:")
	fmt.Println("  For more information go to:")
	Green.Println("    https://github.com/farghul/cerberus.git")
}

// ASCII art generated logo
func logo() {
	Red.Println(" ▗▄▄▖▗▄▄▄▖▗▄▄▖ ▗▄▄▖ ▗▄▄▄▖▗▄▄▖ ▗▖ ▗▖ ▗▄▄▖")
	Red.Println("▐▌   ▐▌   ▐▌ ▐▌▐▌ ▐▌▐▌   ▐▌ ▐▌▐▌ ▐▌▐▌   ")
	Red.Println("▐▌   ▐▛▀▀▘▐▛▀▚▖▐▛▀▚▖▐▛▀▀▘▐▛▀▚▖▐▌ ▐▌ ▝▀▚▖")
	Red.Println("▝▚▄▄▖▐▙▄▄▖▐▌ ▐▌▐▙▄▞▘▐▙▄▄▖▐▌ ▐▌▝▚▄▞▘▗▄▄▞▘")
	Red.Println(bv)
}

// Program tagline and creator credit
func credits() {
	fmt.Println("\nAn interactive CLI tool for S3 Object Storage")
	fmt.Println("Created by Byron Stuike")
	fmt.Println()
}
