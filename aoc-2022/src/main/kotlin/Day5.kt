import utils.Input
import java.util.Stack

object Day5 {
    fun solvePart1(): String {
        val (stackInput, instructionInput) = Input.readInput(day = 5, isSample = false).split("\n\n")
        val stackList = populateStackList(stackInput)

        instructionInput.lines().forEach { line ->
            val instructions = line
                .split(" ")
                .mapNotNull { it.toIntOrNull() }

            val (amount, source, destination) = instructions
            val sourceStack = stackList[source - 1]
            val destinationStack = stackList[destination - 1]

            for (i in 1..amount) {
                val newValue = sourceStack.pop()
                destinationStack.push(newValue)
            }
        }

        return stackList.joinToString { it.peek() }.replace(", ", "")
    }

    fun solvePart2(): String {
        val (stackInput, instructionInput) = Input.readInput(day = 5, isSample = false).split("\n\n")
        val stackList = populateStackList(stackInput)

        instructionInput.lines().forEach { line ->
            val instructions = line
                .split(" ")
                .mapNotNull { it.toIntOrNull() }

            val (amount, source, destination) = instructions
            val sourceStack = stackList[source - 1]
            val destinationStack = stackList[destination - 1]

            val res = mutableListOf<String>()
            for (i in 1..amount) {
                res.add(sourceStack.pop())
            }
            res.reversed().forEach { destinationStack.add(it) }
        }

        return stackList.joinToString { it.pop() }.replace(", ", "")
    }

    private fun populateStackList(stackInput: String): List<Stack<String>> {
        val stackList = mutableListOf<Stack<String>>()
        stackInput.lines().reversed().withIndex().forEach { (idx, it) ->
            when {
                idx == 0 -> {
                    val stacksToCreate = it
                        .split("")
                        .filter { char -> char.isNotEmpty() }
                        .maxOfOrNull { it }
                    for (i in 1..stacksToCreate?.toInt()!!) {
                        stackList.add(Stack<String>())
                    }
                }
                idx >= 1 -> {
                    val letters = mutableListOf<String>()
                    for (i in 1 .. it.length step 4) {
                        letters.add(it[i].toString())
                    }
                    letters.withIndex().forEach { (idx, it) ->
                        if (it.trim().isNotEmpty()) {
                            stackList[idx].push(it)
                        }
                    }
                }
            }
        }

        return stackList
    }
}