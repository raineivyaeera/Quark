package main
import "fmt"
import "encoding/json"
import "os"
import "math/rand/v2"
import "strings"
import "strconv"

var card []string
var emoji []string
var texture []string
var genre []string
var chord []string
var chordProgression []string
var BPM int
var timeSignature []string
var rhythm []string

func setup() {
	fileCard, errCard := os.Open("cards.json")
	if errCard != nil {
		panic(errCard)
	}; defer fileCard.Close()
	errCard = json.NewDecoder(fileCard).Decode(&card)
	if errCard != nil {
		panic(errCard)
	}
	fileEmoji, errEmoji := os.Open("emojis.json")
	if errEmoji != nil {
		panic(errEmoji)
	}; defer fileEmoji.Close()
	errEmoji = json.NewDecoder(fileEmoji).Decode(&emoji)
	if errEmoji != nil {
		panic(errEmoji)
	}
	fileGenre, errGenre := os.Open("genres.json")
	if errGenre != nil {
		panic(errGenre)
	}; defer fileGenre.Close()
	errGenre = json.NewDecoder(fileGenre).Decode(&genre)
	if errGenre != nil {
		panic(errGenre)
	}
	fileTexture, errTexture := os.Open("textures.json")
	if errTexture != nil {
		panic (errTexture)
	}; defer fileTexture.Close()
	errTexture = json.NewDecoder(fileTexture).Decode(&texture)
	if errTexture != nil {
		panic(errTexture)
	}
	fileChords, errChords := os.Open("chords.json")
	if errChords != nil {
		panic (errChords)
	}; defer fileChords.Close()
	errChords = json.NewDecoder(fileChords).Decode(&chord)
	if errChords != nil {
		panic(errChords)
	}
	timeSignature = append(timeSignature, "4/4", "3/4", "6/8", "5/4", "3/8", "5/8", "6/4", "7/8", "7/4", "9/4", "10/4", "10/8", "11/4", "11/8", "13/4", "13/8", "14/4", "14/8", "15/4", "15/8")
}

func colorizeText(text string) string {
    r := rand.IntN(256)
    g := rand.IntN(256)
    b := rand.IntN(256)
    return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, text)
}

func pullCard() {
	var rnCard int = rand.IntN(len(card))
	var rnEmoji int = rand.IntN(len(emoji))
	var rnGenre int = rand.IntN(len(genre))
	var rnTexture int = rand.IntN(len(texture))
	var rnTimesig int = rand.IntN(len(timeSignature))
	for i:=0; i < 4; i++ {
		var rnChord int = rand.IntN(len(chord))
		if i < 3 {
			chordProgression = append(chordProgression, chord[rnChord]);
			chordProgression = append(chordProgression, "=>")
		} else { chordProgression = append(chordProgression, chord[rnChord]) } 
	}
	BPM = (rand.IntN(180)) + 40
	rhythmNumString := strings.Split(timeSignature[rnTimesig], "/")[0]
	rhythmNum, err := strconv.Atoi(rhythmNumString); if err != nil {
		panic(err)
	}
	rhythm = []string{}
	var xCount int = 0
	for i:=0; i < rhythmNum; i++ {
		rn := rand.IntN(2)
		if rn > 0 {
			rhythm = append(rhythm, "x"); xCount++
		} else {
			rhythm = append(rhythm, "-")
		}
	}
	if xCount == 0 { rhythm[len(rhythm)-1] = "x" }
	fmt.Printf(colorizeText("\nQUARK\n\n"))
	fmt.Printf("Card: %s\n", card[rnCard])
	fmt.Printf("Emoji: %s\n", emoji[rnEmoji])
	fmt.Printf("Genre: %s\n", genre[rnGenre])
	fmt.Printf("Texture/Sense: %s\n\n", texture[rnTexture])
	fmt.Printf("Chord Progression: %s\n", chordProgression)
	fmt.Printf("BPM: %d\n", BPM)
	fmt.Printf("Time Signature: %s\n", timeSignature[rnTimesig])
	fmt.Printf("Rhythm: %s\n\n", rhythm)
}

func main() {
	setup()
	pullCard()
}
