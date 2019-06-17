# assetflow

하드웨어, 소프트웨어, 공용계정, 부동산을 관리하기 위한 Back-end commandline tool.
그룹을 운용하면서 필요시 자산을 계산하기 위해서 사용합니다.

현재는 명령어를 사용하면서 기능을 추가, 수정하는 단계입니다.

# 사용법

#### 하드웨어 추가

```bash
$ assetflow -add hw -user lazypic -cost 2100000 -product "imac" -macadress "68:3a:3a:14:59:b6"
```

#### 카메라 추가

```bash
$ assetflow -add camera -user lazypic -cost 150000 -product "Nikon F90X" -description "기증"
```

#### Rig 추가

```bash
$ assetflow -add rig -user lazypic -cost 380000 -product "Sirui t-025x" -description "tripod"
```

#### 렌즈 추가

```bash
$ assetflow -add lens -user lazypic -cost 0 -product "Anam Nikon AF NIKKOR" -description "기증" -focallength "28-70mm" -sn 7005744
```

#### 렌즈필터 추가

```bash
$ assetflow -add other -user lazypic -cost 0 -product "MATIN 52mm UV" -description "lensfilter,기증"
```

#### 소프트웨어 추가
간혹 특정 소프트웨어는 시리얼넘버, 맥어드레스 관리도 필요합니다.

```bash
$ assetflow -add sw -product "Productname" -sn "XXXXX-XXXXX-XXXXX-XXXXX" -macaddress "68:3a:3a:14:59:b6"
```

#### 계정관리
공용으로 사용하는 웹서비스 계정관리

```bash
$ assetflow -add acount -id "id" -pw "password" -url "https://url"
```

### Reference
- https://github.com/Rhymond/go-money