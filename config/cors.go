package config

var allowedOrigins = []string{
	"https://indrariksa.github.io/",
	"http://localhost:5173/",
	"http://localhost:5174/",
	"http://localhost:5175/",
	"http://127.0.0.1:8088/",
	"https://charming-starlight-a44ff5.netlify.app/",
	"https://fe-pemrog3.vercel.app/",
	"http://localhost:4173/",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
