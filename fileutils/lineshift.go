package fileutils

// GetLineending returns control codes for lineshifts in a given file.
// The first encountered line of at least 2 characters
// in length is used to safely determine whether it uses LF or CRLF.
func GetLineending(filename string) string {
    var lineending string
    reader := newBufferedReader(filename)
    for {
        line, err := reader.ReadString(0x0A)
        check(err)
        if len(line) > 1 {
            // Get the first line that is at least 2 characters long...
            lineending = line[len(line)-2:] // ... because the lineending is always at most 2 chars long
            break
        }
    }
    if lineending[0] > 0x1F {
        // If the first character is above 0x1F, it is not a control character,
        // and the last character is always bound to be a \n as that is what the line split by
        return lineending[1:]
    } else {
        return lineending
    }
}
