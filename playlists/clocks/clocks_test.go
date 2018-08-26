package clocks

import (
	"reflect"
	"testing"
	"time"

	"github.com/innovate-technologies/DJ/data"
	"github.com/stretchr/testify/mock"
)

type apiMock struct {
	mock.Mock
}

func (a *apiMock) GetAllClocks() []data.Clock {
	args := a.Called()
	return args.Get(0).([]data.Clock)
}
func (a *apiMock) GetAllSongsForTag(tag string) []data.Song {
	args := a.Called()
	return args.Get(0).([]data.Song)
}

func Test_getCurrentClock(t *testing.T) {
	oneclock := []data.Clock{
		data.Clock{
			Name: "always on",
			Start: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 1,
				Hour:      0,
				Minute:    0,
			},
			End: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 7,
				Hour:      23,
				Minute:    59,
			},
		},
	}

	oneclockEveryDay := []data.Clock{
		data.Clock{
			Name: "Sunday",
			Start: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 7,
				Hour:      0,
				Minute:    0,
			},
			End: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 7,
				Hour:      23,
				Minute:    59,
			},
		},
		data.Clock{
			Name: "Monday",
			Start: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 1,
				Hour:      0,
				Minute:    0,
			},
			End: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 1,
				Hour:      23,
				Minute:    59,
			},
		},
		data.Clock{
			Name: "Tuesday",
			Start: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 2,
				Hour:      0,
				Minute:    0,
			},
			End: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 2,
				Hour:      23,
				Minute:    59,
			},
		},
		data.Clock{
			Name: "Wednesday",
			Start: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 3,
				Hour:      0,
				Minute:    0,
			},
			End: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 3,
				Hour:      23,
				Minute:    59,
			},
		},
		data.Clock{
			Name: "Thursday",
			Start: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 4,
				Hour:      0,
				Minute:    0,
			},
			End: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 4,
				Hour:      23,
				Minute:    59,
			},
		},
		data.Clock{
			Name: "Friday",
			Start: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 5,
				Hour:      0,
				Minute:    0,
			},
			End: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 5,
				Hour:      23,
				Minute:    59,
			},
		},
		data.Clock{
			Name: "Saturday",
			Start: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 6,
				Hour:      0,
				Minute:    0,
			},
			End: struct {
				DayOfWeek int `json:"dayOfWeek"`
				Hour      int `json:"hour"`
				Minute    int `json:"minute"`
			}{
				DayOfWeek: 6,
				Hour:      23,
				Minute:    59,
			},
		},
	}

	tests := []struct {
		name string
		want data.Clock
		data []data.Clock
	}{
		{
			name: "one clock",
			want: oneclock[0],
			data: oneclock,
		},
		{
			name: "today's clock",
			want: oneclockEveryDay[int(time.Now().Weekday())],
			data: oneclockEveryDay,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &apiMock{}
			api = m
			m.On("GetAllClocks").Return(tt.data)
			if got := getCurrentClock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCurrentClock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSongsForTag(t *testing.T) {
	onesong := []data.Song{
		data.Song{
			Album:     "unit",
			Artist:    "test",
			Available: true,
			Song:      "passed",
		},
	}
	fivesongs := []data.Song{
		data.Song{
			Album:     "unit",
			Artist:    "test",
			Available: true,
			Song:      "passed",
		},
		data.Song{
			Album:     "unit",
			Artist:    "test",
			Available: true,
			Song:      "passed",
		},
		data.Song{
			Album:     "unit",
			Artist:    "test",
			Available: true,
			Song:      "passed",
		},
		data.Song{
			Album:     "unit",
			Artist:    "test",
			Available: true,
			Song:      "passed",
		},
		data.Song{
			Album:     "unit",
			Artist:    "test",
			Available: true,
			Song:      "passed",
		},
	}
	type args struct {
		tag string
		num int
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
		data    []data.Song
	}{
		{
			name: "exact amount",
			args: args{
				tag: "5songs",
				num: 5,
			},
			wantNum: 5,
			data:    fivesongs,
		},
		{
			name: "give more",
			args: args{
				tag: "5songs",
				num: 10,
			},
			wantNum: 10,
			data:    fivesongs,
		},
		{
			name: "give less",
			args: args{
				tag: "5songs",
				num: 1,
			},
			wantNum: 1,
			data:    fivesongs,
		},
		{
			name: "one song to 10 songs",
			args: args{
				tag: "5songs",
				num: 10,
			},
			wantNum: 10,
			data:    onesong,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &apiMock{}
			api = m
			m.On("GetAllSongsForTag").Return(tt.data)
			if got := getSongsForTag(tt.args.tag, tt.args.num); len(got) != tt.wantNum {
				t.Errorf("getSongsForTag() gave %d, want %d", len(got), tt.wantNum)
			}
		})
	}
}
