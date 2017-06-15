package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"time"
)

// Image ges json image struct
type Image struct {
	ID       int    `json:"image_id"`
	URL      string `json:"image_url"`
	MimeType string `json:"image_mime_type"`
	Tw       int    `json:"tw"`
	Th       int    `json:"th"`
	Nw       int    `json:"nw"`
	Nh       int    `json:"nh"`
	W        int    `json:"w"`
	H        int    `json:"h"`
}
type Job struct {
	ID    string
	ImgId int
	URL   string
}
type JobResult struct {
	ID        string
	ImgId     int
	ImageConf image.Config
}

// Images ges json images struct
type Images map[string][]Image

var client = http.Client{}

func worker(id int, jobs <-chan Job, results chan<- JobResult) {
	for job := range jobs {
		imageConf := getImgSize("https://image.63pokupki.ru/images/" + job.URL[0:2] + "/" + job.URL)
		results <- JobResult{job.ID, job.ImgId, imageConf}
	}
}
func getImgSize(url string) image.Config {
	reqImg, _ := client.Get(url)
	buffer := make([]byte, reqImg.ContentLength)
	reqImg.Body.Read(buffer)
	im, _, err := image.DecodeConfig(bytes.NewReader(buffer))
	if err != nil {
		fmt.Println(err)
	}
	return im
}

func main() {
	start := time.Now()
	dat, _ := os.Open("./images.json")
	defer dat.Close()

	var imagesRaw map[string]*json.RawMessage
	images := make(Images)

	if err := json.NewDecoder(dat).Decode(&imagesRaw); err != nil {
		fmt.Println(err)
		return
	}

	jobs := make(chan Job, 100)
	results := make(chan JobResult, 100)
	// This starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 4; w++ {
		go worker(w, jobs, results)
	}

	var countJobs = len(imagesRaw)
	for id, imageRaw := range imagesRaw {
		var image []Image
		if error := json.Unmarshal(*imageRaw, &image); error != nil {
			fmt.Println(error)
			return
		}
		if len(image) > 1 {
			countJobs = countJobs + len(image) - 1
		}

		for ImgID, img := range image {
			jobs <- Job{id, ImgID, img.URL}
		}
		images[id] = image
	}

	fmt.Println(countJobs)
	close(jobs)

	for r := 1; r <= countJobs; r++ {
		result := <-results
		img := images[result.ID]
		img[result.ImgId].W = result.ImageConf.Width
		img[result.ImgId].H = result.ImageConf.Height
		images[result.ID] = img
	}

	fmt.Println(images)
	elapsed := time.Since(start)
	fmt.Println("Time:", elapsed)
}
