package membership

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Filter struct {
	Keyword       string
	Category      string
	Status        string
	FromDate      string
	ToDate        string
	FirstName     string
	MemberProfile bool
	Level         string
	OrderId       int
	TransactionId string
}

type TblMstrMembershiplevel struct {
	Id                    int       `gorm:"primaryKey;auto_increment;type:serial"`
	SubscriptionName      string    `gorm:"type:character varying"`
	Description           string    `gorm:"type:character varying"`
	MembergroupLevelId    int       `gorm:"type:integer"`
	InitialPayment        float64   `gorm:"type:decimal(10,2)"`
	RecurrentSubscription int       `gorm:"type:integer"`
	BillingAmount         float64   `gorm:"type:decimal(10,2)"`
	BillingfrequentValue  int       `gorm:"type:integer"`
	BillingfrequentType   int       `gorm:"type:integer"`
	BillingCyclelimit     int       `gorm:"type:integer"`
	CustomTrial           int       `gorm:"type:integer"`
	TrialBillingAmount    float64   `gorm:"type:decimal(10,2)"`
	TrialBillingLimit     int       `gorm:"type:integer"`
	CreatedOn             time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy             int       `gorm:"type:integer"`
	ModifiedOn            time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy             int       `gorm:"DEFAULT:NULL"`
	ModifiedBy            int       `gorm:"DEFAULT:NULL"`
	IsDeleted             int       `gorm:"type:integer"`
	IsActive              int       `gorm:"type:integer"`
	DeletedOn             time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId              int       `gorm:"DEFAULT:NULL"`
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
	DateString  string    `gorm:"-"`
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

type MembershipModel struct {
	Userid     int
	DataAccess int
}

var Membershipmodel MembershipModel

// membership pro models

func (Membershipmodel MembershipModel) GetMembershipGroup(Subscriptiongroup *[]TblMstrMembergrouplevel, offset int, limit int, filter Filter, tenantid int, DB *gorm.DB,isactive int) (Total_MembershipLevelgroup int64, err error) {

	query := DB.Table("tbl_mstr_membergrouplevels").Where("is_deleted=0")

	if isactive==1{
		query = query.Where("is_active=1")
	}

	if limit != 0 {

		query = query.Offset(offset).Limit(limit).Order("id desc")

		query.Find(&Subscriptiongroup)

		return Total_MembershipLevelgroup, nil

	}

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(group_name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")

		query.Find(&Subscriptiongroup)

		fmt.Println("group::", Subscriptiongroup)

		return Total_MembershipLevelgroup, nil
	}

	query.Count(&Total_MembershipLevelgroup).Find(&Subscriptiongroup)

	return Total_MembershipLevelgroup, nil
}

func (Membershipmodel MembershipModel) CreateMembershipGrouplevel(paygroup TblMstrMembergrouplevel, DB *gorm.DB) error {
	if err := DB.Table("tbl_mstr_membergrouplevels").Create(&paygroup).Error; err != nil {
		return err
	}
	return nil

}

func (membershipmodel MembershipModel) EditMembershipGroupLevel(Editmembershipgroup *TblMstrMembergrouplevel, Id int, DB *gorm.DB) error {
	if err := DB.Table("tbl_mstr_membergrouplevels").Debug().Where("id=?", Id).First(&Editmembershipgroup).Error; err != nil {
		return err
	}
	return nil
}

func (membershipmodel MembershipModel) UpdatemembershipGroup(membershipGroup TblMstrMembergrouplevel, tenantid int, DB *gorm.DB) error {
	if err := DB.Table("tbl_mstr_membergrouplevels").Debug().Where("id=? ", membershipGroup.Id).UpdateColumns(map[string]interface{}{"group_name": membershipGroup.GroupName, "description": membershipGroup.Description, "slug": membershipGroup.Slug, "modified_on": membershipGroup.ModifiedOn, "modified_by": membershipGroup.ModifiedBy}).Error; err != nil {
		return err
	}
	return nil
}

func (membershipmodel MembershipModel) DeleteMembershipgroup(membershipGroup TblMstrMembergrouplevel, DB *gorm.DB) error {
	if err := DB.Table("tbl_mstr_membergrouplevels").Where("id=?", membershipGroup.Id).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": membershipGroup.DeletedOn, "deleted_by": membershipGroup.DeletedBy}).Error; err != nil {
		return err
	}
	return nil
}

func (Membershipmodel MembershipModel) MultiselectDeleteMembershipGroup(membershipGroupids []int, DB *gorm.DB, deletedon time.Time, deletedby int) error {
	if err := DB.Table("tbl_mstr_membergrouplevels").Where("id IN (?)", membershipGroupids).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": deletedon, "deleted_by": deletedby}).Error; err != nil {
		return err
	}
	return nil
}

