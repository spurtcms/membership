package membership

import (
	"time"

	"gorm.io/gorm"
)

type TblMembershipOrder struct {
	Id                        int       `gorm:"primaryKey;auto_increment;type:serial"`
	UserId                    int       `gorm:"type:integer"`
	MembershiplevelId         int       `gorm:"type:integer"`
	BillingName               string    `gorm:"type:character varying"`
	BillingStreet             string    `gorm:"type:character varying"`
	BillingStreet2            string    `gorm:"type:character varying"`
	BillingCity               string    `gorm:"type:character varying"`
	BillingState              string    `gorm:"type:character varying"`
	BillingPostalcode         string    `gorm:"type:character varying"`
	BillingCountry            string    `gorm:"type:character varying"`
	BillingPhone              string    `gorm:"type:character varying"`
	SubTotal                  int       `gorm:"type:integer"`
	Tax                       int       `gorm:"type:integer"`
	Total                     int       `gorm:"type:integer"`
	PaymentType               string    `gorm:"type:character varying"`
	Status                    string    `gorm:"type:character varying"`
	Gateway                   string    `gorm:"type:character varying"`
	GatewayEnvironment        string    `gorm:"type:character varying"`
	PaymenttransactionId      int       `gorm:"type:integer"`
	SubscriptiontransactionId int       `gorm:"type:integer"`
	CreatedOn                 time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy                 int       `gorm:"type:integer"`
	IsDeleted                 int       `gorm:"type:integer"`
	DeletedOn                 time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy                 int       `gorm:"DEFAULT:NULL"`
	ModifiedOn                time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy                int       `gorm:"DEFAULT:NULL"`
	TenantId                  int       `gorm:"type:integer"`
	DateString                string    `gorm:"-"`
	SubscriptionName          string    `gorm:"column:subscription_name"`
	FirstName                 string    `gorm:"column:first_name"`
}

func (Membershipmodel MembershipModel) MemberShipOrderList(limit, offset int, filter Filter, tenantid int, DB *gorm.DB) (orderlist []TblMembershipOrder, count int64, err error) {

	var orderlistcount int64

	query := DB.Debug().Table("tbl_membership_orders").
		Select("tbl_membership_orders.*, tbl_mstr_membershiplevels.subscription_name,tbl_membership_members.first_name").
		Joins("inner join tbl_mstr_membershiplevels on tbl_membership_orders.membershiplevel_id=tbl_mstr_membershiplevels.id").Joins("inner join tbl_membership_members on tbl_membership_orders.user_id=tbl_membership_members.id").Where("tbl_membership_orders.is_deleted = 0 and  tbl_membership_orders.tenant_id = ?", tenantid).Order("tbl_membership_orders.id desc")

	if limit != 0 {

		query.Limit(limit).Offset(offset).Find(&orderlist)

		return orderlist, orderlistcount, nil

	}

	query.Find(&orderlist).Count(&orderlistcount)
	if query.Error != nil {

		return []TblMembershipOrder{}, 0, query.Error
	}

	return orderlist, orderlistcount, nil

}

func (Membershipmodel MembershipModel) CreateMemberShipOrder(order TblMembershipOrder, DB *gorm.DB) error {

	if err := DB.Table("tbl_membership_orders").Create(&order).Error; err != nil {

		return err
	}

	return nil

}

func (Membershipmodel MembershipModel) Editorder(id, tenantid int, DB *gorm.DB) (orderlist TblMembershipOrder, err error) {

	if err := DB.Table("tbl_membership_orders").Where("id=? and tenant_id=? and is_deleted=0", id, tenantid).First(&orderlist).Error; err != nil {

		return TblMembershipOrder{}, err
	}

	return orderlist, nil

}

func (Membershipmodel MembershipModel) UpdateOrder(order TblMembershipOrder, id, tenantid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_membership_orders").Where("id=? and tenant_id=? and is_deleted=0", id, tenantid).UpdateColumns(map[string]interface{}{"user_id": order.UserId, "membershiplevel_id": order.MembershiplevelId, "billing_name": order.BillingName, "billing_street": order.BillingStreet, "billing_street2": order.BillingStreet2, "billing_city": order.BillingCity, "billing_state": order.BillingState, "billing_postalcode": order.BillingPostalcode, "billing_country": order.BillingCountry, "billing_phone": order.BillingPhone, "sub_total": order.SubTotal, "tax": order.Tax, "total": order.Total, "payment_type": order.PaymentType, "status": order.Status, "gateway": order.Gateway, "gateway_environment": order.GatewayEnvironment, "paymenttransaction_id": order.PaymenttransactionId, "subscriptiontransaction_id": order.SubscriptiontransactionId, "modified_on": order.ModifiedOn, "modified_by": order.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil

}

func (Membershipmodel MembershipModel) DeleteOrder(id, tenantid, deletedby int, deletedon time.Time, DB *gorm.DB) error {

	if err := DB.Table("tbl_membership_orders").Where("id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": deletedby, "deleted_on": deletedon}).Error; err != nil {

		return err
	}

	return nil
}

func (Membershipmodel MembershipModel) MultiSelectOrderDelete(order *TblMembershipOrder, id []int, DB *gorm.DB) error {

	if err := DB.Table("tbl_membership_orders").Where("id in (?) and tenant_id=?", id, order.TenantId).UpdateColumns(map[string]interface{}{"is_deleted": order.IsDeleted, "deleted_on": order.DeletedOn, "deleted_by": order.DeletedBy}).Error; err != nil {

		return err
	}

	return nil

}
