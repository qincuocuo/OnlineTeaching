package form_resp

type ExercisesResp struct {
	Results []ExerciseResult `json:"results"`
}

type ExerciseResult struct {
	Correct bool `json:"correct"` //判断是否正确
	Answer string `json:"answer"` //正确答案
}
