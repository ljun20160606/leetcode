package main

type Language struct {
	Lang      string
	Extension string
}

var (
	Cpp = Language{
		Lang:      "C++",
		Extension: "cpp",
	}
	C = Language{
		Lang:      "C",
		Extension: "c",
	}
	Java = Language{
		Lang:      "Java",
		Extension: "java",
	}
	Python = Language{
		Lang:      "Python",
		Extension: "py",
	}
	Python3 = Language{
		Lang:      "Python3",
		Extension: "py",
	}
	JavaScript = Language{
		Lang:      "JavaScript",
		Extension: "js",
	}
	Go = Language{
		Lang:      "Go",
		Extension: "go",
	}
)
