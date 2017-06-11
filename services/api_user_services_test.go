package services

import (
	"testing"
	"fmt"
)

func TestFindApiUsers(t *testing.T) {

	result := ListApiUser()

	fmt.Println(result)

	fmt.Println("taille : ", len(result))
	//t, ok := result.Value(models.ApiUser{})
	//
	//if ok != nil {
	//	fmt.Println("ici")
	//}
	//fmt.Println(ok)
	//fmt.Println(t)
	//assert.Equal(t, "Yohan", t.Firstname)
}
