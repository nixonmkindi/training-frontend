package auth

import (
	"training-frontend/server/systems/auth/controllers"

	"github.com/labstack/echo/v4"
)

// WebRouters initialises web routes
func WebRouters(app *echo.Echo) {
	auth := app.Group("/auth")
	{
		auth.GET("/register", controllers.Register.ShowRegistrationForm)
		auth.POST("/register", controllers.Register.CreateAccount)
		auth.GET("/login", controllers.Login.ShowLoginForm)
		auth.POST("/login", controllers.Login.Login)
		auth.GET("/logout", controllers.Login.Logout)
		auth.GET("/forgot-password", controllers.ForgotPassword.ShowForgotPasswordForm)
		auth.POST("/reset-password", controllers.ForgotPassword.ResetPassword)
		auth.POST("/password-reset", controllers.ForgotPassword.PasswordReset)
		auth.POST("/user-reset-password", controllers.ForgotPassword.UserResetPassword)
		auth.GET("/reset-student-password", controllers.ForgotPassword.ResetStudentPassword)
		auth.POST("/reset-password-for-student", controllers.ForgotPassword.ResetPasswordForStudent)
		auth.GET("/reset-email-for-student", controllers.ForgotPassword.ResetStudentEmail)
		auth.POST("/update-student-email", controllers.ForgotPassword.UpdateEmailForStudent)
		auth.POST("/reset-staff-password", controllers.ForgotPassword.ResetStaffPassword)
	}

	// users := app.Group("/auth/users")
	// {
	// 	users.GET("/list", controllers.User.List)
	// 	users.GET("/create", controllers.User.Create)
	// }

	permissions := app.Group("/auth/permissions")
	{
		permissions.GET("/list", controllers.Permission.List)
		permissions.POST("/show", controllers.Permission.Show)
		permissions.POST("/edit", controllers.Permission.Edit)
		permissions.POST("/update", controllers.Permission.Update)
		permissions.GET("/generate-permission", controllers.Permission.GeneratePermission)
		permissions.POST("/soft-delete", controllers.Permission.Delete)
		permissions.POST("/hard-delete", controllers.Permission.HardDelete)
		permissions.GET("/force-delete", controllers.Permission.ForceDelete)
	}

	roles := app.Group("/auth/roles")
	{
		roles.GET("/list", controllers.Role.List)
		roles.GET("/create", controllers.Role.Create)
		roles.POST("/store", controllers.Role.Store)
		roles.POST("/show", controllers.Role.Show)
		roles.POST("/edit", controllers.Role.Edit)
		roles.POST("/update", controllers.Role.Update)
		roles.POST("/delete", controllers.Role.Delete)
	}

	subsystems := app.Group("/auth/subsystems")
	{
		subsystems.GET("/list", controllers.SubSystems.List)
		subsystems.GET("/create", controllers.SubSystems.Create)
		subsystems.POST("/store", controllers.SubSystems.Store)
		subsystems.POST("/show", controllers.SubSystems.Show)
		subsystems.POST("/edit", controllers.SubSystems.Edit)
		subsystems.POST("/update", controllers.SubSystems.Update)
		subsystems.POST("/delete", controllers.SubSystems.Delete)
	}

	userCategories := app.Group("/auth/user-categories")
	{
		userCategories.GET("/list", controllers.UserCategory.List)
		userCategories.GET("/create", controllers.UserCategory.Create)
		userCategories.POST("/store", controllers.UserCategory.Store)
		userCategories.POST("/show", controllers.UserCategory.Show)
		userCategories.POST("/edit", controllers.UserCategory.Edit)
		userCategories.POST("/update", controllers.UserCategory.Update)
		userCategories.POST("/delete", controllers.UserCategory.Delete)
	}

	userHasPermissions := app.Group("/auth/user-permissions")
	{
		userHasPermissions.POST("/list", controllers.UserHasPermission.List)
		userHasPermissions.POST("/create", controllers.UserHasPermission.Create)
		userHasPermissions.POST("/store", controllers.UserHasPermission.Store)
	}

	userHasRoles := app.Group("/auth/user-roles")
	{
		userHasRoles.POST("/list", controllers.UserHasRole.List)
		userHasRoles.POST("/create", controllers.UserHasRole.Create)
		userHasRoles.POST("/store", controllers.UserHasRole.Store)
		userHasRoles.POST("/role-user-permission", controllers.UserHasRole.ListRoleUserPermission)
	}

	roleHasPermissions := app.Group("/auth/role-permissions")
	{
		roleHasPermissions.POST("/list", controllers.RoleHasPermission.List)
		roleHasPermissions.POST("/create", controllers.RoleHasPermission.Create)
		roleHasPermissions.POST("/store", controllers.RoleHasPermission.Store)
	}
}