func (Membershipmodel MembershipModel) MembershipGroupChangeStatus(membershipgroupstatus TblMstrMembergrouplevel, membershipgroupid int, status int, DB *gorm.DB, tenantid int) error {
	if err := DB.Table("tbl_mstr_membergrouplevels").Debug().Where("id=? and tenant_id=?", membershipgroupid, tenantid).UpdateColumns(map[string]interface{}{"is_active": status, "modified_by": membershipgroupstatus.ModifiedBy, "modified_on": membershipgroupstatus.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

// membership level

func (membershipmodel MembershipModel) CreateSubscriptionLevel(subscriptions TblMstrMembershiplevel, DB *gorm.DB) error {

	if err := DB.Table("tbl_mstr_membershiplevels").Create(&subscriptions).Error; err != nil {
		return err
	}
	return nil
}

func (membershipmodel MembershipModel) GetMembershipLevel(offset int, limit int, filter Filter, sublist *[]TblMstrMembershiplevel, tenant_id int, DB *gorm.DB) (Total_MembershipLevel int64, err error) {
	query := DB.Table("tbl_mstr_membershiplevels").Where("tenant_id=? and is_deleted=0", tenant_id)

	if limit != 0 {

		query = query.Offset(offset).Limit(limit).Order("id desc")

		err := query.Find(&sublist).Error

		if err != nil {

			return 0, err
		}

		return Total_MembershipLevel, nil

	}

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(subscription_name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")

		err := query.Find(&sublist).Error

		if err != nil {

			return 0, err
		}

		return Total_MembershipLevel, nil
	}

	query.Count(&Total_MembershipLevel).Find(&sublist)
	fmt.Println("Total_MembershipLevel:", Total_MembershipLevel)
	return Total_MembershipLevel, nil
}

func (membershipmodel MembershipModel) GetdefaultTemplate(Defaultlist *[]TblMstrMembershiplevel, DB *gorm.DB) error {
	if err := DB.Table("tbl_mstr_membershiplevels").Debug().Where("tenant_id IS NULL").Find(&Defaultlist).Error; err != nil {
		return err
	}
	return nil
}

func (membershipmodel MembershipModel) GetMembershiplevelDetails(SelectedMembershiplevel *[]TblMstrMembershiplevel, levelId int, DB *gorm.DB) error {
	if err := DB.Table("tbl_mstr_membershiplevels").Debug().Where("is_deleted=0 and id=?", levelId).First(&SelectedMembershiplevel).Error; err != nil {
		return err
	}
	return nil
}

func (membershipmodel MembershipModel) Editmembershiplevel(Editmembership *TblMstrMembershiplevel, Id int, DB *gorm.DB) error {
	if err := DB.Table("tbl_mstr_membershiplevels").Debug().Where("is_deleted=0 and id=?", Id).First(&Editmembership).Error; err != nil {
		return err
	}
	return nil
}

func (membershipmodel MembershipModel) Subscriptionupdate(SubscriptionUpdate TblMstrMembershiplevel, tenantid int, DB *gorm.DB) error {
	if err := DB.Table("tbl_mstr_membershiplevels").Debug().Where(" id=?", SubscriptionUpdate.Id).UpdateColumns(map[string]interface{}{"subscription_name": SubscriptionUpdate.SubscriptionName, "description": SubscriptionUpdate.Description, "membergroup_level_id": SubscriptionUpdate.MembergroupLevelId, "initial_payment": SubscriptionUpdate.InitialPayment, "recurrent_subscription": SubscriptionUpdate.RecurrentSubscription, "billing_amount": SubscriptionUpdate.BillingAmount, "billingfrequent_value": SubscriptionUpdate.BillingfrequentValue, "billingfrequent_type": SubscriptionUpdate.BillingfrequentType, "billing_cyclelimit": SubscriptionUpdate.BillingCyclelimit, "custom_trial": SubscriptionUpdate.CustomTrial, "trial_billing_amount": SubscriptionUpdate.TrialBillingAmount, "trial_billing_limit": SubscriptionUpdate.TrialBillingLimit, "modified_on": SubscriptionUpdate.ModifiedOn, "modified_by": SubscriptionUpdate.ModifiedBy}).Error; err != nil {
		return err
	}
	return nil
}

func (membershipmodel MembershipModel) DeleteSubscription(SubscriptionDelete *TblMstrMembershiplevel, id int, DB *gorm.DB) error {
	if err := DB.Table("tbl_mstr_membershiplevels").Debug().Where("id=?", id).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": SubscriptionDelete.DeletedOn, "deleted_by": SubscriptionDelete.DeletedBy}).Error; err != nil {
		return err
	}
	return nil
}

func (Membershipmodel MembershipModel) MultiselectDeleteMembershipLevel(membershipLevelids []int, DB *gorm.DB, deletedon time.Time, deletedby int) error {
	if err := DB.Table("tbl_mstr_membershiplevels").Where("id IN (?)", membershipLevelids).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": deletedon, "deleted_by": deletedby}).Error; err != nil {
		return err
	}
	return nil
}

func (Membershipmodel MembershipModel) MembershiplevelChangeStatus(membershiplevelstatus TblMstrMembershiplevel, membershiplevelid int, status int, DB *gorm.DB, tenantid int) error {
	if err := DB.Table("tbl_mstr_membershiplevels").Debug().Where("id=? and tenant_id=?", membershiplevelid, tenantid).UpdateColumns(map[string]interface{}{"is_active": status, "modified_by": membershiplevelstatus.ModifiedBy, "modified_on": membershiplevelstatus.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}
