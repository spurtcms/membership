package membership

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type TblMembershipMembers struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	Uuid             string    `gorm:"type:character varying"`
	FirstName        string    `gorm:"type:character varying"`
	LastName         string    `gorm:"type:character varying"`
	Email            string    `gorm:"type:character varying"`
	MobileNo         string    `gorm:"type:character varying"`
	IsActive         int       `gorm:"type:integer"`
	ProfileImage     string    `gorm:"type:character varying"`
	ProfileImagePath string    `gorm:"type:character varying"`
	LastLogin        int       `gorm:"type:integer"`
	Password         string    `gorm:"type:character varying"`
	Username         string    `gorm:"DEFAULT:NULL"`
	Otp              int       `gorm:"DEFAULT:NULL"`
	OtpExpiry        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	LoginTime        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	StorageType      string    `gorm:"type:character varying"`
	TenantId         string    `gorm:"type:character varying"`
	DateString       string    `gorm:"-"`
	DateStringend    string    `gorm:"-"`

	SubscriptionName string    `gorm:"<-:false"`
	InitialPayment   float64   `gorm:"<-:false"`
	PlanDuration     int       `gorm:"<-:false"`
	MultiplyDuration int       `gorm:"<-:false"`
	EndDate          time.Time `gorm:"<-:false"`
}

type TblMembershipCompanyInfo struct {
	Id          int    `gorm:"primaryKey;auto_increment;type:serial"`
	MemberId    int    `gorm:"type:integer"`
	CompanyName string `gorm:"type:character varying"`
	Position    string `gorm:"type:character varying"`
	TenantId    string `gorm:"type:character varying"`
}

func demo() {
	fmt.Println("ds")
	strings.ToLower("dskjv ")
	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	fmt.Println("t", t)

}

// func (membershipmodel MembershipModel) ListMembers(MembershipMemberList *[]TblMembershipMembers, DB *gorm.DB, offset int, limit int, filter Filter, flag bool, TenantId string) (Total_Members int64, err error) {
// 	// Start building the SQL query
// 	query := DB.Table("tbl_membership_members").
// 		Debug().
// 		Select("tbl_membership_members.*, tbl_mstr_membershiplevels.subscription_name as subscription_name, tbl_mstr_membershiplevels.initial_payment as initial_payment, tbl_mstr_membershiplevels.billingfrequent_type as plan_uration, tbl_mstr_membershiplevels.billingfrequent_value as multiply_duration,tbl_membership_subscriptions.created_on as end_date").
// 		Where("tbl_membership_members.tenant_id = ? AND tbl_membership_members.is_deleted = 0", TenantId)

// 	// Join the tables with proper references
// 	query = query.Joins("left JOIN tbl_membership_subcriptions ON tbl_membership_subcriptions.member_id = tbl_membership_members.id").
// 		Joins("left JOIN tbl_mstr_membershiplevels ON tbl_mstr_membershiplevels.id = tbl_membership_subcriptions.membership_level_id")

// 	// Apply limit and offset if provided
// 	if limit != 0 {
// 		query = query.Offset(offset).Limit(limit).Order("tbl_membership_members.id DESC")

// 		// Execute the query and return the results
// 		query.Find(&MembershipMemberList)

// 		return Total_Members, nil
// 	}

// 	// If a filter keyword is provided, apply the filter
// 	if filter.Keyword != "" {
// 		query = query.Debug().
// 			Where("LOWER(TRIM(tbl_membership_members.first_name)) LIKE LOWER(TRIM(?)) OR LOWER(TRIM(tbl_membership_members.last_name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

// 		// Execute the query and return the results
// 		query.Find(&MembershipMemberList)

// 		return Total_Members, nil
// 	}

// 	// Get the total count of members
// 	query.Count(&Total_Members).Find(&MembershipMemberList)

// 	return Total_Members, nil
// }

