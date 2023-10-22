package rbacdb

var theTableNamePrefix string

////////////////////////////////////////////////////////////////////////////////

// entities ...
func entities() []any {
	// todo:list
	list := make([]any, 0)
	list = append(list, &PermissionEntity{})
	list = append(list, &RoleEntity{})
	list = append(list, &UserEntity{})
	list = append(list, &EmailAddressEntity{})
	list = append(list, &PhoneNumberEntity{})
	return list
}

////////////////////////////////////////////////////////////////////////////////

// TableName ...
func (PermissionEntity) TableName() string {
	return theTableNamePrefix + "permissions"
}

// TableName ...
func (RoleEntity) TableName() string {
	return theTableNamePrefix + "roles"
}

// TableName ...
func (UserEntity) TableName() string {
	return theTableNamePrefix + "users"
}

// TableName ...
func (EmailAddressEntity) TableName() string {
	return theTableNamePrefix + "email_addresses"
}

// TableName ...
func (PhoneNumberEntity) TableName() string {
	return theTableNamePrefix + "phone_numbers"
}
