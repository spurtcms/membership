package postgres

import (
	"time"

	"gorm.io/gorm"
)

// Paid membership tables

type TblMstrMembershiplevel struct {
	Id                     int       `gorm:"primaryKey;auto_increment;type:serial"`
	SubscriptionName       string    `gorm:"type:character varying"`
	Description            string    `gorm:"type:character varying"`
	MembershiplevelDetails string    `gorm:"type:character varying"`
	MembergroupLevelId     int       `gorm:"type:integer"`
	InitialPayment         float64   `gorm:"type:decimal(10,2)"`
	IsDiscount             int       `gorm:"type:integer"`
	DiscountPercentage     int       `gorm:"type:integer"`
	DiscountedAmount       float64   `gorm:"type:decimal(10,2)"`
	RecurrentSubscription  int       `gorm:"type:integer"`
	BillingAmount          float64   `gorm:"type:decimal(10,2)"`
	BillingfrequentValue   int       `gorm:"type:integer"`
	BillingfrequentType    int       `gorm:"type:integer"`
	BillingCyclelimit      int       `gorm:"type:integer"`
	CustomTrial            int       `gorm:"type:integer"`
	TrialBillingAmount     float64   `gorm:"type:decimal(10,2)"`
	TrialBillingLimit      int       `gorm:"type:integer"`
	CreatedOn              time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy              int       `gorm:"type:integer"`
	ModifiedOn             time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy              int       `gorm:"DEFAULT:NULL"`
	ModifiedBy             int       `gorm:"DEFAULT:NULL"`
	IsDeleted              int       `gorm:"type:integer"`
	IsActive               int       `gorm:"type:integer"`
	DeletedOn              time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId               string    `gorm:"type:character varying"`
	GroupName              string    `gorm:"-"`
}

type TblMstrMembergrouplevel struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	GroupName   string    `gorm:"type:character varying"`
	Description string    `gorm:"type:character varying"`
	Slug        string    `gorm:"type:character varying"`
	CreatedOn   time.Time `gorm:"column:created_on;type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:integer"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:integer"`
	IsActive    int       `gorm:"type:integer"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL"`
	TenantId    int       `gorm:"type:integer"`
}

type MemberCheckoutDetails struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	UserName    string    `gorm:"type:character varying"`
	EmailId     string    `gorm:"type:character varying"`
	Password    string    `gorm:"type:character varying"`
	CompanyName string    `gorm:"type:character varying"`
	Position    string    `gorm:"type:character varying"`
	CreatedBy   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone"`
	ModifiedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsActive    int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:integer;DEFAULT:0"`
	DeletedBy   int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId    int       `gorm:"type:integer;DEFAULT:NULL"`
}

// MigrateTable creates this package related tables in your database
func MigrateTables(db *gorm.DB) {

	db.AutoMigrate(&TblMstrMembershiplevel{}, &TblMstrMembergrouplevel{}, &MemberCheckoutDetails{})

}
