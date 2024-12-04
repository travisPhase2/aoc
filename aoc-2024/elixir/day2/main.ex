defmodule Helpers do
  def is_ascending?(report, max_distance) do
    case report do
      [a, b | rest] when a < b and abs(a - b) <= max_distance ->
        is_ascending?([b | rest], max_distance)

      [_] ->
        true

      _ ->
        false
    end
  end

  def is_descending?(report, max_distance) do
    case report do
      [a, b | rest] when a > b and abs(a - b) <= max_distance ->
        is_descending?([b | rest], max_distance)

      [_] ->
        true

      _ ->
        false
    end
  end

  def is_safe?(report, max_distance),
    do: is_ascending?(report, max_distance) or is_descending?(report, max_distance)
end

defmodule Aoc do
  def solve() do
    reports =
      "input.txt"
      |> File.read!()
      |> String.split("\n", trim: true)
      |> Enum.map(fn report ->
        report
        |> String.split()
        |> Enum.map(&String.to_integer/1)
      end)

    part_1 = reports |> Enum.count(&Helpers.is_safe?(&1, 3))

    part_2 =
      reports
      |> Enum.count(fn report ->
        report
        |> Enum.with_index()
        |> Enum.map(fn {_, index} ->
          report
          |> List.delete_at(index)
          |> Helpers.is_safe?(3)
        end)
        |> Enum.any?()
      end)

    {part_1, part_2}
  end
end
