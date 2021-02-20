package handler

import (
	"encoding/json"
	"pubsub/sqlutils"
	"pubsub/publog"
	"strconv"
)

func HandleData(data []byte){
	var t interface{}
	var hotel sqlutils.Hotel
	var room sqlutils.Room
	var rateplan sqlutils.RatePlan
	publog.Logrs.Info(data)
	_ = json.Unmarshal(data, &t)

	m := t.(map[string]interface{})
	l := m["offers"].([]interface{})
	for _, v := range l {
		for j,n :=range v.(map[string]interface{}){
			if j == "hotel"{
				val := n.(map[string]interface{})
				hotel = sqlutils.Hotel{
					Address: val["address"].(string),
					HotelUid:val["hotel_id"].(string),
					Name : 		val["name"].(string),
					Country :	val["country"].(string),
					Latitude  :    	FloatToString(val["latitude"].(float64)),
					Longitude :     	FloatToString(val["longitude"].(float64)),
					Telephone :     	val["telephone"].(string),
					Description :    val["description"].(string),
					Room_count  :    FloatToString(val["room_count"].(float64)),
					Currency   :   	val["currency"].(string),

			}
				amenities := val["amenities"].([]interface{})
				for  _, am := range amenities{
					hotel.Amenities = append(hotel.Amenities, sqlutils.Amenity{Type: am.(string)})
					}
			}
			if j == "room"{
				room_val := n.(map[string]interface{})
				room = sqlutils.Room{RoomId: room_val["room_id"].(string),
					Name: room_val["name"].(string),
					Description: room_val["description"].(string),
					}
				capa := room_val["capacity"].(map[string]interface{})
				room.Capacities = append(room.Capacities, sqlutils.Capacity{MaxAudlts: FloatToString(capa["max_adults"].(float64)),
						ExtraChildren: FloatToString(capa["extra_children"].(float64)),
					})
				hotel.Rooms = append(hotel.Rooms, room)
			}
			if j == "rate_plan"{
				plan := n.(map[string]interface{})
				rateplan = sqlutils.RatePlan{Name: plan["name"].(string),
					PlanId:plan["rate_plan_id"].(string),
					MealPlan: plan["meal_plan"].(string),
					}
				hotel.RatePlans = append(hotel.RatePlans, rateplan)
			}

	}
}
	publog.Logrs.Info("sent Data in database")
	sqlutils.SqliteCreateData(&hotel)
}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
