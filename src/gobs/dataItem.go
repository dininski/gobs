package gobs

import ("time"
)

type DataItem struct {
    Id string
    Owner string
    CreatedBy string
    ModifiedBy string
    CreatedAt time.Time
    ModifiedAt time.Time
    contentType string
    gobsInstance gobs
}