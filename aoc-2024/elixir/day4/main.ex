defmodule Aoc do
  @directions [{1, 0}, {0, 1}, {1, 1}, {1, -1}, {-1, 0}, {0, -1}, {-1, -1}, {-1, 1}]
  @diagonal_dirs [{1, -1}, {1, 1}, {-1, 1}, {-1, -1}]

  def solve() do
    grid =
      "input.txt"
      |> parse_grid()

    xmas_occurences =
      grid
      |> Enum.filter(fn {_, char} -> char == "X" end)
      |> Enum.reduce(0, fn {{x, y}, _}, acc ->
        acc + find_xmas_occurences(grid, {x, y})
      end)

    mas_occurences =
      grid
      |> Enum.filter(fn {_, char} -> char == "A" end)
      |> Enum.reduce(0, fn {{x, y}, _}, acc -> acc + find_diagonals(grid, {x, y}) end)

    {xmas_occurences, mas_occurences}
  end

  defp parse_grid(file_name) do
    file_name
    |> File.read!()
    |> String.codepoints()
    |> Enum.reduce({%{}, 0, 0}, fn char, {acc, x, y} ->
      case char do
        "\n" -> {acc, 0, y + 1}
        char -> {Map.put(acc, {x, y}, char), x + 1, y}
      end
    end)
    |> elem(0)
  end

  defp find_xmas_occurences(grid, {x, y}) do
    Enum.reduce(@directions, 0, fn {dx, dy}, acc ->
      ~w/M A S/
      |> Enum.reduce_while({nil, {x, y}}, fn char, {_, {x, y}} ->
        coord_to_check = {x + dx, y + dy}

        case Map.get(grid, coord_to_check) do
          ^char -> {:cont, {true, coord_to_check}}
          _ -> {:halt, false}
        end
      end)
      |> case do
        {true, _} -> acc + 1
        _ -> acc
      end
    end)
  end

  defp find_diagonals(grid, {x, y}) do
    @diagonal_dirs
    |> Enum.map(fn {dx, dy} -> Map.get(grid, {x + dx, y + dy}) end)
    |> case do
      ~w/S S M M/ -> 1
      ~w/M M S S/ -> 1
      ~w/S M M S/ -> 1
      ~w/M S S M/ -> 1
      _ -> 0
    end
  end
end
