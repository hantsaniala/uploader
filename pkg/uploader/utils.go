package uploader

// truncateMiddle shortens a string to a specified maximum length by retaining
// characters from the start and end and replacing the middle with ellipses ("...").
// If the string's length is less than or equal to the specified maximum length,
// the original string is returned unmodified. The function ensures that the
// resulting string, including the ellipses, does not exceed the specified maximum length.
func truncateMiddle(s string, maxLength int) string {
	// Return the original string if it's short enough
	if len(s) <= maxLength {
		return s
	}

	// Calculate how many characters to keep from the start and end
	keepLength := (maxLength - 3) / 2 // Subtract 3 for the dots
	start := s[:keepLength]
	end := s[len(s)-keepLength:]

	// Handle odd maxLength
	if (maxLength-3)%2 != 0 {
		end = s[len(s)-keepLength-1:]
	}

	return start + "..." + end
}
