package main

import "fmt"

type Observable interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
	sendAll()
}

type Observer interface {
	handleEvent(vacancies []string)
}

type JobSite struct {
	subcribers []Observer
	vacancies  []string
}

type Person struct {
	name string
}

func newPerson(cur_name string) Person {
	return Person{name: cur_name}
}

func newJobSite() JobSite {
	return JobSite{subcribers: make([]Observer, 0), vacancies: make([]string, 0)}
}

func (s *JobSite) addVacancy(vacancy string) {
	s.vacancies = append(s.vacancies, vacancy)
}

func (s *JobSite) removeVacancy(vacancy string) {
	var idx = 0
	for i, vac := range s.vacancies {
		if vac == vacancy {
			idx = i
			break
		}
	}
	s.vacancies[idx] = s.vacancies[len(s.vacancies)-1]
	s.vacancies[len(s.vacancies)-1] = ""
	s.vacancies = s.vacancies[:len(s.vacancies)-1]
}

func (s *JobSite) subscribe(observer Observer) {
	s.subcribers = append(s.subcribers, observer)
}

func (s *JobSite) unsubscribe(observer Observer) {
	var idx = 0
	for i, obs := range s.subcribers {
		if obs == observer {
			idx = i
			break
		}
	}
	s.subcribers[idx] = s.subcribers[len(s.subcribers)-1]
	s.subcribers[len(s.subcribers)-1] = nil
	s.subcribers = s.subcribers[:len(s.subcribers)-1]
}

func (s *JobSite) sendAll() {
	if len(s.subcribers) == 0 {
		fmt.Printf("There are no subscribers\n")
		return
	}
	for _, observer := range s.subcribers {
		observer.handleEvent(s.vacancies)
	}
}

func (p Person) handleEvent(vacancies []string) {
	fmt.Printf("Hello dear %s\n", p.name)
	if len(vacancies) == 0 {
		fmt.Printf("There are no vacancies\n")
		return
	}
	fmt.Printf("Vacancies updated: \n")
	for _, vacancy := range vacancies {
		fmt.Println(vacancy)
	}
}

/*
func main() {
	hhKz := newJobSite()
	bob := newPerson("bob")
	hhKz.addVacancy("Senior HTML Developer")
	hhKz.addVacancy("Senior Back-End developer")
	hhKz.subscribe(bob)
	hhKz.sendAll()
	hhKz.removeVacancy("Senior HTML Developer")
	hhKz.sendAll()
	hhKz.removeVacancy("Senior Back-End developer")
	hhKz.sendAll()
}*/
