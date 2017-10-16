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

// Test ...
/*
{
	httpStatusCode : "200"
	httpBody : ""
}
*/
type Test struct {
	ID         bson.ObjectId `bson:"_id"`
	APIName    string
	URLPath    string
	MethodType string
	URLParam   string
	Body       string
	Expected   string
	CreateAt   time.Time `bson:"createAt"`
	UpdatedAt  time.Time `bson:"updateAt"`
}

// ResultTest ...
type ResultTest struct {
	Name     string
	Expected string
	Actual   string
	Pass     bool
}

// RequestResult ...
type RequestResult struct {
	HTTPStatusCode string
	Body           string
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

// Vaildate ...
func (test *Test) Vaildate(actual string) (ResultTest, error) {
	expectedhttpStatusCoactualde := gjson.Get(test.Expected, "httpStatusCode")
	actualhttpStatusCode := gjson.Get(actual, "httpStatusCode")
	if expectedhttpStatusCoactualde != actualhttpStatusCode {
		result := ResultTest{
			Name:     "httpStausCode invalid",
			Expected: expectedhttpStatusCoactualde.String(),
			Actual:   actualhttpStatusCode.String(),
			Pass:     false,
		}
		return result, nil
	}
	result := ResultTest{
		Name:     "httpStausCode invalid",
		Expected: expectedhttpStatusCoactualde.String(),
		Actual:   actualhttpStatusCode.String(),
		Pass:     true,
	}
	return result, nil
}

// VaildateResult ...
func (test *Test) VaildateResult(actual *RequestResult) (ResultTest, error) {
	expectedhttpStatusCoactualde := gjson.Get(test.Expected, "httpStatusCode")

	if expectedhttpStatusCoactualde.String() != actual.HTTPStatusCode {
		result := ResultTest{
			Name:     "httpStausCode invalid",
			Expected: expectedhttpStatusCoactualde.String(),
			Actual:   actual.HTTPStatusCode,
			Pass:     false,
		}
		return result, nil
	}
	result := ResultTest{
		Name:     "httpStausCode invalid",
		Expected: expectedhttpStatusCoactualde.String(),
		Actual:   actual.HTTPStatusCode,
		Pass:     true,
	}
	return result, nil
}

func httpRequest(url string, methodType string, body string) (*RequestResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var req *http.Request

	if len(body) > 0 {
		req, _ = http.NewRequest(methodType, url, nil)
	} else {
		req, _ = http.NewRequest(methodType, url, strings.NewReader(body))
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	resbody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	log.Println(resp.Status)
	log.Println(resbody)
	m, ok := gjson.Parse(string(resbody[:])).Value().(map[string]interface{})
	if !ok {
		log.Println("- error cannot ...")
	} else {
		// log.Println(m)
		for key, value := range m {
			fmt.Println("Key:", key, "Value:", value)
		}
	}
	rr := RequestResult{
		HTTPStatusCode: resp.Status,
		Body:           string(resbody[:]),
	}
	return &rr, nil
}

// RunTest ...
func RunTest(test *Test) (*ResultTest, error) {
	if test.ID == "" {
		return nil, fmt.Errorf("required id fo test")
	}
	log.Println("Run Test ....")
	log.Println(test)
	log.Println(test.APIName)
	log.Println(test.Body)
	log.Println(test.URLPath)
	log.Println(test.URLParam)
	log.Println(test.MethodType)
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	var rr *RequestResult
	var err error
	if test.MethodType == "GET" {
		rr, err = httpRequest(test.URLPath, test.MethodType, "")
		if err != nil {
			return nil, err
		}
	} else if test.MethodType == "POST" {
		rr, err = httpRequest(test.URLPath, test.MethodType, test.Body)
		if err != nil {
			return nil, err
		}
	}
	rt, err := test.VaildateResult(rr)
	return &rt, err
}
