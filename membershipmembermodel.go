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
	TenantId         int       `gorm:"type:integer"`
	DateString       string    `gorm:"-"`
}

func demo() {
	fmt.Println("ds")
	strings.ToLower("dskjv ")
	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	fmt.Println("t", t)

}

func (membershipmodel MembershipModel) ListMembers(MembershipMemberList *[]TblMembershipMembers, DB *gorm.DB, offset int, limit int, filter Filter, flag bool, TenantId int) (Total_Members int64, err error) {
	query := DB.Table("tbl_membership_members").Where("is_deleted=0")

	if limit != 0 {

		query = query.Offset(offset).Limit(limit).Order("id desc")

		query.Find(&MembershipMemberList)

		return Total_Members, nil

	}

	if filter.Keyword != "" {

		query = query.Debug().Where("LOWER(TRIM(first_name)) LIKE LOWER(TRIM(?))"+" OR LOWER(TRIM(last_name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

		query.Find(&MembershipMemberList)

		return Total_Members, nil
	}

	query.Count(&Total_Members)

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

func (Membershipmodel MembershipModel) MembershipChangeStatus(membershipstatus TblMembershipMembers, membershipid int, status int, DB *gorm.DB, tenantid int) error {
	if err := DB.Table("tbl_membership_members").Where("id=? and tenant_id=?", membershipid, tenantid).UpdateColumns(map[string]interface{}{"is_active": status, "modified_by": membershipstatus.ModifiedBy, "modified_on": membershipstatus.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}
