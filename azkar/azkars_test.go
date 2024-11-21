package azkar

import (
	"testing"

	"github.com/azdaev/azkar-tg-bot/repository/models"
)

func TestWrap(t *testing.T) {
	type args struct {
		config    *models.ConfigInclude
		index     int
		isMorning bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Morning zikr with all config enabled",
			args: args{
				config: &models.ConfigInclude{
					Arabic:        true,
					Russian:       true,
					Transcription: true,
				},
				index:     0,
				isMorning: true,
			},
			want: "Утренний зикр №1\n\n" + MorningAzkar[0].Arabic + "\n\n" + MorningAzkar[0].Russian + "\n\n" + MorningAzkar[0].Transcription,
		},
		{
			name: "Evening zikr with only Arabic",
			args: args{
				config: &models.ConfigInclude{
					Arabic:        true,
					Russian:       false,
					Transcription: false,
				},
				index:     0,
				isMorning: false,
			},
			want: "Вечерний зикр №1\n\n" + EveningAzkar[0].Arabic + "\n\n",
		},
		{
			name: "Morning zikr with Russian and Transcription",
			args: args{
				config: &models.ConfigInclude{
					Arabic:        false,
					Russian:       true,
					Transcription: true,
				},
				index:     1,
				isMorning: true,
			},
			want: "Утренний зикр №2\n\n" + MorningAzkar[1].Russian + "\n\n" + MorningAzkar[1].Transcription,
		},
		{
			name: "Evening zikr with all config disabled",
			args: args{
				config: &models.ConfigInclude{
					Arabic:        false,
					Russian:       false,
					Transcription: false,
				},
				index:     0,
				isMorning: false,
			},
			want: "Вечерний зикр №1\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wrap(tt.args.config, tt.args.index, tt.args.isMorning); got != tt.want {
				t.Errorf("Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
