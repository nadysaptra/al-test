package database

import (
	"log"

	"gorm.io/gorm"

	"azure-labs-test/app/models"
)

var users = []models.User{
	{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	},
	{
		Name:  "Jane Doe",
		Email: "janedoe@example.com",
	},
}

var events = []models.Event{
	{
		Name:        "Jogja Fair 2020",
		TicketPrize: 20000,
		Date:        "2020-01-01 00:00:00",
	},
	{
		Name:        "Jakarta Fair 2020",
		TicketPrize: 30000,
		Date:        "2020-01-01 00:00:00",
	},
}

var tickets = []models.Ticket{
	{
		Status: "ordered",
	},
	{
		Status: "attended",
	},
	{
		Status: "forfeit",
	},
}

func InitSeeder(db *gorm.DB) {
	for i := range events {
		var exists bool
		_ = db.Model(&events[i]).Select("count(*) > 0").Where("name = ?", events[i].Name).Limit(1).Find(&exists).Error
		if !exists {
			err := db.Debug().Model(&models.Event{}).Create(&events[i]).Error
			if err != nil {
				log.Fatalf("cannot seed events table: %v", err)
			}
		}
	}

	for i := range users {
		var exists bool
		_ = db.Model(&users[i]).Select("count(*) > 0").Where("name = ?", users[i].Name).Limit(1).Find(&exists).Error
		if !exists {
			err := db.Debug().Model(&models.User{}).Create(&users[i]).Error
			if err != nil {
				log.Fatalf("cannot seed users table: %v", err)
			}
			if i == 0 {
				for j := range tickets {
					tickets[j].UserRefer = int(users[i].ID)
					tickets[j].EventRefer = int(events[i].ID)
					err := db.Debug().Model(&models.Ticket{}).Create(&tickets[j]).Error
					if err != nil {
						log.Fatalf("cannot seed tickets table: %v", err)
					}
				}
			}
		}
	}
}
