package membership

import (
	"fmt"
	"time"
)

func (membership *Membership) OrderList(limit, offset int, filter Filter, tenantid int) (list []TblMembershipOrders, Count int64, err error) {

	if Autherr := AuthandPermission(membership); Autherr != nil {

		return []TblMembershipOrders{}, 0, Autherr
	}

	orderlist, _, _ := Membershipmodel.MemberShipOrderList(limit, offset, filter, tenantid, membership.DB)

	_, count, err := Membershipmodel.MemberShipOrderList(0, 0, filter, tenantid, membership.DB)
	if err != nil {

		return []TblMembershipOrders{}, 0, err
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

func (membership *Membership) EditMembershipOrder(id, tenantid int) (orderlist TblMembershipOrder, err error) {

	if Autherr := AuthandPermission(membership); Autherr != nil {

		return TblMembershipOrder{}, Autherr
	}

	orderrlist, err := Membershipmodel.Editorder(id, tenantid, membership.DB)
	if err != nil {
		fmt.Println(err)
	}

	return orderrlist, nil

}

func (membership *Membership) UpdateMembershipOrder(orders TblMembershipOrder, id, tenantid int) error {

	if Autherr := AuthandPermission(membership); Autherr != nil {

		return Autherr
	}

	updatedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

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
		ModifiedOn:                updatedon,
		ModifiedBy:                orders.ModifiedBy,
		TenantId:                  orders.TenantId,
	}

	err := Membershipmodel.UpdateOrder(Order, id, tenantid, membership.DB)

	if err != nil {

		return err
	}

	return nil

}

func (membership *Membership) DeleteMembershipOrder(id, userid, tenantid int) error {

	if Autherr := AuthandPermission(membership); Autherr != nil {

		return Autherr
	}

	deletedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	deletedby := userid

	err := Membershipmodel.DeleteOrder(id, tenantid, deletedby, deletedon, membership.DB)

	if err != nil {

		return err
	}

	return nil

}

func (membership *Membership) MultiSelectDeleteOrder(orderids []int, modifiedby int, tenantid int) error {

	if Autherr := AuthandPermission(membership); Autherr != nil {

		return Autherr
	}

	var Order TblMembershipOrder

	Order.DeletedBy = modifiedby

	Order.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Order.IsDeleted = 1

	Order.TenantId = tenantid

	err := Membershipmodel.MultiSelectOrderDelete(&Order, orderids, membership.DB)
	if err != nil {

		return err

	}
	return nil
}
