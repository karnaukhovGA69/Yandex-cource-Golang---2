package main

import (
	"bufio"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

var ready = []string{"Готово", "В работе", "Не будет сделано"}

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

func GetTasks(text string, user *string, status *string) []Ticket {
	tickets := make([]Ticket, 0)
	output := make([]Ticket, 0)
	reader := strings.NewReader(text)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		ticket, err := ParseTicket(text)
		if err != nil {
			continue
		}
		tickets = append(tickets, ticket)
	}
	for _, value := range tickets {
		if user == nil && status != nil {
			if value.Status == *status {
				output = append(output, value)
			}
		}

		if user != nil && status == nil {
			if value.User == *user {
				output = append(output, value)
			}
		}

		if user != nil && status != nil {
			if value.User == *user && value.Status == *status {
				output = append(output, value)
			}
		}

		if user == nil && status == nil {
			output = append(output, value)
		}

	}

	return output
}

func ParseTicket(ticket string) (Ticket, error) {
	ticket = strings.TrimSpace(ticket)
	newTicket := Ticket{}
	fields := strings.Split(ticket, "_")
	if len(fields) != 4 {
		return Ticket{}, errors.New("Не верный тикет")
	}

	if !IsValidName(fields[0]) {
		return Ticket{}, errors.New("Ошибка")
	}
	newTicket.Ticket = fields[0]
	newTicket.User = fields[1]
	if !slices.Contains(ready, fields[2]) {
		return Ticket{}, errors.New("Не коректная готовность")
	}
	newTicket.Status = fields[2]
	if _, err := time.Parse("2006-01-02", fields[3]); err != nil {
		return Ticket{}, err
	}
	newTicket.Date, _ = time.Parse("2006-01-02", fields[3])
	return newTicket, nil
}

func IsValidName(name string) bool {
	const pref = "TICKET-"
	if !strings.HasPrefix(name, pref) {
		return false
	}
	num := name[len(pref):]
	if num == "" {
		return false
	}
	_, err := strconv.Atoi(num)
	return err == nil
}

func (t *Ticket) ToString() string {

	return fmt.Sprintf("%s_%s_%s_%s",
		t.Ticket, t.User, t.Status, t.Date.Format("2006-01-02"))
}

func main() {
	text := "TICKET-12345_Паша Попов_Готово_2024-01-01\nTICKET-12346_Иван Иванов_Вработе_2024-01-02\nTICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03\nTICKET-12348_Паша Попов_В работе_2024-01-04"
	name := "Паша Попов"
	tickets := GetTasks(text, &name, nil)
	fmt.Println(len(tickets))

}
