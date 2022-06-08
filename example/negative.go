//go:generate ../bin/go-enum -f=$GOFILE --marshal --lower --ptr --mustparse

package example

/* ENUM(
Unknown = -1,
Good,
Bad
).
*/
type Status int

/* ENUM(
Unknown = -5,
Good,
Bad,
Ugly
).
*/
type AllNegative int
