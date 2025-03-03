package logging

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"spe_test/utils/constant"
	"time"
)

func LogRequest(ctx *fiber.Ctx, requestData, uuidTr, uuidTc string) time.Time {
	log.Printf("[INFO]\tTR-%s\t%s\t%s\t%s\t%s\tTC-%s\t%s\t[START]",
		uuidTr, constant.APP_NAME, ctx.Route().Path, ctx.Method(), ctx.IP(), uuidTc, requestData)

	return time.Now()
}

// [INFO] ... [STOP]
func LogResponse(ctx *fiber.Ctx, responseData, uuidTr, uuidTc string, start time.Time) time.Time {
	log.Printf("[INFO]\tTR-%s\t%s\t%s\t%s\t%s\t%s\tTC-%s\t%s\t[STOP]",
		uuidTr, constant.APP_NAME, ctx.Route().Path, ctx.Method(), time.Since(start), ctx.IP(), uuidTc, responseData)

	return time.Now()
}

// [TRACE] | [INFO] | [WARN] | [ERROR] | [FATAL]
func Logging(ctx *fiber.Ctx, level, functionName, uniqueCode, transactionId, traceId, notes string) {
	if level != "DEBUG" {
		log.Printf("[%s]\tTR-%s\t%s\t%s\t%s\t%s\tTC-%s\t[%s]\t[%s]\t(%s)",
			level, transactionId, constant.APP_NAME, ctx.Route().Path, ctx.Method(), ctx.IP(), traceId, functionName, uniqueCode, notes)
	} else {
		log.Printf("[%s]\tTR-%s\t%s\t%s\t%s\t%s\tTC-%s\t[%s]\t[%s]\t(%s)",
			"OTHER", transactionId, constant.APP_NAME, ctx.Route().Path, ctx.Method(), ctx.IP(), traceId, functionName, uniqueCode, notes)
	}
}
