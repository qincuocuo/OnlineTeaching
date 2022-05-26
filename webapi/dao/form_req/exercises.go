package form_req

type CreateExercisesReq struct {
	ContentId int             `json:"learning_content_id"` // 学习内容id
	Exercises []ExercisesItem `json:"exercises"`           // 练习题
}

type ExercisesItem struct {
	Question string   `json:"question"` // 题目
	Answer   string   `json:"answer"`   // 答案
	Type     int      `json:"type"`     //类型 1-选择 2-判断
	Options  []string `json:"options"`  //选项
}

type GetExercisesReq struct {
	ContentId int `json:"content_id"` //内容id
}

type ExercisesReq struct {
	ContentId int               `json:"content_id"` //内容id
	Answers   []ExercisesAnswer `json:"answers"`    //答案提交
}

type ExercisesAnswer struct {
	Id     int    `json:"id"`     // 习题id
	Answer string `json:"answer"` // 答案
}
