import utils.Input

object Day6 {
    private fun String.firstUniqueNumOfCharsIndex(number: Int) =
        number + this.windowed(number).indexOfFirst { it.toSet().size === number }

    fun solvePart1() =
        Input.readInput(day = 6, isSample = false).firstUniqueNumOfCharsIndex(4)

    fun solvePart2() =
        Input.readInput(day = 6, isSample = false).firstUniqueNumOfCharsIndex(14)
}