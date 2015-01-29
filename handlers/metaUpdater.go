package handlers

import (
	sl "github.com/djosephsen/slacker/slackerlib"
)

var MetaUpdater = sl.EventHandler{
	Name: `Meta Updater`,
	Usage:`keeps Sbot.Meta up to date using event traffic from Slackhq`,
	Type: `*`,
	Run:		metaUpdaterFunc,
}

func metaUpdaterFunc (hp *sl.HandlerPackage){
	sl.Logger.Debug(`MetaUpdater:: got a: `, hp.Type)
}
