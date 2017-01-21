package data

import "time"

// see http://tech.townsourced.com/post/anatomy-of-a-go-web-app/#the-data-package
type Version struct {
	VersionTag string
	created_at time.Time
	updated_at time.Time
}
