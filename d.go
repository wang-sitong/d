package d

import "time"

func Date(params ...interface{}) string {
	// 定义默认值
	defaultValues := map[int]interface{}{
		0: "Y-m-d H:i:s", // format
		1: 0,             // date
		2: "UTC+8",       // date

	}
	defaultCount := len(defaultValues)
	count := len(params)
	value := defaultCount - count
	if value > 0 {
		for i := defaultCount - value; i < defaultCount; i++ {
			params = append(params, defaultValues[i])
		}
	}

	format := params[0]
	day := params[1].(int)
	result := "format error"
	currentDateTime := time.Unix(time.Now().Unix()+int64(86400*day), 0)
	switch format {
	case "Y-m-d H:i:s":
		result = currentDateTime.Format("2006-01-02 15:04:05")
		break
	case "Y-m-d":
		result = currentDateTime.Format("2006-01-02")
		break
	case "H:i:s":
		result = currentDateTime.Format("15:04:05")
		break
	case "Y":
		result = currentDateTime.Format("2006")
		break
	case "m":
		result = currentDateTime.Format("01")
		break
	case "d":
		result = currentDateTime.Format("02")
		break
	case "H":
		result = currentDateTime.Format("15")
		break
	case "i":
		result = currentDateTime.Format("04")
		break
	case "s":
		result = currentDateTime.Format("05")
		break
	default:
		result = "format error"
	}
	return result
}
