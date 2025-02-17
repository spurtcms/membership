package membership

import (
	"time"

	"gorm.io/gorm"
)

type TblMembershipOrdeddfr struct {
	Id                        int       `gorm:"primaryKey;auto_increment;type:serial"`
	UserId                    int       `gorm:"type:integer"`
	MembershiplevelId         int       `gorm:"type:integer"`
	BillingName               string    `gorm:"type:character varying"`
	BillingStreet             string    `gorm:"type:character varying"`
	BillingStreet2            string    `gorm:"type:character varying"`
	BillingCity               string    `gorm:"type:character varying"`
	BillingState              string    `gorm:"type:character varying"`
	BillingPostalcode         string    `gorm:"type:character varying"`
	BillingCountry            string    `gorm:"type:character varying"`
	BillingPhone              string    `gorm:"type:character varying"`
	SubTotal                  int       `gorm:"type:integer"`
	Tax                       int       `gorm:"type:integer"`
	Total                     int       `gorm:"type:integer"`
	PaymentType               string    `gorm:"type:character varying"`
	Status                    string    `gorm:"type:character varying"`
	Gateway                   string    `gorm:"type:character varying"`
	GatewayEnvironment        string    `gorm:"type:character varying"`
	PaymenttransactionId      int       `gorm:"type:integer"`
	SubscriptiontransactionId int       `gorm:"type:integer"`
	CreatedOn                 time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy                 int       `gorm:"type:integer"`
	IsDeleted                 int       `gorm:"type:integer"`
	DeletedOn                 time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy                 int       `gorm:"DEFAULT:NULL"`
	ModifiedOn                time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy                int       `gorm:"DEFAULT:NULL"`
	TenantId                  int       `gorm:"type:integer"`
	DateString                string    `gorm:"-"`
}


func (Membershipmodel MembershipModel) CreateMembershipGrouplevedl(paygroup TblMstrMembergrouplevel, DB *gorm.DB) error {
	if err := DB.Table("tbl_mstr_membergrouplevels").Create(&paygroup).Error; err != nil {
		return err
	}
	return nil

}