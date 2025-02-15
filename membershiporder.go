package membership

import (
	"time"
)

func (membership *Membership) OrderList(limit, offset int, filter Filter, tenantid int) (list []TblMembershipOrder, Count int64, err error) {

	if Autherr := AuthandPermission(membership); Autherr != nil {

		return []TblMembershipOrder{}, 0, Autherr
	}

	orderlist, count, err := Membershipmodel.MemberShipOrderList(limit, offset, filter, tenantid, membership.DB)
	if err != nil {

		return []TblMembershipOrder{}, 0, err
	}

	return orderlist, count, nil

}

func (membership *Membership) CreateOrder(orders TblMembershipOrder) error {

	if Autherr := AuthandPermission(membership); Autherr != nil {

		return Autherr
	}

	createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Order := TblMembershipOrder{
		UserId:                    orders.UserId,
		MembershiplevelId:         orders.MembershiplevelId,
		BillingName:               orders.BillingName,
		BillingStreet:             orders.BillingStreet,
		BillingStreet2:            orders.BillingStreet2,
		BillingCity:               orders.BillingCity,
		BillingState:              orders.BillingState,
		BillingPostalcode:         orders.BillingPostalcode,
		BillingCountry:            orders.BillingCountry,
		BillingPhone:              orders.BillingPhone,
		SubTotal:                  orders.SubTotal,
		Tax:                       orders.Tax,
		Total:                     orders.Total,
		PaymentType:               orders.PaymentType,
		Status:                    orders.Status,
		Gateway:                   orders.Gateway,
		GatewayEnvironment:        orders.GatewayEnvironment,
		PaymenttransactionId:      orders.PaymenttransactionId,
		SubscriptiontransactionId: orders.SubscriptiontransactionId,
		CreatedOn:                 createdon,
		CreatedBy:                 orders.CreatedBy,
		TenantId:                  orders.TenantId,
	}

	err := Membershipmodel.CreateMemberShipOrder(Order, membership.DB)

	if err != nil {

		return err
	}

	return nil

}
