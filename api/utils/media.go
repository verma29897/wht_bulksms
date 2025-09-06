// utils/media.go
package utils

import "strings"

func GetMediaFormat(extension string) string {
	mediaFormats := map[string]string{
		"jpg": "image/jpeg", "jpeg": "image/jpeg", "png": "image/png",
		"gif": "image/gif", "bmp": "image/bmp", "svg": "image/svg+xml",
		"mp4": "video/mp4", "avi": "video/x-msvideo", "mov": "video/quicktime",
		"flv": "video/x-flv", "mkv": "video/x-matroska", "mp3": "audio/mpeg",
		"aac": "audio/aac", "ogg": "audio/ogg", "wav": "audio/wav",
		"pdf": "application/pdf", "doc": "application/msword", "docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"ppt": "application/vnd.ms-powerpoint", "pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		"xls": "application/vnd.ms-excel", "xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"txt": "text/plain", "csv": "text/csv",
	}
	ext := strings.ToLower(extension)
	if val, ok := mediaFormats[ext]; ok {
		return val
	}
	return "application/octet-stream"
}
