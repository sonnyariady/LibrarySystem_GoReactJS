package models

import "gorm.io/gorm"

// Book represents a book in the library
type Book struct {
    gorm.Model
    Title       string `json:"title" gorm:"unique"`
    Author      string `json:"author"`
    ISBN        string `json:"isbn" gorm:"unique"`
    PublishedAt string `json:"published_at"`
}

// Member represents a library member
type Member struct {
    gorm.Model
    Name      string `json:"name"`
    Email     string `json:"email" gorm:"unique"`
    CreatedAt time.Time `json:"created_at"`
}

// Loan represents a loan record
type Loan struct {
    gorm.Model
    BookID  uint   `json:"book_id" gorm:"not null"`
    MemberID uint   `json:"member_id" gorm:"not null"`
    LoanedAt time.Time `json:"loaned_at"`
    ReturnedAt *time.Time `json:"returned_at,omitempty"`
}

// PaginationResponse is used for paginated responses
type PaginationResponse struct {
    Total     int         `json:"total"`
    PerPage   int         `json:"per_page"`
    CurrentPage int       `json:"current_page"`
    Data      interface{} `json:"data"`
}

// APIResponse is a standard response format
type APIResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}
