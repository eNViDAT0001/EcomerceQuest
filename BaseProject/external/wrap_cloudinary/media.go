package wrap_cloudinary

type MediaFolderType int

const (
	Product MediaFolderType = iota
	ProductDescriptions
	Category
	Comment
)

var presets = []string{"Product", "ProductDescriptions", "Category", "Comment"}

func GetMediaFolder(t MediaFolderType) string {
	return presets[t]
}
