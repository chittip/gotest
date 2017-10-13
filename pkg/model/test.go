package model

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"

	"gopkg.in/mgo.v2/bson"
)

// Test type
type Test struct {
	ID         bson.ObjectId `bson:"_id"`
	APIName    string
	URLPath    string
	MethodType string
	URLParam   string
	Body       string
	CreateAt   time.Time `bson:"createAt"`
	UpdatedAt  time.Time `bson:"updateAt"`
}

// CreateTest ...
func CreateTest(test Test) error {
	test.ID = bson.NewObjectId()
	test.CreateAt = time.Now()
	test.UpdatedAt = test.CreateAt

	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("test").Insert(&test)
	if err != nil {
		return err
	}
	return nil
}

// ListTest ...
func ListTest() ([]*Test, error) {
	s := mongoSession.Copy()
	defer s.Close()
	var test []*Test
	err := s.DB(database).C("test").Find(nil).All(&test)
	if err != nil {
		return nil, err
	}
	return test, nil
}

// GetTest ...
func GetTest(id string) (*Test, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, fmt.Errorf("invalid id")
	}
	objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()
	var t Test
	err := s.DB(database).C("test").FindId(objectID).One(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// DeleteTest a a
func DeleteTest(id string) error {
	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("invalid id")
	}
	objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("test").RemoveId(objectID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTest update news
func UpdateTest(test *Test) error {
	if test.ID == "" {
		return fmt.Errorf("required id fo update")
	}
	test.UpdatedAt = time.Now()
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("test").UpdateId(test.ID, test)
	if err != nil {
		return err
	}
	return nil
}

// RunTest ...
func RunTest(test *Test) error {
	if test.ID == "" {
		return fmt.Errorf("required id fo test")
	}
	log.Println("Run Test ....")
	log.Println(test)
	log.Println(test.APIName)
	log.Println(test.Body)
	log.Println(test.URLPath)
	log.Println(test.URLParam)
	log.Println(test.MethodType)
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/

	if test.MethodType == "GET" {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		req, _ := http.NewRequest(test.MethodType, test.URLPath, nil)
		req = req.WithContext(ctx)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(err.Error())
		}

		log.Println(resp.Status)
		log.Println(body)
		m, ok := gjson.Parse(string(body[:])).Value().(map[string]interface{})
		if !ok {
			log.Println("- error cannot ...")
		} else {
			// log.Println(m)
			for key, value := range m {
				fmt.Println("Key:", key, "Value:", value)
			}
		}
	} else if test.MethodType == "POST" {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		req, _ := http.NewRequest(test.MethodType, test.URLPath, strings.NewReader(test.Body))
		req = req.WithContext(ctx)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(err.Error())
		}

		log.Println(resp.Status)
		log.Println(body)
		m, ok := gjson.Parse(string(body[:])).Value().(map[string]interface{})
		if !ok {
			log.Println("- error cannot ...")
		} else {
			// log.Println(m)
			for key, value := range m {
				fmt.Println("Key:", key, "Value:", value)
			}
		}
	}
	return nil
}
