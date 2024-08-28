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
    apiKey, err := web.AppConfig.String("cat_api_key")
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to get API key"}
        c.ServeJSON()
        return
    }
    url := "https://api.thecatapi.com/v1/images/search"

    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("x-api-key", apiKey)

    resp, err := client.Do(req)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    var cats []CatImage
    json.Unmarshal(body, &cats)

    if len(cats) > 0 {
        c.Data["json"] = cats[0]
    } else {
        c.Data["json"] = map[string]string{"error": "No cats found"}
    }
    c.ServeJSON()
}

func (c *CatAPIController) GetBreeds() {
    apiKey, err := web.AppConfig.String("cat_api_key")
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to get API key"}
        c.ServeJSON()
        return
    }
    url := "https://api.thecatapi.com/v1/breeds"

    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("x-api-key", apiKey)

    resp, err := client.Do(req)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    var breeds []Breed
    json.Unmarshal(body, &breeds)

    c.Data["json"] = breeds
    c.ServeJSON()
}

func (c *CatAPIController) GetBreedInfo() {
    apiKey, err := web.AppConfig.String("cat_api_key")
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to get API key"}
        c.ServeJSON()
        return
    }
    breedID := c.Ctx.Input.Param(":id")
    url := fmt.Sprintf("https://api.thecatapi.com/v1/breeds/%s", breedID)

    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("x-api-key", apiKey)

    resp, err := client.Do(req)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    var breed Breed
    json.Unmarshal(body, &breed)

    c.Data["json"] = breed
    c.ServeJSON()
}

func (c *CatAPIController) VoteCat() {
    apiKey, err := web.AppConfig.String("cat_api_key")
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to get API key"}
        c.ServeJSON()
        return
    }
    url := "https://api.thecatapi.com/v1/votes"

    var vote Vote
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &vote); err != nil {
        c.Data["json"] = map[string]string{"error": "Invalid request body"}
        c.ServeJSON()
        return
    }

    voteJSON, err := json.Marshal(vote)
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to marshal vote data"}
        c.ServeJSON()
        return
    }

    client := &http.Client{}
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(voteJSON))
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to create request"}
        c.ServeJSON()
        return
    }
    req.Header.Set("x-api-key", apiKey)
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to read response body"}
        c.ServeJSON()
        return
    }

    var result map[string]interface{}
    if err := json.Unmarshal(body, &result); err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to parse response"}
        c.ServeJSON()
        return
    }

    c.Data["json"] = result
    c.ServeJSON()
}

func (c *CatAPIController) GetFavourites() {
    apiKey, err := web.AppConfig.String("cat_api_key")
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to get API key"}
        c.ServeJSON()
        return
    }
    url := "https://api.thecatapi.com/v1/favourites"

    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("x-api-key", apiKey)

    resp, err := client.Do(req)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    var favourites []Favourite
    json.Unmarshal(body, &favourites)

    c.Data["json"] = favourites
    c.ServeJSON()
}