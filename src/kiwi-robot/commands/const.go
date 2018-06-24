package commands

const (
	rNamePlaceX   = `place_x`
	rNamePlaceY   = `place_y`
	rNamePlaceDir = `place_dir`
	rNameCmd      = `cmd`

	DIR_EAST  = `EAST`
	DIR_NORTH = `NORTH`
	DIR_WEST  = `WEST`
	DIR_SOUTH = `SOUTH`

	CMD_MOVE   = `MOVE`
	CMD_LEFT   = `LEFT`
	CMD_RIGHT  = `RIGHT`
	CMD_REPORT = `REPORT`

	placeRegexp        = `(?:PLACE (?P<` + rNamePlaceX + `>\d+),(?P<` + rNamePlaceY + `>\d+),(?P<` + rNamePlaceDir + `>` + DIR_EAST + `|` + DIR_NORTH + `|` + DIR_WEST + `|` + DIR_SOUTH + `))`
	mainRegexpTemplate = `^(?P<` + rNameCmd + `>(?:` + CMD_MOVE + `|` + CMD_LEFT + `|` + CMD_RIGHT + `|` + CMD_REPORT + `|` + placeRegexp + `))\s*`
)
