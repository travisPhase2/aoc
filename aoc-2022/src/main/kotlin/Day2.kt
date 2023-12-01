import utils.Input

object Day2 {
    fun solvePart1(): Int {
        val answersMap: HashMap<Pair<String, String>, Int> = hashMapOf(
            Pair("A", "X") to 3 + 1,
            Pair("A", "Y") to 6 + 2,
            Pair("A", "Z") to 0 + 3,

            Pair("B", "X") to 0 + 1,
            Pair("B", "Y") to 3 + 2,
            Pair("B", "Z") to 6 + 3,

            Pair("C", "X") to 6 + 1,
            Pair("C", "Y") to 0 + 2,
            Pair("C", "Z") to 3 + 3,
        )

        return Input.readInput(2, isSample = false).lines().sumOf {
            val (you, me) = it.split(" ")
            answersMap[Pair(you, me)]!!
        }
    }

    fun solvePart2(): Int {
        val answersMap: HashMap<Pair<String, String>, Int> = hashMapOf(
            Pair("A", "X") to 0 + 3,
            Pair("A", "Y") to 3 + 1,
            Pair("A", "Z") to 6 + 2,

            Pair("B", "X") to 0 + 1,
            Pair("B", "Y") to 3 + 2,
            Pair("B", "Z") to 6 + 3,

            Pair("C", "X") to 0 + 2,
            Pair("C", "Y") to 3 + 3,
            Pair("C", "Z") to 6 + 1,
        )

        return Input.readInput(2, isSample = false).lines().sumOf {
            val (you, me) = it.split(" ")
            answersMap[Pair(you, me)]!!
        }
    }
}
