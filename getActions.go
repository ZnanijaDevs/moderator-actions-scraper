package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func GetActions(
	userId int, 
	limit int, 
	offsetPage int,
	authToken string,
) (a []Action, errors []string) {
	actions := []Action{}
	fetchErrors := []string{}

	s := GetScraper(authToken)
	s.OnResponse(func(r *colly.Response) {
		if (strings.Contains(r.Request.URL.Path, "login")) {
			fetchErrors = append(fetchErrors, "invalid auth")
			return
		}

		if (r.StatusCode == http.StatusOK) {
			fmt.Println(fmt.Sprintf("success -> %s", r.Request.URL.Path))
		} else {
			fetchErrors = append(fetchErrors, fmt.Sprintf("request failed with status %d", r.StatusCode))
		}
	})

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
			action.Attachments = FindAllMatchesInText("https:/{2}ru-static.z-dn.net[/.a-z0-9/]+", reason)
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

	url := fmt.Sprintf(
		"%s/moderation_new/view_moderator/%[2]d?_userId=%[2]d&page=%d&limit=%d",
		os.Getenv("BRAINLY_PROXY_HOST"),
		userId,
		offsetPage,
		limit,
	)

	s.Visit(url)
	s.Wait()

	return actions, fetchErrors
}