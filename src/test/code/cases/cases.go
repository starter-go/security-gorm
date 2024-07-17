package cases

// 定义各种 测试用例 的名称
const (
	HelpByMail     = "help-by-mail"
	LoginByMail    = "login-by-mail"
	SendcodeByMail = "send-code-by-mail"

	HelpBySMS     = "help-by-sms"
	LoginBySMS    = "login-by-sms"
	SendcodeBySMS = "send-code-by-sms"

	LoginWithPassword = "login-with-password"

	// for user
	GetUserList = "get-user-list"
	InsertUser  = "insert-user"
	DoUserCRUD  = "do-user-crud"

	// for roles
	GetRoleList = "get-role-list"
	InsertRole  = "insert-role"
	DoRoleCRUD  = "do-role-crud"

	// for permission
	GetPermissionList = "get-permission-list"
	InsertPermission  = "insert-permission"
	DoPermissionCRUD  = "do-permission-crud"
)
