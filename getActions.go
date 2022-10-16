package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func GetActions(userIds []string, limit int, offsetPage int) ([]Action, error) {
	actions := []Action{}

	s := GetScraper()
	s.OnHTML(".activities tr", func(collyElem *colly.HTMLElement) {
		var action Action
		var elem = *collyElem.DOM

		action.ModeratorID = ToNumber(collyElem.Request.URL.Query().Get("_userId"))
		action.TaskID = ToNumber(elem.Find(".dataTime > a").Text())

		content := elem.Find("td:nth-child(2)").First().
			Contents().
			Get(6).
			Data
		action.Content = Trim(content)

		userLink := elem.Find("td:nth-child(2) > a").First()
		userHref, _ := userLink.Attr("href")

		actionUserId := ToNumber(FindInText(`\d+$`, userHref))
		action.User = User{Nick: userLink.Text(), ID: actionUserId}

		// Find date
		date := Trim(elem.Find(".dataTime").Contents().Last().Text())

		action.Date = strings.Join(strings.Split(date, " "), "T")

		reasonSelector := elem.Find("td:nth-child(2)").Contents()
		reason := ""

		if reasonSelector.Length() >= 12 {
			reason = ReplaceInText(`^:\s?`, reasonSelector.Get(12).Data, "")

			action.Reason = Trim(reason)
			action.DeletionReason = GetShortDeleteReason(reason)
		}

		reasonSel := elem.Find(".reason").Text()

		if !StringContains("принятые", reasonSel) && content == "\n" {
			action.Type = "DELETED"
			action.ContentType = "ATTACHMENT"
		} else if StringContains("принятые", reasonSel) {
			action.Type = "ACCEPTED"
		} else if StringContains("удален", reasonSel) {
			action.Type = "DELETED"

			if StringContains("задание", reasonSel) {
				action.ContentType = "QUESTION"
			} else {
				action.ContentType = "ANSWER"
			}
		} else if StringContains("отмечено не верно", reasonSel) {
			action.ContentType = "ANSWER"
			action.Type = "REPORTED_FOR_CORRECTION"
		} else {
			action.ContentType = "COMMENT"
			action.Type = "DELETED"
		}

		actions = append(actions, action)
	})

	for _, userId := range userIds {
		url := fmt.Sprintf(
			"%s/moderation_new/view_moderator/%[2]s?_userId=%[2]s&page=%d&limit=%d",
			os.Getenv("BRAINLY_PROXY_HOST"),
			userId,
			offsetPage,
			limit,
		)

		fmt.Println(fmt.Sprintf("fetching %s", url))

		s.Visit(url)
	}

	s.Wait()

	return actions, nil
}