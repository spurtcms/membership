package membership

import (
	"fmt"
	"time"
)

func (Membership *Membership) MembershipListMembers(offset int, limt int, filter Filter, flag bool, TenantId int)( []TblMembershipMembers , int64) {

	var MembershipMemberList []TblMembershipMembers

	_,err := Membershipmodel.ListMembers(&MembershipMemberList, Membership.DB,offset , limt ,filter, flag , TenantId )

	Totalmembercount,err:= Membershipmodel.ListMembers(&MembershipMemberList, Membership.DB,0, 0 ,filter, flag , TenantId )

	fmt.Println("err", err)

	return MembershipMemberList,Totalmembercount
}

func (Membership *Membership) CreateMembershipMembers(CreateMembershipMember TblMembershipMembers) {

	// uvuid := (uuid.New()).String()

	var Createmember TblMembershipMembers
	// Createmember.Uuid = uvuid
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
