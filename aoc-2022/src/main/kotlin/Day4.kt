import utils.Input

object Day4 {
    private fun parseInput() = Input.readInput(day = 4, isSample = false)
        .lines()
        .map { it.split(",") }

    private fun String.toIntRange() =
        this.substringBefore("-").toInt() .. this.substringAfter("-").toInt()

    private infix fun IntRange.contains(other: IntRange) =
        this.contains(other.first) && this.contains(other.last)

    private infix fun IntRange.intersect(other: IntRange) =
        (this.toSet()).intersect(other.toSet())

    fun solvePart1() = parseInput().filter { (firstGroup, secondGroup) ->
        firstGroup.toIntRange() contains secondGroup.toIntRange()
            || secondGroup.toIntRange() contains firstGroup.toIntRange()
    }.size

    fun solvePart2() = parseInput().filter { (firstGroup, secondGroup) ->
        (firstGroup.toIntRange() intersect secondGroup.toIntRange()).isNotEmpty()
    }.size
}
