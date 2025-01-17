package dto

import "time"

type Path struct {
        ID          int       `json:"id"`
        Name        string    `json:"name"`
        GUID        string    `json:"guid"`
        Description string    `json:"description"`
        CreatedAt   time.Time `json:"created_at"`
        UpdatedAt   time.Time `json:"updated_at"`
        DeletedAt   *time.Time `json:"deleted_at"`
}

type Content struct {
        ID          int       `json:"id"`
        Name        string    `json:"name"`
        GUID        string    `json:"guid"`
        Description string    `json:"description"`
        CreatedAt   time.Time `json:"created_at"`
        UpdatedAt   time.Time `json:"updated_at"`
        DeletedAt   *time.Time `json:"deleted_at"`
}

type ContentPath struct {
        ID        int       `json:"id"`
        ContentID int       `json:"content_id"`
        PathID    int       `json:"path_id"`
        CreatedAt time.Time `json:"created_at"`
        UpdatedAt time.Time `json:"updated_at"`
        DeletedAt *time.Time `json:"deleted_at"`
}