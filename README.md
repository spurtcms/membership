# Membership Package

The Membership Package for SpurtCMS that enables user-based access control and monetization features. It provides a powerful set of tools to manage user subscriptions, restrict content based on membership levels, and offer tiered access to exclusive content.

## Features

- Membership Levels: Create and manage multiple membership plans (Free, Basic, Premium, etc.)
- Access Control: Restrict pages, posts, or media based on user membership
- Payment Integration: Supports integration with payment gateways for paid memberships
- Subscription Management: Allow users to upgrade, downgrade, or cancel their subscriptions


# Installation

``` bash
go get github.com/spurtcms/membership
```


# Usage Example


``` bash
import (
	"github.com/spurtcms/auth"
	"github.com/spurtcms/membership"
)

func main() {

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "SecretKey@123",
		DB:         &gorm.DB{},
		RoleId:     1,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Membership", auth.CRUD)

	MembershipConfig = memship.MembershipSetup(memship.Config{
		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: false,
		Auth:             NewAuth,
	})

	//Membership
	if permisison {

		//list Members
		MembershipMemberList, TolatmemberCount := MembershipConfig.MembershipListMembers(0, 10, memship.Filter(filter), 1, 1)

		//list MembershipLevel Group
		subscriptiongroup, TotalGroupcount := MembershipConfig.MembershipGroupList(offset, limt, memship.Filter(filter), TenantId, 0)
		fmt.Println("")

		//list MembershipLevel
		MembershiplevelLists, _, _ := MembershipConfig.MembershipLevelsList(0, 10, memship.Filter{}, 1)

		if err != nil {
			log.Fatal("membership level list error", err)
			c.AbortWithError(500, err)
		}

		//list Subscription
		SubscriptionList, SubscriptioCount, err := MembershipConfig.SubscriptionList(offset, limt, memship.Filter(filter), TenantId)

		if err != nil {
			log.Fatal("Subscription List Error :", err)
			c.AbortWithError(500, err)
		}

		//list Orders
		orderlist, count, err := MembershipConfig.OrderList(limt, offset, filter, TenantId)
		if err != nil {
			fmt.Println(err)
		}

	}
}

```


# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/membership/issues]. 
