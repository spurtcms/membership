package membership

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func (Membershipmodel MembershipModel) GetMembershipGdroup(DB *gorm.DB) ([]TblMstrMembergrouplevel, error) {
	createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	fmt.Println("",createdon)
	var Subscriptiongroup []TblMstrMembergrouplevel

	if err := DB.Table("tbl_mstr_membergrouplevels").Where("is_deleted=0").Find(&Subscriptiongroup).Error; err != nil {
		return []TblMstrMembergrouplevel{}, err
	}
	return Subscriptiongroup, nil
}
