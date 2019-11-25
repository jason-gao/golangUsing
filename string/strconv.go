package string

import "strconv"

func StrToInt(s string) (int, error) {
i, err := strconv.Atoi(s)

return i, err
}


//string转成int：
//int, err := strconv.Atoi(string)
//string转成int64：
//int64, err := strconv.ParseInt(string, 10, 64)
//int转成string：
//string := strconv.Itoa(int)
//int64转成string：
//string := strconv.FormatInt(int64,10)


