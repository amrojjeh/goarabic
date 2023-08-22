package goarabic

// Harf holds the Arabic character with its different representation forms (glyphs).
type Harf struct {
	Base, Isolated, Beginning, Medial, Final rune
}

// Vowels (Tashkeel) characters.
const (
	FATHA            = '\u064e'
	FATHATAN         = '\u064b'
	DAMMA            = '\u064f'
	DAMMATAN         = '\u064c'
	KASRA            = '\u0650'
	KASRATAN         = '\u064d'
	SHADDA           = '\u0651'
	SUKUN            = '\u0652'
	SUPERSCRIPT_ALEF = '\u0670'
)

// Arabic Alphabet using the Harf type.
var (
	ALEF_WASLA = Harf{ // ٱ
		Base:      '\u0671',
		Isolated:  '\ufb50',
		Beginning: '\u0671',
		Medial:    '\u0671',
		Final:     '\ufb61'}

	ALEF_HAMZA_ABOVE = Harf{ // أ
		Base:      '\u0623',
		Isolated:  '\ufe83',
		Beginning: '\u0623',
		Medial:    '\ufe84',
		Final:     '\ufe84'}

	ALEF = Harf{ // ا
		Base:      '\u0627',
		Isolated:  '\ufe8d',
		Beginning: '\u0627',
		Medial:    '\ufe8e',
		Final:     '\ufe8e'}

	ALEF_MADDA_ABOVE = Harf{ // آ
		Base:      '\u0622',
		Isolated:  '\ufe81',
		Beginning: '\u0622',
		Medial:    '\ufe82',
		Final:     '\ufe82'}

	HAMZA = Harf{ // ء
		Base:      '\u0621',
		Isolated:  '\ufe80',
		Beginning: '\u0621',
		Medial:    '\u0621',
		Final:     '\u0621'}

	WAW_HAMZA_ABOVE = Harf{ // ؤ
		Base:      '\u0624',
		Isolated:  '\ufe85',
		Beginning: '\u0624',
		Medial:    '\ufe86',
		Final:     '\ufe86'}

	ALEF_HAMZA_BELOW = Harf{ // أ
		Base:      '\u0625',
		Isolated:  '\ufe87',
		Beginning: '\u0625',
		Medial:    '\ufe88',
		Final:     '\ufe88'}

	YEH_HAMZA_ABOVE = Harf{ // ئ
		Base:      '\u0626',
		Isolated:  '\ufe89',
		Beginning: '\ufe8b',
		Medial:    '\ufe8c',
		Final:     '\ufe8a'}

	BEH = Harf{ // ب
		Base:      '\u0628',
		Isolated:  '\ufe8f',
		Beginning: '\ufe91',
		Medial:    '\ufe92',
		Final:     '\ufe90'}

	PEH = Harf{ // پ
		Base:      '\u067e',
		Isolated:  '\ufb56',
		Beginning: '\ufb58',
		Medial:    '\ufb59',
		Final:     '\ufb57'}

	TEH = Harf{ // ت
		Base:      '\u062A',
		Isolated:  '\ufe95',
		Beginning: '\ufe97',
		Medial:    '\ufe98',
		Final:     '\ufe96'}

	TEH_MARBUTA = Harf{ // ة
		Base:      '\u0629',
		Isolated:  '\ufe93',
		Beginning: '\u0629',
		Medial:    '\u0629',
		Final:     '\ufe94'}

	THEH = Harf{ // ث
		Base:      '\u062b',
		Isolated:  '\ufe99',
		Beginning: '\ufe9b',
		Medial:    '\ufe9c',
		Final:     '\ufe9a'}

	JEEM = Harf{ // ج
		Base:      '\u062c',
		Isolated:  '\ufe9d',
		Beginning: '\ufe9f',
		Medial:    '\ufea0',
		Final:     '\ufe9e'}

	TCHEH = Harf{ // چ
		Base:      '\u0686',
		Isolated:  '\ufb7a',
		Beginning: '\ufb7c',
		Medial:    '\ufb7d',
		Final:     '\ufb7b'}

	HAH = Harf{ // ح
		Base:      '\u062d',
		Isolated:  '\ufea1',
		Beginning: '\ufea3',
		Medial:    '\ufea4',
		Final:     '\ufea2'}

	KHAH = Harf{ // خ
		Base:      '\u062e',
		Isolated:  '\ufea5',
		Beginning: '\ufea7',
		Medial:    '\ufea8',
		Final:     '\ufea6'}

	DAL = Harf{ // د
		Base:      '\u062f',
		Isolated:  '\ufea9',
		Beginning: '\u062f',
		Medial:    '\ufeaa',
		Final:     '\ufeaa'}

	THAL = Harf{ // ذ
		Base:      '\u0630',
		Isolated:  '\ufeab',
		Beginning: '\u0630',
		Medial:    '\ufeac',
		Final:     '\ufeac'}

	REH = Harf{ // ر
		Base:      '\u0631',
		Isolated:  '\ufead',
		Beginning: '\u0631',
		Medial:    '\ufeae',
		Final:     '\ufeae'}

	JEH = Harf{
		Base:      '\u0698',
		Isolated:  '\ufb8a',
		Beginning: '\u0698',
		Medial:    '\ufb8b',
		Final:     '\ufb8b',
	}

	ZAIN = Harf{ // ز
		Base:      '\u0632',
		Isolated:  '\ufeaf',
		Beginning: '\u0632',
		Medial:    '\ufeb0',
		Final:     '\ufeb0'}

	SEEN = Harf{ // س
		Base:      '\u0633',
		Isolated:  '\ufeb1',
		Beginning: '\ufeb3',
		Medial:    '\ufeb4',
		Final:     '\ufeb2'}

	SHEEN = Harf{ // ش
		Base:      '\u0634',
		Isolated:  '\ufeb5',
		Beginning: '\ufeb7',
		Medial:    '\ufeb8',
		Final:     '\ufeb6'}

	SAD = Harf{ // ص
		Base:      '\u0635',
		Isolated:  '\ufeb9',
		Beginning: '\ufebb',
		Medial:    '\ufebc',
		Final:     '\ufeba'}

	DAD = Harf{ // ض
		Base:      '\u0636',
		Isolated:  '\ufebd',
		Beginning: '\ufebf',
		Medial:    '\ufec0',
		Final:     '\ufebe'}

	TAH = Harf{ // ط
		Base:      '\u0637',
		Isolated:  '\ufec1',
		Beginning: '\ufec3',
		Medial:    '\ufec4',
		Final:     '\ufec2'}

	ZAH = Harf{ // ظ
		Base:      '\u0638',
		Isolated:  '\ufec5',
		Beginning: '\ufec7',
		Medial:    '\ufec8',
		Final:     '\ufec6'}

	AIN = Harf{ // ع
		Base:      '\u0639',
		Isolated:  '\ufec9',
		Beginning: '\ufecb',
		Medial:    '\ufecc',
		Final:     '\ufeca'}

	GHAIN = Harf{ // غ
		Base:      '\u063a',
		Isolated:  '\ufecd',
		Beginning: '\ufecf',
		Medial:    '\ufed0',
		Final:     '\ufece'}

	FEH = Harf{ // ف
		Base:      '\u0641',
		Isolated:  '\ufed1',
		Beginning: '\ufed3',
		Medial:    '\ufed4',
		Final:     '\ufed2'}

	QAF = Harf{ // ق
		Base:      '\u0642',
		Isolated:  '\ufed5',
		Beginning: '\ufed7',
		Medial:    '\ufed8',
		Final:     '\ufed6'}

	VEH = Harf{ // ڤ
		Base:      '\u06a4',
		Isolated:  '\ufb6a',
		Beginning: '\ufb6c',
		Medial:    '\ufb6d',
		Final:     '\ufb6b'}

	KAF = Harf{ // ك
		Base:      '\u0643',
		Isolated:  '\ufed9',
		Beginning: '\ufedb',
		Medial:    '\ufedc',
		Final:     '\ufeda'}

	KEHEH = Harf{ // ک
		Base:      '\u06a9',
		Isolated:  '\ufb8e',
		Beginning: '\ufb90',
		Medial:    '\ufb91',
		Final:     '\ufb8f',
	}

	GAF = Harf{ // گ
		Base:      '\u06af',
		Isolated:  '\ufb92',
		Beginning: '\ufb94',
		Medial:    '\ufb95',
		Final:     '\ufb93'}

	LAM = Harf{ // ل
		Base:      '\u0644',
		Isolated:  '\ufedd',
		Beginning: '\ufedf',
		Medial:    '\ufee0',
		Final:     '\ufede'}

	MEEM = Harf{ // م
		Base:      '\u0645',
		Isolated:  '\ufee1',
		Beginning: '\ufee3',
		Medial:    '\ufee4',
		Final:     '\ufee2'}

	NOON = Harf{ // ن
		Base:      '\u0646',
		Isolated:  '\ufee5',
		Beginning: '\ufee7',
		Medial:    '\ufee8',
		Final:     '\ufee6'}

	HEH = Harf{ // ه
		Base:      '\u0647',
		Isolated:  '\ufee9',
		Beginning: '\ufeeb',
		Medial:    '\ufeec',
		Final:     '\ufeea'}

	WAW = Harf{ // و
		Base:      '\u0648',
		Isolated:  '\ufeed',
		Beginning: '\u0648',
		Medial:    '\ufeee',
		Final:     '\ufeee'}

	FARSIYEH = Harf{ // ی
		Base:      '\u06cc',
		Isolated:  '\ufbfc',
		Beginning: '\ufbfe',
		Medial:    '\ufbff',
		Final:     '\ufbfd'}

	YEH = Harf{ // ي
		Base:      '\u064a',
		Isolated:  '\ufef1',
		Beginning: '\ufef3',
		Medial:    '\ufef4',
		Final:     '\ufef2'}

	ALEF_MAKSURA = Harf{ // ى
		Base:      '\u0649',
		Isolated:  '\ufeef',
		Beginning: '\u0649',
		Medial:    '\ufef0',
		Final:     '\ufef0'}

	TATWEEL = Harf{ // ـ
		Base:      '\u0640',
		Isolated:  '\u0640',
		Beginning: '\u0640',
		Medial:    '\u0640',
		Final:     '\u0640'}

	LAM_ALEF = Harf{ // لا
		Base:      '\ufefb',
		Isolated:  '\ufefb',
		Beginning: '\ufefb',
		Medial:    '\ufefc',
		Final:     '\ufefc'}

	LAM_ALEF_HAMZA_ABOVE = Harf{ // ﻷ
		Base:      '\ufef7',
		Isolated:  '\ufef7',
		Beginning: '\ufef7',
		Medial:    '\ufef8',
		Final:     '\ufef8'}
)

