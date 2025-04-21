package words

import (
	"context"
	"proxima-vocabulary/app/word/internal/dao"
	"proxima-vocabulary/app/word/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

// 修改 Create 方法签名
func Create(ctx context.Context, word *entity.Words) (id uint, err error) {
    // 设置创建和更新时间
    word.CreatedAt = gtime.Now()
    word.UpdatedAt = gtime.Now()
    
    // 使用 DAO 层插入数据
    result, err := dao.Words.Ctx(ctx).Insert(word)
    if err != nil {
        return 0, err
    }
    
    lastId, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
    
    return uint(lastId), nil
}

func Get(ctx context.Context) (word *entity.Words, err error) {
	return &entity.Words{
		Id:                 1,
		Uid:                1,
		Word:               "hello",
		Definition:         "used as a greeting when you meet somebody.",
		ExampleSentence:    "Hello, I am oldme!",
		ChineseTranslation: "你好",
		Pronunciation:      "həˈləʊ",
		CreatedAt:          gtime.New("2024-12-05 22:00:00"),
		UpdatedAt:          gtime.New("2024-12-05 22:00:00"),
	}, nil
}