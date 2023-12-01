import utils.Input
import kotlin.math.absoluteValue
import kotlin.math.sign

private data class Point(val x: Int = 0, val y: Int = 0) {
    fun move(direction: Char): Point =
        when(direction) {
            'U' -> copy(y = y - 1)
            'D'-> copy(y = y + 1)
            'L'-> copy(x = x - 1)
            'R'-> copy(x = x + 1)
            // this should never happen
            else -> { Point() }
        }

    fun moveToward(other: Point): Point =
        Point(
            (other.x - x).sign + x,
            (other.y - y).sign + y,
        )

    fun touches(other: Point): Boolean =
        (x - other.x).absoluteValue <= 1 && (y - other.y).absoluteValue <= 1
}

object Day9 {
    fun solvePart1(): Int {
        val input = Input.readInput(day = 9, isSample = false).lines()
        val moves = parseInput(input)
        return countTailVisits(moves, knotSize = 2)
    }

    fun solvePart2(): Int {
        val input = Input.readInput(day = 9, isSample = false).lines()
        val moves = parseInput(input)
        return countTailVisits(moves, knotSize = 10)
    }

    private fun countTailVisits(moves: String, knotSize: Int): Int {
        // each rope has 2 knots
        val rope = Array(knotSize) { Point() }
        val tailVisits = mutableSetOf(Point())

        moves.forEach { direction ->
            // move direction on rope
            rope[0] = rope[0].move(direction)

            rope.indices.windowed(2, 1) { (head, tail) ->
                // if head and tail do not touch move tail towards head
                if (!rope[head].touches(rope[tail])) {
                    rope[tail] = rope[tail].moveToward(rope[head])
                }
            }

            // update visit count
            tailVisits += rope.last()
        }

        return tailVisits.size
    }

    private fun parseInput(input: List<String>): String =
        input.joinToString("") { row ->
            val direction = row.substringBefore(" ")
            val numberOfMoves = row.substringAfter(" ").toInt()
            direction.repeat(numberOfMoves)
        }
}
