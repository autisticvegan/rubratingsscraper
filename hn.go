package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)
var citiesToAds = make(map[string]map[string]bool)
func main() {

	listOfCities := [...]string{
		"https://birmingham.rubratings.com",
		"https://mobile.rubratings.com",
		"https://montgomery.rubratings.com",
		"https://anchorage.rubratings.com",
		"https://flagstaff.rubratings.com",
		"https://phoenix.rubratings.com",
		"https://tucson.rubratings.com",
		"https://littlerock.rubratings.com",
		"https://fresno.rubratings.com",
		"https://inlandempire.rubratings.com",
		"https://losangeles.rubratings.com",
		"https://orangecounty.rubratings.com",
		"https://palmsprings.rubratings.com",
		"https://sacramento.rubratings.com",
		"https://sandiego.rubratings.com",
		"https://sanjose.rubratings.com",
		"https://sf.rubratings.com",
		"https://boulder.rubratings.com",
		"https://coloradosprings.rubratings.com",
		"https://denver.rubratings.com",
		"https://hartford.rubratings.com",
		"https://newhaven.rubratings.com",
		"https://dc.rubratings.com",
		"https://fortmyers.rubratings.com",
		"https://gainesville.rubratings.com",
		"https://jacksonville.rubratings.com",
		"https://miami.rubratings.com",
		"https://orlando.rubratings.com",
		"https://panamacity.rubratings.com",
		"https://tallahassee.rubratings.com",
		"https://tampa.rubratings.com",
		"https://atlanta.rubratings.com",
		"https://augusta.rubratings.com",
		"https://macon.rubratings.com",
		"https://savannah.rubratings.com",
		"https://honolulu.rubratings.com",
		"https://boise.rubratings.com",
		"https://twinfalls.rubratings.com",
		"https://chicago.rubratings.com",
		"https://peoria.rubratings.com",
		"https://rockford.rubratings.com",
		"https://evansville.rubratings.com",
		"https://fortwayne.rubratings.com",
		"https://indianapolis.rubratings.com",
		"https://cedarrapids.rubratings.com",
		"https://desmoines.rubratings.com",
		"https://siouxcity.rubratings.com",
		"https://topeka.rubratings.com",
		"https://wichita.rubratings.com",
		"https://bowlinggreen.rubratings.com",
		"https://lexington.rubratings.com",
		"https://louisville.rubratings.com",
		"https://batonrouge.rubratings.com",
		"https://lafayette.rubratings.com",
		"https://neworleans.rubratings.com",
		"https://shreveport.rubratings.com",
		"https://maine.rubratings.com",
		"https://baltimore.rubratings.com",
		"https://boston.rubratings.com",
		"https://lowell.rubratings.com",
		"https://springfield.rubratings.com",
		"https://worcester.rubratings.com",
		"https://annarbor.rubratings.com",
		"https://detroit.rubratings.com",
		"https://grandrapids.rubratings.com",
		"https://lansing.rubratings.com",
		"https://minneapolis.rubratings.com",
		"https://jackson.rubratings.com",
		"https://kc.rubratings.com",
		"https://stlouis.rubratings.com",
		"https://billings.rubratings.com",
		"https://missoula.rubratings.com",
		"https://lincoln.rubratings.com",
		"https://omaha.rubratings.com",
		"https://lasvegas.rubratings.com",
		"https://reno.rubratings.com",
		"https://newhampshire.rubratings.com",
		"https://cnj.rubratings.com",
		"https://newjersey.rubratings.com",
		"https://southjersey.rubratings.com",
		"https://albuquerque.rubratings.com",
		"https://lascruces.rubratings.com",
		"https://santafe.rubratings.com",
		"https://albany.rubratings.com",
		"https://buffalo.rubratings.com",
		"https://newyork.rubratings.com",
		"https://rochester.rubratings.com",
		"https://syracuse.rubratings.com",
		"https://asheville.rubratings.com",
		"https://charlotte.rubratings.com",
		"https://fayetteville.rubratings.com",
		"https://greensboro.rubratings.com",
		"https://raleigh.rubratings.com",
		"https://wilmington.rubratings.com",
		"https://bismarck.rubratings.com",
		"https://fargo.rubratings.com",
		"https://grandforks.rubratings.com",
		"https://minot.rubratings.com",
		"https://cincinnati.rubratings.com",
		"https://cleveland.rubratings.com",
		"https://columbus.rubratings.com",
		"https://toledo.rubratings.com",
		"https://durant.rubratings.com",
		"https://oklahomacity.rubratings.com",
		"https://tulsa.rubratings.com",
		"https://eugene.rubratings.com",
		"https://portland.rubratings.com",
		"https://salem.rubratings.com",
		"https://allentown.rubratings.com",
		"https://erie.rubratings.com",
		"https://philadelphia.rubratings.com",
		"https://pittsburgh.rubratings.com",
		"https://scranton.rubratings.com",
		"https://providence.rubratings.com",
		"https://charleston.rubratings.com",
		"https://columbia.rubratings.com",
		"https://greenville.rubratings.com",
		"https://clarksville.rubratings.com",
		"https://knoxville.rubratings.com",
		"https://memphis.rubratings.com",
		"https://nashville.rubratings.com",
		"https://tricities.rubratings.com",
		"https://austin.rubratings.com",
		"https://corpuschristi.rubratings.com",
		"https://dallas.rubratings.com",
		"https://elpaso.rubratings.com",
		"https://houston.rubratings.com",
		"https://midland.rubratings.com",
		"https://sanantonio.rubratings.com",
		"https://saltlakecity.rubratings.com",
		"https://charlottesville.rubratings.com",
		"https://richmond.rubratings.com",
		"https://roanoke.rubratings.com",
		"https://virginiabeach.rubratings.com",
		"https://seattle.rubratings.com",
		"https://spokane.rubratings.com",
		"https://tacoma.rubratings.com",
		"https://charlestonwv.rubratings.com",
		"https://appleton.rubratings.com",
		"https://greenbay.rubratings.com",
		"https://madison.rubratings.com",
		"https://milwaukee.rubratings.com",
		"https://racine.rubratings.com",
	}

	regexCityPart := regexp.MustCompile(`\/\/.+.rub`)

	for _,c := range listOfCities {
		cc := getCityNameOutOfLink(c, *regexCityPart)
		citiesToAds[cc] = make(map[string]bool)
	}

	// Instantiate default collector
	c := colly.NewCollector(
		// MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(1),
	)

	regexP := `https:\/\/.+\.rubratings\.com\/\d+`
	regexJustNumbers := regexp.MustCompile(`\d+`)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		cc := getCityNameOutOfLink(link, *regexCityPart)
		// Print link
		// Visit link found on page
		doesMatch, _ := regexp.Match(regexP, []byte(link))

		if doesMatch {
			if _, ok := citiesToAds[cc]; ok {

				adId := regexJustNumbers.FindString(link)
				if adId != "" {
					//addStr := fmt.Sprintf("adding %s to ads for %s", adId, link)
					//fmt.Println(addStr)
					citiesToAds[cc][adId] = true
				}
			}
		}

		e.Request.Visit(link)
	})

	for _,k := range listOfCities {
		c.Visit(k)
	}

	//get total number of ads
	//get total for each city
	//put it out to a file
	totalAds := 0
	cityToAdCount := make(map[string]int)
	adCountToCity := make(map[int][]string)
	var listOfAdCounts []int
	for k,v := range citiesToAds {

		fmt.Println(k)
		adsForThisCity := 0
		for adId,_ := range v {
			fmt.Println(adId)
			adsForThisCity += 1
		}
		cityToAdCount[k] = adsForThisCity
		totalAds += adsForThisCity
		adCountToCity[adsForThisCity] = append(adCountToCity[adsForThisCity], k)
		listOfAdCounts = append(listOfAdCounts, adsForThisCity)
		fmt.Println(fmt.Sprintf("ads for %s:%s", k, strconv.Itoa(adsForThisCity)))
	}

	sort.Ints(listOfAdCounts)
	listOfAdCounts = unique(listOfAdCounts)
	for _,x := range listOfAdCounts {
		str := fmt.Sprintf("cities with %s number of ads: ", strconv.Itoa(x))
		fmt.Println(str + strings.Join(adCountToCity[x], ", "))
	}

	fmt.Println("total ads " + strconv.Itoa(totalAds))
	fmt.Println("done")
}

func getCityNameOutOfLink(link string, regexCityPart regexp.Regexp) string {
	cityName := regexCityPart.FindString(link)
	if cityName != "" {
		chopOffFront := cityName[2:]
		return chopOffFront[:(len(chopOffFront)-4)]
	}
	return ""
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}