package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//Struct for JSON

type Curs struct {
	Date         time.Time `json:"Date"`
	PreviousDate time.Time `json:"PreviousDate"`
	PreviousURL  string    `json:"PreviousURL"`
	Timestamp    time.Time `json:"Timestamp"`
	Valute       struct {
		Aud struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"AUD"`
		Azn struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"AZN"`
		Gbp struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"GBP"`
		Amd struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"AMD"`
		Byn struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"BYN"`
		Bgn struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"BGN"`
		Brl struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"BRL"`
		Huf struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"HUF"`
		Hkd struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"HKD"`
		Dkk struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"DKK"`
		Usd struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"USD"`
		Eur struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"EUR"`
		Inr struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"INR"`
		Kzt struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"KZT"`
		Cad struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"CAD"`
		Kgs struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"KGS"`
		Cny struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"CNY"`
		Mdl struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"MDL"`
		Nok struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"NOK"`
		Pln struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"PLN"`
		Ron struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"RON"`
		Xdr struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"XDR"`
		Sgd struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"SGD"`
		Tjs struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"TJS"`
		Try struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"TRY"`
		Tmt struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"TMT"`
		Uzs struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"UZS"`
		Uah struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"UAH"`
		Czk struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"CZK"`
		Sek struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"SEK"`
		Chf struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"CHF"`
		Zar struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"ZAR"`
		Krw struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"KRW"`
		Jpy struct {
			Name  string  `json:"Name"`
			Value float64 `json:"Value"`
		} `json:"JPY"`
	} `json:"Valute"`
}

func main() {

	resp, _ := http.Get("https://www.cbr-xml-daily.ru/daily_json.js")
	body, _ := ioutil.ReadAll(resp.Body) // response body is []byte

	var value Curs
	json.Unmarshal(body, &value)

	var sum float64

	var valName string

	var min float64

	sum, valName = MaxValueCzk(sum, value)
	sum, valName = MaxValueAud(sum, value)
	sum, valName = MaxValueAzn(sum, value)
	sum, valName = MaxValueAmd(sum, value)
	sum, valName = MaxValueJpy(sum, value)
	sum, valName = MaxValueCad(sum, value)
	sum, valName = MaxValueUah(sum, value)
	sum, valName = MaxValueZar(sum, value)
	sum, valName = MaxValueKzt(sum, value)
	sum, valName = MaxValueKgs(sum, value)
	sum, valName = MaxValueGbp(sum, value)
	sum, valName = MaxValueBgn(sum, value)
	sum, valName = MaxValueByn(sum, value)
	sum, valName = MaxValueBrl(sum, value)
	sum, valName = MaxValueXdr(sum, value)
	sum, valName = MaxValueKrw(sum, value)
	sum, valName = MaxValueUzs(sum, value)
	sum, valName = MaxValueUsd(sum, value)
	sum, valName = MaxValueHuf(sum, value)
	sum, valName = MaxValueTry(sum, value)
	sum, valName = MaxValueTmt(sum, value)
	sum, valName = MaxValueTjs(sum, value)
	sum, valName = MaxValueSgd(sum, value)
	sum, valName = MaxValueSek(sum, value)
	sum, valName = MaxValueRon(sum, value)
	sum, valName = MaxValuePln(sum, value)
	sum, valName = MaxValueMdl(sum, value)
	sum, valName = MaxValueInr(sum, value)
	sum, valName = MaxValueHkd(sum, value)
	sum, valName = MaxValueEur(sum, value)
	sum, valName = MaxValueDkk(sum, value)
	sum, valName = MaxValueCny(sum, value)
	sum, valName = MaxValueChf(sum, value)
	sum, valName = MaxValueNok(sum, value)

	if value.Valute.Jpy.Value == sum {
		valName = value.Valute.Jpy.Name
	} else if value.Valute.Cad.Value == sum {
		valName = value.Valute.Cad.Name
	} else if value.Valute.Xdr.Value == sum {
		valName = value.Valute.Xdr.Name
	} else if value.Valute.Krw.Value == sum {
		valName = value.Valute.Krw.Name
	} else if value.Valute.Nok.Value == sum {
		valName = value.Valute.Nok.Name
	} else if value.Valute.Chf.Value == sum {
		valName = value.Valute.Chf.Name
	} else if value.Valute.Cny.Value == sum {
		valName = value.Valute.Cny.Name
	} else if value.Valute.Dkk.Value == sum {
		valName = value.Valute.Dkk.Name
	} else if value.Valute.Eur.Value == sum {
		valName = value.Valute.Eur.Name
	} else if value.Valute.Inr.Value == sum {
		valName = value.Valute.Inr.Name
	} else if value.Valute.Mdl.Value == sum {
		valName = value.Valute.Mdl.Name
	} else if value.Valute.Pln.Value == sum {
		valName = value.Valute.Pln.Name
	} else if value.Valute.Ron.Value == sum {
		valName = value.Valute.Ron.Name
	} else if value.Valute.Sek.Value == sum {
		valName = value.Valute.Sek.Name
	} else if value.Valute.Sgd.Value == sum {
		valName = value.Valute.Sgd.Name
	} else if value.Valute.Tjs.Value == sum {
		valName = value.Valute.Tjs.Name
	} else if value.Valute.Tmt.Value == sum {
		valName = value.Valute.Tmt.Name
	} else if value.Valute.Try.Value == sum {
		valName = value.Valute.Try.Name
	} else if value.Valute.Huf.Value == sum {
		valName = value.Valute.Huf.Name
	} else if value.Valute.Usd.Value == sum {
		valName = value.Valute.Usd.Name
	} else if value.Valute.Uzs.Value == sum {
		valName = value.Valute.Uzs.Name
	} else if value.Valute.Brl.Value == sum {
		valName = value.Valute.Brl.Name
	} else if value.Valute.Byn.Value == sum {
		valName = value.Valute.Byn.Name
	} else if value.Valute.Bgn.Value == sum {
		valName = value.Valute.Bgn.Name
	} else if value.Valute.Gbp.Value == sum {
		valName = value.Valute.Gbp.Name
	} else if value.Valute.Kgs.Value == sum {
		valName = value.Valute.Kgs.Name
	} else if value.Valute.Kzt.Value == sum {
		valName = value.Valute.Kzt.Name
	} else if value.Valute.Zar.Value == sum {
		valName = value.Valute.Zar.Name
	} else if value.Valute.Hkd.Value == sum {
		valName = value.Valute.Hkd.Name
	} else if value.Valute.Czk.Value == sum {
		valName = value.Valute.Czk.Name
	} else if value.Valute.Aud.Value == sum {
		valName = value.Valute.Aud.Name
	} else if value.Valute.Azn.Value == sum {
		valName = value.Valute.Azn.Name
	} else if value.Valute.Amd.Value == sum {
		valName = value.Valute.Amd.Name
	} else if value.Valute.Uah.Value == sum {
		valName = value.Valute.Uah.Name
	}

	fmt.Printf("Наибольшее значение курса валюты <%s> на сегодняшний день: %f\n", valName, sum)

	min = value.Valute.Aud.Value

	min, valName = MinValueCzk(min, value)
	min, valName = MinValueAud(min, value)
	min, valName = MinValueAzn(min, value)
	min, valName = MinValueAmd(min, value)
	min, valName = MinValueJpy(min, value)
	min, valName = MinValueCad(min, value)
	min, valName = MinValueUah(min, value)
	min, valName = MinValueZar(min, value)
	min, valName = MinValueKzt(min, value)
	min, valName = MinValueKgs(min, value)
	min, valName = MinValueGbp(min, value)
	min, valName = MinValueBgn(min, value)
	min, valName = MinValueByn(min, value)
	min, valName = MinValueBrl(min, value)
	min, valName = MinValueXdr(min, value)
	min, valName = MinValueKrw(min, value)
	min, valName = MinValueUzs(min, value)
	min, valName = MinValueUsd(min, value)
	min, valName = MinValueHuf(min, value)
	min, valName = MinValueTry(min, value)
	min, valName = MinValueTmt(min, value)
	min, valName = MinValueTjs(min, value)
	min, valName = MinValueSgd(min, value)
	min, valName = MinValueSek(min, value)
	min, valName = MinValueRon(min, value)
	min, valName = MinValuePln(min, value)
	min, valName = MinValueMdl(min, value)
	min, valName = MinValueInr(min, value)
	min, valName = MinValueHkd(min, value)
	min, valName = MinValueEur(min, value)
	min, valName = MinValueDkk(min, value)
	min, valName = MinValueCny(min, value)
	min, valName = MinValueChf(min, value)
	min, valName = MinValueNok(min, value)

	if value.Valute.Jpy.Value == min {
		valName = value.Valute.Jpy.Name
	} else if value.Valute.Cad.Value == min {
		valName = value.Valute.Cad.Name
	} else if value.Valute.Xdr.Value == min {
		valName = value.Valute.Xdr.Name
	} else if value.Valute.Krw.Value == min {
		valName = value.Valute.Krw.Name
	} else if value.Valute.Nok.Value == min {
		valName = value.Valute.Nok.Name
	} else if value.Valute.Chf.Value == min {
		valName = value.Valute.Chf.Name
	} else if value.Valute.Cny.Value == min {
		valName = value.Valute.Cny.Name
	} else if value.Valute.Dkk.Value == min {
		valName = value.Valute.Dkk.Name
	} else if value.Valute.Eur.Value == min {
		valName = value.Valute.Eur.Name
	} else if value.Valute.Inr.Value == min {
		valName = value.Valute.Inr.Name
	} else if value.Valute.Mdl.Value == min {
		valName = value.Valute.Mdl.Name
	} else if value.Valute.Pln.Value == min {
		valName = value.Valute.Pln.Name
	} else if value.Valute.Ron.Value == min {
		valName = value.Valute.Ron.Name
	} else if value.Valute.Sek.Value == min {
		valName = value.Valute.Sek.Name
	} else if value.Valute.Sgd.Value == min {
		valName = value.Valute.Sgd.Name
	} else if value.Valute.Tjs.Value == min {
		valName = value.Valute.Tjs.Name
	} else if value.Valute.Tmt.Value == min {
		valName = value.Valute.Tmt.Name
	} else if value.Valute.Try.Value == min {
		valName = value.Valute.Try.Name
	} else if value.Valute.Huf.Value == min {
		valName = value.Valute.Huf.Name
	} else if value.Valute.Usd.Value == min {
		valName = value.Valute.Usd.Name
	} else if value.Valute.Uzs.Value == min {
		valName = value.Valute.Uzs.Name
	} else if value.Valute.Brl.Value == min {
		valName = value.Valute.Brl.Name
	} else if value.Valute.Byn.Value == min {
		valName = value.Valute.Byn.Name
	} else if value.Valute.Bgn.Value == min {
		valName = value.Valute.Bgn.Name
	} else if value.Valute.Gbp.Value == min {
		valName = value.Valute.Gbp.Name
	} else if value.Valute.Kgs.Value == min {
		valName = value.Valute.Kgs.Name
	} else if value.Valute.Kzt.Value == min {
		valName = value.Valute.Kzt.Name
	} else if value.Valute.Zar.Value == min {
		valName = value.Valute.Zar.Name
	} else if value.Valute.Hkd.Value == min {
		valName = value.Valute.Hkd.Name
	} else if value.Valute.Czk.Value == min {
		valName = value.Valute.Czk.Name
	} else if value.Valute.Aud.Value == min {
		valName = value.Valute.Aud.Name
	} else if value.Valute.Azn.Value == min {
		valName = value.Valute.Azn.Name
	} else if value.Valute.Amd.Value == min {
		valName = value.Valute.Amd.Name
	} else if value.Valute.Uah.Value == min {
		valName = value.Valute.Uah.Name
	}

	fmt.Printf("Наименьшее значение курса валюты <%s> на сегодняшний день: %f\n", valName, min)

	averageValue := (value.Valute.Jpy.Value + value.Valute.Xdr.Value + value.Valute.Krw.Value +
		value.Valute.Cad.Value + value.Valute.Nok.Value + value.Valute.Chf.Value + value.Valute.Cny.Value + value.Valute.Dkk.Value +
		value.Valute.Eur.Value + value.Valute.Inr.Value + value.Valute.Mdl.Value + value.Valute.Pln.Value + value.Valute.Ron.Value +
		value.Valute.Sek.Value + value.Valute.Sgd.Value + value.Valute.Tjs.Value + value.Valute.Tmt.Value + value.Valute.Try.Value +
		value.Valute.Huf.Value + value.Valute.Usd.Value + value.Valute.Uzs.Value + value.Valute.Brl.Value + value.Valute.Byn.Value +
		value.Valute.Bgn.Value + value.Valute.Gbp.Value + value.Valute.Kgs.Value + value.Valute.Kzt.Value + value.Valute.Zar.Value +
		value.Valute.Hkd.Value + value.Valute.Czk.Value + value.Valute.Aud.Value + value.Valute.Amd.Value + value.Valute.Uah.Value +
		value.Valute.Azn.Value) / 34

	fmt.Printf("Среднее значение курса рубля по всем валютам: %f\n", averageValue)

}

func MaxValueCzk(sum float64, value Curs) (float64, string) {
	if value.Valute.Czk.Value > sum {
		return value.Valute.Czk.Value, value.Valute.Czk.Name
	} else {
		return sum, ""
	}
}

func MaxValueAud(sum float64, value Curs) (float64, string) {
	if value.Valute.Aud.Value > sum {
		return value.Valute.Aud.Value, value.Valute.Aud.Name
	} else {
		return sum, ""
	}
}

func MaxValueAzn(sum float64, value Curs) (float64, string) {
	if value.Valute.Azn.Value > sum {
		return value.Valute.Azn.Value, value.Valute.Azn.Name
	} else {
		return sum, ""
	}
}

func MaxValueAmd(sum float64, value Curs) (float64, string) {
	if value.Valute.Amd.Value > sum {
		return value.Valute.Amd.Value, value.Valute.Amd.Name
	} else {
		return sum, ""
	}
}

func MaxValueJpy(sum float64, value Curs) (float64, string) {
	if value.Valute.Jpy.Value > sum {
		return value.Valute.Jpy.Value, value.Valute.Jpy.Name
	} else {
		return sum, ""
	}
}

func MaxValueCad(sum float64, value Curs) (float64, string) {
	if value.Valute.Cad.Value > sum {
		return value.Valute.Cad.Value, value.Valute.Cad.Name
	} else {
		return sum, ""
	}
}

func MaxValueUah(sum float64, value Curs) (float64, string) {
	if value.Valute.Uah.Value > sum {
		return value.Valute.Uah.Value, value.Valute.Uah.Name
	} else {
		return sum, ""
	}
}

func MaxValueZar(sum float64, value Curs) (float64, string) {
	if value.Valute.Zar.Value > sum {
		return value.Valute.Zar.Value, value.Valute.Zar.Name
	} else {
		return sum, ""
	}
}

func MaxValueKzt(sum float64, value Curs) (float64, string) {
	if value.Valute.Kzt.Value > sum {
		return value.Valute.Kzt.Value, value.Valute.Kzt.Name
	} else {
		return sum, ""
	}
}

func MaxValueKgs(sum float64, value Curs) (float64, string) {
	if value.Valute.Kgs.Value > sum {
		return value.Valute.Kgs.Value, value.Valute.Kgs.Name
	} else {
		return sum, ""
	}
}

func MaxValueGbp(sum float64, value Curs) (float64, string) {
	if value.Valute.Gbp.Value > sum {
		return value.Valute.Gbp.Value, value.Valute.Gbp.Name
	} else {
		return sum, ""
	}
}

func MaxValueBgn(sum float64, value Curs) (float64, string) {
	if value.Valute.Bgn.Value > sum {
		return value.Valute.Bgn.Value, value.Valute.Bgn.Name
	} else {
		return sum, ""
	}
}

func MaxValueByn(sum float64, value Curs) (float64, string) {
	if value.Valute.Byn.Value > sum {
		return value.Valute.Byn.Value, value.Valute.Byn.Name
	} else {
		return sum, ""
	}
}

func MaxValueBrl(sum float64, value Curs) (float64, string) {
	if value.Valute.Brl.Value > sum {
		return value.Valute.Brl.Value, value.Valute.Brl.Name
	} else {
		return sum, ""
	}
}

func MaxValueXdr(sum float64, value Curs) (float64, string) {
	if value.Valute.Xdr.Value > sum {
		return value.Valute.Xdr.Value, value.Valute.Xdr.Name
	} else {
		return sum, ""
	}
}

func MaxValueKrw(sum float64, value Curs) (float64, string) {
	if value.Valute.Krw.Value > sum {
		return value.Valute.Krw.Value, value.Valute.Krw.Name
	} else {
		return sum, ""
	}
}

func MaxValueUzs(sum float64, value Curs) (float64, string) {
	if value.Valute.Uzs.Value > sum {
		return value.Valute.Uzs.Value, value.Valute.Uzs.Name
	} else {
		return sum, ""
	}
}

func MaxValueUsd(sum float64, value Curs) (float64, string) {
	if value.Valute.Usd.Value > sum {
		return value.Valute.Usd.Value, value.Valute.Usd.Name
	} else {
		return sum, ""
	}
}

func MaxValueHuf(sum float64, value Curs) (float64, string) {
	if value.Valute.Huf.Value > sum {
		return value.Valute.Huf.Value, value.Valute.Huf.Name
	} else {
		return sum, ""
	}
}

func MaxValueTry(sum float64, value Curs) (float64, string) {
	if value.Valute.Try.Value > sum {
		return value.Valute.Try.Value, value.Valute.Try.Name
	} else {
		return sum, ""
	}
}

func MaxValueTmt(sum float64, value Curs) (float64, string) {
	if value.Valute.Tmt.Value > sum {
		return value.Valute.Tmt.Value, value.Valute.Tmt.Name
	} else {
		return sum, ""
	}
}

func MaxValueTjs(sum float64, value Curs) (float64, string) {
	if value.Valute.Tjs.Value > sum {
		return value.Valute.Tjs.Value, value.Valute.Tjs.Name
	} else {
		return sum, ""
	}
}

func MaxValueSgd(sum float64, value Curs) (float64, string) {
	if value.Valute.Sgd.Value > sum {
		return value.Valute.Sgd.Value, value.Valute.Sgd.Name
	} else {
		return sum, ""
	}
}

func MaxValueSek(sum float64, value Curs) (float64, string) {
	if value.Valute.Sek.Value > sum {
		return value.Valute.Sek.Value, value.Valute.Sek.Name
	} else {
		return sum, ""
	}
}

func MaxValueRon(sum float64, value Curs) (float64, string) {
	if value.Valute.Ron.Value > sum {
		return value.Valute.Ron.Value, value.Valute.Ron.Name
	} else {
		return sum, ""
	}
}

func MaxValuePln(sum float64, value Curs) (float64, string) {
	if value.Valute.Pln.Value > sum {
		return value.Valute.Pln.Value, value.Valute.Pln.Name
	} else {
		return sum, ""
	}
}

func MaxValueMdl(sum float64, value Curs) (float64, string) {
	if value.Valute.Mdl.Value > sum {
		return value.Valute.Mdl.Value, value.Valute.Mdl.Name
	} else {
		return sum, ""
	}
}

func MaxValueInr(sum float64, value Curs) (float64, string) {
	if value.Valute.Inr.Value > sum {
		return value.Valute.Inr.Value, value.Valute.Inr.Name
	} else {
		return sum, ""
	}
}

func MaxValueHkd(sum float64, value Curs) (float64, string) {
	if value.Valute.Hkd.Value > sum {
		return value.Valute.Hkd.Value, value.Valute.Eur.Name
	} else {
		return sum, ""
	}
}

func MaxValueEur(sum float64, value Curs) (float64, string) {
	if value.Valute.Eur.Value > sum {
		return value.Valute.Eur.Value, value.Valute.Eur.Name
	} else {
		return sum, ""
	}
}

func MaxValueDkk(sum float64, value Curs) (float64, string) {
	if value.Valute.Dkk.Value > sum {
		return value.Valute.Dkk.Value, value.Valute.Dkk.Name
	} else {
		return sum, ""
	}
}

func MaxValueCny(sum float64, value Curs) (float64, string) {
	if value.Valute.Cny.Value > sum {
		return value.Valute.Cny.Value, value.Valute.Cny.Name
	} else {
		return sum, ""
	}
}

func MaxValueChf(sum float64, value Curs) (float64, string) {
	if value.Valute.Chf.Value > sum {
		return value.Valute.Chf.Value, value.Valute.Chf.Name
	} else {
		return sum, ""
	}
}

func MaxValueNok(sum float64, value Curs) (float64, string) {
	if value.Valute.Nok.Value > sum {
		return value.Valute.Nok.Value, value.Valute.Nok.Name
	} else {
		return sum, ""
	}
}

//

func MinValueCzk(sum float64, value Curs) (float64, string) {
	if value.Valute.Czk.Value < sum {
		return value.Valute.Czk.Value, value.Valute.Czk.Name
	} else {
		return sum, ""
	}
}

func MinValueAud(sum float64, value Curs) (float64, string) {
	if value.Valute.Aud.Value < sum {
		return value.Valute.Aud.Value, value.Valute.Aud.Name
	} else {
		return sum, ""
	}
}

func MinValueAzn(sum float64, value Curs) (float64, string) {
	if value.Valute.Azn.Value < sum {
		return value.Valute.Azn.Value, value.Valute.Azn.Name
	} else {
		return sum, ""
	}
}

func MinValueAmd(sum float64, value Curs) (float64, string) {
	if value.Valute.Amd.Value < sum {
		return value.Valute.Amd.Value, value.Valute.Amd.Name
	} else {
		return sum, ""
	}
}

func MinValueJpy(sum float64, value Curs) (float64, string) {
	if value.Valute.Jpy.Value < sum {
		return value.Valute.Jpy.Value, value.Valute.Jpy.Name
	} else {
		return sum, ""
	}
}

func MinValueCad(sum float64, value Curs) (float64, string) {
	if value.Valute.Cad.Value < sum {
		return value.Valute.Cad.Value, value.Valute.Cad.Name
	} else {
		return sum, ""
	}
}

func MinValueUah(sum float64, value Curs) (float64, string) {
	if value.Valute.Uah.Value < sum {
		return value.Valute.Uah.Value, value.Valute.Uah.Name
	} else {
		return sum, ""
	}
}

func MinValueZar(sum float64, value Curs) (float64, string) {
	if value.Valute.Zar.Value < sum {
		return value.Valute.Zar.Value, value.Valute.Zar.Name
	} else {
		return sum, ""
	}
}

func MinValueKzt(sum float64, value Curs) (float64, string) {
	if value.Valute.Kzt.Value < sum {
		return value.Valute.Kzt.Value, value.Valute.Kzt.Name
	} else {
		return sum, ""
	}
}

func MinValueKgs(sum float64, value Curs) (float64, string) {
	if value.Valute.Kgs.Value < sum {
		return value.Valute.Kgs.Value, value.Valute.Kgs.Name
	} else {
		return sum, ""
	}
}

func MinValueGbp(sum float64, value Curs) (float64, string) {
	if value.Valute.Gbp.Value < sum {
		return value.Valute.Gbp.Value, value.Valute.Gbp.Name
	} else {
		return sum, ""
	}
}

func MinValueBgn(sum float64, value Curs) (float64, string) {
	if value.Valute.Bgn.Value < sum {
		return value.Valute.Bgn.Value, value.Valute.Bgn.Name
	} else {
		return sum, ""
	}
}

func MinValueByn(sum float64, value Curs) (float64, string) {
	if value.Valute.Byn.Value < sum {
		return value.Valute.Byn.Value, value.Valute.Byn.Name
	} else {
		return sum, ""
	}
}

func MinValueBrl(sum float64, value Curs) (float64, string) {
	if value.Valute.Brl.Value < sum {
		return value.Valute.Brl.Value, value.Valute.Brl.Name
	} else {
		return sum, ""
	}
}

func MinValueXdr(sum float64, value Curs) (float64, string) {
	if value.Valute.Xdr.Value < sum {
		return value.Valute.Xdr.Value, value.Valute.Xdr.Name
	} else {
		return sum, ""
	}
}

func MinValueKrw(sum float64, value Curs) (float64, string) {
	if value.Valute.Krw.Value < sum {
		return value.Valute.Krw.Value, value.Valute.Krw.Name
	} else {
		return sum, ""
	}
}

func MinValueUzs(sum float64, value Curs) (float64, string) {
	if value.Valute.Uzs.Value < sum {
		return value.Valute.Uzs.Value, value.Valute.Uzs.Name
	} else {
		return sum, ""
	}
}

func MinValueUsd(sum float64, value Curs) (float64, string) {
	if value.Valute.Usd.Value < sum {
		return value.Valute.Usd.Value, value.Valute.Usd.Name
	} else {
		return sum, ""
	}
}

func MinValueHuf(sum float64, value Curs) (float64, string) {
	if value.Valute.Huf.Value < sum {
		return value.Valute.Huf.Value, value.Valute.Huf.Name
	} else {
		return sum, ""
	}
}

func MinValueTry(sum float64, value Curs) (float64, string) {
	if value.Valute.Try.Value < sum {
		return value.Valute.Try.Value, value.Valute.Try.Name
	} else {
		return sum, ""
	}
}

func MinValueTmt(sum float64, value Curs) (float64, string) {
	if value.Valute.Tmt.Value < sum {
		return value.Valute.Tmt.Value, value.Valute.Tmt.Name
	} else {
		return sum, ""
	}
}

func MinValueTjs(sum float64, value Curs) (float64, string) {
	if value.Valute.Tjs.Value < sum {
		return value.Valute.Tjs.Value, value.Valute.Tjs.Name
	} else {
		return sum, ""
	}
}

func MinValueSgd(sum float64, value Curs) (float64, string) {
	if value.Valute.Sgd.Value < sum {
		return value.Valute.Sgd.Value, value.Valute.Sgd.Name
	} else {
		return sum, ""
	}
}

func MinValueSek(sum float64, value Curs) (float64, string) {
	if value.Valute.Sek.Value < sum {
		return value.Valute.Sek.Value, value.Valute.Sek.Name
	} else {
		return sum, ""
	}
}

func MinValueRon(sum float64, value Curs) (float64, string) {
	if value.Valute.Ron.Value < sum {
		return value.Valute.Ron.Value, value.Valute.Ron.Name
	} else {
		return sum, ""
	}
}

func MinValuePln(sum float64, value Curs) (float64, string) {
	if value.Valute.Pln.Value < sum {
		return value.Valute.Pln.Value, value.Valute.Pln.Name
	} else {
		return sum, ""
	}
}

func MinValueMdl(sum float64, value Curs) (float64, string) {
	if value.Valute.Mdl.Value < sum {
		return value.Valute.Mdl.Value, value.Valute.Mdl.Name
	} else {
		return sum, ""
	}
}

func MinValueInr(sum float64, value Curs) (float64, string) {
	if value.Valute.Inr.Value < sum {
		return value.Valute.Inr.Value, value.Valute.Inr.Name
	} else {
		return sum, ""
	}
}

func MinValueHkd(sum float64, value Curs) (float64, string) {
	if value.Valute.Hkd.Value < sum {
		return value.Valute.Hkd.Value, value.Valute.Eur.Name
	} else {
		return sum, ""
	}
}

func MinValueEur(sum float64, value Curs) (float64, string) {
	if value.Valute.Eur.Value < sum {
		return value.Valute.Eur.Value, value.Valute.Eur.Name
	} else {
		return sum, ""
	}
}

func MinValueDkk(sum float64, value Curs) (float64, string) {
	if value.Valute.Dkk.Value < sum {
		return value.Valute.Dkk.Value, value.Valute.Dkk.Name
	} else {
		return sum, ""
	}
}

func MinValueCny(sum float64, value Curs) (float64, string) {
	if value.Valute.Cny.Value < sum {
		return value.Valute.Cny.Value, value.Valute.Cny.Name
	} else {
		return sum, ""
	}
}

func MinValueChf(sum float64, value Curs) (float64, string) {
	if value.Valute.Chf.Value < sum {
		return value.Valute.Chf.Value, value.Valute.Chf.Name
	} else {
		return sum, ""
	}
}

func MinValueNok(sum float64, value Curs) (float64, string) {
	if value.Valute.Nok.Value < sum {
		return value.Valute.Nok.Value, value.Valute.Nok.Name
	} else {
		return sum, ""
	}
}
