package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestProcessMeasurements(t *testing.T) {
	expected := []string{
		"Abidjan=22.5/22.5/22.5",
		"Abéché=37.6/37.6/37.6",
		"Accra=36.6/36.6/36.6",
		"Assab=39.2/39.2/39.2",
		"Baltimore=11.6/11.6/11.6",
		"Banjul=12.8/14.5/16.1",
		"Belize City=27.8/27.8/27.8",
		"Benghazi=18.1/18.1/18.1",
		"Berlin=25.8/25.8/25.8",
		"Bishkek=7.0/7.0/7.0",
		"Bordeaux=26.5/26.5/26.5",
		"Bucharest=10.2/10.2/10.2",
		"Burnie=15.5/15.5/15.5",
		"Cairns=36.6/36.6/36.6",
		"Cairo=26.1/26.1/26.1",
		"Calgary=9.8/9.8/9.8",
		"Canberra=11.0/11.0/11.0",
		"Chiang Mai=27.4/37.2/47.0",
		"Cotonou=17.5/17.5/17.5",
		"Cracow=0.7/0.7/0.7",
		"Dampier=18.6/18.6/18.6",
		"Darwin=23.9/23.9/23.9",
		"Denpasar=21.7/21.7/21.7",
		"Dodoma=22.1/22.1/22.1",
		"El Paso=12.9/19.7/27.1",
		"Flores,  Petén=22.3/22.3/22.3",
		"Frankfurt=30.8/30.8/30.8",
		"Ho Chi Minh City=28.2/30.6/33.1",
		"Hobart=7.9/7.9/7.9",
		"Indianapolis=17.5/17.5/17.5",
		"Istanbul=10.9/15.7/20.5",
		"Kampala=21.0/21.0/21.0",
		"Kansas City=11.4/11.4/11.4",
		"Karonga=33.1/33.1/33.1",
		"Kingston=39.9/39.9/39.9",
		"Kunming=14.7/14.7/14.7",
		"Kuopio=11.5/11.5/11.5",
		"Kyiv=10.8/10.8/10.8",
		"La Ceiba=32.0/32.0/32.0",
		"Lagos=26.4/26.4/26.4",
		"Lake Havasu City=19.2/21.0/22.8",
		"Las Vegas=-1.8/-1.8/-1.8",
		"Libreville=26.3/27.4/28.5",
		"Ljubljana=9.6/9.6/9.6",
		"Luxembourg City=4.0/8.3/12.6",
		"Lviv=9.7/9.7/9.7",
		"Manila=18.1/18.1/18.1",
		"Marrakesh=28.6/28.6/28.6",
		"Marseille=-9.1/2.4/13.8",
		"Mexicali=12.6/12.6/12.6",
		"Miami=26.7/26.7/26.7",
		"Moscow=18.6/18.6/18.6",
		"Nairobi=10.3/22.2/34.1",
		"Napier=5.2/5.2/5.2",
		"Napoli=34.3/34.3/34.3",
		"Nouakchott=17.5/21.1/24.6",
		"Odienné=20.4/20.4/20.4",
		"Oranjestad=23.0/29.8/36.6",
		"Oslo=24.9/24.9/24.9",
		"Ouagadougou=30.1/30.1/30.1",
		"Palembang=19.2/21.3/23.4",
		"Panama City=26.8/26.8/26.8",
		"Port Sudan=25.5/25.5/25.5",
		"Rangpur=37.4/37.4/37.4",
		"Riga=6.3/6.3/6.3",
		"Rome=19.5/19.5/19.5",
		"Roseau=30.0/30.0/30.0",
		"Sacramento=41.7/41.7/41.7",
		"San Antonio=21.0/21.0/21.0",
		"Shanghai=12.3/12.3/12.3",
		"Split=14.5/14.5/14.5",
		"Suwałki=7.8/7.8/7.8",
		"Tabriz=2.0/2.0/2.0",
		"Tamanrasset=25.5/25.5/25.5",
		"Tashkent=18.8/18.8/18.8",
		"Tel Aviv=16.1/16.1/16.1",
		"Thiès=21.8/21.8/21.8",
		"Tijuana=-6.4/-6.4/-6.4",
		"Tirana=9.9/9.9/9.9",
		"Tripoli=16.4/16.4/16.4",
		"Upington=8.4/8.4/8.4",
		"Valencia=17.3/17.3/17.3",
		"Wellington=12.3/12.3/12.3",
		"Yaoundé=28.2/29.3/30.4",
		"Yellowknife=-21.3/-21.3/-21.3",
	}

	var w bytes.Buffer
	processMeasurements("measurements-100.txt", &w)
	result := w.String()
	got := strings.Split(result, "\n")

	for i, expectedLine := range expected {
		fmt.Println("expected: ", expectedLine)
		if expectedLine != got[i] {
			t.Errorf("Line %d expected %q, got %q", i+1, expectedLine, got[i])
		}
	}
}
