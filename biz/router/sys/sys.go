// Code generated by hertz generator. DO NOT EDIT.

package Sys

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	sys "hertz_admin/biz/handler/sys"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		_api.POST("/menu_delete", append(_menudeleteMw(), sys.MenuDelete)...)
		_api.POST("/menu_list", append(_menulistMw(), sys.MenuList)...)
		_api.POST("/menu_save", append(_menusaveMw(), sys.MenuSave)...)
		_api.POST("/menu_update", append(_menuupdateMw(), sys.MenuUpdate)...)
		_api.POST("/query_role_menu", append(_queryrolemenuMw(), sys.QueryRoleMenu)...)
		_api.POST("/query_user_menu", append(_queryusermenuMw(), sys.QueryUserMenu)...)
		_api.POST("/role_delete", append(_roledeleteMw(), sys.RoleDelete)...)
		_api.POST("/role_list", append(_rolelistMw(), sys.RoleList)...)
		_api.POST("/role_save", append(_rolesaveMw(), sys.RoleSave)...)
		_api.POST("/role_update", append(_roleupdateMw(), sys.RoleUpdate)...)
		_api.POST("/update_role_menu", append(_updaterolemenuMw(), sys.UpdateRoleMenu)...)
		_api.POST("/user_delete", append(_userdeleteMw(), sys.UserDelete)...)
		_api.POST("/user_list", append(_userlistMw(), sys.UserList)...)
		_api.POST("/user_save", append(_usersaveMw(), sys.UserSave)...)
		_api.POST("/user_update", append(_userupdateMw(), sys.UserUpdate)...)
	}
}
