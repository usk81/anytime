# Anytime

This package is **EXPERIMENTAL** now !!

## Overview

anytime is a library which provides a set of extensions on go's standard `time` library.

## Install

```
go get github.com/usk81/anytime
```

## Goal

- Be compatible with standard `time` library
- Can be replaced from the standard `time` library

## Progress

### Time

| Function                  | Implementation     | Note         |
|---------------------------|--------------------|--------------|
| Date                      | :construction:     |              |
| Now                       | :heavy_check_mark: |              |
| Parse                     | :heavy_check_mark: |              |
| ParseInLocation           | :construction:     |              |
| Unix                      | :heavy_check_mark: |              |
| (t Time) Add              | :x:                |              |
| (t Time) AddDate          | :x:                |              |
| (t Time) After            | :x:                |              |
| (t Time) AppendFormat     | :heavy_check_mark: | no overrdide |
| (t Time) Before           | :x:                |              |
| (t Time) Clock            | :heavy_check_mark: | no overrdide |
| (t Time) Date             | :construction:     |              |
| (t Time) Day              | :heavy_check_mark: | no overrdide |
| (t Time) Equal            | :x:                |              |
| (t Time) Format           | :heavy_check_mark: | no overrdide |
| (t *Time) GobDecode       | :x:                |              |
| (t Time) GobEncode        | :x:                |              |
| (t Time) Hour             | :heavy_check_mark: | no overrdide |
| (t Time) ISOWeek          | :heavy_check_mark: | no overrdide |
| (t Time) In               | :x:                |              |
| (t Time) IsZero           | :heavy_check_mark: | no overrdide |
| (t Time) Local            | :x:                |              |
| (t Time) Location         | :x:                |              |
| (t Time) MarshalBinary    | :x:                |              |
| (t Time) MarshalJSON      | :heavy_check_mark: |              |
| (t Time) MarshalText      | :heavy_check_mark: |              |
| (t Time) Minute           | :heavy_check_mark: | no overrdide |
| (t Time) Month            | :x:                |              |
| (t Time) Nanosecond       | :heavy_check_mark: | no overrdide |
| (t Time) Round            | :x:                |              |
| (t Time) Second           | :heavy_check_mark: | no overrdide |
| (t Time) String           | :heavy_check_mark: | no overrdide |
| (t Time) Sub              | :x:                |              |
| (t Time) Truncate         | :x:                |              |
| (t Time) UTC              | :x:                |              |
| (t Time) Unix             | :heavy_check_mark: | no overrdide |
| (t Time) UnixNano         | :heavy_check_mark: | no overrdide |
| (t *Time) UnmarshalBinary | :x:                |              |
| (t *Time) UnmarshalJSON   | :heavy_check_mark: |              |
| (t *Time) UnmarshalText   | :heavy_check_mark: |              |
| (t Time) Weekday          | :heavy_check_mark: | no overrdide |
| (t Time) Year             | :heavy_check_mark: | no overrdide |
| (t Time) YearDay          | :heavy_check_mark: | no overrdide |
| (t Time) Zone             | :heavy_check_mark: | no overrdide |