package exeWords

import (
	"gin/common/function"
	"gin/core"
	"gin/core/redis"
	"gin/model"
	"github.com/syyongx/php2go"
	"github.com/tidwall/gjson"
	"github.com/yanyiwu/gojieba"
	"unicode/utf8"
)

type ExeWords struct {
}

func (ExeWords ExeWords) Run(topic string, pushTime int64, data string) {

	content := gjson.Get(data, "content")
	rankKey := gjson.Get(data, "rankKey")

	if !content.Exists() || !rankKey.Exists() {
		return
	}

	contentString := content.String()
	rankKeyString := rankKey.String()

	contentString = function.GetZhAndLetter(contentString)

	//处理词云
	var s string
	var words []string
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()

	var BsKeywordsSlice []model.BsKeywords
	core.Db().Where("flag = ?", 1).Find(&BsKeywordsSlice)

	if len(BsKeywordsSlice) > 0 {
		for _, item := range BsKeywordsSlice {
			x.AddWord(item.Name)
		}
	}

	var BsXueqiuUserSlice []model.BsXueqiuUser

	core.Db().Where("flag = ?", 1).Find(&BsXueqiuUserSlice)

	//查询需要排除的关键词
	var BsKeywordsExclude []model.BsKeywordsExclude
	core.Db().Where("flag = ?", 1).Find(&BsKeywordsExclude)

	var needExcludeWord []string

	if len(BsKeywordsExclude) > 0 {
		for _, item := range BsKeywordsExclude {
			needExcludeWord = append(needExcludeWord, item.Name)
		}
	}

	if len(BsXueqiuUserSlice) > 0 {
		for _, item := range BsXueqiuUserSlice {
			needExcludeWord = append(needExcludeWord, item.Username)
		}
	}

	s = contentString
	words = x.Cut(s, use_hmm)
	for _, item := range words {
		len := utf8.RuneCountInString(item)
		if len == 1 || php2go.InArray(item, needExcludeWord) {
			continue
		}
		res := redis.ZRANK(rankKeyString, item)

		if res < 0 {
			redis.ZADD(rankKeyString, item, 0)
		} else {
			redis.ZINCRBY(rankKeyString, 1, item)
		}
	}
}
