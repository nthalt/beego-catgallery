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
	Wikipedia	string `json:"wikipedia_url"`
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
	Value   bool   `json:"value"`
	SubID   string `json:"sub_id"`
}

type Favourite struct {
	ID     int      `json:"id"`
	SubID  string   `json:"sub_id"`
	ImageID string   `json:"image_id"`
	Image  CatImage `json:"image"`
}

type FavoriteRequest struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
}

func getUserID() string {
	userID, err := web.AppConfig.String("user_id")
	if err != nil {
		// If there's an error reading the user_id, return a default value
		return "default-user"
	}
	return userID
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

// v1
// func (c *CatAPIController) GetFavourites() {
// 	apiKey, _ := web.AppConfig.String("cat_api_key")
// 	// subID := c.GetString("sub_id")
// 	subID := getUserID()
// 	println("user id: ", subID)
// 	url := fmt.Sprintf("https://api.thecatapi.com/v1/favourites?sub_id=%s", subID)
// 	print("url: ", url)
// 	println()
// 	favouritesChan := make(chan []Favourite)
// 	go fetchFavourites(url, apiKey, favouritesChan)

// 	favourites := <-favouritesChan
// 	c.Data["json"] = favourites
// 	c.ServeJSON()
// }

// func fetchFavourites(url, apiKey string, ch chan<- []Favourite) {
// 	client := &http.Client{}
// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.Header.Set("x-api-key", apiKey)

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		ch <- []Favourite{}
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	var favourites []Favourite
// 	json.Unmarshal(body, &favourites)

// 	ch <- favourites
// }


// v1
// func (c *CatAPIController) AddFavourite() {
// 	apiKey, _ := web.AppConfig.String("cat_api_key")
// 	url := "https://api.thecatapi.com/v1/favourites"

// 	var favorite struct {
// 		ImageID string `json:"image_id"`
// 		SubID   string `json:"sub_id"`
// 	}

// 	// Read the raw request body
// 	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
// 	if err != nil {
// 		c.Data["json"] = map[string]string{"error": "Failed to read request body", "details": err.Error()}
// 		c.ServeJSON()
// 		return
// 	}

// 	// Log the raw request body for debugging
// 	fmt.Printf("Raw request body: %s\n", string(body))

// 	// Attempt to unmarshal the JSON
// 	if err := json.Unmarshal(body, &favorite); err != nil {
// 		c.Data["json"] = map[string]string{"error": "Invalid JSON in request body", "details": err.Error()}
// 		c.ServeJSON()
// 		return
// 	}

// 	// Validate the required fields
// 	if favorite.ImageID == "" {
// 		c.Data["json"] = map[string]string{"error": "image_id is required"}
// 		c.ServeJSON()
// 		return
// 	}

// 	// Use the user_id from the configuration
// 	favorite.SubID = getUserID()

// 	favoriteChan := make(chan map[string]interface{})
// 	go submitFavorite(url, apiKey, favorite.ImageID, favorite.SubID, favoriteChan)

// 	result := <-favoriteChan
// 	c.Data["json"] = result
// 	c.ServeJSON()
// }

// func submitFavorite(url, apiKey, imageID, subID string, ch chan<- map[string]interface{}) {
// 	favoriteJSON, _ := json.Marshal(map[string]string{
// 		"image_id": imageID,
// 		"sub_id":   subID,
// 	})
// 	client := &http.Client{}
// 	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(favoriteJSON))
// 	req.Header.Set("x-api-key", apiKey)
// 	req.Header.Set("Content-Type", "application/json")


// 	resp, err := client.Do(req)
// 	if err != nil {
// 		ch <- map[string]interface{}{"error": "Failed to submit favorite", "details": err.Error()}
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	var result map[string]interface{}
// 	json.Unmarshal(body, &result)

// 	// Log the response from the Cat API
// 	fmt.Printf("Cat API response: %s\n", string(body))

// 	ch <- result
// }


// v2
// func (c *CatAPIController) AddFavourite() {
// 	apiKey, _ := web.AppConfig.String("cat_api_key")
// 	userID, _ := web.AppConfig.String("user_id")
// 	url := "https://api.thecatapi.com/v1/favourites"

// 	var favoriteRequest struct {
// 		ImageID string `json:"image_id"`
// 		SubID   string `json:"sub_id"`
// 	}

// 	// Read the raw request body
// 	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
// 	if err != nil {
// 		c.Data["json"] = map[string]string{"error": "Failed to read request body", "details": err.Error()}
// 		c.ServeJSON()
// 		return
// 	}

// 	// Log the raw request body for debugging
// 	fmt.Printf("Raw request body: %s\n", string(body))

// 	// Attempt to unmarshal the JSON
// 	if err := json.Unmarshal(body, &favoriteRequest); err != nil {
// 		c.Data["json"] = map[string]string{"error": "Invalid JSON in request body", "details": err.Error()}
// 		c.ServeJSON()
// 		return
// 	}

// 	// Validate the required fields
// 	if favoriteRequest.ImageID == "" {
// 		c.Data["json"] = map[string]string{"error": "image_id is required"}
// 		c.ServeJSON()
// 		return
// 	}

// 	// Use the user_id from the configuration if not provided in the request
// 	if favoriteRequest.SubID == "" {
// 		favoriteRequest.SubID = userID
// 	}

// 	favoriteChan := make(chan map[string]interface{})
// 	go submitFavorite(url, apiKey, favoriteRequest.ImageID, favoriteRequest.SubID, favoriteChan)

// 	result := <-favoriteChan
// 	c.Data["json"] = result
// 	c.ServeJSON()
// }

// func submitFavorite(url, apiKey, imageID, subID string, ch chan<- map[string]interface{}) {
// 	favoriteJSON, _ := json.Marshal(map[string]string{
// 		"image_id": imageID,
// 		"sub_id":   subID,
// 	})
// 	client := &http.Client{}
// 	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(favoriteJSON))
// 	req.Header.Set("x-api-key", apiKey)
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		ch <- map[string]interface{}{"error": "Failed to submit favorite", "details": err.Error()}
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	var result map[string]interface{}
// 	json.Unmarshal(body, &result)

// 	// Log the response from the Cat API
// 	fmt.Printf("Cat API response: %s\n", string(body))

// 	ch <- result
// }






// v3
func (c *CatAPIController) AddFavourite() {
	apiKey, _ := web.AppConfig.String("cat_api_key")
	userID, _ := web.AppConfig.String("user_id")
	url := "https://api.thecatapi.com/v1/favourites"

	var favoriteRequest struct {
		ImageID string `json:"image_id"`
	}

	// Read the raw request body
	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to read request body", "details": err.Error()}
		c.ServeJSON()
		return
	}

	// Log the raw request body for debugging
	fmt.Printf("Raw request body: %s\n", string(body))

	// Attempt to unmarshal the JSON
	if err := json.Unmarshal(body, &favoriteRequest); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid JSON in request body", "details": err.Error()}
		c.ServeJSON()
		return
	}

	// Validate the required fields
	if favoriteRequest.ImageID == "" {
		c.Data["json"] = map[string]string{"error": "image_id is required"}
		c.ServeJSON()
		return
	}

	favoriteChan := make(chan map[string]interface{})
	go submitFavorite(url, apiKey, favoriteRequest.ImageID, userID, favoriteChan)

	result := <-favoriteChan
	c.Data["json"] = result
	c.ServeJSON()
}

