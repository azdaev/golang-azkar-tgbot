package audio

import (
	"fmt"
	"strconv"
)

type AudioInfo struct {
	FilePath string
	Title    string
}

// GetAudioInfo returns the file path and title for an azkar audio file
func GetAudioInfo(index int, isMorning bool) AudioInfo {
	audioFilePath := "media/"
	audioTitle := ""

	if isMorning {
		audioFilePath += "morning/"
		audioTitle += "Утренний зикр №"
	} else {
		audioFilePath += "evening/"
		audioTitle += "Вечерний зикр №"
	}

	audioFilePath += strconv.Itoa(index)
	audioFilePath += ".mp3"
	audioTitle += strconv.Itoa(index + 1)

	return AudioInfo{
		FilePath: audioFilePath,
		Title:    audioTitle,
	}
}

// GetAudioFilePath returns just the file path for an azkar audio file
func GetAudioFilePath(index int, isMorning bool) string {
	return GetAudioInfo(index, isMorning).FilePath
}

// GetAudioTitle returns just the title for an azkar audio file
func GetAudioTitle(index int, isMorning bool) string {
	return GetAudioInfo(index, isMorning).Title
}

// GetFormattedAudioFilePath returns a formatted file path string
func GetFormattedAudioFilePath(index int, isMorning bool) string {
	timeOfDay := "evening"
	if isMorning {
		timeOfDay = "morning"
	}
	return fmt.Sprintf("media/%s/%d.mp3", timeOfDay, index)
}