func (membershipmodel MembershipModel) ListMembers(MembershipMemberList *[]TblMembershipMembers, DB *gorm.DB, offset int, limit int, filter Filter, flag bool, TenantId string) (Total_Members int64, err error) {
	// Start building the SQL query
	query := DB.Table("tbl_membership_members").
		Debug().
		Select("tbl_membership_members.*, tbl_mstr_membershiplevels.subscription_name as subscription_name, tbl_mstr_membershiplevels.initial_payment as initial_payment, tbl_mstr_membershiplevels.billingfrequent_type as plan_duration, tbl_mstr_membershiplevels.billingfrequent_value as multiply_duration, tbl_membership_subcriptions.created_on as end_date").
		Where("tbl_membership_members.tenant_id = ? AND tbl_membership_members.is_deleted = 0", TenantId)

	// Join the tables with proper references
	query = query.Joins("left JOIN tbl_membership_subcriptions ON tbl_membership_subcriptions.member_id = tbl_membership_members.id").
		Joins("left JOIN tbl_mstr_membershiplevels ON tbl_mstr_membershiplevels.id = tbl_membership_subcriptions.membership_level_id")

	// Apply limit and offset if provided
	if limit != 0 {
		query = query.Offset(offset).Limit(limit).Order("tbl_membership_members.id DESC")

	}

	// If a filter keyword is provided, apply the filter
	if filter.Keyword != "" {
		query = query.Debug().
			Where("LOWER(TRIM(tbl_membership_members.first_name)) LIKE LOWER(TRIM(?)) OR LOWER(TRIM(tbl_membership_members.email)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

	}

	if filter.Level != "" {

		query = query.Where("LOWER(TRIM(tbl_mstr_membershiplevels.subscription_name)) like LOWER(TRIM(?))", "%"+filter.Level+"%")

	}
	if filter.FromDate != "" {
		query = query.Where("tbl_membership_subcriptions.created_on >= ? AND tbl_membership_subcriptions.created_on < ?",
			filter.FromDate+" 00:00:00",
			filter.FromDate+" 23:59:59")
	}

	// Get the total count of members
	query.Count(&Total_Members).Find(&MembershipMemberList)

	return Total_Members, nil
}

func (Membershipmodel MembershipModel) MemberCreateMembership(membercreate *TblMembershipMembers, DB *gorm.DB) error {

	if err := DB.Table("tbl_membership_members").Create(&membercreate).Error; err != nil {
		return err
	}
	return nil

}

func (Membershipmodel MembershipModel) MembershipEditMember(MembershipMember *TblMembershipMembers, Id int, DB *gorm.DB) error {
	if err := DB.Table("tbl_membership_members").Debug().Where("id=?", Id).First(&MembershipMember).Error; err != nil {
		return err
	}
	return nil
}

func (Membershipmodel MembershipModel) MembershipUpdateMember(UpdatedMember TblMembershipMembers, DB *gorm.DB) error {
	if err := DB.Table("tbl_membership_members").Debug().Where(" id=?", UpdatedMember.Id).UpdateColumns(map[string]interface{}{"first_name": UpdatedMember.FirstName, "last_name": UpdatedMember.LastName, "email": UpdatedMember.Email, "mobile_no": UpdatedMember.MobileNo, "is_active": UpdatedMember.IsActive, "password": UpdatedMember.Password, "username": UpdatedMember.Username, "modified_on": UpdatedMember.ModifiedOn, "modified_by": UpdatedMember.ModifiedBy}).Error; err != nil {
		return err
	}
	return nil
}

func (Membershipmodel MembershipModel) MembershipDeleteMember(memberid int, DB *gorm.DB, deletedon time.Time, deletedby int) error {
	if err := DB.Table("tbl_membership_members").Debug().Where("id=?", memberid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": deletedon, "deleted_by": deletedby}).Error; err != nil {
		return err
	}
	return nil
}

func (Membershipmodel MembershipModel) MultiselectDeleteMember(memberids []int, DB *gorm.DB, deletedon time.Time, deletedby int) error {
	if err := DB.Table("tbl_membership_members").Where("id IN (?)", memberids).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": deletedon, "deleted_by": deletedby}).Error; err != nil {
		return err
	}
	return nil
}

// Membership  IsActive Function

func (Membershipmodel MembershipModel) MembershipChangeStatus(membershipstatus TblMembershipMembers, membershipid int, status int, DB *gorm.DB, tenantid string) error {
	if err := DB.Table("tbl_membership_members").Where("id=? and tenant_id=?", membershipid, tenantid).UpdateColumns(map[string]interface{}{"is_active": status, "modified_by": membershipstatus.ModifiedBy, "modified_on": membershipstatus.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

func (Membershipmodel MembershipModel) CheckoutCreate(Checkout *TblMembershipMembers, DB *gorm.DB, companyname, position string) (bool, error) {

	if err := DB.Table("tbl_membership_members").Create(&Checkout).Error; err != nil {
		return false, nil
	}

	var companyinfo = TblMembershipCompanyInfo{
		MemberId:    Checkout.Id,
		CompanyName: companyname,
		Position:    position,
		TenantId:    Checkout.TenantId,
	}

	if err := DB.Table("tbl_membership_company_infos").Create(&companyinfo).Error; err != nil {
		return false, nil
	}

	return true, nil
}
