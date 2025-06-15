package middleware

import (
	"log"
	"runtime/debug"

	"github.com/getsentry/sentry-go"
	"gopkg.in/telebot.v4"
)

func Recover() telebot.MiddlewareFunc{
	return func(hf telebot.HandlerFunc) telebot.HandlerFunc {
		return func(ctx telebot.Context) error {
			defer func ()  {
				if r := recover(); r != nil {
					log.Printf("捕获 panic %v\n%s", r, debug.Stack())
					sentry.CurrentHub().Recover(r)
				}
			}()
			return hf(ctx)
		}
	}
}