package membership

import (
	"time"

	"gorm.io/gorm"
)


type TblMembershipOrder struct {
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

func (Membershipmodel MembershipModel) MemberShipOrderList(limit, offset int, filter Filter, tenantid int, DB *gorm.DB) (orderlist []TblMembershipOrder, count int64, err error) {

	var orderlistcount int64

	query := DB.Table("tbl_membership_orders").Where("is_deleted = 0 and  tenant_id = ?", tenantid).Order("tbl_membership_orders.created_on desc")

	query.Find(&orderlist).Count(&orderlistcount)
	if query.Error != nil {

		return []TblMembershipOrder{}, 0, query.Error
	}

	return orderlist, orderlistcount, nil

}

func (Membershipmodel MembershipModel) CreateMemberShipOrder(order TblMembershipOrder, DB *gorm.DB) error {

	if err := DB.Table("tbl_membership_orders").Create(&order).Error; err != nil {

		return err
	}

	return nil

}
