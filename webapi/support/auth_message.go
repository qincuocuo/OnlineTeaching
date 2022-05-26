package support

const (
	UserIsExist            string = "账户已经存在"
	UserNotExist           string = "账户不存在"
	UpdatePasswordFailed   string = "修改账户密码失败"
	PasswordFailed         string = "账户或者密码错误"
	PasswordNotConfirm     string = "输入的密码不一致"
	PasswordStrengthFailed string = "密码强度过低"
	CreateUserFailed       string = "创建用户失败"
	UserLockFailed         string = "用户被锁定"
	UserNoPermission       string = "用户没有操作权限"

	CreateCourseFailed string = "创建课程失败"
	CourseNotExists    string = "课程不存在"
	UpdateCourseFailed string = "修改课程信息失败"
	DeleteCourseFailed string = "删除课程失败"
	EnterCourseFailed  string = "加入课程失败"

	GetLearningContentListFailed string = "获取学习内容信息失败"
	GetCourseInfoFailed          string = "根据学习内容获取课程信息失败"

	RegisterNotFount     string = "签到任务不存在"
	CreateRegisterFailed string = "创建签到任务失败"
	JoinInRegisterFailed string = "参与签到失败"
	RegisterExpired      string = "签到任务已过期"

	CreateTalkFailed string = "创建讨论失败"
)
