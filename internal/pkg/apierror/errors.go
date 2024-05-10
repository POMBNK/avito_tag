package apierror

var ErrNotFound = New(404, "Запись не найдена", nil)
var ErrIvalidData = New(422, "Невалидные данные", nil)
