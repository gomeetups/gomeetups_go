package fixtures

import (
	"time"

	"github.com/gomeetups/gomeetups/models"
)

// Users Contains user fixtures for memory service
var Users = []models.User{
	{
		UserID:      "d468bd91-39a2-46a1-99c3-4c4b0f20e78a",
		Email:       "aydin@gomeetups.com",
		DisplayName: "Aydin U. - /dev/mach",
		Password:    "$2a$12$1wBYhA0W0SATMwCLpzfyg.w9QJFBkunTQFjzUdJb7zVwdyZxoFt1G", //gomeetups
		IsVerified:  true,
		IsActive:    true,

		CreatedAt: time.Date(2016, time.January, 28, 2, 30, 0, 0, newYork),
		CreatedBy: "d468bd91-39a2-46a1-99c3-4c4b0f20e78a",
		UpdatedAt: time.Date(2016, time.January, 28, 2, 30, 0, 0, newYork),
		UpdatedBy: "d468bd91-39a2-46a1-99c3-4c4b0f20e78a",
	},
	{
		UserID:      "0c8e35a5-4c1c-4294-8927-d23f0b1b4969",
		Email:       "kaylyn@gomeetups.com",
		DisplayName: "Kaylyn G.",
		Password:    "$2a$12$1wBYhA0W0SATMwCLpzfyg.w9QJFBkunTQFjzUdJb7zVwdyZxoFt1G", //gomeetups
		IsVerified:  true,
		IsActive:    true,

		CreatedAt: time.Date(2016, time.January, 28, 2, 30, 0, 0, newYork),
		CreatedBy: "0c8e35a5-4c1c-4294-8927-d23f0b1b4969",
		UpdatedAt: time.Date(2016, time.January, 28, 2, 30, 0, 0, newYork),
		UpdatedBy: "0c8e35a5-4c1c-4294-8927-d23f0b1b4969",
	},
	{
		UserID:      "478f4b8e-b231-4efe-828b-f63b877fbbe3",
		Email:       "arc@gomeetups.com",
		DisplayName: "Archana K.",
		Password:    "$2a$12$1wBYhA0W0SATMwCLpzfyg.w9QJFBkunTQFjzUdJb7zVwdyZxoFt1G", //gomeetups
		IsVerified:  true,
		IsActive:    true,

		CreatedAt: time.Date(2016, time.January, 28, 2, 30, 0, 0, newYork),
		CreatedBy: "478f4b8e-b231-4efe-828b-f63b877fbbe3",
		UpdatedAt: time.Date(2016, time.January, 28, 2, 30, 0, 0, newYork),
		UpdatedBy: "478f4b8e-b231-4efe-828b-f63b877fbbe3",
	},
}
