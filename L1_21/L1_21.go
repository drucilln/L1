package main

import (
	"fmt"
)

// Интерфейс MediaPlayer - целевой интерфейс, ожидаемый клиентским кодом
// Он определяет метод Play, который принимает тип аудио и имя файла
type MediaPlayer interface {
	Play(audioType, fileName string)
}

// Интерфейс AdvancedMediaPlayer - расширенный интерфейс для новых форматов
// Он определяет методы для воспроизведения MP4 и VLC файлов
type AdvancedMediaPlayer interface {
	PlayMP4(fileName string)
	PlayVLC(fileName string)
}

// Структура MP4Player - реализует AdvancedMediaPlayer и может воспроизводить MP4 файлы
type MP4Player struct{}

// Метод PlayMP4 - реализация воспроизведения MP4 файла
func (p *MP4Player) PlayMP4(fileName string) {
	fmt.Printf("Playing mp4 file. Name: %s\n", fileName)
}

// Метод PlayVLC - пустая реализация, так как MP4Player не поддерживает VLC формат
func (p *MP4Player) PlayVLC(fileName string) {
	// Ничего не делаем, так как MP4Player не поддерживает VLC формат.
}

// Структура VLCPlayer - реализует AdvancedMediaPlayer и может воспроизводить VLC файлы
type VLCPlayer struct{}

// Метод PlayVLC - реализация воспроизведения VLC файла
func (p *VLCPlayer) PlayVLC(fileName string) {
	fmt.Printf("Playing vlc file. Name: %s\n", fileName)
}

// Метод PlayMP4 - пустая реализация, так как VLCPlayer не поддерживает MP4 формат
func (p *VLCPlayer) PlayMP4(fileName string) {
	// Ничего не делаем, так как VLCPlayer не поддерживает MP4 формат.
}

// Структура MediaAdapter - реализует интерфейс MediaPlayer и использует AdvancedMediaPlayer для воспроизведения новых форматов
type MediaAdapter struct {
	advancedPlayer AdvancedMediaPlayer
}

// Функция-конструктор NewMediaAdapter - создает новый адаптер в зависимости от типа аудио
func NewMediaAdapter(audioType string) *MediaAdapter {
	if audioType == "mp4" {
		// Если формат mp4, используем MP4Player
		return &MediaAdapter{advancedPlayer: &MP4Player{}}
	} else if audioType == "vlc" {
		// Если формат vlc, используем VLCPlayer
		return &MediaAdapter{advancedPlayer: &VLCPlayer{}}
	}
	// Если формат не поддерживается, возвращаем nil
	return nil
}

// Метод Play - реализация метода из интерфейса MediaPlayer
// Внутри вызывает соответствующий метод AdvancedMediaPlayer
func (adapter *MediaAdapter) Play(audioType, fileName string) {
	if audioType == "mp4" {
		// Воспроизводим MP4 файл с помощью AdvancedMediaPlayer
		adapter.advancedPlayer.PlayMP4(fileName)
	} else if audioType == "vlc" {
		// Воспроизводим VLC файл с помощью AdvancedMediaPlayer
		adapter.advancedPlayer.PlayVLC(fileName)
	}
}

// Основной медиаплеер, который использует MediaAdapter для воспроизведения новых форматов
// Структура AudioPlayer - реализует интерфейс MediaPlayer
// Может воспроизводить mp3 файлы напрямую и использует MediaAdapter для других форматов
type AudioPlayer struct {
	mediaAdapter *MediaAdapter
}

// Метод Play - реализация метода из интерфейса MediaPlayer.
// Определяет, как воспроизводить файлы в зависимости от их формата.
func (player *AudioPlayer) Play(audioType, fileName string) {
	if audioType == "mp3" {
		// Если формат mp3, воспроизводим напрямую
		fmt.Printf("Playing mp3 file. Name: %s\n", fileName)
	} else if audioType == "mp4" || audioType == "vlc" {
		// Если формат mp4 или vlc, используем адаптер
		player.mediaAdapter = NewMediaAdapter(audioType)
		if player.mediaAdapter != nil {
			player.mediaAdapter.Play(audioType, fileName)
		} else {
			fmt.Printf("Invalid media. %s format not supported\n", audioType)
		}
	} else {
		// Формат не поддерживается
		fmt.Printf("Invalid media. %s format not supported\n", audioType)
	}
}

func main() {
	audioPlayer := &AudioPlayer{}

	audioPlayer.Play("mp3", "song.mp3")
	audioPlayer.Play("mp4", "video.mp4")
	audioPlayer.Play("vlc", "movie.vlc")
	audioPlayer.Play("avi", "clip.avi")
}
