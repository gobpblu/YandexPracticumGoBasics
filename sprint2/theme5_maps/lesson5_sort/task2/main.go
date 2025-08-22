package main

import "fmt"

func main() {
	titles := map[string][]string{
		"Что делать?":               {"планы", "думы"},
		"Где отдохнуть в выходные":  {"отдых", "планы"},
		"Кто виноват?":              {"поиск", "думы", "общество"},
		"Лучшие рестораны города":   {"еда", "отдых"},
		"С заботой о народе":        {"думы", "общество"},
		"Как попасть на дегустацию": {"еда", "планы", "отдых"},
	}
	tags := make(map[string][]string)

	for title, tagsList := range titles {
		for _, tag := range tagsList {
			_, ok := tags[tag]
			if !ok {
				tags[tag] = make([]string, 0)
			}
			tags[tag] = append(tags[tag], title)
		}
	}

	fmt.Println(len(tags["думы"]), len(tags["общество"]), tags["поиск"])
}
