package membership

import (
	"strings"
	"time"

)


func (memsership *Membership) MembershipGroupList() []TblMstrMembergrouplevel {

	list, _ := Membershipmodel.GetMembershipGroup(memsership.DB)

	return list

}

func (membership *Membership) MembershipGroupLevelCreate(namae string, desc string, is_active int, tenantid int, userid int) {

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

func (membership *Membership) MembershipGrupUpdate(namae string, desc string, is_active int, tenantid int, userid int, id int) {

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

func (Membership *Membership) MembershipGroupDelete(id int, userid int, tenantid int) {
	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Groupupdate := TblMstrMembergrouplevel{
		Id:        id,
		DeletedOn: t,
		DeletedBy: userid,
		TenantId:  tenantid,
	}
	Membershipmodel.DeleteMembershipgroup(Groupupdate, Membership.DB)
}
