package swiss

import (
	"fmt"
	"os/exec"
	"time"
)

var AyanamsaOptions = map[string]string{
	"0":  "Fagan/Bradley",
	"1":  "Lahiri",
	"2":  "De Luce",
	"3":  "Raman",
	"4":  "Usha/Shashi",
	"5":  "Krishnamurti",
	"6":  "Djwhal Khul",
	"7":  "Yukteshwar",
	"8":  "J.N. Bhasin",
	"9":  "Babylonian/Kugler 1",
	"10": "Babylonian/Kugler 2",
	"11": "Babylonian/Kugler 3",
	"12": "Babylonian/Huber",
	"13": "Babylonian/Eta Piscium",
	"14": "Babylonian/Aldebaran = 15 Tau",
	"15": "Hipparchos",
	"16": "Sassanian",
	"17": "Galact. Center = 0 Sag",
	"18": "J2000",
	"19": "J1900",
	"20": "B1950",
	"21": "Suryasiddhanta",
	"22": "Suryasiddhanta, mean Sun",
	"23": "Aryabhata",
	"24": "Aryabhata, mean Sun",
	"25": "SS Revati",
	"26": "SS Citra",
	"27": "True Citra",
	"28": "True Revati",
	"29": "True Pushya (PVRN Rao)",
	"30": "Galactic (Gil Brand)",
	"31": "Galactic Equator (IAU1958)",
	"32": "Galactic Equator",
	"33": "Galactic Equator mid-Mula",
	"34": "Skydram (Mardyks)",
	"35": "True Mula (Chandra Hari)",
	"36": "Dhruva/Gal.Center/Mula (Wilhelm)",
	"37": "Aryabhata 522",
	"38": "Babylonian/Britton",
	"39": "Vedic/Sheoran",
	"40": "Cochrane (Gal.Center = 0 Cap)",
	"41": "Galactic Equator (Fiorenza)",
	"42": "Vettius Valens",
	"43": "Lahiri 1940",
	"44": "Lahiri VP285 (1980)",
	"45": "Krishnamurti VP291",
	"46": "Lahiri ICRC",
}

var BhavaSystemOptions = map[string]string{
	"A": "equal",
	"B": "Alcabitius",
	"C": "Campanus",
	"D": "equal / MC",
	"E": "equal = A",
	"F": "Carter poli-equatorial",
	"G": "36 Gauquelin sectors",
	"H": "horizon / azimuth",
	"I": "Sunshine",
	"K": "Koch",
	"L": "Pullen S-delta",
	"M": "Morinus",
	"N": "Whole sign, Aries = 1st house",
	"O": "Porphyry",
	"P": "Placidus",
	"Q": "Pullen S-ratio",
	"R": "Regiomontanus",
	"S": "Sripati",
	"T": "Polich/Page (topocentric)",
	"U": "Krusinski-Pisa-Goelzer",
	"V": "equal Vehlow",
	"W": "equal, whole sign",
	"X": "axial rotation system/ Meridian houses",
	"Y": "APC houses",
}

type SwissOptions struct {
	DateTime    time.Time // -b: data (ex: 23.12.2025)
	BhavaSystem string    // -hsy: ex: 0.00,51.50,p
	Geopos      LatLng    // -geopos: ex: 0.00,51.50,0
	Ayanamsa    string    // -ay: ex: ay0, ay1, etc
	TrueNode    bool      // -true
}
type SwissResult struct {
	RawOutput string
	Planets   []Graha
}

type Graha struct {
	Name      string
	Longitude string
	Latitude  string
	Distance  string
	Speed     string
}

type Bhava struct {
	Number int
	Start  float64
	End    float64
}

type Extra struct {
	MC            float64
	Vertex        float64
	EquatorialAsc float64
	Coasw         float64
	Coasm         float64
	PolarAsc      float64
}

type Result struct {
	Bhavas    []Bhava
	Grahas    []Graha
	Info      Extra
	RawOutput string
}

// Monta os argumentos do swetest conforme as opções
func (opt *SwissOptions) Args() []string {
	args := []string{}

	if opt.BhavaSystem == "" {
		opt.BhavaSystem = "P"
	}

	if opt.Ayanamsa != "" {
		opt.Ayanamsa = "0"
	}

	args = append(args, "-house"+fmt.Sprintf("%f,%f,%s", opt.Geopos.Lat, opt.Geopos.Lng, opt.BhavaSystem))
	args = append(args, "-b"+opt.DateTime.UTC().Format("2.1.2006"))
	args = append(args, "-ut"+opt.DateTime.UTC().Format("15:04:05"))
	args = append(args, "-hsy"+opt.BhavaSystem)
	// args = append(args, "-ay"+opt.Ayanamsa)

	if opt.TrueNode {
		args = append(args, "-true")
	}

	return args
}

func parseSwissOutputToResult(output string) Result {
	var result Result
	result.RawOutput = output
	result.Info = Extra{}
	var grahas []Graha
	var bhavas []Bhava
	for _, line := range splitLines(output) {
		// Parse planetas
		grahaFound, err := parseGraha(line)

		if err != nil {
			grahas = append(grahas, grahaFound)
		}

		bhavaFound, err := parseBhava(line)

		if err == nil {
			bhavas = append(bhavas, bhavaFound)
		}
	}
	result.Grahas = grahas
	result.Bhavas = bhavas
	return result
}

func parseGraha(line string) (Graha, error) {

	var swissGrahas = []string{
		"Sun",
		"Moon",
		"Mercury",
		"Venus",
		"Mars",
		"Jupiter",
		"Saturn",
		"mean Node",
		"true Node",
	}

	if len(line) > 0 {

		for _, grahaName := range swissGrahas {

			if line[0:len(grahaName)] == grahaName {
				var graha Graha

				_, err := fmt.Sscanf(line, "%s %s %s %s %s", &graha.Name, &graha.Longitude, &graha.Latitude, &graha.Distance, &graha.Speed)

				if err != nil {
					return Graha{}, err
				}
				return graha, nil
			}
		}
	}
	return Graha{}, fmt.Errorf("Not Found")
}

func parseBhava(line string) (Bhava, error) {
	var swissBhavas = []string{
		"house  1",
		"house  2",
		"house  3",
		"house  4",
		"house  5",
		"house  6",
		"house  7",
		"house  8",
		"house  9",
		"house  10",
		"house  11",
		"house  12",
	}

	if len(line) > 0 {
		for _, bhavaName := range swissBhavas {
			if line[0:len(bhavaName)] == bhavaName {
				var num int
				var start, end float64
				_, err := fmt.Sscanf(line, "%d %f %f", &num, &start, &end)

				if err != nil {
					return Bhava{}, err
				}
				return Bhava{Number: num, Start: start, End: end}, nil
			}
		}
	}
	return Bhava{}, fmt.Errorf("Not Found")
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i, c := range s {
		if c == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

// Executa o swetest e retorna Result (com Bhavas e Grahas)
func ExecSwiss(opt *SwissOptions) (Result, error) {
	args := opt.Args()
	cmd := exec.Command("../swetest", args...)
	output, err := cmd.CombinedOutput()
	res := parseSwissOutputToResult(string(output))
	if err != nil {
		return res, fmt.Errorf("erro ao executar swetest: %w", err)
	}
	return res, nil
}
