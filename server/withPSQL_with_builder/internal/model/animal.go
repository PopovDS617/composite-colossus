package model

import "time"

type Animal struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	Age            int       `json:"age"`
	Gender         string    `json:"gender"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Region         string    `json:"region"`
	LastTimeSeenAt time.Time `json:"last_time_seen_at"`
	SeenByDevice   int       `json:"seen_by_device"`
}

func (a *Animal) ValidateCreate() (map[string]string, bool) {
	inputErrors := map[string]string{}
	ok := true

	if a.Name == "" {
		inputErrors["name"] = "should not be empty"
	}
	if a.Age < 0 {
		inputErrors["age"] = "should not be a negative number"
	}
	if a.Type == "" {
		inputErrors["type"] = "should not be empty"
	}
	if a.Gender == "" {
		inputErrors["gender"] = "should not be empty"
	}
	if a.Region == "" {
		inputErrors["region"] = "should not be empty"
	}

	if len(inputErrors) != 0 {
		ok = false
	}

	return inputErrors, ok
}

func (stored *Animal) ValidateAndUpdate(received *Animal) {

	if stored.Name != received.Name && received.Name != "" {
		stored.Name = received.Name
	}
	if stored.Type != received.Type && received.Type != "" {
		stored.Type = received.Type
	}
	if stored.Age != received.Age && received.Age != 0 {
		stored.Age = received.Age
	}
	if stored.Gender != received.Gender && received.Gender != "" {
		stored.Gender = received.Gender
	}
	if stored.LastTimeSeenAt != received.LastTimeSeenAt && !received.LastTimeSeenAt.IsZero() {
		stored.LastTimeSeenAt = received.LastTimeSeenAt
	}

}
