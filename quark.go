package main

import "encoding/json"
import "fmt"
import "math/rand/v2"
import "os"
import "strconv"
import "strings"

type Quark struct {
	Entry []string
	Index int
	Randomize bool
}

var card, sensory, emoji, genre, chord, bpm, time, rhythm, color Quark
var chordProg []string
var input string

func loadJSON(filename string, target *[]string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(target); err != nil {
		panic(err)
	}
}

func setup() {
	loadJSON("cards.json", &card.Entry)
	loadJSON("textures.json", &sensory.Entry)
	loadJSON("emojis.json", &emoji.Entry)
	loadJSON("genres.json", &genre.Entry)
	loadJSON("chords.json", &chord.Entry)

	time.Entry = []string{"4/4", "3/4", "6/8", "5/4", "3/8", "5/8", "6/4", "7/8", "7/4", "9/4", "10/4", "10/8", "11/4", "11/8", "13/4", "13/8", "14/4", "14/8", "15/4", "15/8"}

	card.Randomize, sensory.Randomize, emoji.Randomize, genre.Randomize, chord.Randomize, bpm.Randomize, time.Randomize, rhythm.Randomize, color.Randomize = true, true, true, true, true, true, true, true, true
}

func pickRandom(q *Quark) {
	if q.Randomize && len(q.Entry) > 0 {
		q.Index = rand.IntN(len(q.Entry))
	}
}

func pickBPM(q *Quark) {
	if q.Randomize {
		q.Index = 40 + rand.IntN(160)
		rng := rand.IntN(10)
		if q.Index <= 100 && rng >= 3 {
			q.Index += rand.IntN(55) + 25
		}
		if q.Index >= 150 && rng >= 3 {
			q.Index -= rand.IntN(55) + 25
		}
	}
}

func pickColor(q *Quark) {
	if q.Randomize {
		r := rand.IntN(256)
		g := rand.IntN(256)
		b := rand.IntN(256)
		colorStr := fmt.Sprintf("\033[38;2;%d;%d;%dmQUARK\033[0m", r, g, b)
		q.Entry = []string{colorStr}
		q.Index = 0
	}
}

func colorizeText(q *Quark) string {
	return q.Entry[q.Index]
}

func pullCard() {
	pickRandom(&card)
	pickRandom(&sensory)
	pickRandom(&emoji)
	pickRandom(&genre)
	pickRandom(&chord)
	pickRandom(&time)
	pickRandom(&rhythm)
	pickBPM(&bpm)
	pickColor(&color)

	if chord.Randomize {
		chordProg = nil
		for i := 0; i < 4; i++ {
			chordProg = append(chordProg, chord.Entry[rand.IntN(len(chord.Entry))])
			if i < 3 {
				chordProg = append(chordProg, "=>")
			}
		}
	}

	if rhythm.Randomize {
		numStr := strings.Split(time.Entry[time.Index], "/")[0]
		count, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		rhythm.Entry = nil
		xCount := 0
		for i := 0; i < count; i++ {
			if rand.IntN(2) > 0 {
				rhythm.Entry = append(rhythm.Entry, "x")
				xCount++
			} else {
				rhythm.Entry = append(rhythm.Entry, "-")
			}
		}
		if xCount == 0 && len(rhythm.Entry) > 0 {
			rhythm.Entry[len(rhythm.Entry)-1] = "x"
		}
	}

	fmt.Printf("\n%s\n\n", colorizeText(&color))
	fmt.Printf("Card: %s\n", card.Entry[card.Index])
	fmt.Printf("Sensory: %s\n", sensory.Entry[sensory.Index])
	fmt.Printf("Emoji: %s\n", emoji.Entry[emoji.Index])
	fmt.Printf("Tone–Genre: %s\n\n", genre.Entry[genre.Index])
	fmt.Printf("Chord Progression: %s\n", strings.Join(chordProg, " "))
	fmt.Printf("BPM: %d\n", bpm.Index)
	fmt.Printf("Time Signature: %s\n", time.Entry[time.Index])
	fmt.Printf("Rhythm: %s\n\n", strings.Join(rhythm.Entry, " "))
}

func disableAll() {
	card.Randomize, sensory.Randomize, emoji.Randomize, genre.Randomize, chord.Randomize, bpm.Randomize, time.Randomize, rhythm.Randomize, color.Randomize = false, false, false, false, false, false, false, false, false
}

func main() {
	setup()
	pullCard()
	for {
		fmt.Scanln(&input)
		switch input {
		case "n":
			setup()
			pullCard()
		case "e", "q":
			os.Exit(0)
		case "1":
			disableAll(); card.Randomize = true; pullCard()
		case "2":
			disableAll(); sensory.Randomize = true; pullCard()
		case "3":
			disableAll(); emoji.Randomize = true; pullCard()
		case "4":
			disableAll(); genre.Randomize = true; pullCard()
		case "5":
			disableAll(); chord.Randomize = true; pullCard()
		case "6":
			disableAll(); bpm.Randomize = true; pullCard()
		case "7":
			disableAll(); time.Randomize, rhythm.Randomize = true, true; pullCard()
		case "8":
			disableAll(); rhythm.Randomize = true; pullCard()
		case "9":
			disableAll(); color.Randomize = true; pullCard()
		default:
			fmt.Println("give me a number or press n for a new set")
		}
	}
}
