import utils.Input

object Day3 {
    fun solvePart1(): Int {
        return Input.readInput(day = 3, isSample = false).lines().sumOf { line ->
            val backpack = line
                .split("")
                .filter { it.isNotBlank() }
            val (firstCompartment, secondCompartment) = backpack
                .chunked(backpack.size / 2)
                .map { it.toSet() }

            firstCompartment.intersect(secondCompartment).sumOf {
                val code = it.toCharArray().first().code
                when {
                    code > 90 -> code - 96
                    code <= 90 -> (code + 26) - 64
                    else -> 0
                }
            }
        }
    }

    fun solvePart2(): Int {
        val groups = Input
            .readInput(day = 3, isSample = false)
            .lines()
            .chunked(3)

        return groups.sumOf { backpack ->
            val (first, second, third) = backpack.map { items ->
                items.split("")
                    .filter { it.isNotBlank() }
                    .toSet()
            }

            first.intersect(second).intersect(third).sumOf {
                val code = it.toCharArray().first().code
                when {
                    code > 90 -> code - 96
                    code <= 90 -> (code + 26) - 64
                    else -> 0
                }
            }
        }
    }
}