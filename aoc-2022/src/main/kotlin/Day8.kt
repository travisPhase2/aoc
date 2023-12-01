import utils.Input

object Day8 {

    fun countVisibleTrees(input: List<List<Int>>): Int {
        val rows = input.size
        val columns = input[0].size
        val visibleTrees = Array(rows) { BooleanArray(columns) }

        for (i in 0 until rows) {
            var tallest = -1
            for (j in 0 until columns) {
                if (input[i][j] > tallest) {
                    tallest = input[i][j]
                    visibleTrees[i][j] = true
                }
            }
            tallest = -1
            for (j in columns - 1 downTo 0) {
                if (input[i][j] > tallest) {
                    tallest = input[i][j]
                    visibleTrees[i][j] = true
                }
            }
        }
        for (i in 0 until columns) {
            var tallest = -1
            for (j in 0 until rows) {
                if (input[j][i] > tallest) {
                    tallest = input[j][i]
                    visibleTrees[j][i] = true
                }
            }
            tallest = -1
            for (j in rows - 1 downTo 0) {
                if (input[j][i] > tallest) {
                    tallest = input[j][i]
                    visibleTrees[j][i] = true
                }
            }
        }
        return visibleTrees.sumOf { it.count { it } }
    }

    fun highestScenicScore(input: List<List<Int>>): Int {
        val rows = input.size
        val columns = input[0].size
        val visibleTrees = Array(rows) { BooleanArray(columns) }

        var score = 0

        for (i in 0 until rows) {
            for (j in 0 until columns) {
                var d1 = j
                for (k in j - 1 downTo 0) {
                    if (input[i][k] >= input[i][j]) {
                        d1 = j - k
                        break
                    }
                }
                var d2 = columns - j - 1
                for (k in j + 1 until columns) {
                    if (input[i][k] >= input[i][j]) {
                        d2 = k - j
                        break
                    }
                }
                var d3 = i
                for (k in i - 1 downTo 0) {
                    if (input[k][j] >= input[i][j]) {
                        d3 = i - k
                        break
                    }
                }
                var d4 = rows - i - 1
                for (k in i + 1 until rows) {
                    if (input[k][j] >= input[i][j]) {
                        d4 = k - i
                        break
                    }
                }
                val f = d1 * d2 * d3 * d4
                if (f > score) score = f
            }
        }

        return score
    }

    fun solvePart1(): Int {
        val input = Input.readInput(day = 8, isSample = false).lines().map { line ->
            line.map { it.digitToInt() }
        }

        return countVisibleTrees(input)
    }

    fun solvePart2(): Int {
        val input = Input.readInput(day = 8, isSample = false).lines().map { line ->
            line.map { it.digitToInt() }
        }

        return highestScenicScore(input)
    }
}