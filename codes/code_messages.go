package codes

import (
	"net/http"

	"github.com/adiatma85/own-go-sdk/language"
)

type Message struct {
	StatusCode int
	TitleEN    string
	TitleID    string
	BodyEN     string
	BodyID     string
}

var (
	// 4xx
	ErrMsgBadRequest = Message{
		StatusCode: http.StatusBadRequest,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusBadRequest),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusBadRequest),
		BodyEN:     "Invalid input. Please validate your input.",
		BodyID:     "Input data tidak valid. Mohon cek kembali input data anda.",
	}
	ErrMsgUnauthorized = Message{
		StatusCode: http.StatusUnauthorized,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusUnauthorized),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusUnauthorized),
		BodyEN:     "Unauthorized access. You are not authorized to access this resource.",
		BodyID:     "Akses ditolak. Anda tidak memiliki izin untuk mengakses laman ini.",
	}
	ErrMsgInvalidToken = Message{
		StatusCode: http.StatusUnauthorized,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusUnauthorized),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusUnauthorized),
		BodyEN:     "Invalid token. Please renew your session by reloading.",
		BodyID:     "Token tidak valid. Mohon perbarui sesi dengan memuat ulang laman.",
	}
	ErrMsgRefreshTokenExpired = Message{
		StatusCode: http.StatusUnauthorized,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusUnauthorized),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusUnauthorized),
		BodyEN:     "Session refresh token has expired. Please renew your session by reloading.",
		BodyID:     "Token pembaruan sudah tidak berlaku. Mohon perbarui sesi dengan memuat ulang laman.",
	}
	ErrMsgAccessTokenExpired = Message{
		StatusCode: http.StatusUnauthorized,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusUnauthorized),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusUnauthorized),
		BodyEN:     "Session access token has expired. Please renew your session by reloading.",
		BodyID:     "Token akses sudah tidak berlaku. Mohon perbarui sesi dengan memuat ulang laman.",
	}
	ErrMsgForbidden = Message{
		StatusCode: http.StatusForbidden,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusForbidden),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusForbidden),
		BodyEN:     "Forbidden. You don't have permission to access this resource.",
		BodyID:     "Terlarang. Anda tidak memiliki izin untuk mengakses laman ini.",
	}
	ErrMsgRevokeRefreshTokenFailed = Message{
		StatusCode: http.StatusInternalServerError,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusInternalServerError),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusInternalServerError),
		BodyEN:     "Failed revoking refresh token.",
		BodyID:     "Gagal mencabut refresh token.",
	}
	ErrMsgNotFound = Message{
		StatusCode: http.StatusNotFound,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusNotFound),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusNotFound),
		BodyEN:     "Record does not exist. Please validate your input or contact the administrator.",
		BodyID:     "Data tidak ditemukan. Mohon cek kembali input data anda atau hubungi administrator.",
	}
	ErrMsgContextTimeout = Message{
		StatusCode: http.StatusRequestTimeout,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusRequestTimeout),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusRequestTimeout),
		BodyEN:     "Request time has been exceeded.",
		BodyID:     "Waktu permintaan telah habis.",
	}
	ErrMsgConflict = Message{
		StatusCode: http.StatusConflict,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusConflict),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusConflict),
		BodyEN:     "Record has existed. Please validate your input or contact the administrator.",
		BodyID:     "Data sudah ada. Mohon cek kembali input data anda atau hubungi administrator.",
	}
	ErrMsgTooManyRequest = Message{
		StatusCode: http.StatusTooManyRequests,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusTooManyRequests),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusTooManyRequests),
		BodyEN:     "Too many requests. Please wait and try again after a few moments.",
		BodyID:     "Terlalu banyak permintaan. Mohon tunggu dan coba lagi setelah beberapa saat.",
	}

	// 5xx
	ErrMsgInternalServerError = Message{
		StatusCode: http.StatusInternalServerError,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusInternalServerError),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusInternalServerError),
		BodyEN:     "Internal server error. Please contact the administrator.",
		BodyID:     "Terjadi kendala pada server. Mohon hubungi administrator.",
	}
	ErrMsgNotImplemented = Message{
		StatusCode: http.StatusNotImplemented,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusNotImplemented),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusNotImplemented),
		BodyEN:     "Not Implemented. Please contact the administrator.",
		BodyID:     "Layanan tidak tersedia. Mohon hubungi administrator.",
	}
	ErrMsgServiceUnavailable = Message{
		StatusCode: http.StatusServiceUnavailable,
		TitleEN:    language.HTTPStatusText(language.English, http.StatusServiceUnavailable),
		TitleID:    language.HTTPStatusText(language.Indonesian, http.StatusServiceUnavailable),
		BodyEN:     "Service is unavailable. Please contact the administrator.",
		BodyID:     "Layanan sedang tidak tersedia. Mohon hubungi administrator.",
	}
)
