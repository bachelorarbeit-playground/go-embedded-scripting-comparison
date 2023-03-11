# go-embedded-scripting-comparison

This is a comparison between the execution speed of example scripts written in Lua and Jsonnet and embedded inside Go host programs with the help of [gopher-lua](https://github.com/yuin/gopher-lua) and [go-jsonnet](https://github.com/google/go-jsonnet). Measurements were taken on desktop machine having a Intel(R) Core(TM) i5-6600K CPU @ 3.50GHz with 16.0 GB of RAM running WSL2 and on a MacBook Pro 2016 15" with 16.0 GB of RAM running MacOS.

## Lua script

```
function processing(event)
  local valueInKwh = event.value * 1000
  local hour = ''
  if event.interval <= 10 then
    hour = '0' .. tostring(event.interval - 1)
  else
    hour = tostring(event.interval - 1)
  end
  local timestamp = event.date .. " " .. hour .. ":00:00" .. event.timezone
  local regionUppercase = string.upper(event.region)
  local availabilityPercentage = event.availability / 100

  if valueInKwh >= 50 then
    return {
      parkId = event.parkId,
      timestamp = timestamp,
      value = valueInKwh
    }
  else
    return {
      parkId = event.parkId,
      timestamp = timestamp,
      value = valueInKwh,
      availability = availabilityPercentage,
      region = regionUppercase,
    }
  end
end
```

## Jsonnet script

```
local valueInKwh = stdin.value * 1000;
local hour =
  if stdin.interval <= 10
  then '0' + std.toString(stdin.interval - 1)
  else std.toString(stdin.interval - 1);
local timestamp = stdin.date + ' ' + hour + ':00:00 ' + stdin.timezone;
local regionUppercase = std.asciiUpper(stdin.region);
local availabilityPercentage = stdin.availability / 100;

if stdin.value >= 50 then {
  park_id: stdin.park_id,
  timestamp: timestamp,
  value: valueInKwh,
} else {
  park_id: stdin.park_id,
  region: regionUppercase,
  timestamp: timestamp,
  value: valueInKwh,
  availability: availabilityPercentage,
}
```
## Result
![image](https://user-images.githubusercontent.com/23280777/224491979-1577e97f-ea37-470b-bbbd-4c8834be4f56.png)
