package membership

import (
	"fmt"
	"strings"
	"time"
)

func (memsership *Membership) MembershipGroupList(offset int, limt int, filter Filter, tenantid string, isactive int) ([]TblMstrMembergrouplevel, int64) {
	fmt.Println("")

	var Subscriptiongroup []TblMstrMembergrouplevel

	Membershipmodel.GetMembershipGroup(&Subscriptiongroup, offset, limt, filter, tenantid, memsership.DB, isactive)

	totalgroup, _ := Membershipmodel.GetMembershipGroup(&Subscriptiongroup, 0, 0, filter, tenantid, memsership.DB, isactive)

	return Subscriptiongroup, totalgroup

}

func (membership *Membership) MembershipGroupLevelCreate(namae string, desc string, is_active int, tenantid string, userid int) {

	// createdOnStr := time.Now().UTC().Format("2006-01-02 15:04:05")

	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	groupdata := TblMstrMembergrouplevel{
		GroupName:   namae,
		Description: desc,
		TenantId:    tenantid,
		Slug:        strings.ToLower(namae),
		IsActive:    is_active,
		CreatedOn:   t,
		CreatedBy:   userid,
	}

	Membershipmodel.CreateMembershipGrouplevel(groupdata, membership.DB)

}

func (membership *Membership) MembershipGroupLevelEdit(MembershipGroupId int) TblMstrMembergrouplevel {

	var MembershipGroupLevelEdit TblMstrMembergrouplevel

	Membershipmodel.EditMembershipGroupLevel(&MembershipGroupLevelEdit, MembershipGroupId, membership.DB)

	return MembershipGroupLevelEdit

}

func (membership *Membership) MembershipGrupUpdate(namae string, desc string, is_active int, tenantid string, userid int, id int) {

	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Groupupdate := TblMstrMembergrouplevel{
		Id:          id,
		GroupName:   namae,
		Description: desc,
		TenantId:    tenantid,
		Slug:        strings.ToLower(namae),
		IsActive:    is_active,
		ModifiedOn:  t,
		ModifiedBy:  userid,
	}

	Membershipmodel.UpdatemembershipGroup(Groupupdate, tenantid, membership.DB)

}

func (Membership *Membership) MembershipGroupDelete(id int, userid int, tenantid string) {
	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Groupupdate := TblMstrMembergrouplevel{
		Id:        id,
		DeletedOn: t,
		DeletedBy: userid,
		TenantId:  tenantid,
	}
	Membershipmodel.DeleteMembershipgroup(Groupupdate, Membership.DB)
}

func (Membership *Membership) DeleteMultiselectMembershipGroup(MembershipGroupids []int, userid int) {

	deletedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Membershipmodel.MultiselectDeleteMembershipGroup(MembershipGroupids, Membership.DB, deletedon, userid)

}

func (Membership *Membership) ChangeMembershipGroupStatus(membershipGroupid int, status int, modifiedby int, tenantid string) (bool, error) {
	var membershipGroupstatus TblMstrMembergrouplevel
	membershipGroupstatus.ModifiedBy = modifiedby
	membershipGroupstatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Membershipmodel.MembershipGroupChangeStatus(membershipGroupstatus, membershipGroupid, status, Membership.DB, tenantid)
	if err != nil {
		return false, err
	}

	return true, nil
}
