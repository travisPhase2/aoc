defmodule Aoc do
  @mul_regex ~r/mul\((\d{1,3}),(\d{1,3})\)/
  @do_dont_regex ~r/mul\((-?\d+),(-?\d+)\)|do(?:n't)?\(\)/

  def solve() do
    input = "input.txt" |> File.read!()

    part_1 =
      Regex.scan(@mul_regex, input)
      |> Enum.reduce(0, fn [_text, a, b], total ->
        product = String.to_integer(a) * String.to_integer(b)
        total + product
      end)

    part_2 =
      Regex.scan(@do_dont_regex, input)
      |> Enum.reduce({0, :do}, fn
        ["don't()"], {count, _status} ->
          {count, :dont}

        ["do()"], {count, _status} ->
          {count, :do}

        [_text, a, b], {count, :do} ->
          product = String.to_integer(a) * String.to_integer(b)
          {product + count, :do}

        [_text, _a, _b], {count, :dont} ->
          {count, :dont}
      end)
      |> elem(0)

    {part_1, part_2}
  end
end
