package producer

import (
	"github.com/zp857/goutil/errorx"
	"github.com/zp857/goutil/kafka"
	"github.com/zp857/goutil/structs"
	"go.uber.org/zap"
	"net/http"
)

func SendJSON(topic string, data interface{}, p *kafka.Producer) {
	p.SendJSON(topic, data)
}

func SendWithResponse(topic string, code int, data interface{}, msg string, p *kafka.Producer) {
	response := structs.Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	SendJSON(topic, response, p)
}

func SendOk(topic string, data interface{}, msg string, p *kafka.Producer) {
	SendWithResponse(topic, http.StatusOK, data, msg, p)
}

func SendOkWithMessage(topic string, msg string, p *kafka.Producer) {
	SendWithResponse(topic, http.StatusOK, struct{}{}, msg, p)
}

func SendErrorWithMessage(topic string, msg string, err error, p *kafka.Producer) {
	if err != nil {
		zap.L().Named("[kafka-producer]").Error(
			msg,
			zap.Error(err),
			zap.Any("stack", string(errorx.GetStack(2, 5))),
		)
	}
	SendWithResponse(topic, http.StatusInternalServerError, struct{}{}, msg, p)
}

func BadRequestWithMessage(topic string, msg string, p *kafka.Producer) {
	SendWithResponse(topic, http.StatusBadRequest, struct{}{}, msg, p)
}
