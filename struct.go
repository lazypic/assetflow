package main

// Hw 는 데스크탑, 노트북등의 하드웨어 관리를 위한 자료구조이다.
type Hw struct {
	Typ           string  // 타입
	CreateDate    string  // 생성일
	Product       string  // 제품명
	ProductStatus string  // 상태
	ProductUser   string  // 사용자
	Sn            string  // 시리얼넘버
	Price         float64 // 가격
	PurchaseDate  string  // 구매일
	MonetaryUnit  string  // 가격단위
	Description   string  // 설명
}

// Sw 는 소프트웨어를 관리하기 위한 자료구조이다.
type Sw struct {
	Typ                  string  // 타입
	CreateDate           string  // 생성일
	Sn                   string  // 시리얼넘버
	Product              string  // 제품명
	MonthMaintenanceCost float64 // 월유지비
	MonetaryUnit         string  // 가격단위
	Description          string  // 설명
}

// Account 는 공용계정을 관리하기 위한 자료구조이다.
type Account struct {
	Typ                  string  // 타입
	CreateDate           string  // 생성일
	Product              string  // 제품명
	ProductURL           string  // 주소
	ID                   string  // ID
	Password             string  // 패스워드
	MonthMaintenanceCost float64 // 월유지비
	MonetaryUnit         string  // 가격단위
	PurchaseDate         string  // 구매일
	Description          string  // 설명
}

// RealEstate 는 부동산을 관리하기 위한 자료구조이다.
type RealEstate struct {
	Typ                  string  // 타입
	CreateDate           string  // 생성일
	ContractDate         string  // 구매일
	MonthMaintenanceCost float64 // 월유지비
	MonetaryUnit         string  // 가격단위
	Address              string  // 주소
	Description          string  // 설명
}

// Other 는 위 분류에 들어가지 않는 항목이다. 예) 안마의자
type Other struct {
	Typ                  string  // 타입
	CreateDate           string  // 생성일
	PurchaseDate         string  // 구매일
	Product              string  // 제품명
	MonthMaintenanceCost float64 // 월유지비
	MonetaryUnit         string  // 가격단위
	Description          string  // 설명
}
