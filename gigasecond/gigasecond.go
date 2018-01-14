
package gigasecond

// import path for the time package from the standard library
import "time"

const GIGASECOND = time.Duration(1e9) * time.Second

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(GIGASECOND)
}
