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

func (membershipmodel MembershipModel) ListMembers(MembershipMemberList *[]TblMembershipMembers, DB *gorm.DB) error {
	if err := DB.Table("tbl_membership_members").Where("is_deleted=0").Debug().Find(&MembershipMemberList).Error; err != nil {
		return err
	}
	return nil
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

func (Membershipmodel MembershipModel)MembershipDeleteMember(memberid int, DB *gorm.DB,deletedon time.Time,deletedby int) error {
	if err := DB.Table("tbl_membership_members").Debug().Where("id=?", memberid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": deletedon, "deleted_by": deletedby}).Error; err != nil {
		return err
	}
	return nil
}
