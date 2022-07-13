package conf

import "os"

const (
	// EnvDev 開発環境
	EnvDev = "dev"
	// EnvTest テスト環境
	EnvTest = "test"
	// EnvQa QA環境
	EnvQa = "qa"
	// EnvProd 本番環境
	EnvProd = "prod"
)

//  Env 環境情報の取得
func Env() string {
	switch os.Getenv("GO_ENV") {
	case EnvDev:
		return EnvDev
	case EnvTest:
		return EnvTest
	case EnvQa:
		return EnvQa
	case EnvProd:
		return EnvProd
	default:
		return EnvDev
	}
}

// IsDev 開発環境であるか
func IsDev() bool {
	return Env() == EnvDev
}
