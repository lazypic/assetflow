package main

import "fmt"

// Hw 는 데스크탑, 노트북등의 하드웨어 관리를 위한 자료구조이다.
type Hw struct {
	Typ            string // 타입
	CreateDate     string // 생성일
	Product        string // 제품명
	ProductStatus  string // 상태
	ProductUser    string // 사용자
	Cost           int64  // 가격
	PurchaseDate   string // 구매일
	MonetaryUnit   string // 가격단위
	Description    string // 설명
	MonthlyPayment bool   // 월결제
}

// Sound 는 사운드 장비 관리를 위한 자료구조이다.
type Sound struct {
	Typ           string // 타입
	CreateDate    string // 생성일
	Product       string // 제품명
	ProductStatus string // 상태
	ProductUser   string // 사용자
	Cost          int64  // 가격
	PurchaseDate  string // 구매일
	MonetaryUnit  string // 가격단위
	Description   string // 설명
}

// Camera 는 카메라 장비 관리를 위한 자료구조이다.
type Camera struct {
	Typ           string // 타입
	CreateDate    string // 생성일
	Product       string // 제품명
	ProductStatus string // 상태
	ProductUser   string // 사용자
	Cost          int64  // 가격
	PurchaseDate  string // 구매일
	MonetaryUnit  string // 가격단위
	Description   string // 설명
}

// Lens 는 카메라 렌즈 관리를 위한 자료구조이다.
type Lens struct {
	Typ           string // 타입
	CreateDate    string // 생성일
	Product       string // 제품명
	ProductStatus string // 상태
	ProductUser   string // 사용자
	Cost          int64  // 가격
	PurchaseDate  string // 구매일
	MonetaryUnit  string // 가격단위
	Description   string // 설명
	Sn            string // 시리얼넘버
	FocalLength   string // 화각
}

// Rig 는 카메라 리그장비 관리를 위한 자료구조이다.
type Rig struct {
	Typ           string // 타입
	CreateDate    string // 생성일
	Product       string // 제품명
	ProductStatus string // 상태
	ProductUser   string // 사용자
	Cost          int64  // 가격
	PurchaseDate  string // 구매일
	MonetaryUnit  string // 가격단위
	Description   string // 설명
	Sn            string // 시리얼넘버
}

// Sw 는 소프트웨어를 관리하기 위한 자료구조이다.
type Sw struct {
	Typ            string // 타입
	CreateDate     string // 생성일
	Sn             string // 시리얼넘버
	Product        string // 제품명
	Cost           int64  // 유지비
	MonetaryUnit   string // 가격단위
	Description    string // 설명
	MonthlyPayment bool   // 월결제
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
	Cost           int64  // 유지비
	MonetaryUnit   string // 가격단위
	Description    string // 설명
	MonthlyPayment bool   // 월결제
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
	)
}
