package models

import "time"

type User struct {
	Id        int64      `bson:"_id,omitempty"`
	FirstName string     `bson:"first_name"`
	LastName  string     `bson:"last_name"`
	DateBirth *time.Time `bson:"date_of_birth"`
}

type Discount struct {
	Ptc          float32 `bson:"pct"`
	ValueInCents int64   `bson:"value_in_cents"`
}

type Product struct {
	Id          int64    `bson:"_id,omitenmpty"`
	PriceCents  int64    `bson:"price_in_cents"`
	Title       string   `bson:"title"`
	Description string   `bson:"description"`
	Discount    Discount `bson:"discount"`
}
