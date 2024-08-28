package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type CatAPIController struct {
	web.Controller
}

type CatImage struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Breed struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Temperament string `json:"temperament"`
	Origin      string `json:"origin"`
	LifeSpan    string `json:"life_span"`
	Weight      struct {
		Metric string `json:"metric"`
	} `json:"weight"`
	Image struct {
		URL string `json:"url"`
	} `json:"image"`
}

type Vote struct {
	ImageID string `json:"image_id"`
	Value   int    `json:"value"`
}

type Favourite struct {
	ID      int      `json:"id"`
	UserID  string   `json:"user_id"`
	ImageID string   `json:"image_id"`
	Image   CatImage `json:"image"`
}

func (c *CatAPIController) GetRandomCat() {
	apiKey, _ := web.AppConfig.String("cat_api_key")
	url := "https://api.thecatapi.com/v1/images/search"

	catChan := make(chan CatImage)
	go fetchCatImage(url, apiKey, catChan)

	cat := <-catChan
	c.Data["json"] = cat
	c.ServeJSON()
}

func (c *CatAPIController) GetBreeds() {
	apiKey, _ := web.AppConfig.String("cat_api_key")
	url := "https://api.thecatapi.com/v1/breeds"

	breedsChan := make(chan []Breed)
	go fetchBreeds(url, apiKey, breedsChan)

	breeds := <-breedsChan
	c.Data["json"] = breeds
	c.ServeJSON()
}

func (c *CatAPIController) GetBreedInfo() {
	apiKey, _ := web.AppConfig.String("cat_api_key")
	breedID := c.Ctx.Input.Param(":id")
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=10", breedID)

	breedImagesChan := make(chan []CatImage)
	go fetchBreedImages(url, apiKey, breedImagesChan)

	breedImages := <-breedImagesChan
	c.Data["json"] = breedImages
	c.ServeJSON()
}

func (c *CatAPIController) VoteCat() {
	apiKey, _ := web.AppConfig.String("cat_api_key")
	url := "https://api.thecatapi.com/v1/votes"

	var vote Vote
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &vote); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	voteChan := make(chan bool)
	go submitVote(url, apiKey, vote, voteChan)

	success := <-voteChan
	if success {
		c.Data["json"] = map[string]string{"message": "Vote submitted successfully"}
	} else {
		c.Data["json"] = map[string]string{"error": "Failed to submit vote"}
	}
	c.ServeJSON()
}

func (c *CatAPIController) GetFavourites() {
	apiKey, _ := web.AppConfig.String("cat_api_key")
	url := "https://api.thecatapi.com/v1/favourites"

	favouritesChan := make(chan []Favourite)
	go fetchFavourites(url, apiKey, favouritesChan)

	favourites := <-favouritesChan
	c.Data["json"] = favourites
	c.ServeJSON()
}

func fetchCatImage(url, apiKey string, ch chan<- CatImage) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		ch <- CatImage{}
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var cats []CatImage
	json.Unmarshal(body, &cats)

	if len(cats) > 0 {
		ch <- cats[0]
	} else {
		ch <- CatImage{}
	}
}

func fetchBreeds(url, apiKey string, ch chan<- []Breed) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		ch <- []Breed{}
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var breeds []Breed
	json.Unmarshal(body, &breeds)

	ch <- breeds
}

func fetchBreedImages(url, apiKey string, ch chan<- []CatImage) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		ch <- []CatImage{}
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var images []CatImage
	json.Unmarshal(body, &images)

	ch <- images
}

func submitVote(url, apiKey string, vote Vote, ch chan<- bool) {
	voteJSON, _ := json.Marshal(vote)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(voteJSON))
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		ch <- false
		return
	}
	defer resp.Body.Close()

	ch <- resp.StatusCode == http.StatusOK
}

func fetchFavourites(url, apiKey string, ch chan<- []Favourite) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		ch <- []Favourite{}
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var favourites []Favourite
	json.Unmarshal(body, &favourites)

	ch <- favourites
}

// new
func (c *CatAPIController) AddFavourite() {
    apiKey, _ := web.AppConfig.String("cat_api_key")
    url := "https://api.thecatapi.com/v1/favourites"

    var favorite struct {
        ImageID string `json:"image_id"`
    }
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &favorite); err != nil {
        c.Data["json"] = map[string]string{"error": "Invalid request body"}
        c.ServeJSON()
        return
    }

    favoriteChan := make(chan bool)
    go submitFavorite(url, apiKey, favorite.ImageID, favoriteChan)

    success := <-favoriteChan
    if success {
        c.Data["json"] = map[string]string{"message": "Image favorited successfully"}
    } else {
        c.Data["json"] = map[string]string{"error": "Failed to favorite image"}
    }
    c.ServeJSON()
}

func submitFavorite(url, apiKey, imageID string, ch chan<- bool) {
    favoriteJSON, _ := json.Marshal(map[string]string{"image_id": imageID})
    client := &http.Client{}
    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(favoriteJSON))
    req.Header.Set("x-api-key", apiKey)
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        ch <- false
        return
    }
    defer resp.Body.Close()

    ch <- resp.StatusCode == http.StatusOK
}