package utils

import (
	"log"
	"testing"
)

// TestEquals ....
func TestEquals(t *testing.T) {
	s1 := `{
	  "a": "1",
	  "b": [
		{"x":"v1", "y":"v2"},
		{"x2":"v2", "y2":"v2"}
	        ]
	}`

	s2 := `{
	  "a": "1",
	  "b": [
		{"x2":"v2", "y2":"v2"},
		{"x":"v1", "y":"v2"}
	        ]
	   }`
	// s1 := `{"a":"b"}`
	// s2 := `{"a":"b"}`
	result := JSONEquals(s1, s2)
	log.Println(result)
	if !result {
		t.Errorf("expected faile %v", result)
	}
}

// TestEquals2 ....
func TestEquals2(t *testing.T) {
	s1 := `{
	  "b": [
		{"x":"v1", "y":"v2"},
		{"x2":"v2", "y2":"v2"}
	        ],
		"a": "1"
	}`

	s2 := `{
	  "a": "1",
	  "b": [
		{"x2":"v2", "y2":"v2"},
		{"x":"v1", "y":"v2"}
	        ]
	   }`
	// s1 := `{"a":"b"}`
	// s2 := `{"a":"b"}`
	result := JSONEquals(s1, s2)
	if !result {
		t.Errorf("expected faile %v", result)
	}
}

// TestEquals3 ....
func TestEquals3(t *testing.T) {
	s1 := `{
	  "b": [
		{"x":"v1", "y":"v2"},
		{"x2":"v2", "y2":"v2"}
	        ],
		"a": "2"
	}`

	s2 := `{
	  "a": "1",
	  "b": [
		{"x2":"v2", "y2":"v2"},
		{"x":"v1", "y":"v2"}
	        ]
	   }`
	// s1 := `{"a":"b"}`
	// s2 := `{"a":"b"}`
	result := JSONEquals(s1, s2)
	log.Println(result)
	if result {
		t.Errorf("expected faile %v", result)
	}
}

// TestEquals4 ....
func TestEquals4(t *testing.T) {
	s1 := `{
	  "b": [
		{"x":"v1", "y":"v2"},
		{"x2":"v2", "y2":"v2"}
	        ],
		"a": "2",
		"x": "v"
	}`

	s2 := `{
	  "a": "1",
	  "b": [
		{"x2":"v2", "y2":"v2"},
		{"x":"v1", "y":"v2"}
	        ]
	   }`
	// s1 := `{"a":"b"}`
	// s2 := `{"a":"b"}`
	result := JSONEquals(s1, s2)
	log.Println(result)
	if result {
		t.Errorf("expected faile %v", result)
	}
}

// TestEquals5 ....
func TestEquals5(t *testing.T) {
	s1 := `{
	  "b": [
		{"x":"v1", "y":"v2"},
		{"x2":"v2", "y2":"v2"}
	        ],
		"a": "2",
	}`

	s2 := `{
	  "a": "1",
		"x": "v",		
	  "b": [
		{"x2":"v2", "y2":"v2"},
		{"x":"v1", "y":"v2"}
	        ]
	   }`
	// s1 := `{"a":"b"}`
	// s2 := `{"a":"b"}`
	result := JSONEquals(s1, s2)
	log.Println(result)
	if result {
		t.Errorf("expected faile %v", result)
	}
}

// TestEquals6 ....
func TestEquals6(t *testing.T) {
	s1 := `{
	  "b": [
		{"x":"v1", "y":"v2"},
		{"x2":"v2", "y2":"v2"}
	        ]
		"a": "2",
	}`

	s2 := `{
	  "a": "1",
		"x": "v",		
	  "b": [
		{"x2":"v2", "y2":"v2"},
		{"x":"v1", "y":"v2"}
	        ]
	   }`
	// s1 := `{"a":"b"}`
	// s2 := `{"a":"b"}`
	result := JSONEquals(s1, s2)
	log.Println(result)
	if result {
		t.Errorf("expected faile %v", result)
	}
}
