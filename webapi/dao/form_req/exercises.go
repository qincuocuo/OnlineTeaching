package form_req

type CreateExercisesReq struct {
	ContentId int             `json:"learning_content_id" validate:"required"` // 学习内容id
	Exercises []ExercisesItem `json:"exercises" validate:"required"`           // 练习题
}

type ExercisesItem struct {
	Question string   `json:"question" validate:"required"` // 题目
	Answer   string   `json:"answer" validate:"required"`   // 答案
	Type     int      `json:"type" validate:"required"`     //类型 1-选择 2-判断
	Options  []string `json:"option"`                       //选项
}

type GetExercisesReq struct {
	ContentId int `form:"content_id" json:"content_id" validate:"required"` //内容id
}

type ExercisesReq struct {
	ContentId int               `json:"content_id" validate:"required"` //内容id
	Answers   []ExercisesAnswer `json:"answers" validate:"required"`    //答案提交
}

type ExercisesAnswer struct {
	Id     int    `json:"id" validate:"required"`     // 习题id
	Answer string `json:"answer" validate:"required"` // 答案
}
