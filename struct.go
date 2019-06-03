package main

// Hw 는 데스크탑, 노트북등의 하드웨어 관리를 위한 자료구조이다.
type Hw struct {
	Typ          string  // 타입
	PurchaseDate string  // 구매일
	Product      string  // 제품명
	Owner        string  // 관리자
	State        string  // 상태
	User         string  // 사용자
	Sn           string  // 시리얼넘버
	Price        float64 // 가격
	MonetaryUnit string  // 가격단위
	Description  string  // 설명
}

// Sw 는 소프트웨어를 관리하기 위한 자료구조이다.
type Sw struct {
	Typ                  string  // 타입
	PurchaseDate         string  // 구매일
	Sn                   string  // 시리얼넘버
	Product              string  // 제품명
	MonthMaintenanceCost float64 // 월유지비
	MonetaryUnit         string  // 가격단위
	Description          string  // 설명
}

// Account 는 공용계정을 관리하기 위한 자료구조이다.
type Account struct {
	Typ                  string // 타입
	PurchaseDate         string // 구매일
	Product              string // 제품명
	URL                  string // 주소
	ID                   string
	Password             string
	MonthMaintenanceCost float64 // 월유지비
	MonetaryUnit         string  // 가격단위
	Description          string  // 설명
}

// RealEstate 는 부동산을 관리하기 위한 자료구조이다.
type RealEstate struct {
	Typ                  string  // 타입
	PurchaseDate         string  // 구매일
	MonthMaintenanceCost float64 // 월유지비
	MonetaryUnit         string  // 가격단위
	Address              string  // 주소
	Description          string  // 설명
}

// Other 는 위 분류에 들어가지 않는 항목이다. 예) 안마의자
type Other struct {
	Typ                  string  // 타입
	PurchaseDate         string  // 구매일
	Product              string  // 제품명
	MonthMaintenanceCost float64 // 월유지비
	MonetaryUnit         string  // 가격단위
	Description          string  // 설명
}
