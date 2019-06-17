# assetflow

하드웨어, 소프트웨어, 공용계정, 부동산을 관리하기 위한 Back-end commandline tool.
그룹을 운용하면서 필요시 자산을 계산하기 위해서 사용합니다.

현재는 명령어를 사용하면서 기능을 추가, 수정하는 단계입니다.

# 사용법

#### 카메라 추가

```bash
$ assetflow -add camera -user lazypic -cost 150000 -product "Nikon F90X" -description "기증"
```

#### 렌즈 추가

```bash
$ assetflow -add lens -user lazypic -cost 0 -product "Anam Nikon AF NIKKOR" -description "기증" -focallength "28-70mm" -sn 7005744
```

#### 렌즈필터 추가

```bash
$ assetflow -add other -user lazypic -cost 0 -product "MATIN 52mm UV" -description "lensfilter,기증"
```

### Reference
- https://github.com/Rhymond/go-money