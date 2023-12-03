local utils = require "utils"

-- local function part_1()
--     local example = [[
--     1abc2
--     pqr3stu8vwx
--     a1b2c3d4e5f
--     treb7uchet
--     ]]

--     -- local lines = utils.split_str_by(example, "\n")

--     local numbers_list = {}
--     for line in io.lines("input") do
--         local numbers = line:gsub("%D", "")
--         local first = string.sub(numbers, 1, 1)
--         local last = string.sub(numbers, #numbers)
--         local combined = first .. last
--         local result = tonumber(combined)
--         if result then
--             table.insert(numbers_list, result)
--         end
--     end

--     local sum = 0
--     for _, v in ipairs(numbers_list) do
--         sum = sum + v
--     end

--     print(sum)
-- end

local example = [[
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
]]

local number_map = {
    one = 1,
    two = 2,
    three = 3,
    four = 4,
    five = 5,
    six = 6,
    seven = 7,
    eight = 8,
    nine = 9
}

---records each number's position in a given string
---@param str string
---@return table
local function search_for_nums(str)
    local results = {}
    for k, v in pairs(number_map) do
        local start, ending = str:find(k)
        if start ~= nil and ending ~= nil then
            results[v] = start
        end

        start, ending = str:find(v)
        if start ~= nil and ending ~= nil then
            results[v] = start
        end
    end

    return results
end

local lines = utils.split_str_by(example, "\n")
local nums = {}
for i = 1, #lines do
-- for line in io.lines("input") do
    local line = lines[i]
    local results = search_for_nums(line)

    for key, value in pairs(results) do
        print(key, value)
    end

    break
end

-- print('\n')

-- local sum = 0
-- for _, value in ipairs(nums) do
--     sum = sum + value
-- end

-- print(sum)