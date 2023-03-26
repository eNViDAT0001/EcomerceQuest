package wrap_cloudinary

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
)

var (
	mediaServer *cloudinary.Cloudinary
	vp          = wrap_viper.GetViper()
)

func GetMediaServer() *cloudinary.Cloudinary {
	if mediaServer != nil {
		return mediaServer
	}

	mediaServer = newCloudinaryServer(
		vp.GetString("CLOUDINARY.NAME"),
		vp.GetString("CLOUDINARY.API_KEY"),
		vp.GetString("CLOUDINARY.API_SECRET"),
		vp.GetBool("CLOUDINARY.SECURE"))

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
