package membership

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TblMembershipSubcriptionslist struct {
	Id                        int       `gorm:"primaryKey;auto_increment;type:serial"`
	MemberId                  int       `gorm:"type:integer"`
	Gateway                   string    `gorm:"type:character varying"`
	GatewayEnvironment        string    `gorm:"type:character varying"`
	SubscriptionTransactionId string    `gorm:"type:character varying"`
	IsDeleted                 int       `gorm:"type:integer"`
	DeletedOn                 time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy                 int       `gorm:"DEFAULT:NULL"`
	CreatedOn                 time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy                 int       `gorm:"type:integer"`
	ModifiedOn                time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy                int       `gorm:"DEFAULT:NULL"`
	TenantId                  int       `gorm:"type:integer"`
	MembershipLevelId         int       `gorm:"type:integer"`
	MemberName                string    `gorm:"type:character varying"`
	IsActive                  int       `gorm:"type:integer"`
	DateString                string    `gorm:"-"`
	FirstName                 string    `gorm:"type:character varying"`
	LastName                  string    `gorm:"type:character varying"`
	SubscriptionName          string    `gorm:"type:character varying"`
	InitialPayment            float64   `gorm:"type:decimal(10,2)"`
}

type TblMembershipSubcriptions struct {
	Id                        int       `gorm:"primaryKey;auto_increment;type:serial"`
	MemberId                  int       `gorm:"type:integer"`
	Gateway                   string    `gorm:"type:character varying"`
	GatewayEnvironment        string    `gorm:"type:character varying"`
	SubscriptionTransactionId string    `gorm:"type:character varying"`
	IsDeleted                 int       `gorm:"type:integer"`
	DeletedOn                 time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy                 int       `gorm:"DEFAULT:NULL"`
	CreatedOn                 time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy                 int       `gorm:"type:integer"`
	ModifiedOn                time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy                int       `gorm:"DEFAULT:NULL"`
	TenantId                  int       `gorm:"type:integer"`
	MembershipLevelId         int       `gorm:"type:integer"`
	MemberName                string    `gorm:"type:character varying"`
	IsActive                  int       `gorm:"type:integer"`
	DateString                string    `gorm:"-"`

}
func (membershipmodel MembershipModel) ListSubscription(offset int, limit int, filter Filter, subscriptionlist *[]TblMembershipSubcriptionslist, tenant_id int, DB *gorm.DB) (Total_Subscription int64, err error) {
	query := DB.Table("tbl_membership_subcriptions").
		Select("tbl_membership_subcriptions.*, tbl_mstr_membershiplevels.subscription_name,tbl_mstr_membershiplevels.initial_payment, tbl_membership_members.first_name, tbl_membership_members.last_name").
		Where("tbl_membership_subcriptions.tenant_id=? and tbl_membership_subcriptions.is_deleted=0", tenant_id)

	query = query.Joins("INNER JOIN tbl_mstr_membershiplevels ON tbl_mstr_membershiplevels.id = tbl_membership_subcriptions.membership_level_id")

	query = query.Joins("INNER JOIN tbl_membership_members ON tbl_membership_members.id = tbl_membership_subcriptions.member_id")

	if limit != 0 {
		query = query.Offset(offset).Limit(limit).Order("tbl_membership_subcriptions.id desc")

		err := query.Debug().Find(&subscriptionlist).Error
		if err != nil {
			return 0, err
		}

		return Total_Subscription, nil
	}

	if filter.Keyword != "" {
		query = query.Where("LOWER(TRIM(tbl_membership_subcriptions.subscription_name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")

		err := query.Find(&subscriptionlist).Error
		if err != nil {
			return 0, err
		}

		return Total_Subscription, nil
	}

	query.Count(&Total_Subscription)

	return Total_Subscription, nil
}

func (Membershipmodel MembershipModel) CreateMembershipSubscription(SubscriptionCreate TblMembershipSubcriptions, DB *gorm.DB) error {
	fmt.Println("")
	if err := DB.Table("tbl_membership_subcriptions").Debug().Create(&SubscriptionCreate).Error; err != nil {
		return err
	}
	return nil
}



func (membershipmodel MembershipModel) EditSubscription(subscriptionId *TblMembershipSubcriptions, Id int, DB *gorm.DB) error {
	if err := DB.Table("tbl_membership_subcriptions").Debug().Where("id=?", Id).First(&subscriptionId).Error; err != nil {
		return err
	}
	return nil
}
