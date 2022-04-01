package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	bazi "github.com/damoncoo/BaziGo"
)

// GetBazi GetBazi
func GetBazi(year, month, day, hour, minutes, seconds, sex int) interface{} {

	pbazi := bazi.GetBazi(year, month, day, hour, minutes, seconds, sex)

	data := map[string]interface{}{}
	_ = json.Unmarshal([]byte(pbazi.Data()), &data)

	baziMap := map[string]interface{}{
		"solar": fmt.Sprintf("%d年%02d月%02d日 %02d:%02d:%02d",
			pbazi.PSolarDate.NYear, pbazi.PSolarDate.NMonth, pbazi.PSolarDate.NDay, pbazi.PSolarDate.NHour, pbazi.PSolarDate.NMinute, pbazi.PSolarDate.NSecond),
		"lunar": strings.Replace(pbazi.PLunarDate.String(), "农历: ", "", 1),
		"sizhu": fmt.Sprintf("%s %s %s %s",
			pbazi.PSiZhu.PYearZhu.GanZhi(), pbazi.PSiZhu.PMonthZhu.GanZhi(), pbazi.PSiZhu.PDayZhu.GanZhi(), pbazi.PSiZhu.PHourZhu.GanZhi()),
		"bazi": pbazi.PBaziDate.String(),
		"qiYun": fmt.Sprintf("%d年%02d月%02d日 %02d:%02d:%02d",
			pbazi.PQiYunDate.NYear, pbazi.PQiYunDate.NMonth, pbazi.PQiYunDate.NDay, pbazi.PQiYunDate.NHour, pbazi.PQiYunDate.NMinute, pbazi.PQiYunDate.NSecond),
		"daYun": strings.Replace(pbazi.PDaYun.String(), "大运:\n", "", 1),
		"sex":   pbazi.NSex,
		"data":  data,
	}
	return baziMap
}
