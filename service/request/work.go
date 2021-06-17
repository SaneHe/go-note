package request

import (
	"strconv"
	"time"
	"work-wechat/config"
)

type LocalDateTimeString time.Time

func (l LocalDateTimeString) String() string {
	return time.Time(l).Format(config.App.Logger.LogTimestampFormat)
}

func (l *LocalDateTimeString) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		return nil
	}
	now, err := time.Parse(strconv.Quote(config.App.Logger.LogTimestampFormat), string(data))

	if err != nil {
		return err
	}
	*l = LocalDateTimeString(now)

	return nil
}

func (l LocalDateTimeString) MarshalJSON() ([]byte, error) {
	return ([]byte)(strconv.FormatInt(time.Time(l).Unix(), 10)), nil
}

/**
 * @Description: 日报、月报打卡参数
 */
type PunchBody struct {
	StartTime  LocalDateTimeString `json:"starttime" binding:"required" time_format:"2006-01-02 15:04:05" label:"开始时间"`
	EndTime    LocalDateTimeString `json:"endtime" binding:"required,gtfield=StartDate" time_format:"2006-01-02 15:04:05" label:"结束时间"`
	UserIdList []string            `json:"useridlist" binding:"required" label:"人员列表"`
}

/**
 * @Description: 打卡记录参数
 */
type PunchRecord struct {
	PunchBody
	OpenCheckinDataType uint8 `json:"opencheckindatatype" binding:"required,oneof=1 2 3" label:"打卡类型"`
}
