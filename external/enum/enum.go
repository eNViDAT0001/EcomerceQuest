package enum

type MediaType string

const (
	MediaFullImage MediaType = "IMAGE_FULL"
	MediaItemImage MediaType = "IMAGE_ITEM"
	MediaVideo     MediaType = "VIDEO"
	MediaImage     MediaType = "IMAGE"
)
