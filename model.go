package fxtabs

type jsonData struct {
	Windows []struct {
		Tabs []struct {
			Entries []struct {
				Title string
				URL   string
			}
			Index uint
		}
	}
}

// FirefoxTab represents a Firefox open tab.
type FirefoxTab struct {
	Title string
	URL   string
}
