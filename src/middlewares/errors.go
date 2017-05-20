package middlewares

/*
USAGE:
	import (
		mw "github.com/djviolin/lanti-mvc/src/middlewares"
	)

	mw.CheckErr(err)
*/

// CheckErr : error checking middleware
func CheckErr(err error) {
	if err != nil {
		//log.Fatal("ERROR: ", err)
		panic(err)
	}
}
