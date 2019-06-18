package main

import (
	"fmt"
	"log"
)

const (
	// LimitKRW 자산분류 최소금액
	LimitKRW = 500000
	// LimitAssetMonth 자산보관 최대월 60개월, 5년
	LimitAssetMonth = 60
	// LimitAssetMonth 자산보관 최대월 240개월, 20년(240)~40년(480)
	LimitRealEstateMonth = 240
)

// Hw 는 데스크탑, 노트북등의 하드웨어 관리를 위한 자료구조이다.
type Hw struct {
	Typ            string // 타입
	CreateDate     string // 생성일
	Product        string // 제품명
	ProductUser    string // 사용자
	ProductStatus  string // 상태
	Cost           int64  // 가격
	PurchaseDate   string // 구매일
	MonetaryUnit   string // 가격단위
	Description    string // 설명
	MonthlyPayment bool   // 월결제
}

// Sound 는 사운드 장비 관리를 위한 자료구조이다.
type Sound struct {
	Typ            string // 타입
	CreateDate     string // 생성일
	Product        string // 제품명
	ProductStatus  string // 상태
	ProductUser    string // 사용자
	Cost           int64  // 가격
	PurchaseDate   string // 구매일
	MonetaryUnit   string // 가격단위
	Description    string // 설명
	MonthlyPayment bool   // 월대여 형태
}

// Camera 는 카메라 장비 관리를 위한 자료구조이다.
type Camera struct {
	Typ            string // 타입
	CreateDate     string // 생성일
	Product        string // 제품명
	ProductStatus  string // 상태
	ProductUser    string // 사용자
	Cost           int64  // 가격
	PurchaseDate   string // 구매일
	MonetaryUnit   string // 가격단위
	Description    string // 설명
	MonthlyPayment bool   // 월대여 형태
}

// Lens 는 카메라 렌즈 관리를 위한 자료구조이다.
type Lens struct {
	Typ            string // 타입
	CreateDate     string // 생성일
	Product        string // 제품명
	ProductStatus  string // 상태
	ProductUser    string // 사용자
	Cost           int64  // 가격
	PurchaseDate   string // 구매일
	MonetaryUnit   string // 가격단위
	Description    string // 설명
	Sn             string // 시리얼넘버
	FocalLength    string // 화각
	MonthlyPayment bool   // 월대여 형태
}

// Rig 는 카메라 리그장비 관리를 위한 자료구조이다.
type Rig struct {
	Typ            string // 타입
	CreateDate     string // 생성일
	Product        string // 제품명
	ProductStatus  string // 상태
	ProductUser    string // 사용자
	Cost           int64  // 가격
	PurchaseDate   string // 구매일
	MonetaryUnit   string // 가격단위
	Description    string // 설명
	Sn             string // 시리얼넘버
	MonthlyPayment bool   // 월대여 형태
}

// Sw 는 소프트웨어를 관리하기 위한 자료구조이다. 소프트웨어는 무형자산. 양도할 수 없다.
type Sw struct {
	Typ            string // 타입
	CreateDate     string // 생성일
	Sn             string // 시리얼넘버
	Product        string // 제품명
	Cost           int64  // 유지비
	MonetaryUnit   string // 가격단위
	Description    string // 설명
	MonthlyPayment bool   // 월결제
	MacAdress      string // 맥어드레스, 간혹 특정소프트웨어는 맥어드레스가 필요하다.
}

// Account 는 공용계정을 관리하기 위한 자료구조이다.
type Account struct {
	Typ            string // 타입
	CreateDate     string // 생성일
	Product        string // 제품명
	ProductURL     string // 주소
	ID             string // ID
	Password       string // 패스워드
	Cost           int64  // 월유지비
	MonetaryUnit   string // 가격단위
	PurchaseDate   string // 구매일
	Description    string // 설명
	MonthlyPayment bool   // 월결제
}

// RealEstate 는 부동산을 관리하기 위한 자료구조이다.
type RealEstate struct {
	Typ            string // 타입
	CreateDate     string // 생성일
	ContractDate   string // 구매일
	Cost           int64  // 월유지비
	MonetaryUnit   string // 가격단위
	Address        string // 주소
	Description    string // 설명
	MonthlyPayment bool   // 월결제
}

// Other 는 위 분류에 들어가지 않는 항목이다. 예) 안마의자
type Other struct {
	Typ            string // 타입
	CreateDate     string // 생성일
	PurchaseDate   string // 구매일
	Product        string // 제품명
	ProductStatus  string // 상태
	ProductUser    string // 사용자
	Cost           int64  // 유지비
	MonetaryUnit   string // 가격단위
	Description    string // 설명
	MonthlyPayment bool   // 월결제
}

// Depreciation 메소드 Hw 자료구조의 감가상각누계액 모델
func (hw Hw) Depreciation(model string) int64 {
	// 월결제 모델이라면 감가상각비에 넣지 않는다.
	if hw.MonthlyPayment {
		return 0
	}
	// 비용이 500000 만원 미만이라면 감가상각비에 넣지않는다.
	if hw.MonetaryUnit == "KRW" && hw.Cost < int64(LimitKRW) {
		return 0
	}
	// 구매일이 60개월 이상이라면 감가상각비에 넣지않는다.
	m, err := PeriodOfUse(hw.PurchaseDate)
	if err != nil {
		log.Fatal(err)
	}
	if m > LimitAssetMonth {
		return 0
	}
	switch model {
	case "year":
		// 감가상각누계액
		return hw.Cost - (hw.Cost * (int64(m) / 12) / 5)
	case "month":
		// 월별 세부 감가상각액
		return (60*hw.Cost - int64(m)*hw.Cost) / int64(LimitAssetMonth)
	default:
		// 기본적으로 연도별로 감가상각을 계산한다.
		return hw.Cost - (hw.Cost * (int64(m) / 12) / 5)
	}
}

