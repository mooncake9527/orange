package stringUtil

func Limit(str string, limit int) string {
	if len(str) > limit {
		return str[:limit]
	}
	return str
}
