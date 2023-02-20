package wrap_cloudinary

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_viper"
)

var (
	mediaServer *cloudinary.Cloudinary
	vp          = wrap_viper.GetViper()
)

func GetMediaServer() *cloudinary.Cloudinary {
	if mediaServer != nil {
		return mediaServer
	}

	//mediaServer = newCloudinaryServer(
	//	vp.GetString("CLOUDINARY_NAME"),
	//	vp.GetString("CLOUDINARY_API_KEY"),
	//	vp.GetString("CLOUDINARY_API_SECRET"))
	mediaServer = newCloudinaryServer(
		"damzcas3k",
		"332611833886276",
		"CApeUOYt-JIgUSj9LvSAs6rO610",
		true)
	return mediaServer
}
func newCloudinaryServer(cloud, key, secret string, secure bool) *cloudinary.Cloudinary {

	cld, err := cloudinary.NewFromParams(cloud, key, secret)
	if err != nil {
		panic(err)
	}

	cld.Config.URL.Secure = secure
	return cld
}
