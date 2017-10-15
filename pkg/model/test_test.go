package model

import (
	"log"
	"testing"
)

// TestValidate ...
func TestValidate(t *testing.T) {
	var test = Test{
		Expected: ` { "httpStatusCode" : "200" , "httpBody" : "" }`,
	}
	log.Println(test)
	resultTest, _ := test.Vaildate(` { "httpStatusCode" : "200" , "httpBody" : "" }`)
	log.Println(resultTest.Pass)
	if resultTest.Pass == false {
		t.Errorf("expected faile %v", resultTest)
		t.Errorf("expected faile %v", resultTest.Expected)
		t.Errorf("expected faile %v", resultTest.Actual)
	}
}

// TestValidate ...
func TestValidate2(t *testing.T) {
	var test = Test{
		Expected: ` { "httpStatusCode" : "200" , "httpBody" : "" }`,
	}
	log.Println(test)
	resultTest, _ := test.Vaildate(` { "httpStatusCode" : "201" , "httpBody" : "" }`)
	log.Println(resultTest.Pass)
	if resultTest.Pass == true {
		t.Errorf("expected faile %v", resultTest)
		t.Errorf("expected faile %v", resultTest.Expected)
		t.Errorf("expected faile %v", resultTest.Actual)
	}
}
