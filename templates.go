package main

import (
	"fmt"
)

func PageWrapper(contents []byte) []byte {
	return []byte(fmt.Sprintf(`
<html>
  <head>
  </head>
  <body>%s</body>
</html>
`, contents))
}
