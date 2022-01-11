package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/HDG-Gabriel/goyt/utils"
)

type Video struct {
	Title      string
	Cleantitle string
	Thumb      string
	Length     string
	Lengthsec  string
	Author     string
	Link       map[string][]string
}

func main() {
	id := flag.String("i", "", "(id) takes the videos' url.")
	url := flag.String("u", "", "(url) uri from video.")
	isHD := flag.Bool("hd", false, "(hd) download video in 720p. Default is 360p.")
	filename := flag.String("n", "", "(filename) file's name after downloaded.")
	help := flag.Bool("h", false, "(help) prints help menu")
	flag.Parse()

	if *help {
		utils.ShowBanner()
	}

	startSearch(*id, *url, *filename, *isHD)
}

// Starts doing requests to the api
func startSearch(id, urlVideo, filename string, isHD bool) {
	var base string
	if id != "" {
		base = "https://ytstream-download-youtube-videos.p.rapidapi.com/dl?id=" + id + "&geo=BR"
	} else {
		id = utils.GetID(urlVideo)
		base = "https://ytstream-download-youtube-videos.p.rapidapi.com/dl?id=" + id + "&geo=BR"
	}

	req, _ := http.NewRequest("GET", base, nil)

	req.Header.Add("x-rapidapi-host", "ytstream-download-youtube-videos.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "32116b9979msh9d3a64a4f678dcfp153c38jsn2a2aeb7bc25d")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("I'm sorry, something went wrong:", err)
		os.Exit(1)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	video := &Video{}
	json.Unmarshal(body, video)

	utils.ShowDownloadInfo(video.Title, video.Author)

	var keyQuality string
	if isHD {
		keyQuality = "22"
	} else {
		keyQuality = "18"
	}

	if filename == "" {
		filename = video.Cleantitle
	}

	err = downloadFile(video.Link[keyQuality][0], "/"+filename+".mp4")
	if err != nil {
		panic(err)
	}
	fmt.Println("Sucess to download!")
}

func downloadFile(url, filename string) error {
	res, err := http.Get(url)
	if err != nil {
		os.Exit(1)
		return err
	}
	defer res.Body.Close()

	out, err := os.Create(utils.GetDownloadsPath() + filename)
	if err != nil {
		os.Exit(1)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	return err
}
