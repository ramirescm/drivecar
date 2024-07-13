package users

import (
	"database/sql/driver"
	"time"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type AnyPassword struct{}

func (a AnyPassword) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}
