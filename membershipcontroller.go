package membership

import (
	"time"

	"github.com/spurtcms/membership/migration"
)

// MemberSetup used initialize member configruation
func MembershipSetup(config Config) *Membership {

	migration.AutoMigration(config.DB, config.DataBaseType)

	return &Membership{
		AuthEnable:       config.AuthEnable,
		Permissions:      config.Permissions,
		PermissionEnable: config.PermissionEnable,
		Auth:             config.Auth,
		DB:               config.DB,
	}

}

// func (memsership *Membership) MembershipGroupList() []TblMstrMembergrouplevel {

// 	list, _ := Membershipmodel.GetMembershipGroup(memsership.DB)

// 	return list

// }

// func (membership *Membership) MembershipGroupLevelCreate(namae string, desc string, is_active int, tenantid int, userid int) {

// 	// createdOnStr := time.Now().UTC().Format("2006-01-02 15:04:05")

// 	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

// 	groupdata := TblMstrMembergrouplevel{
// 		GroupName:   namae,
// 		Description: desc,
// 		TenantId:    tenantid,
// 		Slug:        strings.ToLower(namae),
// 		IsActive:    is_active,
// 		CreatedOn:   t,
// 		CreatedBy:   userid,
// 	}

// 	Membershipmodel.CreateMembershipGrouplevel(groupdata, membership.DB)

// }

// func (membership *Membership) MembershipGrupUpdate(namae string, desc string, is_active int, tenantid int, userid int, id int) {

// 	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

// 	Groupupdate := TblMstrMembergrouplevel{
// 		Id:          id,
// 		GroupName:   namae,
// 		Description: desc,
// 		TenantId:    tenantid,
// 		Slug:        strings.ToLower(namae),
// 		IsActive:    is_active,
// 		ModifiedOn:  t,
// 		ModifiedBy:  userid,
// 	}

// 	Membershipmodel.UpdatemembershipGroup(Groupupdate, tenantid, membership.DB)

// }

// func (Membership *Membership) MembershipGroupDelete(id int, userid int, tenantid int) {
// 	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

// 	Groupupdate := TblMstrMembergrouplevel{
// 		Id:        id,
// 		DeletedOn: t,
// 		DeletedBy: userid,
// 		TenantId:  tenantid,
// 	}
// 	Membershipmodel.DeleteMembershipgroup(Groupupdate, Membership.DB)
// }

func (Membership *Membership) MembershipLevelsList(tenantid int) []TblMstrMembershiplevel {

	var subscriptionlist []TblMstrMembershiplevel

	Membershipmodel.GetMembershipLevel(&subscriptionlist, tenantid, Membership.DB)

	return subscriptionlist

}

func (Membership *Membership) GetdefaultMembershiplevelTemplate() []TblMstrMembershiplevel {

	var DefaultMembershipLevelList []TblMstrMembershiplevel

	Membershipmodel.GetdefaultTemplate(&DefaultMembershipLevelList, Membership.DB)

	return DefaultMembershipLevelList

}
func (Membership *Membership) MembershiplevelDetails(membershiplevelId int) []TblMstrMembershiplevel {

	var SelectedMembershipData []TblMstrMembershiplevel

	Membershipmodel.GetMembershiplevelDetails(&SelectedMembershipData, membershiplevelId, Membership.DB)

	return SelectedMembershipData

}

func (Membership *Membership) MembershiplevelEdit(membershipid int) TblMstrMembershiplevel {

	var Editmembership TblMstrMembershiplevel

	Membershipmodel.Editmembershiplevel(&Editmembership, membershipid, Membership.DB)

	return Editmembership

}

func (Membership *Membership) MembershipLevelsCreate(sd TblMstrMembershiplevel, tenantid int) {

	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	var subscriptiondata TblMstrMembershiplevel

	subscriptiondata.SubscriptionName = sd.SubscriptionName
	subscriptiondata.Description = sd.Description
	subscriptiondata.MembergroupLevelId = sd.MembergroupLevelId
	subscriptiondata.InitialPayment = sd.InitialPayment
	subscriptiondata.RecurrentSubscription = sd.RecurrentSubscription
	subscriptiondata.BillingAmount = sd.BillingAmount
	subscriptiondata.BillingfrequentValue = sd.BillingfrequentValue
	subscriptiondata.BillingfrequentType = sd.BillingfrequentType
	subscriptiondata.CustomTrial = sd.CustomTrial
	subscriptiondata.TrialBillingAmount = sd.TrialBillingAmount
	subscriptiondata.TrialBillingLimit = sd.TrialBillingLimit
	subscriptiondata.CreatedOn = t
	subscriptiondata.TenantId = tenantid
	subscriptiondata.IsActive = sd.IsActive

	Membershipmodel.CreateSubscriptionLevel(subscriptiondata, Membership.DB)

}

func (Membership *Membership) UpdateSubscription(subscriptionNewdata TblMstrMembershiplevel, tenantid int) {

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	var Updatesubscription TblMstrMembershiplevel
	Updatesubscription.Id = subscriptionNewdata.Id
	Updatesubscription.SubscriptionName = subscriptionNewdata.SubscriptionName
	Updatesubscription.Description = subscriptionNewdata.Description
	Updatesubscription.MembergroupLevelId = subscriptionNewdata.MembergroupLevelId
	Updatesubscription.InitialPayment = subscriptionNewdata.InitialPayment
	Updatesubscription.RecurrentSubscription = subscriptionNewdata.RecurrentSubscription
	Updatesubscription.BillingAmount = subscriptionNewdata.BillingAmount
	Updatesubscription.BillingfrequentValue = subscriptionNewdata.BillingfrequentValue
	Updatesubscription.BillingfrequentType = subscriptionNewdata.BillingfrequentType
	Updatesubscription.BillingCyclelimit = subscriptionNewdata.BillingCyclelimit
	Updatesubscription.CustomTrial = subscriptionNewdata.CustomTrial
	Updatesubscription.ModifiedBy = subscriptionNewdata.ModifiedBy
	Updatesubscription.TrialBillingAmount = subscriptionNewdata.TrialBillingAmount
	Updatesubscription.TrialBillingLimit = subscriptionNewdata.TrialBillingLimit
	Updatesubscription.IsActive = subscriptionNewdata.IsActive
	Updatesubscription.ModifiedOn = time

	Membershipmodel.Subscriptionupdate(Updatesubscription, tenantid, Membership.DB)

}

func (Membership *Membership) SubscriptionDelete(tenantid int, id int, userid int) {

	var subscriptionlist TblMstrMembershiplevel

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	subscriptionlist.DeletedOn = time
	subscriptionlist.DeletedBy = userid
	subscriptionlist.TenantId = tenantid

	Membershipmodel.DeleteSubscription(&subscriptionlist, id, Membership.DB)

}
