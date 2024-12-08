defmodule Vec2d do
  def add({x1, y1}, {x2, y2}), do: {x1 + x2, y1 + y2}
  def rotate_clockwise({x, y}), do: {-y, x}
end

defmodule Aoc do
  def solve() do
    input =
      "sample-input.txt"
      |> File.read!()
      |> String.split("\n", trim: true)
      |> Enum.map(&String.to_charlist/1)

    wall_positions = find_walls(input)

    guard_y = Enum.find_index(input, &(?^ in &1))
    guard_x = Enum.find_index(Enum.at(input, guard_y), &(&1 == ?^))
    guard_pos = {guard_x, guard_y}

    max_x = (hd(input) |> length()) - 1
    max_y = length(input) - 1

    distinct_positions =
      {guard_pos, {0, -1}}
      |> Stream.iterate(&move_direction(&1, wall_positions))
      |> Enum.reduce_while(MapSet.new(), fn {{x, y}, _direction} = acc, seen ->
        cond do
          x > max_x or y > max_y -> {:halt, seen}
          acc in seen -> {:halt, seen}
          true -> {:cont, MapSet.put(seen, acc)}
        end
      end)
      |> Enum.map(&elem(&1, 0))
      |> Enum.uniq()

    # endless_loops =
    #   distinct_positions
    #   |> List.delete(guard_pos)
    #   |> Task.async_stream(fn new_wall ->
    #     walls = MapSet.put(wall_positions, new_wall)

    #     {guard_pos, {0, -1}}
    #     |> Stream.iterate(&move_direction(&1, walls))
    #     |> Enum.reduce_while(MapSet.new(), fn {{x, y}, _direction} = acc, seen ->
    #       cond do
    #         x > max_x or y > max_y -> {:halt, 0}
    #         acc in seen -> {:halt, 1}
    #         true -> {:cont, MapSet.put(seen, acc)}
    #       end
    #     end)
    #   end, ordered: false, timeout: :infinity)
    #   |> Stream.map(&elem(&1, 1))
    #   |> Enum.sum()

    %{
      part1: distinct_positions |> length(),
      part2: "i give up"
    }
  end

  defp find_walls(rows) when is_list(rows) do
    rows
    |> Enum.with_index()
    |> Enum.flat_map(fn {row, y} ->
      row
      |> Enum.with_index()
      |> Enum.filter(fn {char, _x} -> char == ?# end)
      |> Enum.map(fn {_char, x} -> {x, y} end)
    end)
    |> MapSet.new()
  end

  defp move_direction({current_pos, direction}, wall_positions) do
    new_pos = Vec2d.add(current_pos, direction)

    if new_pos in wall_positions do
      {current_pos, Vec2d.rotate_clockwise(direction)}
    else
      {new_pos, direction}
    end
  end
end