var alphabet = []Harf{
	ALEF_WASLA,
	ALEF_HAMZA_ABOVE,
	ALEF,
	ALEF_MADDA_ABOVE,
	HAMZA,
	WAW_HAMZA_ABOVE,
	ALEF_HAMZA_BELOW,
	YEH_HAMZA_ABOVE,
	BEH,
	PEH,
	TEH,
	TEH_MARBUTA,
	THEH,
	JEEM,
	TCHEH,
	HAH,
	KHAH,
	DAL,
	THAL,
	REH,
	JEH,
	ZAIN,
	SEEN,
	SHEEN,
	SAD,
	DAD,
	TAH,
	ZAH,
	AIN,
	GHAIN,
	FEH,
	QAF,
	VEH,
	KAF,
	KEHEH,
	GAF,
	LAM,
	MEEM,
	NOON,
	HEH,
	WAW,
	YEH,
	FARSIYEH,
	ALEF_MAKSURA,
	TATWEEL,
	LAM_ALEF,
	LAM_ALEF_HAMZA_ABOVE,
}

// use map for faster lookups.
var tashkeel = map[rune]bool{FATHA: true, FATHATAN: true, DAMMA: true,
	DAMMATAN: true, KASRA: true, KASRATAN: true,
	SHADDA: true, SUKUN: true}

// use map for faster lookups.
// var special_char = map[rune]bool{"": true, ' ': true, '?': true,
//	'؟': true, '.': true, KASRATAN: true,
//	SHADDA: true, SUKUN: true}

// use map for faster lookups.
var beggining_after = map[Harf]bool{
	ALEF_HAMZA_ABOVE: true,
	ALEF_MADDA_ABOVE: true,
	ALEF:             true,
	HAMZA:            true,
	WAW_HAMZA_ABOVE:  true,
	ALEF_HAMZA_BELOW: true,
	TEH_MARBUTA:      true,
	DAL:              true,
	THAL:             true,
	REH:              true,
	ZAIN:             true,
	WAW:              true,
	ALEF_MAKSURA:     true}
