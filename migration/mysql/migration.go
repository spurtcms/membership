package mysql

import (
	"time"

	
	"gorm.io/gorm"
)

// Paid membership tables

type TblMstrMembershiplevel struct {
	Id                    int       `gorm:"primaryKey;auto_increment;type:int"`
	SubscriptionName      string    `gorm:"type:varchar(255)"`
	Description           string    `gorm:"type:varchar(255)"`
	MembergroupLevelId    int       `gorm:"type:int"`
	InitialPayment        float64   `gorm:"type:decimal(10,2)"`
	RecurrentSubscription int       `gorm:"type:int"`
	BillingAmount         float64   `gorm:"type:decimal(10,2)"`
	BillingfrequentValue  int       `gorm:"type:int"`
	BillingfrequentType   int       `gorm:"type:int"`
	BillingCyclelimit     int       `gorm:"type:int"`
	CustomTrial           int       `gorm:"type:int"`
	TrialBillingAmount    float64   `gorm:"type:decimal(10,2)"`
	TrialBillingLimit     int       `gorm:"type:int"`
	CreatedOn             time.Time `gorm:"type:timestamp"`
	CreatedBy             int       `gorm:"type:int"`
	ModifiedOn            time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	DeletedBy             int       `gorm:"DEFAULT:NULL"`
	ModifiedBy            int       `gorm:"DEFAULT:NULL"`
	IsDeleted             int       `gorm:"type:int"`
	IsActive              int       `gorm:"type:int"`
	DeletedOn             time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	TenantId              int       `gorm:"DEFAULT:NULL"`
}

type TblMstrMembergrouplevel struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:int"`
	GroupName   string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	Slug        string    `gorm:"type:varchar(255)"`
	CreatedOn   time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:int"`
	IsActive    int       `gorm:"type:int"`
	DeletedOn   time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL"`
	TenantId    int       `gorm:"type:int"`
}

type MemberCheckoutDetails struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:int"`
	UserName    string    `gorm:"type:varchar(255)"`
	EmailId     string    `gorm:"type:varchar(255)"`
	Password    string    `gorm:"type:varchar(255)"`
	CompanyName string    `gorm:"type:varchar(255)"`
	Position    string    `gorm:"type:varchar(255)"`
	CreatedBy   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:timestamp"`
	ModifiedBy  int       `gorm:"type:int;DEFAULT:NULL"`
	ModifiedOn  time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	IsActive    int       `gorm:"type:int;DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:int;DEFAULT:0"`
	DeletedBy   int       `gorm:"type:int;DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	TenantId    int       `gorm:"type:int;DEFAULT:NULL"`
}


// MigrateTable creates this package related tables in your database

func MigrateTables(db *gorm.DB) {

	db.AutoMigrate(&TblMstrMembershiplevel{}, &TblMstrMembergrouplevel{}, &MemberCheckoutDetails{})

}
