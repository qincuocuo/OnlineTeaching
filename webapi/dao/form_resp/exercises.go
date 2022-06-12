package form_resp

type GetExercisesResp struct {
	Count     int             `json:"count"`   // 习题数
	Exercises []ExercisesItem `json:"results"` // 练习题
}

type ExercisesItem struct {
	Id       int      `json:"id"`       //课后练习id
	Question string   `json:"question"` //题目
	Type     int      `json:"type"`     //类型 0-选择 1-判断
	Options  []string `json:"options"`  //选项
	Answer   string   `json:"answer"`
}

type ExercisesResp struct {
	Results []ExerciseResult `json:"results"`
}

type ExerciseResult struct {
	Id      int    `json:"id"`      //课后练习id
	Correct bool   `json:"correct"` //判断是否正确
	Answer  string `json:"answer"`  //正确答案
}
