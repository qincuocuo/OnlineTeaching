package support

const (
	UserIsExist            string = "账户已经存在"
	UserNotExist           string = "账户不存在"
	UpdatePasswordFailed   string = "修改账户密码失败"
	PasswordFailed         string = "账户或者密码错误"
	PasswordNotConfirm     string = "输入的密码不一致"
	PasswordWrong          string = "密码错误，请重新输入"
	PasswordStrengthFailed string = "密码强度过低"
	CreateUserFailed       string = "创建用户失败"
	UserLockFailed         string = "用户被锁定"
	UserNoPermission       string = "用户没有操作权限"

	CourseIsExists     string = "课程已存在"
	CreateCourseFailed string = "创建课程失败"
	CourseNotExists    string = "课程不存在"
	UpdateCourseFailed string = "修改课程信息失败"
	DeleteCourseFailed string = "删除课程失败"
	EnterCourseFailed  string = "加入课程失败"
	GetCourseFailed    string = "获取课程信息失败"
	CourseNotFound     string = "本班级还没有创建课程"

	LearningContentNotFound      string = "章节不存在"
	GetLearningContentListFailed string = "获取学习章节信息失败"
	GetCourseInfoFailed          string = "根据学习章节获取课程信息失败"
	UploadLearningContentFailed  string = "上传学习章节失败"
	CreateLearningContentFailed  string = "创建学习章节失败"
	UpdateContentFailed          string = "更新学习结果失败"

	OpenFileFailed string = "打开文件失败"

	RegisterExists       string = "签到任务已存在"
	RegisterNotFount     string = "签到任务不存在"
	RegisterFinished     string = "用户已完成签到，请勿重复提交"
	CreateRegisterFailed string = "创建签到任务失败"
	JoinInRegisterFailed string = "参与签到失败"
	RegisterExpired      string = "签到任务已过期"

	CreateTalkFailed string = "创建讨论失败"

	CreateExercisesFailed string = "创建课后练习失败"
	ExercisesNotExists    string = "课后练习任务不存在"
)
