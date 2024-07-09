package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	BUILD_DATE_TIME = "null"
	DEPLOY_VERSION  = "0.0.0"
	ALL_DATA_REMOVE = GenerateRandomHex(16)
)

func GenerateRandomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}

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

func (client *Client) Clear() error {
	// Message 구조체에 해당하는 모든 데이터를 삭제합니다.
	if err := client.database.Where("1 = 1").Delete(&Message{}).Error; err != nil {
		return fmt.Errorf("failed to delete all messages: %v", err)
	}
	return nil
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

func (client *Client) GetText(limit int) []Message {
	var results []Message

	err := client.database.Model(&Message{}).
		Select("*").
		Order("created_at DESC").
		Limit(limit).
		Scan(&results).Error

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}

	return results
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

func (client *Client) Process(text string, env Environment) string {
	if len(text) == 0 {
		return ""
	}
	if text[0] == '/' {
		switch strings.Split(strings.ToLower(text[1:]), " ")[0] {
		case "chat":
			// jsonData := fmt.Sprintf(`{"size": %d, "jubu_id": %d}`, 100, env.TelegramJubuId)
			message := client.GetText(env.PreviousTextSize)
			result, err := MakeChat(message, env.GeminiApiKey, env.TelegramJubuId)
			if err != nil {
				return err.Error()
			} else {
				// client.Insert("차차핑-봇",
				// 	"차차핑-봇",
				// 	"차차핑-봇",
				// 	result,
				// 	int64(0),
				// 	int64(0),
				// )
				return result
			}
		case "on":
			client.mode = true
			return "데이터 수집을 재개 합니다."
		case "off":
			client.mode = false
			return "데이터 수집을 중지 합니다."
		case "version":
			return fmt.Sprintf("Version : %s\nBuild Time : %s", DEPLOY_VERSION, BUILD_DATE_TIME)
		case "clear":
			data := strings.Split(text, " ")
			if len(data) != 2 {
				return fmt.Sprintf("아래의 명령어를 사용하시면 데이터베이스에 있는 모든 기록을 삭제합니다.\n/clear %s", ALL_DATA_REMOVE)
			}
			if data[1] == ALL_DATA_REMOVE {
				err := client.Clear()
				return fmt.Sprintf("Error : %v\n모든 처리를 수행하였습니다.", err)
			} else {
				return fmt.Sprintf("아래의 명령어를 사용하시면 데이터베이스에 있는 모든 기록을 삭제합니다.\n/clear %s", ALL_DATA_REMOVE)
			}

		case "count":
			result := ""
			data := client.Count()
			for i, v := range data {
				result += fmt.Sprintf("%d 등 | %d 회 : [%s, %s, %s]\n", i+1, v.Count, v.FirstName, v.LastName, v.UserName)
			}
			if result == "" {
				return "어떠한 데이터도 없습니다."
			} else {
				return result
			}
		default:
			return ""
		}
	}
	return ""
}
