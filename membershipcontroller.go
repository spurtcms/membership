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

func (Membership *Membership) MembershipLevelsList(offset int, limt int, filter Filter, tenantid int) ([]TblMstrMembershiplevel, int64, error) {

	
	var subscriptionlist []TblMstrMembershiplevel

	TotalMemebrshipCount, err := Membershipmodel.GetMembershipLevel(0, 0, filter, &subscriptionlist, tenantid, Membership.DB)


     Membershipmodel.GetMembershipLevel(offset, limt, filter, &subscriptionlist, tenantid, Membership.DB)


	return subscriptionlist, TotalMemebrshipCount, err

}

func (Membership *Membership) GetdefaultMembershiplevelTemplate() []TblMstrMembershiplevel {

	var DefaultMembershipLevelList []TblMstrMembershiplevel

	Membershipmodel.GetdefaultTemplate(&DefaultMembershipLevelList, Membership.DB)

	return DefaultMembershipLevelList

}
func (Membership *Membership) MembershiplevelDetails(membershiplevelId int) ([]TblMstrMembershiplevel, error) {

	var SelectedMembershipData []TblMstrMembershiplevel

	err := Membershipmodel.GetMembershiplevelDetails(&SelectedMembershipData, membershiplevelId, Membership.DB)

	if err != nil {
		return []TblMstrMembershiplevel{}, err
	}

	return SelectedMembershipData, nil

}

func (Membership *Membership) MembershiplevelEdit(membershipid int) (TblMstrMembershiplevel, error) {

	var Editmembership TblMstrMembershiplevel

	err := Membershipmodel.Editmembershiplevel(&Editmembership, membershipid, Membership.DB)

	if err != nil {
		return TblMstrMembershiplevel{}, err
	}

	return Editmembership, nil

}

func (Membership *Membership) MembershipLevelsCreate(sd TblMstrMembershiplevel, tenantid int) error {

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

	err := Membershipmodel.CreateSubscriptionLevel(subscriptiondata, Membership.DB)

	if err != nil {

		return err
	}

	return nil

}

func (Membership *Membership) UpdateSubscription(subscriptionNewdata TblMstrMembershiplevel, tenantid int) error {
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

	err := Membershipmodel.Subscriptionupdate(Updatesubscription, tenantid, Membership.DB)

	if err != nil {

		return err
	}

	return nil

}

func (Membership *Membership) SubscriptionDelete(tenantid int, id int, userid int) error {

	var subscriptionlist TblMstrMembershiplevel

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	subscriptionlist.DeletedOn = time
	subscriptionlist.DeletedBy = userid
	subscriptionlist.TenantId = tenantid

	err := Membershipmodel.DeleteSubscription(&subscriptionlist, id, Membership.DB)
	if err != nil {

		return err
	}

	return nil
}

func (Membership *Membership) DeleteMultiselectMembershipLevel(MembershipLevelids []int, userid int) error {

	deletedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Membershipmodel.MultiselectDeleteMembershipLevel(MembershipLevelids, Membership.DB, deletedon, userid)

	if err != nil {

		return err
	}

	return nil
}

func (Membership *Membership) ChangesMembershipLevelIsactive(membershiplevelid int, status int, modifiedby int, tenantid int) (bool, error) {
	var membershiplevelstatus TblMstrMembershiplevel
	membershiplevelstatus.ModifiedBy = modifiedby
	membershiplevelstatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Membershipmodel.MembershiplevelChangeStatus(membershiplevelstatus, membershiplevelid, status, Membership.DB, tenantid)
	if err != nil {
		return false, err
	}

	return true, nil
}
