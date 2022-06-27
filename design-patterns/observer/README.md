# Observer Pattern

Statement:

Write a weather station program that updates the temperature every 1 second. There are multiple clients connected to the station. Each client should be notified once the temperature is updated.

Run: `go run *.go`

Output:
```
❯ go run *.go
2022/06/27 11:43:03 mobileObserver  : updated temperature is  22  degree Celcius
2022/06/27 11:43:03 desktopObserver  : updated temperature is  22  degree Celcius
2022/06/27 11:43:04 mobileObserver  : updated temperature is  10  degree Celcius
2022/06/27 11:43:04 desktopObserver  : updated temperature is  10  degree Celcius
2022/06/27 11:43:05 mobileObserver  : updated temperature is  12  degree Celcius
2022/06/27 11:43:05 desktopObserver  : updated temperature is  12  degree Celcius
2022/06/27 11:43:06 mobileObserver  : updated temperature is  15  degree Celcius
2022/06/27 11:43:06 desktopObserver  : updated temperature is  15  degree Celcius
2022/06/27 11:43:07 mobileObserver  : updated temperature is  14  degree Celcius
2022/06/27 11:43:07 desktopObserver  : updated temperature is  14  degree Celcius
5 seconds! End Ticker
```