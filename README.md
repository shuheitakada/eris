# eris
CLI to create a csv file with random strings and numbers

# Installation
This package can be installed with the following command:
```
go get github.com/shuheitakada/eris
```

# Usage
## create a csv with 10-digit random numbers

Command:
```
eris create "/\d{10}/" -n 5
```

Output: (example.csv)
```
7371333677
3841619178
5868282416
2434605750
4842472718
```

## create a csv with random strings of 10 characters

Command:
```
eris create "/\w{10}/" -n 5
```

Output: (example.csv)
```
4CFzOONl2X
VXlEioZosD
iJ4QKc7nVM
fPo5Yxlcy5
hmVb_ru0fo
```

## create a csv with fixed strings

Command:
```
eris create "fixed-string" -n 5
```

Output: (example.csv)
```
fixed-string
fixed-string
fixed-string
fixed-string
fixed-string
```

## create a csv that combines a fixed string and a random string

Command:
```
eris create "prefix-" "/\w{10}/" -n 5
```

Output: (example.csv)
```
prefix-h12aaxqqH4
prefix-PfW1Paw_my
prefix-fOUN1WYotz
prefix-pYdcFVRpXh
prefix-EKKP2pBVgk
```

## create a multi-column csv

Command:
```
eris create "/\w{10}/" "," "column2" "," "column3" -n 5
```

Output: (example.csv)
```
dUIN7sJfSt,column2,column3
zI2dZXFmAa,column2,column3
cKHqGic4TO,column2,column3
uqT5MtEyfl,column2,column3
bSHpN3iVxs,column2,column3
```

# Flags
```
-f, --file: file path where the csv file is created (default "example.csv")
-n, --line: line number (default: 100)
```


# License
haumea is released under the Apache 2.0 license. See [LICENSE](https://github.com/shuheitakada/haumea/blob/main/LICENSE)
