
defmodule Aoc do
  def solve() do
    equations =
      "input.txt"
      |> File.read!()
      |> parse_input()
      |> Enum.filter(&is_valid?/1)
      |> Enum.map(fn {result, numbers} -> result end)
      |> Enum.sum()
  end

  defp parse_input(input) do
    input
    |> String.split("\n")
    |> Enum.map(fn line ->
      [result, numbers] = String.split(line, ": ")
      numbers = String.split(numbers, " ") |> Enum.map(&String.to_integer/1)
      {String.to_integer(result), numbers}
    end)
  end

  def is_valid?({result, numbers}), do: check(result, numbers)

  defp check(result, numbers, total \\ nil)
  defp check(result, [], total), do: total == result
  defp check(result, [number | tail], nil), do: check(result, tail, number)
  defp check(result, _numbers, total) when total > result, do: false

  defp check(result, [number | tail], total) do
    check(result, tail, total + number) or check(result, tail, total * number)
  end
end
