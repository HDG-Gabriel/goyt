package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// Get id from url's video
func GetID(url string) string {
	return url[32 : 32+11]
}

// Returns user downloads path format: /home/user/Downloads/
func GetDownloadsPath() string {
	return filepath.Join(os.Getenv("HOME"), "/Downloads")
}

func ShowBanner() {
	banner := `________                __   
/  _____/  ____ ___.__._/  |_ 
/   \  ___ /  _ <   |  |\   __\
\    \_\  (  <_> )___  | |  |  
\______  /\____// ____| |__|  
       \/       \/  

Simple program that downloads videos from Youtube

Usage: goyt [OPTIONS]
	
Options:
    -u,    Takes in single URL as input and download video
    -i,    Takes in single ID as input and download video 
    -hd,   Downloads video in HD quality [DEFAULT 360p]
    -n,    Renames file after downloaded
    -h,    Prints help menu

Example:
- Download video:
$ goyt -u https://www.youtube.com/watch?v=Isim0ysZ6X4
- HD quality:
$ goyt -hd true -u https://www.youtube.com/watch?v=Isim0ysZ6X4
- Renamed file:
$ goyt -n hawking -u https://www.youtube.com/watch?v=Isim0ysZ6X4`

	fmt.Println(banner)
	os.Exit(0)
}

func ShowDownloadInfo(name, author string) {
	fmt.Println("==================================================")
	fmt.Println("Video:", name)
	fmt.Println("Author:", author)
	fmt.Println("==================================================")
}
