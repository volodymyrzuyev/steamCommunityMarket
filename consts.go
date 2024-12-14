package steamcommunityapi

import "fmt"

const BaseWebApiUrl = "https://www.steamcommunity.com"

var (
	CooldownNotPassed         = fmt.Errorf("Cooldown since last request has not passed")
	SteamRateLimitExceeded    = fmt.Errorf("Steam rate limit exceeded")
	UnexpectedSteamStatusCode = fmt.Errorf("Received unexpected status code from Steam")
	InvalidParameter          = fmt.Errorf("Invalid parameter input")
)

type language struct {
	fullname     string
	abbreviation string
}

// source https://partner.steamgames.com/doc/store/localization/languages
var (
	Arabic     = language{"arabic", "ar"}
	Bulgarian  = language{"bulgarian", "bg"}
	SChinese   = language{"schinese", "zh-CN"}
	TChinese   = language{"tchinese", "zh-TW"}
	Czech      = language{"czech", "cs"}
	Danish     = language{"danish", "da"}
	Dutch      = language{"dutch", "nl"}
	English    = language{"english", "en"}
	Finnish    = language{"finnish", "fi"}
	French     = language{"french", "fr"}
	German     = language{"german", "de"}
	Greek      = language{"greek", "el"}
	Hungarian  = language{"hungarian", "hu"}
	Indonesian = language{"indonesian", "id"}
	Italian    = language{"italian", "it"}
	Japanese   = language{"japanese", "ja"}
	Korean     = language{"koreana", "ko"}
	Norwegian  = language{"norwegian", "no"}
	Polish     = language{"polish", "pl"}
	Portuguese = language{"portuguese", "pt"}
	// Brazil mentioned
	PortugueseBrazil = language{"brazilian", "pt-BR"}
	Romanian         = language{"romanian", "ro"}
	SpanishSpain     = language{"spanish", "es"}
	SpanishLA        = language{"latam", "es-419"}
	Swedish          = language{"swedish", "sv"}
	Thai             = language{"thai", "th"}
	Turkish          = language{"turkish", "tr"}
	Ukrainian        = language{"ukrainian", "uk"}
	Vietnamese       = language{"vietnamese", "vn"}
	blankLang        = language{}
)

type currency struct {
	name   string
	code   string
	symbol string
}

// source https://partner.steamgames.com/doc/store/pricing/currencies
var (
	USD      = currency{"United States Dollar", "1", "$"}
	GBP      = currency{"United Kingdom Pound", "2", "£"}
	EUR      = currency{"European Union Euro", "3", "€"}
	CHF      = currency{"Swiss Francs", "4", "₣"}
	PLN      = currency{"Polish Złoty", "6", "zł"}
	BRL      = currency{"Brazilian Reals", "7", "R$"}
	JPY      = currency{"Japanese Yen", "8", "¥"}
	NOK      = currency{"Norwegian Krone", "9", "Kr"}
	IDR      = currency{"Indonesian Rupiah", "10", "Rp"}
	MYR      = currency{"Malaysian Ringgit", "11", "RM"}
	PHP      = currency{"Philippine Peso", "12", "₱"}
	SGD      = currency{"Singapore Dollar", "13", "$"}
	THB      = currency{"Thai Baht", "14", "฿"}
	VND      = currency{"Vietnamese Dong", "15", "đ"}
	KRW      = currency{"South Korean Won", "16", "₩"}
	UAH      = currency{"Ukrainian Hryvnia", "18", "₴"}
	MXN      = currency{"Mexican Peso", "19", "¢"}
	CAD      = currency{"Canadian Dollars", "20", "$"}
	AUD      = currency{"Australian Dollars", "21", "$"}
	NZD      = currency{"New Zealand Dollar", "22", "$"}
	CNY      = currency{"Chinese Renminbi", "23", "¥"}
	INR      = currency{"Indian Rupee", "24", "₹"}
	CLP      = currency{"Chilean Peso", "25", "$"}
	PEN      = currency{"Peruvian Sol", "26", "S/"}
	COP      = currency{"Colombian Peso", "27", "$"}
	ZAR      = currency{"South African Rand", "28", "R"}
	HKD      = currency{"Hong Kong Dollar", "29", "$"}
	TWD      = currency{"New Taiwan Dollar", "30", "$"}
	SAR      = currency{"Saudi Riyal", "31", "SR"}
	AED      = currency{"United Arab Emirates Dirham", "32", "Dh"}
	ILS      = currency{"Israeli New Shekel", "35", "₪"}
	KZT      = currency{"Kazakhstani Tenge", "37", "₸"}
	KWD      = currency{"Kuwaiti Dinar", "38", "د"}
	QAR      = currency{"Qatari Riyal", "39", "QR"}
	CRC      = currency{"Costa Rican Colón", "40", "₡"}
	UYU      = currency{"Uruguayan Peso", "41", "$"}
	blankCur = currency{}
)

type country struct {
	fullName string
	code     string
}

var (
	Argentina     = country{"Argentina", "AR"}
	Brazil        = country{"Brazil", "BR"}
	Belgium       = country{"Belgium", "BE"}
	China         = country{"China", "CN"}
	Chile         = country{"Chile", "CL"}
	Canada        = country{"Canada", "CA"}
	Germany       = country{"Germany", "DE"}
	GreatBritain  = country{"Great Britain", "GB"}
	Italy         = country{"Italy", "IT"}
	Netherlands   = country{"Netherlands", "NL"}
	UnitedStates  = country{"United States", "US"}
	Sweden        = country{"Sweden", "SE"}
	Switzerland   = country{"Switzerland", "CH"}
	Taiwan        = country{"Taiwan", "TW"}
	Thailand      = country{"Thailand", "TH"}
	Ukraine       = country{"Ukraine", "UA"}
	UnitedKingdom = country{"United Kingdon", "UK"}
	blankCountry  = country{}
)
