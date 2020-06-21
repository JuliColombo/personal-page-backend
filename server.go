package main

import (
	"encoding/json";
	"net/http";
	"fmt";
	"time";
)

type education struct {
	Name string `json:"name"`
	Description string `json:"description"`
	InstitutionName string `json:"institutionName"`
	InstitutionThumbnail string `json:"institutionThumbnail"`
	FromDate time.Time `json:"fromDate"`
	ToDate time.Time `json:"toDate"`
}

func get_education(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	
	var education_list = []education {
		education {
			Name: "Engineering in Information Systems", 
			Description: "Description", 
			InstitutionName: "Universidad Tecnológica Nacional", 
			InstitutionThumbnail: "https://upload.wikimedia.org/wikipedia/commons/6/67/UTN_logo.jpg",
			FromDate: time.Date(2012, time.April, 1, 0, 0, 0, 0, time.UTC),
			ToDate: time.Date(2019, time.August, 2, 0, 0, 0, 0, time.UTC),
		},
	}
	if err := json.NewEncoder(w).Encode(education_list); err != nil {
        panic(err)
    }
}

type job struct {
	Name string `json:"name"`
	Description string `json:"description"`
	CompanyName string `json:"companyName"`
	CompanyThumbnail string `json:"companyThumbnail"`
	FromDate time.Time `json:"fromDate"`
	ToDate time.Time `json:"toDate"`
}

func get_jobs(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	
	var job_list = []job {
		job {
			Name: "Assistant Professor", 
			Description: "I was involved in the creation and correction of assignments, exercises and exams. I have been also in charge of the explanation of some topics in class and I have adviced several students by email. Subjects: Programming Paradigms and System Design. ", 
			CompanyName: "Universidad Tecnológica Nacional", 
			CompanyThumbnail: "https://upload.wikimedia.org/wikipedia/commons/6/67/UTN_logo.jpg",
			FromDate: time.Date(2014, time.April, 1, 0, 0, 0, 0, time.UTC),
			ToDate: time.Date(2015, time.December, 1, 0, 0, 0, 0, time.UTC),
		},
		job {
			Name: "Fullstack Developer", 
			Description: "I took part of a team which developed a web system to upload data and audiovisual content, for its later reproduction in any site. The main technologies I'm using are Django (python) and VueJs (Javascript). I have also worked in an ecommerce project where I worked as a developer using Ruby on Rails (ruby) and ReactJs (Javascript). My tasks include both programming and the revision of code and functionalities developed by my colleagues.", 
			CompanyName: "devartis", 
			CompanyThumbnail: "https://pbs.twimg.com/profile_images/1039514458282844161/apKQh1fu_400x400.jpg",
			FromDate: time.Date(2016, time.March, 1, 0, 0, 0, 0, time.UTC),
			ToDate: time.Date(2019, time.December, 1, 0, 0, 0, 0, time.UTC),
		},
		job {
			Name: "Fullstack Developer", 
			Description: "Currently I find myself developing the backend of an ecommerce with Django, taking care of the infrastructure, which involves the CI/CD with Docker in AWS. Likewise I developed the integration with a payment system (Stripe), mailing (Sendgrid) and stock manager (Tailorpad).", 
			CompanyName: "IbisDev", 
			CompanyThumbnail: "https://media-exp1.licdn.com/dms/image/C4E0BAQE-TeSM6YFhxw/company-logo_200_200/0?e=2159024400&v=beta&t=xaXuwF-XcYS5kbeN3_FgErTUtmlgxS0WhWr9J2mUjMc",
			FromDate: time.Date(2019, time.December, 1, 0, 0, 0, 0, time.UTC),
			ToDate: time.Time{},
		},
	}
	if err := json.NewEncoder(w).Encode(job_list); err != nil {
        panic(err)
    }
}

type post struct {
	Title string `json:"title"`
	PublicationDate time.Time `json:"publicationDate"`
	Link string `json:"link"`
	Thumbnail string `json:"thumbnail"`
}

func getJSON(url string, result interface{}) error {
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("cannot fetch URL %q: %v", url, err)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected http GET status: %s", resp.Status)
    }
    err = json.NewDecoder(resp.Body).Decode(result)
    if err != nil {
        return fmt.Errorf("cannot decode JSON: %v", err)
    }
    return nil
}


func get_posts(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	var mediumPosts map[string]interface{}
	getJSON("https://api.rss2json.com/v1/api.json?rss_url=https://medium.com/feed/@julietanataliacolombo", &mediumPosts)
	
	var posts []post
	for _, mediumPost := range mediumPosts["items"].([]interface{}) {
		if len(mediumPost.(map[string]interface{})["categories"].([]interface{})) > 0 {
			date, err := time.Parse(time.RFC3339, mediumPost.(map[string]interface{})["pubDate"].(string))
			if err != nil {
				date = time.Time{}
			}
			posts = append(posts, post {
				Title: mediumPost.(map[string]interface{})["title"].(string),
				PublicationDate: date,
				Link: mediumPost.(map[string]interface{})["link"].(string),
				Thumbnail: mediumPost.(map[string]interface{})["thumbnail"].(string),
			})
		}
	}
	if err := json.NewEncoder(w).Encode(posts); err != nil {
        panic(err)
    }
}

func main() {
	http.HandleFunc("/education/", get_education)
	http.HandleFunc("/jobs/", get_jobs)
	http.HandleFunc("/posts/", get_posts)
	http.ListenAndServe(":8090", nil)
}