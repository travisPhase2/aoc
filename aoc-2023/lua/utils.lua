local M = {}

---splits a string by a character into a list
---@param str string the string we want to split
---@param char string the delimeter
---@return table
M.split_str_by = function(str, char)
    assert(type(str) == "string", "str needs to be a string")
    assert(type(char) == "string", "char needs to be a string")

    local ret = {}
    for s in str:gmatch("[^" .. char .. "]+") do
        table.insert(ret, s)
    end
    return ret
end

---prints a table 1 level deep
---@param tbl table
M.inspect = function (tbl)
    for index, value in ipairs(tbl) do
        print(index, value)
    end
end

return M