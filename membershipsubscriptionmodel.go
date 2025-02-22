package membership

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

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
	FirstName                 string    `gorm:"<-:false"`
	LastName                  string    `gorm:"<-:false"`
	SubscriptionName          string    `gorm:"<-:false"`
	InitialPayment            float64   `gorm:"<-:false"`
	UserId                    string    `gorm:"<-:false"`
	MembershiplevelId         string    `gorm:"<-:false"`
}

func (membershipmodel MembershipModel) ListSubscription(offset int, limit int, filter Filter, subscriptionlist *[]TblMembershipSubcriptions, tenant_id int, DB *gorm.DB) (Total_Subscription int64, err error) {
	query := DB.Table("tbl_membership_subcriptions").
		Select("tbl_membership_subcriptions.*, tbl_mstr_membershiplevels.subscription_name as subscription_name,tbl_mstr_membershiplevels.initial_payment as initial_payment, tbl_membership_members.first_name as first_name, tbl_membership_members.last_name as last_name").
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
		query = query.Where("LOWER(TRIM(tbl_membership_members.first_name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")

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

func (membershipmodel MembershipModel) EditSubscription(subscriptionId *TblMembershipSubcriptions, Id int, DB *gorm.DB, tenant_id int) error {

	query := DB.Table("tbl_membership_subcriptions").Select("tbl_membership_subcriptions.*, tbl_mstr_membershiplevels.subscription_name as subscription_name,tbl_membership_members.id as user_id,tbl_mstr_membershiplevels.id as membershiplevel_id,tbl_membership_members.first_name as first_name, tbl_membership_members.last_name as last_name").
		Where("tbl_membership_subcriptions.tenant_id=? and tbl_membership_subcriptions.id=? and  tbl_membership_subcriptions.is_deleted=0", tenant_id, Id)

	query = query.Joins("INNER JOIN tbl_mstr_membershiplevels ON tbl_mstr_membershiplevels.id = tbl_membership_subcriptions.membership_level_id")

	query = query.Joins("INNER JOIN tbl_membership_members ON tbl_membership_members.id = tbl_membership_subcriptions.member_id")

	err := query.Debug().First(&subscriptionId).Error

	if err != nil {
		return err
	}
	return nil
}

func (Membershipmodel MembershipModel) UpdateSubscription(UpdatedSubscription TblMembershipSubcriptions, DB *gorm.DB) error {
	if err := DB.Table("tbl_membership_subcriptions").Debug().Where(" id=?", UpdatedSubscription.Id).UpdateColumns(map[string]interface{}{"subscription_transaction_id": UpdatedSubscription.SubscriptionTransactionId, "gateway": UpdatedSubscription.Gateway, "gateway_environment": UpdatedSubscription.GatewayEnvironment, "membership_level_id": UpdatedSubscription.MembershipLevelId, "member_id": UpdatedSubscription.MemberId, "modified_on": UpdatedSubscription.ModifiedOn, "modified_by": UpdatedSubscription.ModifiedBy}).Error; err != nil {
		return err
	}
	return nil
}

func (Membershipmodel MembershipModel) DeleteSubscriptions(id, tenantid, deletedby int, deletedon time.Time, DB *gorm.DB) error {

	if err := DB.Table("tbl_membership_subcriptions").Where("id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": deletedby, "deleted_on": deletedon}).Error; err != nil {

		return err
	}

	return nil
}

func (Membershipmodel MembershipModel) SubscriptionChangeStatus(subscriptionstatus TblMembershipSubcriptions, subscriptionid int, status int, DB *gorm.DB, tenantid int) error {

	fmt.Println("status:",status,subscriptionid)
	if err := DB.Table("tbl_membership_subcriptions").Debug().Where("id=? and tenant_id=?", subscriptionid, tenantid).UpdateColumns(map[string]interface{}{"is_active": status, "modified_by": subscriptionstatus.ModifiedBy, "modified_on": subscriptionstatus.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}




func (Membershipmodel MembershipModel) MultiselectDeleteSubscription(SubscriptionIds []int, DB *gorm.DB, deletedon time.Time, deletedby int) error {
	if err := DB.Table("tbl_membership_subcriptions").Where("id IN (?)", SubscriptionIds).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": deletedon, "deleted_by": deletedby}).Error; err != nil {
		return err
	}
	return nil
}
