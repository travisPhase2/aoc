import utils.Input

object Day10 {
    private fun Int.isSignalStrengthCycle() = this % 40 == 20

    private fun printCRTValue(index: Int, x: Int) {
        val adjustedIndex = index % 40
        if (adjustedIndex == 0) println()
        if ((adjustedIndex) in x - 1..x + 1) {
            print("#")
        } else {
            print(".")
        }
    }

    fun solvePart1(): Int {
        val instructions = Input.readInput(day = 10, isSample = false).lines()

        var x = 1
        var currentCycle = 1
        var signalStrength = 0

        instructions.forEach { instruction ->
            currentCycle += 1

            if (currentCycle.isSignalStrengthCycle()) {
                signalStrength += (currentCycle * x)
            }

            if (instruction != "noop") {
                currentCycle += 1
                x += instruction.substringAfter(" ").toInt()

                if (currentCycle.isSignalStrengthCycle()) {
                    signalStrength += (currentCycle * x)
                }
            }
        }

        return signalStrength
    }

    fun solvePart2() {
        val instructions = Input.readInput(day = 10, isSample = false).lines()
        var x = 1
        var currentCycle = 1

        instructions.forEach { instruction ->
            printCRTValue(currentCycle - 1, x)
            currentCycle += 1

            if (instruction != "noop") {
                printCRTValue(currentCycle - 1, x)
                currentCycle += 1
                x += instruction.substringAfter(" ").toInt()
            }
        }
    }
}