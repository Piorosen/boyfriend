package main

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FirstName string
	LastName  string
	UserName  string
	UserId    int64
	TextId    int64
	Text      string
}

type DBCount struct {
	FirstName string
	LastName  string
	UserName  string
	UserId    int64
	Count     int
}

type Client struct {
	database *gorm.DB
	mode     bool
}

func NewClient() *Client {
	return &Client{
		database: nil,
		mode:     false,
	}
}

func (client *Client) Connect(ip string, port int, db, id, pw string) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Seoul", ip, id, pw, db, port)
	gorm_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	gorm_db.AutoMigrate(&Message{})

	client.database = gorm_db

	return err
}

func (client *Client) Insert(first_name, last_name, user_name, text string, text_id, user_id int64) error {
	result := client.database.Create(&Message{
		FirstName: first_name,
		LastName:  last_name,
		UserName:  user_name,
		Text:      text,
		TextId:    text_id,
		UserId:    user_id,
	})
	if result.Error != nil {
		log.Fatalf("Error creating message: %v", result.Error)
	}
	return result.Error
}

func (client *Client) Count() []DBCount {
	var get_user_id []struct {
		UserId int64
		Count  int
	}

	err := client.database.Model(&Message{}).
		Select("user_id, COUNT(*) AS count").
		Group("user_id").
		Order("count DESC").
		Scan(&get_user_id).Error

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	results := []DBCount{}

	for _, v := range get_user_id {
		search := Message{
			UserId: v.UserId,
		}
		result := client.database.First(&search)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				log.Fatalf("User with ID %d not found", v.UserId)
			} else {
				log.Fatalf("Error retrieving user: %v", result.Error)
			}
		}
		results = append(results, DBCount{
			FirstName: search.FirstName,
			LastName:  search.LastName,
			UserName:  search.UserName,
			UserId:    v.UserId,
			Count:     v.Count,
		})
	}

	return results
}

func (client *Client) Run() bool {
	return client.mode
}

func (client *Client) Process(text string) string {
	if len(text) == 0 {
		return ""
	}
	if text[0] == '/' {
		switch strings.ToLower(text[1:]) {
		case "on":
			client.mode = true
			return "데이터 수집을 재개 합니다."
		case "off":
			client.mode = false
			return "데이터 수집을 중지 합니다."
		case "count":
			result := ""
			data := client.Count()
			for i, v := range data {
				result += fmt.Sprintf("%d 등 | %d 회 : [%s, %s, %s]\n", i+1, v.Count, v.FirstName, v.LastName, v.UserName)
			}
			return result
		default:
			return ""
		}
	}
	return ""
}
