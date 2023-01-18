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


