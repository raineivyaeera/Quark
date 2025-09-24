package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

var cardable, emojiable, genreable, texturable, timeable, chordable, bpmable, colorable, rhythmable bool
var card []string
var emoji []string
var texture []string
var genre []string
var chord []string
var chordProgression []string
var timeSignature []string
var rhythm []string
var BPM int
var rnCard int 
var rnEmoji int 
var rnGenre int
var rnTexture int
var rnTimesig int
var r, g, b int
var firstDraw bool = true
var input string

func setup() {
	setBools()
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

func clearBools() {
	cardable, emojiable, genreable, texturable, timeable, chordable, bpmable, colorable, rhythmable = false, false, false, false, false, false, false, false, false;
}

func setBools() {
	cardable, emojiable, genreable, texturable, timeable, chordable, bpmable, colorable, rhythmable = true, true, true, true, true, true, true, true, true;
}

func colorizeText(text string) string {
	if colorable {
    	r = rand.IntN(256)
    	g = rand.IntN(256)
    	b = rand.IntN(256)
	}
    return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, text)
}

func pullCard() {
	if cardable { rnCard = rand.IntN(len(card)) }
	if emojiable { rnEmoji = rand.IntN(len(emoji)) }
	if genreable { rnGenre = rand.IntN(len(genre)) }
	if texturable { rnTexture = rand.IntN(len(texture)) }
	if timeable { rnTimesig = rand.IntN(len(timeSignature)) }	
	if chordable {
		chordProgression = nil
		for i:=0; i < 4; i++ {
			var rnChord int = rand.IntN(len(chord))
			if i < 3 {
				chordProgression = append(chordProgression, chord[rnChord]);
				chordProgression = append(chordProgression, "=>")
			} else { chordProgression = append(chordProgression, chord[rnChord]) } 
		} 
	}
	if bpmable {
		BPM = (rand.IntN(160)) + 40
		rng := rand.IntN(10)
		if BPM <= 100 && rng >= 3 {
			BPM += (rand.IntN(50)+25)
		}
		if BPM >= 150 && rng >= 3 {
			BPM -= (rand.IntN(50)+25)
		}
	}
	if rhythmable {
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
	}

	fmt.Printf(colorizeText("\nQUARK\n\n"))
	fmt.Printf("Card: %s\n", card[rnCard])
	fmt.Printf("Sensory: %s\n", texture[rnTexture])
	fmt.Printf("Emoji: %s\n", emoji[rnEmoji])
	fmt.Printf("Tone–Genre: %s\n\n", genre[rnGenre])
	fmt.Printf("Chord Progression: %s\n", chordProgression)
	fmt.Printf("BPM: %d\n", BPM)
	fmt.Printf("Time Signature: %s\n", timeSignature[rnTimesig])
	fmt.Printf("Rhythm: %s\n\n", rhythm)
}

func main() {
	for {
		if firstDraw {
			setup()
			pullCard()
			firstDraw = false
		}
		fmt.Scanln(&input)
		switch input {
		case "n":
			setBools()
			pullCard()
		case "e":
			os.Exit(1)
		case "1":
			clearBools()
			cardable = true
			pullCard()
		case "2":
			clearBools()
			texturable = true
			pullCard()
		case "3":
			clearBools()
			emojiable = true
			pullCard()
		case "4":
			clearBools()
			genreable = true
			pullCard()
		case "5":
			clearBools()
			chordable = true
			pullCard()
		case "6":
			clearBools()
			bpmable = true;
			pullCard()
		case "7":
			clearBools()
			timeable = true; rhythmable = true
			pullCard()
		case "8":
			clearBools()
			rhythmable = true
			pullCard()
		case "9":
			clearBools()
			colorable = true
			pullCard()
		default: fmt.Println("give me a number or press n for a new set")
		}
	}
}
