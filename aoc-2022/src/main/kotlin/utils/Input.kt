package utils

import java.io.File
import kotlin.io.path.Path
import kotlin.io.path.forEachLine

object Input {
    fun readInputByLine(day: Int, isSample: Boolean = false, work: (str: String) -> Unit) {
        val fileName = "src/main/resources/" + (if(isSample) "sample-" else "") + "input-day-$day.txt"
        Path(fileName).forEachLine {
            work(it) /* make it, do it, makes us, harder, better, faster, stronger */ }
    }

    fun readInput(day: Int, isSample: Boolean = false): String {
        val fileName = "src/main/resources/" + (if(isSample) "sample-" else "") + "input-day-$day.txt"
        return File(fileName).readText()
    }
}
