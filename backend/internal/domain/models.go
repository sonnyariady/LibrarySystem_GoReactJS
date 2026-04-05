package models

import (
	"gorm.io/gorm"
)

// Book represents a book model
type Book struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"`
	Author      string `gorm:"size:255;not null"`
	ISBN        string `gorm:"size:13;unique;not null"`
	Loans       []Loan `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
}

// Member represents a library member model
type Member struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:255;not null"`
	Email      string `gorm:"size:255;unique;not null"`
	Loans      []Loan `gorm:"foreignKey:MemberID;constraint:OnDelete:CASCADE"`
}

// Loan represents a loan model
type Loan struct {
	ID       uint      `gorm:"primaryKey"`
	BookID   uint      `gorm:"not null"`
	MemberID uint      `gorm:"not null"`
	LoanDate string     `gorm:"not null"`
	ReturnDate *string  `gorm:"null"`
	Book     Book      `gorm:"foreignKey:BookID"`
	Member   Member    `gorm:"foreignKey:MemberID"`
}