package main

var deletionReasons = []DeletionReason{
	// default deletion reasons
	{
		Name:         "Культура (вопрос)",
		ID:           40,
		Regex:        "вопрос задан в грубой",
		TakePoints:   true,
		ReturnPoints: false,
		WithWarn:     true,
	}, {
		Name:         "Спам (вопрос)",
		ID:           2,
		Regex:        "задавайте осмысленные вопросы",
		TakePoints:   true,
		ReturnPoints: false,
		WithWarn:     true,
	}, {
		Name:         "Контрольные",
		ID:           38,
		Regex:        "вопросы из действующих контрольных",
		TakePoints:   true,
		ReturnPoints: false,
		WithWarn:     true,
	}, {
		Name:              "Реклама (вопрос)",
		ID:                47,
		Regex:             "рекламной информации",
		TakePoints:        true,
		ReturnPoints:      false,
		WithWarn:          true,
		UserMustBeDeleted: true,
	}, {
		Name:         "Неполное условие",
		ID:           3,
		Regex:        "не хватает важных условий",
		TakePoints:   true,
		ReturnPoints: true,
		WithWarn:     false,
	}, {
		Name:         "Банальный",
		ID:           4,
		Regex:        "является слишком простым",
		TakePoints:   true,
		ReturnPoints: false,
		WithWarn:     false,
	}, {
		Name:         "Нет приложения",
		ID:           17,
		Regex:        "качественное приложение",
		TakePoints:   true,
		ReturnPoints: true,
		WithWarn:     false,
	}, {
		Name:         "Много вопросов",
		ID:           32,
		Regex:        "нужно разделить на несколько",
		TakePoints:   true,
		ReturnPoints: true,
		WithWarn:     false,
	}, {
		Name:         "Проверить ответ",
		ID:           53,
		Regex:        "проверить Ваше решение",
		TakePoints:   true,
		ReturnPoints: true,
		WithWarn:     false,
	}, {
		Name:         "Изменить язык",
		ID:           59,
		Regex:        "не относящимся к литературе и языку",
		TakePoints:   true,
		ReturnPoints: true,
		WithWarn:     false,
	}, {
		Name:         "Оскорбление",
		ID:           6,
		Regex:        "из-за оскорбительного содержания",
		TakePoints:   true,
		ReturnPoints: false,
		WithWarn:     true,
	}, {
		Name:         "Ссылки",
		ID:           23,
		Regex:        "понятный и полный вопрос без ссылок",
		TakePoints:   true,
		ReturnPoints: false,
		WithWarn:     false,
	}, {
		Name:         "Учебник",
		ID:           58,
		Regex:        "перепишите свой вопрос из учебника",
		TakePoints:   true,
		ReturnPoints: true,
		WithWarn:     false,
	}, {
		Name:         "По просьбе",
		ID:           37,
		Regex:        "по просьбе пользователя",
		TakePoints:   false,
		ReturnPoints: true,
		WithWarn:     false,
	}, {
		Name:         "Не тот предмет",
		ID:           39,
		Regex:        "соответствующем ему предмете",
		TakePoints:   true,
		ReturnPoints: true,
		WithWarn:     false,
	}, {
		Name:         "Не школа",
		ID:           61,
		Regex:        "отношение к сфере образования",
		TakePoints:   true,
		ReturnPoints: false,
		WithWarn:     false,
	}, {
		Name:       "Культура (ответ)",
		ID:         41,
		Regex:      "ответ дан в грубой форме",
		TakePoints: true,
		WithWarn:   true,
	}, {
		Name:       "Пропало приложение",
		ID:         52,
		Regex:      "из-за сбоя в работе Сервиса решение",
		TakePoints: false,
		WithWarn:   false,
	}, {
		Name:       "Спам (ответ)",
		ID:         15,
		Regex:      "запись не является ответом на заданный вопрос",
		TakePoints: true,
		WithWarn:   true,
	}, {
		Name:       "Ссылки (ответ)",
		ID:         48,
		Regex:      "ссылки на другие сайты или издания в ответах",
		TakePoints: true,
		WithWarn:   true,
	}, {
		Name:       "Не ответ",
		ID:         51,
		Regex:      "комментарий не является ответом",
		TakePoints: true,
		WithWarn:   true,
	}, {
		Name:              "Реклама (ответ)",
		ID:                62,
		Regex:             "реклама в вопросах",
		TakePoints:        true,
		WithWarn:          true,
		UserMustBeDeleted: true,
	}, {
		Name:       "Копия из калькулятора",
		ID:         63,
		Regex:      "скопирован с онлайн калькулятора",
		TakePoints: true,
		WithWarn:   true,
	}, {
		Name:       "Плагиат",
		ID:         55,
		Regex:      "на другие сайты или издания и текстов",
		TakePoints: true,
		WithWarn:   true,
	}, {
		Name:       "Неверный ответ",
		ID:         7,
		Regex:      "ответ содержит ошибки",
		TakePoints: true,
		WithWarn:   false,
	}, {
		Name:       "Нет расчётов",
		ID:         9,
		Regex:      "не содержит шагов решения",
		TakePoints: true,
		WithWarn:   false,
	}, {
		Name:       "Неполный ответ",
		ID:         35,
		Regex:      "не является полным и исчерпывающим",
		TakePoints: true,
		WithWarn:   false,
	}, {
		Name:       "Переводчик",
		ID:         42,
		Regex:      "с помощью онлайн-переводчика",
		TakePoints: true,
		WithWarn:   false,
	}, {
		Name:       "Нет приложения (ответ)",
		ID:         56,
		Regex:      "рисунка, схемы или чертежа",
		TakePoints: true,
		WithWarn:   false,
	}, {
		Name:       "Не тот ответ",
		ID:         60,
		Regex:      "соответствует заданному вопросу",
		TakePoints: true,
		WithWarn:   false,
	}, {
		Name:       "Оскорбления (ответ)",
		ID:         16,
		Regex:      "ненормативная лексика недопустимы",
		TakePoints: true,
		WithWarn:   true,
	},
	{
		Name:       "Комментарий",
		ID:         44,
		Regex:      "вынуждены удалить Ваш комментарий",
		TakePoints: true,
		WithWarn:   true,
	},
	// extra deletion reasons
	{
		Name:              "Деньги за решение",
		ID:                100,
		Regex:             "денег за решение задачи",
		TakePoints:        true,
		WithWarn:          true,
		UserMustBeDeleted: true,
	},
	{
		Name:       "Время на исправление истекло",
		ID:         101,
		Regex:      "на исправления задания истекло",
		TakePoints: true,
		WithWarn:   false,
	},
	// Brainly Tools deletion reasons
	{
		Name:         "Не тот раздел (BT)",
		ID:           1112,
		Regex:        "опубликовал свой контент не в том разделе",
		TakePoints:   true,
		WithWarn:     false,
		ReturnPoints: true,
	},
	{
		Name:       "Вредоносный ответ (BT)",
		ID:         2103,
		Regex:      "#BT2103",
		TakePoints: true,
		WithWarn:   true,
	},
	{
		Name:       "Нарушение (вопрос, BT)",
		ID:         1101,
		Regex:      "#BT1101",
		TakePoints: true,
	},
	{
		Name:       "Нарушение (ответ, BT)",
		ID:         2102,
		Regex:      "#BT2102",
		TakePoints: true,
	},
}

func GetShortDeleteReason(longReason string) DeletionReason {
	for _, reason := range deletionReasons {
		if StringContains(reason.Regex, longReason) {
			return reason
		}
	}

	return DeletionReason{}
}