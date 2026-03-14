package kpmapper_test

import (
	"testing"

	"github.com/AliyaVV/MovieHub/internal/mapper/kpmapper"
	"github.com/AliyaVV/MovieHub/internal/proxy/kinopoisk"
)

func TestMapKPTitleToEntity(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		wantId          int
		wantMovieType   string
		wantYear        int
		wantDescription string
		wantTop250      int
		respTitle       kinopoisk.KPSearchTitle
		wantname        string
	}{
		{
			name: "Avengers",
			respTitle: kinopoisk.KPSearchTitle{
				ID:        263531,
				Name:      "Avengers2",
				MovieType: "movie",
				Year:      2012,
				Description: `Ник Фьюри созывает сильнейших супергероев планеты, 
				которыми оказываются Железный человек, Черная вдова, Капитан Америка, Халк и Тор, 
				чтобы дать противнику достойный отпор`,
				Top250: 12,
			},
			wantId:        263531,
			wantname:      "Avengers",
			wantMovieType: "movie",
			wantYear:      2012,
			wantDescription: `Ник Фьюри созывает сильнейших супергероев планеты, 
			которыми оказываются Железный человек, Черная вдова, Капитан Америка, Халк и Тор, 
			чтобы дать противнику достойный отпор`,
			wantTop250: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := kpmapper.MapKPTitleToEntity(tt.respTitle)
			if got.Id != tt.wantId {
				t.Errorf("ID = %d, want %d", got.Id, tt.wantId)
			}
			if got.Runame != tt.name {
				t.Errorf("name = %s, want %s", got.Runame, tt.name)
			}
			if got.MovieType != tt.wantMovieType {
				t.Errorf("MovieType = %s, want %s", got.MovieType, tt.wantMovieType)
			}
			if got.MovieYear != tt.wantYear {
				t.Errorf("year = %d, want %d", got.MovieYear, tt.wantYear)
			}

			// if true {
			// 	t.Errorf("MapKPTitleToEntity() = %v, want %v", got, tt.want)
			// }
		})
	}
}
