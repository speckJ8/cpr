package langs

import "strings"

type Lang struct {
    WrapInComments func(string) string
}

var Langs map[string]*Lang

func init() {
    var cWrap = Lang { WrapInComments: c }
    Langs = map[string]*Lang {
        "go":    &cWrap,
        "c":     &cWrap,
        "h":     &cWrap,
        "cpp":   &cWrap,
        "cc":    &cWrap,
        "hpp":   &cWrap,
        "rs":    &cWrap,
        "js":    &cWrap,
        "ts":    &cWrap,
        "java":  &cWrap,
        "py":    &Lang { WrapInComments: python },
    }
}

func c(contents string) string {
    var lines = strings.Split(contents, "\n")
    var comment = "/**\n"
    for _, line := range lines {
        comment += " * " + line + "\n"
    }
    comment += " */\n\n"
    return comment
}

func python(contents string) string {
    var lines = strings.Split(contents, "\n")
    var comment = "\"\"\"\n"
    for _, line := range lines {
        comment += "" + line + "\n"
    }
    comment += "\"\"\"\n\n"
    return comment
}
