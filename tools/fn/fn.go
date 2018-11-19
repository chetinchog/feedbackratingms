package fn

import (
	"fmt"

	"github.com/chetinchog/feedbackratingms/rates"
	"github.com/chetinchog/feedbackratingms/rules"
)

func CalculateRates(feedRate *rates.Rate) (float64, int) {
	feedAmount := (feedRate.Ra1 +
		feedRate.Ra2 +
		feedRate.Ra3 +
		feedRate.Ra4 +
		feedRate.Ra5)

	rating := (float64(feedRate.Ra1)*1 +
		float64(feedRate.Ra2)*2 +
		float64(feedRate.Ra3)*3 +
		float64(feedRate.Ra4)*4 +
		float64(feedRate.Ra5)*5) / float64(feedAmount)
	return rating, feedAmount
}

func AddRate(rate *rates.Rate, rateNF int) *rates.Rate {
	switch rateNF {
	case 1:
		rate.Ra1++
		break
	case 2:
		rate.Ra2++
		break
	case 3:
		rate.Ra3++
		break
	case 4:
		rate.Ra4++
		break
	case 5:
		rate.Ra5++
		break
	default:
	}
	return rate
}

func Classify(rate *rates.Rate) *rates.Rate {
	prom, amount := CalculateRates(rate)
	if amount == 0 {
	}

	dao, err := rules.GetDao()
	if err != nil {
		fmt.Println(" ---------------------------- ")
		fmt.Println(err)
		fmt.Println(" ---------------------------- ")
		return rate
	}

	articleId := rate.ArticleId

	rule, err := dao.FindByArticleID(articleId)
	if err != nil {
	}

	if rule == nil {
		return rate
	}

	if prom <= float64(rule.LowRate) {
		rate.BadRate = true
	} else {
		rate.BadRate = false
	}
	if prom >= float64(rule.HighRate) {
		rate.GoodRate = true
	} else {
		rate.GoodRate = false
	}

	return rate
}
