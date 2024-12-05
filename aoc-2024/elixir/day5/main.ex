defmodule Aoc do
  def solve() do
    [rules, pages] =
      "input.txt"
      |> File.read!()
      |> String.split("\n\n")

    pages =
      pages
      |> String.split("\n", trim: true)
      |> Enum.map(fn line ->
        line |> String.split(",") |> Enum.map(&String.to_integer/1)
      end)

    correct_pages =
      Enum.reduce(pages, 0, fn page, acc ->
        sorted =
          Enum.sort_by(page, &Function.identity/1, fn l, r ->
            String.contains?(rules, "#{l}|#{r}")
          end)

        middle = Enum.at(sorted, div(length(sorted), 2))

        if page == sorted do
          acc + middle
        else
          acc
        end
      end)

    fixed_pages =
      Enum.reduce(pages, 0, fn page, acc ->
        sorted =
          Enum.sort_by(page, &Function.identity/1, fn l, r ->
            String.contains?(rules, "#{l}|#{r}")
          end)

        middle = Enum.at(sorted, div(length(sorted), 2))

        if page == sorted do
          acc
        else
          acc + middle
        end
      end)

    {correct_pages, fixed_pages}
  end
end
