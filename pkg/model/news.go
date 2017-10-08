package model

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// News type
type News struct {
	ID        bson.ObjectId `bson:"_id"`
	Title     string
	Image     string
	Detail    string
	CreateAt  time.Time `bson:"createAt"`
	UpdatedAt time.Time `bson:"updateAt"`
}

var (
	//newsStorage []*News
	newsStorage []News
	mutexNews   sync.RWMutex
)

func generateID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
}

// CreateNews ..
func CreateNews(news News) error {
	news.ID = bson.NewObjectId()
	news.CreateAt = time.Now()
	news.UpdatedAt = news.CreateAt

	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").Insert(&news)
	if err != nil {
		return err
	}
	return nil
}

// ListNews ...
func ListNews() ([]*News, error) {
	s := mongoSession.Copy()
	defer s.Close()
	var news []*News
	err := s.DB(database).C("news").Find(nil).All(&news)
	if err != nil {
		return nil, err
	}
	return news, nil
}

// GetNews ...
func GetNews(id string) (*News, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, fmt.Errorf("invalid id")
	}
	objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()
	var n News
	err := s.DB(database).C("news").FindId(objectID).One(&n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

// DeleteNews a a
func DeleteNews(id string) error {
	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("invalid id")
	}
	objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").RemoveId(objectID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateNews update news
func UpdateNews(news *News) error {
	if news.ID == "" {
		return fmt.Errorf("required id fo update")
	}
	news.UpdatedAt = time.Now()
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").UpdateId(news.ID, news)
	if err != nil {
		return err
	}
	return nil
}
