package membership

import (
	"fmt"
	"time"
)

func (Membership *Membership) SubscriptionList(offset int, limt int, filter Filter, tenantid string) ([]TblMembershipSubcriptions, int64, error) {

	var subscriptionlist []TblMembershipSubcriptions

	_, err := Membershipmodel.ListSubscription(offset, limt, filter, &subscriptionlist, tenantid, Membership.DB)

	TotalSubscription, err := Membershipmodel.ListSubscription(0, 0, filter, &subscriptionlist, tenantid, Membership.DB)

	return subscriptionlist, TotalSubscription, err

}

func (membership *Membership) MembershipCreateSubscription(CreateSubscriptioninput TblMembershipSubcriptions, tenantid string, userid int) error {
	createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	fmt.Println("")

	var CreateSubscription TblMembershipSubcriptions

	CreateSubscription.MemberId = CreateSubscriptioninput.MemberId
	CreateSubscription.MembershipLevelId = CreateSubscriptioninput.MembershipLevelId
	CreateSubscription.GatewayEnvironment = CreateSubscriptioninput.GatewayEnvironment
	CreateSubscription.Gateway = CreateSubscriptioninput.Gateway
	CreateSubscription.SubscriptionTransactionId = CreateSubscriptioninput.SubscriptionTransactionId
	CreateSubscription.TenantId = tenantid
	CreateSubscription.CreatedOn = createdon
	CreateSubscription.CreatedBy = userid

	err := Membershipmodel.CreateMembershipSubscription(CreateSubscription, membership.DB)
	if err != nil {
		return err
	}
	return nil

}

func (membership *Membership) SubscriptionEdit(SubscriptionId int, tenantid string) (TblMembershipSubcriptions, error) {

	var EditSubscription TblMembershipSubcriptions

	err := Membershipmodel.EditSubscription(&EditSubscription, SubscriptionId, membership.DB, tenantid)

	return EditSubscription, err

}

func (membership *Membership) SubscriptionUpdate(UpdateSubscriptioninput TblMembershipSubcriptions, userid int, tenantid string) error {

	ModifiedOn, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	var UpdateSubscription TblMembershipSubcriptions
	UpdateSubscription.Id = UpdateSubscriptioninput.Id
	UpdateSubscription.MemberId = UpdateSubscriptioninput.MemberId
	UpdateSubscription.MembershipLevelId = UpdateSubscriptioninput.MembershipLevelId
	UpdateSubscription.GatewayEnvironment = UpdateSubscriptioninput.GatewayEnvironment
	UpdateSubscription.Gateway = UpdateSubscriptioninput.Gateway
	UpdateSubscription.SubscriptionTransactionId = UpdateSubscriptioninput.SubscriptionTransactionId
	UpdateSubscription.TenantId = tenantid
	UpdateSubscription.ModifiedOn = ModifiedOn
	UpdateSubscription.ModifiedBy = userid

	err := Membershipmodel.UpdateSubscription(UpdateSubscription, membership.DB)

	return err
}

func (membership *Membership) SubscriptionsDelete(id, userid int, tenantid string) error {

	deletedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	deletedby := userid

	err := Membershipmodel.DeleteSubscriptions(id, tenantid, deletedby, deletedon, membership.DB)

	return err

}

func (Membership *Membership) ChangesSubscriptionIsactive(subscriptionid int, status int, modifiedby int, tenantid string) (bool, error) {

	fmt.Println("reachh:2:")

	var subscriptionstatus TblMembershipSubcriptions
	subscriptionstatus.ModifiedBy = modifiedby
	subscriptionstatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Membershipmodel.SubscriptionChangeStatus(subscriptionstatus, subscriptionid, status, Membership.DB, tenantid)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (Membership *Membership) DeleteMultiSelectSubscription(SubscriptionIds []int, userid int) error {

	deletedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Membershipmodel.MultiselectDeleteSubscription(SubscriptionIds, Membership.DB, deletedon, userid)

	return err
}
