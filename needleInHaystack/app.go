package needleinhaystack

import (
    s "strings"
)

func strStr(haystack string, needle string) int {
     return s.Index(haystack, needle)
}