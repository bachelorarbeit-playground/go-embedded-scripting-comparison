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
            value =evalueInKwh
        }
    else
        return {
            parkId = event.parkId,
            timestamp = timestamp,
            value = valueInKwh,
            region = regionUppercase,
            availability = availabilityPercentage
        }
    end
end
