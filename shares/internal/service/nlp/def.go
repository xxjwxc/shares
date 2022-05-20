package nlp

type TencentVoiceResp struct {
	Response TencentVoiceResponse `json:"Response"`
}

type TencentVoiceResponse struct {
	Data      VoiceData `json:"Data"`
	RequestId string    `json:"RequestId"`
}

type VoiceData struct {
	TaskId int64  `json:"TaskId"`
	Status int64  `json:"Status"` // 任务状态 0：任务等待，1：任务执行中，2：任务成功，3：任务失败。
	Result string `json:"Result"` //  识别结果
}
