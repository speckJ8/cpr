package licenses

type License struct {
    Verbose string
    Simple string
}

type LicenseNoticeTemplateArgs struct {
    CurrentYear string
    Author string
    Email string
}

var Licenses = make(map[string]*License)
