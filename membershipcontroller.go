package membership

import (
	"fmt"
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



func (Membership *Membership) MembershipLevelsList(offset int, limt int, filter Filter, tenantid int) ([]TblMstrMembershiplevel, int64) {

	var subscriptionlist []TblMstrMembershiplevel

	Membershipmodel.GetMembershipLevel(offset, limt, filter, &subscriptionlist, tenantid, Membership.DB)

	TotalMemebrshipCount, _ := Membershipmodel.GetMembershipLevel(0, 0, filter, &subscriptionlist, tenantid, Membership.DB)

	return subscriptionlist, TotalMemebrshipCount

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
	fmt.Println("")

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
