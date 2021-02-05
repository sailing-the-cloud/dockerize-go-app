package helper

/**
 * Created by Manggala Pramuditya Wiryawan on 05/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

// ResponseDetail data structure
type ResponseDetail struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseDetailOutput function for contruct output response detail
func ResponseDetailOutput(message string, data interface{}) ResponseDetail {
	res := ResponseDetail{
		Message: message,
		Data:    data,
	}
	return res
}
