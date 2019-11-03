package acl

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"myserver/pkg/config"
)

var acl *casbin.Enforcer

func PolicyLoad() *casbin.Enforcer {
	if acl == nil {
		// Initialize a Xorm adapter and use it in a Casbin enforcer:
		// The adapter will use the MySQL database named "casbin".
		// If it doesn't exist, the adapter will create it automatically.
		dbHost := config.GetConfig().DB.Host
		dbUser := config.GetConfig().DB.UserName
		dbPasswd := config.GetConfig().DB.PassWord

		var err error
		connectString := fmt.Sprintf("%s:%s@tcp(%s)/", dbUser, dbPasswd, dbHost)
		db, err := xormadapter.NewAdapter("mysql", connectString) // Your driver and data source.

		if err !=nil {
			panic(err)
		}
		// Or you can use an existing DB "abc" like this:
		// The adapter will use the table named "casbin_rule".
		// If it doesn't exist, the adapter will create it automatically.
		// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

		acl, err = casbin.NewEnforcer("./rbac_model.conf", db)
		if err != nil {
			panic(err)
		}

		// Load the policy from DB.
		acl.LoadPolicy()

		// Check the permission.
		//acl.Enforce("alice", "data1", "read")

		// Modify the policy.
		acl.AddPolicy("alice", "data1", "read")
		// e.RemovePolicy(...)

		// Save the policy back to DB.
		acl.SavePolicy()
	}
	return acl
}
