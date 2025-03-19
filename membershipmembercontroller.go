package membership

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (Membership *Membership) MembershipListMembers(offset int, limt int, filter Filter, flag bool, TenantId int) ([]TblMembershipMembers, int64) {

	var MembershipMemberList []TblMembershipMembers

	Totalmembercount, err := Membershipmodel.ListMembers(&MembershipMemberList, Membership.DB, 0, 0, filter, flag, TenantId)

	Membershipmodel.ListMembers(&MembershipMemberList, Membership.DB, offset, limt, filter, flag, TenantId)

	fmt.Println("err", err)

	return MembershipMemberList, Totalmembercount
}

func (Membership *Membership) CreateMembershipMembers(CreateMembershipMember TblMembershipMembers) {

	uvuid := (uuid.New()).String()

	customString := "M-"+uvuid[:5]

	var Createmember TblMembershipMembers
	Createmember.Uuid = customString
	Createmember.ProfileImage = CreateMembershipMember.ProfileImage
	Createmember.ProfileImagePath = CreateMembershipMember.ProfileImagePath
	Createmember.FirstName = CreateMembershipMember.FirstName
	Createmember.LastName = CreateMembershipMember.LastName
	Createmember.Email = CreateMembershipMember.Email
	Createmember.MobileNo = CreateMembershipMember.MobileNo
	Createmember.IsActive = CreateMembershipMember.IsActive
	Createmember.Username = CreateMembershipMember.Username
	if CreateMembershipMember.Password != "" {
		hash_pass := hashingPassword(CreateMembershipMember.Password)
		Createmember.Password = hash_pass
	}
	// Createmember.Password = CreateMembershipMember.Password
	Createmember.CreatedBy = CreateMembershipMember.CreatedBy
	Createmember.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	Createmember.StorageType = CreateMembershipMember.StorageType
	Createmember.TenantId = CreateMembershipMember.TenantId
	err := Membershipmodel.MemberCreateMembership(&Createmember, Membership.DB)

	fmt.Println("err", err)

}

func (Membership *Membership) EditMembershipMember(MemberId int) TblMembershipMembers {

	var MembershipMember TblMembershipMembers

	Membershipmodel.MembershipEditMember(&MembershipMember, MemberId, Membership.DB)

	return MembershipMember

}

func (Membership *Membership) UpdateMembershipMember(updatedMember TblMembershipMembers) {

	var MemberUpdate TblMembershipMembers
	// MemberUpdate.ProfileImage = updatedMember.ProfileImage
	// MemberUpdate.ProfileImagePath = updatedMember.ProfileImagePath
	MemberUpdate.Id = updatedMember.Id
	MemberUpdate.FirstName = updatedMember.FirstName
	MemberUpdate.LastName = updatedMember.LastName
	MemberUpdate.Email = updatedMember.Email
	MemberUpdate.MobileNo = updatedMember.MobileNo
	MemberUpdate.IsActive = updatedMember.IsActive
	MemberUpdate.Username = updatedMember.Username
	MemberUpdate.Password = updatedMember.Password
	MemberUpdate.ModifiedBy = updatedMember.ModifiedBy
	MemberUpdate.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	MemberUpdate.StorageType = updatedMember.StorageType
	MemberUpdate.TenantId = updatedMember.TenantId

	Membershipmodel.MembershipUpdateMember(MemberUpdate, Membership.DB)

}

func (Membership *Membership) DeleteMembershipMember(memberid int, userid int) {

	deletedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Membershipmodel.MembershipDeleteMember(memberid, Membership.DB, deletedon, userid)

}

func (Membership *Membership) DeleteMultiselectMember(memberids []int, userid int) {

	deletedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Membershipmodel.MultiselectDeleteMember(memberids, Membership.DB, deletedon, userid)

}

func (Membership *Membership) ChangeMembershipStatus(membershipid int, status int, modifiedby int, tenantid int) (bool, error) {
	var membershipstatus TblMembershipMembers
	membershipstatus.ModifiedBy = modifiedby
	membershipstatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Membershipmodel.MembershipChangeStatus(membershipstatus, membershipid, status, Membership.DB, tenantid)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (Membership *Membership) CreateCheckOut(name string, mail string, pass string, phonenumber string, companyname string, position string, tenant int, createdby int) (bool, error) {

	var checkoutdata TblMembershipMembers
	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	checkoutdata.FirstName = name
	checkoutdata.Username = name
	checkoutdata.Email = mail
	if pass != "" {
		hash_pass := hashingPassword(checkoutdata.Password)
		checkoutdata.Password = hash_pass
	}
	checkoutdata.MobileNo = phonenumber
	checkoutdata.TenantId = tenant
	checkoutdata.CreatedOn = time
	checkoutdata.IsActive = 0
	checkoutdata.CreatedBy = createdby

	_, err := Membershipmodel.CheckoutCreate(&checkoutdata, Membership.DB, companyname, position)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil

}
