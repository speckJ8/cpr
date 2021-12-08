package licenses

import "strings"

var GPL3Simple = `Copyright (c) {{.CurrentYear}} {{.Author}} <{{.Email}}>

SPDX-License-Identifier: GPL-3.0`

var GPL3Verbose =`Copyright (c) {{.CurrentYear}} {{.Author}} <{{.Email}}>

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free Software
Foundation, version 3.

This program is distributed in the hope that it will be useful, but WITHOUT
ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with
this program. If not, see <https://www.gnu.org/licenses/>.
`

var gpl3 = License {
    Verbose: GPL3Verbose,
    Simple: GPL3Simple,
}

func init() {
    Licenses["gpl3"] = &gpl3
}

func IsGPL3(license string) *License {
    if strings.Contains(license, "GPL") ||
        strings.Contains(license, "GNU GENERAL PUBLIC LICENSE") ||
        strings.Contains(license, "GNU General Pulic License") {
        if strings.Contains(license, "v3") ||
            strings.Contains(license, "3.0") ||
            strings.Contains(license, "Version 3") ||
            strings.Contains(license, "version 3") {
            return &gpl3
        }
    }
    return nil
}
