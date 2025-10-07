package routenames

import (
	"fmt"
)

const (
	Home                  = "home"
	Welcome               = "welcome"
	Dashboard             = "dashboard"
	AdminDashboard        = "admin_dashboard"
	AdminUserAdd          = "admin_user_add"
	AdminUserEdit         = "admin_user_edit"
	AdminUserDelete       = "admin_user_delete"
	About                 = "about"
	Contact               = "contact"
	ContactSubmit         = "contact.submit"
	Login                 = "login"
	LoginSubmit           = "login.submit"
	Register              = "register"
	RegisterSubmit        = "register.submit"
	ForgotPassword        = "forgot_password"
	ForgotPasswordSubmit  = "forgot_password.submit"
	Logout                = "logout"
	VerifyEmail           = "verify_email"
	ResetPassword         = "reset_password"
	ResetPasswordSubmit   = "reset_password.submit"
	Search                = "search"
	Task                  = "task"
	TaskSubmit            = "task.submit"
	Cache                 = "cache"
	CacheSubmit           = "cache.submit"
	Files                 = "files"
	FilesSubmit           = "files.submit"
	AdminTasks            = "admin:tasks"
	ProfileEdit           = "profile.edit"
	ProfileUpdate         = "profile.update"
	ProfileDestroy        = "profile.destroy"
	ProfileAppearance     = "profile.appearance"
	ProfilePassword       = "profile.password"
	ProfileUpdatePassword = "profile.update_password"
	Plans                 = "plans"
	PlansSubscribe        = "plans.subscribe"
	Products              = "products"
	ProductsPurchase      = "products.purchase"
	Premium               = "premium"
	Billing               = "billing"
	BillingCancel         = "billing.cancel"
	Forms                 = "forms"
	FormsCreate           = "forms.create"
	FormsStore            = "forms.store"
	FormsEdit             = "forms.edit"
	FormsUpdate           = "forms.update"
	FormsDelete           = "forms.delete"
	FormsShow             = "forms.show"
)

func AdminEntityList(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_list", entityTypeName)
}

func AdminEntityAdd(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_add", entityTypeName)
}

func AdminEntityEdit(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_edit", entityTypeName)
}

func AdminEntityDelete(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_delete", entityTypeName)
}

func AdminEntityAddSubmit(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_add.submit", entityTypeName)
}

func AdminEntityEditSubmit(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_edit.submit", entityTypeName)
}

func AdminEntityDeleteSubmit(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_delete.submit", entityTypeName)
}
