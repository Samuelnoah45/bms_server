package handler


import (
	"net/http"
	"fmt"
  //   "log"
  "github.com/gin-gonic/gin"

)
// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	// // body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(c.Request.Header)
    // id := c.Param("id")

    // // Loop over the list of albums, looking for
    // // an album whose ID value matches the parameter.
    // for _, a := range albums {
    //     if a.ID == id {
    //         c.IndentedJSON(http.StatusOK, a)
    //         return
    //     }
    // }
    // c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

	c.JSON(http.StatusOK,gin.H{"data":"hello world"})
}