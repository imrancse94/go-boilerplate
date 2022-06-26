package constant

func Format(key string) string {
	var status = make(map[string]string)

	status["Y-m-d"] = "2006-01-02"
	status["Y/m/d"] = "2006/01/02"
	status["m-Y-d"] = "01-2006-02"

	return status[key]
}