func (hw Hw) String() string {
	return fmt.Sprintf(`
	Type: %s
	CreateDate: %s
	Product: %s(%s)
	User: %s
	Cost: %d %s
	MonthlyPayment: %t
	PurchaseDate: %s
	Description: %s
	Depreciation: %d
	`,
		hw.Typ,
		hw.CreateDate,
		hw.Product,
		hw.ProductStatus,
		hw.ProductUser,
		hw.Cost,
		hw.MonetaryUnit,
		hw.MonthlyPayment,
		hw.PurchaseDate,
		hw.Description,
		hw.Depreciation("year"),
	)
}

// Depreciation 메소드 Camera 자료구조의 감가상각을 계산한다.
func (cmr Camera) Depreciation() int64 {
	// 월결제 모델이라면 감가상각비에 넣지 않는다.
	if cmr.MonthlyPayment {
		return 0
	}
	// 비용이 500000 만원 미만이라면 감가상각비에 넣지않는다.
	if cmr.MonetaryUnit == "KRW" && cmr.Cost < int64(LimitKRW) {
		return 0
	}
	// 구매일이 60개월 이상이라면 감가상각비에 넣지않는다.
	m, err := PeriodOfUse(cmr.PurchaseDate)
	if err != nil {
		log.Fatal(err)
	}
	if m > LimitAssetMonth {
		return 0
	}
	return (60*cmr.Cost - int64(m)*cmr.Cost) / int64(LimitAssetMonth)
}

// Depreciation 메소드 Sound 자료구조의 감가상각을 계산한다.
func (snd Sound) Depreciation() int64 {
	// 월결제 모델이라면 감가상각비에 넣지 않는다.
	if snd.MonthlyPayment {
		return 0
	}
	// 비용이 500000 만원 미만이라면 감가상각비에 넣지않는다.
	if snd.MonetaryUnit == "KRW" && snd.Cost < int64(LimitKRW) {
		return 0
	}
	// 구매일이 60개월 이상이라면 감가상각비에 넣지않는다.
	m, err := PeriodOfUse(snd.PurchaseDate)
	if err != nil {
		log.Fatal(err)
	}
	if m > LimitAssetMonth {
		return 0
	}
	return (60*snd.Cost - int64(m)*snd.Cost) / int64(LimitAssetMonth)
}

// Depreciation 메소드 Rig 자료구조의 감가상각을 계산한다.
func (rig Rig) Depreciation() int64 {
	// 월결제 모델이라면 감가상각비에 넣지 않는다.
	if rig.MonthlyPayment {
		return 0
	}
	// 비용이 500000 만원 미만이라면 감가상각비에 넣지않는다.
	if rig.MonetaryUnit == "KRW" && rig.Cost < int64(LimitKRW) {
		return 0
	}
	// 구매일이 60개월 이상이라면 감가상각비에 넣지않는다.
	m, err := PeriodOfUse(rig.PurchaseDate)
	if err != nil {
		log.Fatal(err)
	}
	if m > LimitAssetMonth {
		return 0
	}
	return (60*rig.Cost - int64(m)*rig.Cost) / int64(LimitAssetMonth)
}

// Depreciation 메소드는 Lens 자료구조의 감가상각을 계산한다.
func (lns Lens) Depreciation() int64 {
	// 월결제 모델이라면 감가상각비에 넣지 않는다.
	if lns.MonthlyPayment {
		return 0
	}
	// 비용이 500000 만원 미만이라면 감가상각비에 넣지않는다.
	if lns.MonetaryUnit == "KRW" && lns.Cost < int64(LimitKRW) {
		return 0
	}
	// 구매일이 60개월 이상이라면 감가상각비에 넣지않는다.
	m, err := PeriodOfUse(lns.PurchaseDate)
	if err != nil {
		log.Fatal(err)
	}
	if m > LimitAssetMonth {
		return 0
	}
	return (60*lns.Cost - int64(m)*lns.Cost) / int64(LimitAssetMonth)
}

// Depreciation 메소드는 Other 자료구조의 감가상각을 계산한다.
func (oth Other) Depreciation() int64 {
	// 월결제 모델이라면 감가상각비에 넣지 않는다.
	if oth.MonthlyPayment {
		return 0
	}
	// 비용이 500000 만원 미만이라면 감가상각비에 넣지않는다.
	if oth.MonetaryUnit == "KRW" && oth.Cost < int64(LimitKRW) {
		return 0
	}
	// 구매일이 60개월 이상이라면 감가상각비에 넣지않는다.
	m, err := PeriodOfUse(oth.PurchaseDate)
	if err != nil {
		log.Fatal(err)
	}
	if m > LimitAssetMonth {
		return 0
	}
	return (60*oth.Cost - int64(m)*oth.Cost) / int64(LimitAssetMonth)
}
