package sqlutils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pubsub/publog"
)
type Hotel	 struct {
	gorm.Model
	HotelUid  string `sql:"index"`
	Name    		string
	Country 	string
	Address      	string
	Latitude      	string
	Longitude      	string
	Telephone      	string
	Description     string
	Room_count      string
	Currency      	string
	Amenities []Amenity `gorm:"foreignkey:HotelName;references:hotel_uid"`
	Rooms   []Room `gorm:"foreignkey:HotelName;references:hotel_uid"`
	RatePlans []RatePlan `gorm:"foreignkey:HotelName;references:hotel_uid"`
}

type Amenity struct {
	gorm.Model
	Type string
	HotelName string
}

type Room	 struct {
	gorm.Model
	HotelName  		string
	RoomId	   		string `sql:"index"`
	Description 		string
	Name      	string
	Capacities []Capacity `gorm:"foreignkey:RoomName;references:room_id"`

}
type Capacity struct {
	gorm.Model
	MaxAudlts string
	ExtraChildren string
	RoomName string
}

type RatePlan struct {
	gorm.Model
	PlanId string `sql:"index"`
	HotelName  string
	Name string
	MealPlan string

}
const dbstr = "eastern.db"
func SqliteMigrate() {

	db, err := gorm.Open(sqlite.Open(dbstr), &gorm.Config{})
	publog.Logrs.Info("creating database "+dbstr)
	if err != nil {
		publog.Logrs.Info("Migrate "+err.Error())
		publog.Logrs.Info("failed to connect database")
		panic("failed to connect database")
	}
	publog.Logrs.Info( "created database "+dbstr)

	err = db.AutoMigrate(&RatePlan{},&Capacity{},&Amenity{},&Room{},&Hotel{})
	if err != nil {
		publog.Logrs.Info( "failed to create  table")
		panic("failed to create table")
	}


}

func SqliteCreateData(table interface{}) bool {
	db, err := gorm.Open(sqlite.Open(dbstr), &gorm.Config{})
	if err != nil {
		publog.Logrs.Info("failed to connect database")
		panic("failed to connect database")
	}
	publog.Logrs.Info("CreateData__ connected database")
	t := db.Create(table)
	if t.Error != nil{
		return false
	}
	publog.Logrs.Info("Record Inserted in database")
	return true
}
