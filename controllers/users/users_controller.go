package users

import (
	"net/http"

	"github.com/cota-eng/go-users-api/domain/users"
	"github.com/cota-eng/go-users-api/services"
	"github.com/cota-eng/go-users-api/utils/errors"
	"github.com/gin-gonic/gin"
)


var counter int

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented,"implement")
}
 
func CreateUser(c *gin.Context) {
	var user users.User
	// fmt.Println(user)
	// bytes,err := ioutil.ReadAll(c.Request.Body)
	// if err != nil{
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// if err := json.Unmarshal(bytes,&user);err != nil{
	// 	fmt.Println(err.Error())
	// 	return
	// }

	if err := c.ShouldBindJSON(&user);err!=nil{
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		// TODO: handle json err
		// log.Fatalln(err)
		return
	}

	result,saveErr:=services.CreateUser(user)
	if saveErr !=nil{
		c.JSON(saveErr.Status,saveErr)
		return
	}
	// fmt.Println(string(bytes))
	// fmt.Println(user)
	// c.String(http.StatusNotImplemented,"implement")
	c.JSON(http.StatusCreated,result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented,"implement")
}