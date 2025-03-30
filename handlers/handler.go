package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"wikiroute/core"

	"github.com/labstack/echo/v4"
)

type RouteRequest struct {
	FromID    int    `json:"from_id"`
	ToID      int    `json:"to_id"`
}

type RouteResponse struct {
	Found bool `json:"found"`
	Route []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"route,omitempty"`
	Error     string `json:"error,omitempty"`
	ReadCount int    `json:"read_count"`
}

func RegisterRoutes(e *echo.Echo, pagesMap *core.PageData, pageLinksMap *core.PageLinksData) {
	var validPageIDs []int32
	for id := range pagesMap.IDToTitle {
		validPageIDs = append(validPageIDs, id)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/api/random", func(c echo.Context) error {
		randomIndex := rand.Intn(len(validPageIDs))
		randomPageID := validPageIDs[randomIndex]
		title := pagesMap.IDToTitle[randomPageID]
		return c.JSON(http.StatusOK, map[string]interface{}{
			"id":    randomPageID,
			"title": title,
		})
	})

	e.GET("/api/pages", func(c echo.Context) error {
		title := c.QueryParam("title")
		if title != "" {
			id := pagesMap.TitleToID[title]
			return c.JSON(http.StatusOK, id)
		}

		id := c.QueryParam("id")
		if id != "" {
			idInt, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID format"})
			}
			title := pagesMap.IDToTitle[int32(idInt)]
			return c.JSON(http.StatusOK, title)
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Either title or id must be provided"})
	})

	e.POST("/api/route", func(c echo.Context) error {
		return FindRouteHandler(c, pagesMap, pageLinksMap)
	})
}

func FindRouteHandler(c echo.Context, pagesMap *core.PageData, pageLinksMap *core.PageLinksData) error {
	var req RouteRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, RouteResponse{Error: "Invalid request format"})
	}

	if req.FromID == 0 || req.ToID == 0 {
		return c.JSON(http.StatusBadRequest, RouteResponse{Error: "Either from_id and to_id must be provided"})
	}

	if !core.PageExists(req.FromID, pagesMap, pageLinksMap) {
		return c.JSON(http.StatusNotFound, RouteResponse{Error: "From page ID does not exist"})
	}
	if !core.PageExists(req.ToID, pagesMap, pageLinksMap) {
		return c.JSON(http.StatusNotFound, RouteResponse{Error: "To page ID does not exist"})
	}

	path, found, readCount := core.BidirectionalBFS(req.FromID, req.ToID, pageLinksMap)

	var route []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}
	for _, id := range path {
		title, err := core.GetPageTitleByID(id, pagesMap)
		if err != nil {
			title = ""
		}
		route = append(route, struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		}{ID: id, Title: title})
	}

	return c.JSON(http.StatusOK, RouteResponse{Found: found, Route: route, ReadCount: readCount})
}
