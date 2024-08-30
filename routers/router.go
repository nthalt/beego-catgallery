package routers

import (
	"beego-catgallery/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
    web.Router("/", &controllers.MainController{})
    web.Router("/api/cats/random", &controllers.CatAPIController{}, "get:GetRandomCat")
    web.Router("/api/breeds", &controllers.CatAPIController{}, "get:GetBreeds")
    web.Router("/api/breeds/:id", &controllers.CatAPIController{}, "get:GetBreedInfo")
    web.Router("/api/votes", &controllers.CatAPIController{}, "post:VoteCat")
    web.Router("/api/favourites", &controllers.CatAPIController{}, "get:GetFavourites")
    web.Router("/api/favourites", &controllers.CatAPIController{}, "post:AddFavourite")
}