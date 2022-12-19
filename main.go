package main
import (
    "net/http"
	  "fmt"
	//   "log"
    "reflect"
    "github.com/gin-gonic/gin"
	// "server/BMS-server/handler"
)
// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}


// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	 
    c.IndentedJSON(http.StatusOK, albums)
}
// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}




func main() {
    album1:=album  {
        ID: "1", 
        Title: "Blue Train", 
        Artist: "John Coltrane", 
        Price: 56.99}


        values := reflect.ValueOf(album1)
        types := values.Type()
        fmt.Println(values ,types)

        for i := 0; i < values.NumField(); i++ {
            fmt.Println(types.Field(i), types.Field(i).Name, values.Field(i))
        }

    // router := gin.Default()
    // router.GET("/albums", getAlbums)
	// router.POST("/albums", postAlbums)
	// router.POST("/albums/:id", handler.GetAlbumByID)
    // router.Run("localhost:3000")
}
// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album
    
	c.BindJSON(newAlbum); 
    // Call BindJSON to bind the received JSON to
    // newAlbum.
    // if err := c.BindJSON(&newAlbum); err != nil {
    //     return
    // }
     fmt.Println("sky")
    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