func submitFavorite(url, apiKey, imageID, subID string, ch chan<- map[string]interface{}) {
	favoriteJSON, _ := json.Marshal(map[string]string{
		"image_id": imageID,
		"sub_id":   subID,
	})
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(favoriteJSON))
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Print request headers
	fmt.Println("Request Headers:")
	for name, values := range req.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		ch <- map[string]interface{}{"error": "Failed to submit favorite", "details": err.Error()}
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	// Log the response from the Cat API
	fmt.Printf("Cat API response: %s\n", string(body))

	ch <- result
}

func (c *CatAPIController) GetFavourites() {
	apiKey, _ := web.AppConfig.String("cat_api_key")
	userID, _ := web.AppConfig.String("user_id")
	url := fmt.Sprintf("https://api.thecatapi.com/v1/favourites?sub_id=%s", userID)

	favouritesChan := make(chan []Favourite)
	go fetchFavourites(url, apiKey, favouritesChan)

	favourites := <-favouritesChan
	c.Data["json"] = favourites
	c.ServeJSON()
}

func fetchFavourites(url, apiKey string, ch chan<- []Favourite) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-api-key", apiKey)

	// Print request headers
	fmt.Println("Request Headers:")
	for name, values := range req.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

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







func (c *CatAPIController) VoteCat() {
	apiKey, _ := web.AppConfig.String("cat_api_key")
	url := "https://api.thecatapi.com/v1/votes"

	var vote Vote
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &vote); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	// Set the sub_id to the user_id from the configuration
	vote.SubID = getUserID()

	voteChan := make(chan map[string]interface{})
	go submitVote(url, apiKey, vote, voteChan)

	result := <-voteChan
	c.Data["json"] = result
	c.ServeJSON()
}

func submitVote(url, apiKey string, vote Vote, ch chan<- map[string]interface{}) {
	voteJSON, _ := json.Marshal(vote)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(voteJSON))
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		ch <- map[string]interface{}{"error": "Failed to submit vote", "details": err.Error()}
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	// Log the response from the Cat API
	fmt.Printf("Cat API vote response: %s\n", string(body))

	ch <- result
}