package server

import (
	"example/server/model"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// albums slice to seed record album data.
var albums = []model.Album{
	{ID: "", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "", Title: "Smt", Artist: "Gerry Mulligan", Price: 17.99},
}

var user string

func FindUserAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		user = c.GetHeader("User-Agent")
		log.Println(user)
		// Before calling handler
		c.Next()
		// After calling handler
	}
}

var i int = 0

func appType(types string) {
	app := "application/"
	switch types {
	case app + "json":
		i = 1
	case app + "xml":
		i = 2
	case app + "html":
		i = 3
	default:
		i = 3
	}
}

func findApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		appType(c.GetHeader("accept"))
	}
}

func StartServer(listenAddr string) {
	giveIDs()
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.LoadHTMLGlob("pkg/templates/*.html")
	router.Use(FindUserAgent())
	router.Use(findApplication())

	router.GET("/albums", getAlbums)
	router.GET("/albums/:ID", getAlbumByID)
	router.GET("/albums/last", getLastAlbum)
	router.GET("/albums/titles/:Title", getAlbumByTitle)
	router.POST("/albums", postAlbums)
	router.Run(listenAddr)
}

func getAlbums(c *gin.Context) {
	if i == 1 {
		c.IndentedJSON(http.StatusOK, albums)
	} else if i == 2 {
		c.XML(http.StatusOK, albums)
	} else if i == 3 {
		data := gin.H{
			"title":  "HTML Page",
			"albums": albums,
		}
		c.HTML(http.StatusOK, "index.html", data)
	}
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

	getAddID()
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("ID")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	b := 0
	for _, a := range albums {
		if a.ID == id {
			if i == 1 {
				c.IndentedJSON(http.StatusOK, a)
			} else if i == 2 {
				c.XML(http.StatusOK, a)
			} else if i == 3 {
				data := gin.H{
					"id":     albums[b].ID,
					"title":  albums[b].Title,
					"artist": albums[b].Artist,
					"price":  albums[b].Price,
				}
				c.HTML(http.StatusOK, "album.html", data)
			}
			return
		}
		b++
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getAlbumByTitle(c *gin.Context) {
	title := c.Param("Title")
	b := 0
	for _, a := range albums {
		if a.Title == title {
			if i == 1 {
				c.IndentedJSON(http.StatusOK, a)
			} else if i == 2 {
				c.XML(http.StatusOK, a)
			} else if i == 3 {
				data := gin.H{
					"id":     albums[b].ID,
					"title":  albums[b].Title,
					"artist": albums[b].Artist,
					"price":  albums[b].Price,
				}
				c.HTML(http.StatusOK, "album.html", data)
			}
			return
		}
		b++
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getLastAlbum(c *gin.Context) {
	a := albums[len(albums)-1]
	c.IndentedJSON(http.StatusOK, a)
	return
}

// Puts in order the albums
func giveIDs() {
	for i := 0; i < len(albums); i++ {
		newUUID := uuid.New()
		albums[i].ID = newUUID.String()
	}
}

// It adds id
func getAddID() {
	newUUID := uuid.New()
	albums[len(albums)-1].ID = newUUID.String()
}
