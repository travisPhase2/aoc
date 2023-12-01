import utils.Input

object Day1 {
    private fun parseInput(input: String) =
        input.split("\n\n").map { elf ->
            elf.lines().map { it.toInt() }
        }


    fun solvePart1(): Int {
        return parseInput(Input.readInput(day = 1, isSample = false))
            .maxOf { it.sum() }
    }

    fun solvePart2(): Int {
        return parseInput(Input.readInput(day = 1, isSample = false))
            .map { it.sum() }
            .sortedDescending()
            .take(3)
            .sum()
    }
}