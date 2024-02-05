package pkg

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category string

const (
	EMAIL Category = "Email"
	SMS  Category = "SMS"
)



type Manager struct {
	Db *gorm.DB
}

var dbRepo Manager


func GetRepo() Manager {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	if dbRepo.Db == nil {
		db,err := gorm.Open(sqlite.Open(pwd+"/tmp/chistadotdev.sqlite"),&gorm.Config{})
		if err != nil {
			panic("failed to connet to the database")
		}

		dbRepo.Db = db
	}

	return dbRepo
}


type User struct {
	gorm.Model
	Name string
	Email string 
	Password string
	ApiKeys []ApiKey
	Services []Service `gorm:"many2many:user_services;"`
}

type ApiKey struct {
	gorm.Model
	ApiKey string
	ExpireDate string
	UserID uint
}

type Service struct {
	gorm.Model
	Name string `json:"name"`
	Logo string `json:"logo"`
	Category  Category  `json:"category"`
	Users []User `gorm:"many2many:user_services"`
}

type UserService struct {
	gorm.Model
	CreadA string `json:"creadA"`
	CreadB string `json:"creadB"`
	CreadC string `json:"creadC"`
	CreadD string `json:"creadD"`
	UserID uint
	ServiceId uint
}




//seed
func (manager Manager) Seed(){

	var user User
	user.Name = "siraj"
	user.Email = "sirajyesuf762@gmail.com"
	user.Password = "password"

	manager.Db.Create(&user).Clauses(clause.Returning{})

	var service1 Service
	service1.Name = "resend"
	service1.Category = "email"
	service1.Logo  = "img path"

	var service2 Service
	service2.Name = "mailjet"
	service2.Category = "email"
	service2.Logo  = "img path"

	manager.Db.Create(&service1).Clauses(clause.Returning{})
	manager.Db.Create(&service2).Clauses(clause.Returning{})




	var apikey ApiKey
	apikey.ApiKey = "1234567890"
	apikey.ExpireDate = "2/01/2024"
	apikey.UserID = user.ID
	manager.Db.Create(&apikey)
	

	var userservice UserService
	var userservice2 UserService

	userservice.CreadA = "qwedfyhjikl"
	userservice.UserID = user.ID
	userservice.ServiceId = service1.ID

	manager.Db.Create(&userservice)


	userservice2.CreadA = "09875432"
	userservice2.CreadB = "olollllll"
	userservice2.UserID = user.ID
	userservice2.ServiceId = service2.ID

	manager.Db.Create(&userservice2)



}


// func (manager *Manager) GetApiKey(token string) (ApiKey,error)  {

// 	var apikey ApiKey 

// 	manager.Db.Where("apikey = ?",token).First(&apikey)

// 	return apikey,nil
	
// }First

func (manager Manager) GetApiKey(token string) (ApiKey,error)  {

	var apiKey ApiKey

	err := manager.Db.Where("api_key = ?",token).First(&apiKey).Error

	return apiKey,err
	
}


func (manager  Manager) GetUserServices(userId uint) (User, error) {
	var user User
	err := manager.Db.Model(&User{}).Preload("Services").First(&user,userId).Error
	return user,err
}


func(manager Manager) GetUserService(serviceId uint) (UserService,error) {
	var userservice UserService
	err := manager.Db.Model(&UserService{}).Where("service_id = ?",serviceId).First(&userservice).Error
	return userservice,err
}