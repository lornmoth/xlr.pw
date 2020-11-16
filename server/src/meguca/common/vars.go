package common

import (
	"regexp"
)

var (
	// GetVideoNames is a forwarded function from "meguca/assets" to avoid circular imports
	GetVideoNames func() []string
	// Recompile is a forwarded function from "meguca/templates" to avoid circular imports
	Recompile func() error
)

// Maximum lengths of various input fields
const (
	MaxLenName         = 50
	MaxLenAuth         = 50
	MaxLenPostPassword = 100
	MaxLenSubject      = 100
	MaxLenBody         = 2000
	MaxLinesBody       = 100
	MaxLenPassword     = 50
	MaxLenUserID       = 20
	MaxLenBoardID      = 10
	MaxLenBoardTitle   = 100
	MaxLenNotice       = 500
	MaxLenRules        = 5000
	MaxLenEightball    = 2000
	MaxLenReason       = 100
	MaxNumBanners      = 20
	MaxAssetSize       = 100 << 10
	MaxDiceSides       = 10000
	BumpLimit          = 5000
)

// Various cryptographic token exact lengths
const (
	LenSession    = 171
	LenImageToken = 86
)

// Available language packs and themes. Change this, when adding any new ones.
var (
	Langs = []string{
		"en_GB", "es_ES", "fr_FR", "pl_PL", "pt_BR", "sk_SK", "tr_TR", "uk_UA",
		"ru_RU",
	}
	Themes = []string{
	 	"console", "erowid", "fauux", "terminal", "tachibana",
		"moon", "a_xr", "tea", "yots_b",
	}
)

// Common Regex expressions
var (
	CommandRegexp = regexp.MustCompile(`^#(flip|\d*d\d+|8ball|pyu|pcount|sw(?:\d+:)?\d+:\d+(?:[+-]\d+)?|roulette|rcount)$`)
	DiceRegexp    = regexp.MustCompile(`(\d*)d(\d+)`)
)
