package tools

import (
	"crypto/rand"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
	"strconv"
)

var Snowflake = new(snowflake.Snowflake)

func init() {
	var err error
	Snowflake, err = snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		log.Errorf("init snowflake failed, err: %v", err)
		return
	}
	log.Debugln("init snowflake success...")
}

/*
	GenerateUniqueIdStr 生成唯一 ID 字符串
*/
func GenerateUniqueIdStr() string {
	return strconv.FormatInt(Snowflake.NextVal(), 10)
}

/*
	RangeRand 生成区间[min, max)的安全随机数
*/
func RangeRand(min, max int64) int64 {
	if min > max {
		return 0
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}
