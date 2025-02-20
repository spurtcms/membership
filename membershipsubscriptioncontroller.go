package membership

import (
	"fmt"
	"time"

)


func (Membership *Membership)SubscriptionList(offset int, limt int, filter Filter, tenantid int)([]TblMembershipSubcriptionslist, int64, error){

	var subscriptionlist []TblMembershipSubcriptionslist

	_, err := Membershipmodel.ListSubscription(offset, limt, filter, &subscriptionlist, tenantid, Membership.DB)

	TotalSubscription, err := Membershipmodel.ListSubscription(0, 0, filter, &subscriptionlist, tenantid, Membership.DB)

	return subscriptionlist, TotalSubscription, err

}



func (membership *Membership)MembershipCreateSubscription(CreateSubscriptioninput TblMembershipSubcriptions,tenantid int,userid int)error {
	createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	fmt.Println("")

	var CreateSubscription TblMembershipSubcriptions

	CreateSubscription.MemberId=CreateSubscriptioninput.MemberId
	CreateSubscription.MembershipLevelId=CreateSubscriptioninput.MembershipLevelId
	CreateSubscription.GatewayEnvironment=CreateSubscriptioninput.GatewayEnvironment
	CreateSubscription.Gateway=CreateSubscriptioninput.Gateway
	CreateSubscription.SubscriptionTransactionId=CreateSubscriptioninput.SubscriptionTransactionId
	CreateSubscription.TenantId=tenantid
	CreateSubscription.CreatedOn=createdon
	CreateSubscription.CreatedBy=userid



	err:=Membershipmodel.CreateMembershipSubscription(CreateSubscription,membership.DB)
	if err != nil{
		return err
	}
	return nil

}


func (membership *Membership)SubscriptionEdit(SubscriptionId int)(TblMembershipSubcriptions,error){

	var EditSubscription TblMembershipSubcriptions

	err:=Membershipmodel.EditSubscription(&EditSubscription,SubscriptionId,membership.DB)

	return EditSubscription, err

}
