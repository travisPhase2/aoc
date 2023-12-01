local utils = require "utils"

local example = [[
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
]]

local lines = utils.split_str_by(example, "\n")

local numbers_list = {}
for line in io.lines("input") do
    local numbers = line:gsub("%D", "")
    local first = string.sub(numbers, 1, 1)
    local last = string.sub(numbers, #numbers)
    local combined = first .. last
    local result = tonumber(combined)
    if result then
        table.insert(numbers_list, result)
    end
end

local sum = 0
for _, v in ipairs(numbers_list) do
    sum = sum + v
end

print(sum)
