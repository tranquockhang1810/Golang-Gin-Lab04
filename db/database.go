// Create a map of students and add some data
package db

import (
	"github.com/tranquockhang1810/Golang-Gin-Lab04/models"
)

var Students = make(map[string]models.Student)

func init() {
	Students["1"] = models.Student{Id: 1, Name: "Nguyen Van A", Birthday: "01/01/2000", Gender: "Male", Address: "HN", Phone: "0123456789", Email: "XjXtG@example.com"}
	Students["2"] = models.Student{Id: 2, Name: "Tran Van B", Birthday: "02/01/2001", Gender: "Male", Address: "HCM", Phone: "0987654321", Email: "XjXtG@example.com"}
	Students["3"] = models.Student{Id: 3, Name: "Nguyen Thi C", Birthday: "03/01/2002", Gender: "Female", Address: "DN", Phone: "0192837465", Email: "XjXtG@example.com"}
	Students["4"] = models.Student{Id: 4, Name: "Pham Van D", Birthday: "10/01/2003", Gender: "Male", Address: "LD", Phone: "0981237645", Email: "XjXtG@example.com"}
	Students["5"] = models.Student{Id: 5, Name: "Le Thi E", Birthday: "11/01/2005", Gender: "Female", Address: "HCM", Phone: "0345678912", Email: "XjXtG@example.com"}
}