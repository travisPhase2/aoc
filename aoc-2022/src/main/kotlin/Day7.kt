import utils.Input

object Day7 {

    private abstract class Child {
        abstract val name: String
    }

    private data class File(override val name: String, val size: Int) : Child()

    private data class Directory(override val name: String) : Child() {
        private val children = mutableListOf<Child>()

        fun addChildrenToDirectory(children: List<Child>) {
            children.forEach { this.children.add(it) }
        }

        fun containsChildDirectory(dest: String): Boolean {
            return this.children.any { it.name == dest }
        }
    }

    private class Session {
        var fileSystem = FileSystem()

        private var cwd = "/"
        private var currentDirectory: Directory = Directory(this.cwd)

        private fun changeDirectory(argument: String?) {
            argument?.let {
                when {
                    it == "/" -> this.cwd = "/"
                    it == ".." -> {
                        val destination = this.cwd.reversed().substringAfter("/")
                        if (destination.trim().isNullOrEmpty()) {
                            this.cwd = "/"
                        } else {
                            this.cwd = destination.reversed()
                        }
                        this.currentDirectory = this.fileSystem.getDirectory(this.cwd)
                    }
                    this.currentDirectory.containsChildDirectory(it) -> {
                        if (this.cwd == "/") {
                            this.cwd = "${this.cwd}$it"
                        } else {
                            this.cwd = "${this.cwd}/$it"
                        }
                    }
                }
            }
        }

        private fun createFiles(files: List<String>) {
            val children = files.map {
                val (info, name) = it.split(" ")
                if (info.trim() == "dir") {
                    Directory(name)
                } else {
                    File(name, info.toInt())
                }
            }
            fileSystem.addFiles(this.cwd, children)
            this.currentDirectory = this.fileSystem.getDirectory(this.cwd)
        }

        fun initializeFileSystem(input: List<String>) {
            input.withIndex().forEach { (idx, output) ->
                val command = output.split(" ")
                when (command.size) {
                    // cd
                    3 -> this.changeDirectory(command[2])
                    // ls
                    2 -> {
                        val files = mutableListOf<String>()
                        for (i in idx + 1 until input.size) {
                            val currentLine = input[i]
                            if (!currentLine.startsWith("$")) {
                                files.add(currentLine)
                            } else if (currentLine.startsWith("$")) {
                                break
                            }
                        }
                        this.createFiles(files)
                    }
                }
            }

        }
    }

    private class FileSystem {

        private val files = mutableMapOf<String, List<Child>>()

        fun addFiles(dir: String, files: List<Child>) {
            this.files[dir] = files
        }

        fun getDirectory(dir: String): Directory {
            val directory = Directory(dir)
            directory.addChildrenToDirectory(this.files[dir]!!)
            return directory
        }

        fun getAllDirectorySizes(): List<Map<String, Int>> {
            val sizes: List<Map<String, Int>> = this.files.keys.map { directory ->
                val sum = this.getDirectorySize(directory)

                mapOf(directory to sum)
            }

            return sizes
        }

        private fun getDirectorySize(dir: String): Int {
            return this.files.keys
                .filter { it.startsWith(dir) }
                .sumOf { path ->
                    this.files[path]!!.sumOf {
                        when (it) {
                            is File -> it.size
                            else -> 0
                        }
                    }
                }
        }
    }

    fun solvePart1(): Int {
        val input = Input.readInput(day = 7, isSample = true).lines()
        val session = Session()

        session.initializeFileSystem(input)

        val fileSizes = session.fileSystem.getAllDirectorySizes()

        return fileSizes
            .flatMap { it.values }
            .filter { it <= 100000 }
            .sum()
    }

    fun solvePart2(): Int {
        val input = Input.readInput(day = 7, isSample = false).lines()
        val session = Session()
        session.initializeFileSystem(input)

        val fileSizes = session.fileSystem.getAllDirectorySizes()

        val rootSize = fileSizes.first().values.first()
        val totalSpace = 70000000
        val spaceRequired = 30000000
        val spaceLeft = totalSpace - rootSize
        val spaceNeeded = spaceRequired - spaceLeft

        return fileSizes
            .flatMap { it.values }
            .filter { it >= spaceNeeded }
            .min()
    }
}