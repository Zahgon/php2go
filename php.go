// php2go functions

package php2go

import (
	"archive/zip"
	"encoding/binary"
	"net/url"
	"os"
)

//////////// Date/Time Functions ////////////

// Time time()
func Time() int64 { _ = "STUB: not implemented"; return 0 }

// Strtotime strtotime()
// Strtotime("02/01/2006 15:04:05", "02/01/2016 15:04:05") == 1451747045
// Strtotime("3 04 PM", "8 41 PM") == -62167144740
func Strtotime(format, strtime string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Date date()
// Date("02/01/2006 15:04:05 PM", 1524799394)
// Note: the behavior is inconsistent with php's date function
func Date(format string, timestamp int64) string { _ = "STUB: not implemented"; return "" }

// Checkdate checkdate()
// Validate a Gregorian date
func Checkdate(month, day, year int) bool { _ = "STUB: not implemented"; return false }

// leap year

// Sleep sleep()
func Sleep(t int64) { _ = "STUB: not implemented"; return }

// Usleep usleep()
func Usleep(t int64) { _ = "STUB: not implemented"; return }

//////////// String Functions ////////////

// Strpos strpos()
func Strpos(haystack, needle string, offset int) int { _ = "STUB: not implemented"; return 0 }

// Stripos stripos()
func Stripos(haystack, needle string, offset int) int { _ = "STUB: not implemented"; return 0 }

// Strrpos strrpos()
func Strrpos(haystack, needle string, offset int) int { _ = "STUB: not implemented"; return 0 }

// Strripos strripos()
func Strripos(haystack, needle string, offset int) int { _ = "STUB: not implemented"; return 0 }

// StrReplace str_replace()
func StrReplace(search, replace, subject string, count int) string {
	_ = "STUB: not implemented"
	return ""
}

// Strtoupper strtoupper()
func Strtoupper(str string) string { _ = "STUB: not implemented"; return "" }

// Strtolower strtolower()
func Strtolower(str string) string { _ = "STUB: not implemented"; return "" }

// Ucfirst ucfirst()
func Ucfirst(str string) string { _ = "STUB: not implemented"; return "" }

// Lcfirst lcfirst()
func Lcfirst(str string) string { _ = "STUB: not implemented"; return "" }

// Ucwords ucwords()
func Ucwords(str string) string { _ = "STUB: not implemented"; return "" }

// Substr substr()
func Substr(str string, start uint, length int) string { _ = "STUB: not implemented"; return "" }

// Strrev strrev()
func Strrev(str string) string { _ = "STUB: not implemented"; return "" }

// ParseStr parse_str()
// f1=m&f2=n -> map[f1:m f2:n]
// f[a]=m&f[b]=n -> map[f:map[a:m b:n]]
// f[a][a]=m&f[a][b]=n -> map[f:map[a:map[a:m b:n]]]
// f[]=m&f[]=n -> map[f:[m n]]
// f[a][]=m&f[a][]=n -> map[f:map[a:[m n]]]
// f[][]=m&f[][]=n -> map[f:[map[]]] // Currently does not support nested slice.
// f=m&f[a]=n -> error // This is not the same as PHP.
// a .[[b=c -> map[a___[b:c]
func ParseStr(encodedString string, result map[string]interface{}) error {
	_ = "STUB: not implemented"
	// build nested map.
	return nil
}

// trim ',"

// The end is slice. like f[], f[a][]

// todo nested slice

// The end is slice + map. like f[][a]

// map. like f[a], f[a][b]

// split encodedString.

// split into multiple keys

// first key

// build nested map

// NumberFormat number_format()
// decimals: Sets the number of decimal points.
// decPoint: Sets the separator for the decimal point.
// thousandsSep: Sets the thousands' separator.
func NumberFormat(number float64, decimals uint, decPoint, thousandsSep string) string {
	_ = "STUB: not implemented"
	return ""
}

// Will round off

// thousands sep num

// ChunkSplit chunk_split()
func ChunkSplit(body string, chunklen uint, end string) string {
	_ = "STUB: not implemented"
	return ""
}

// StrWordCount str_word_count()
func StrWordCount(str string) []string { _ = "STUB: not implemented"; return nil }

// Wordwrap wordwrap()
func Wordwrap(str string, width uint, br string, cut bool) string {
	_ = "STUB: not implemented"
	return ""
}

// Strlen strlen()
func Strlen(str string) int {
	_ = "STUB: not implemented"

	// StrRepeat str_repeat()
	return 0
}

func StrRepeat(input string, multiplier int) string { _ = "STUB: not implemented"; return "" }

// Strstr strstr()
func Strstr(haystack string, needle string) string { _ = "STUB: not implemented"; return "" }

// Strtr strtr()
//
// If the parameter length is 1, type is: map[string]string
// Strtr("baab", map[string]string{"ab": "01"}) will return "ba01"
// If the parameter length is 2, type is: string, string
// Strtr("baab", "ab", "01") will return "1001", a => 0; b => 1.
func Strtr(haystack string, params ...interface{}) string { _ = "STUB: not implemented"; return "" }

// trlen != 1

// StrShuffle str_shuffle()
func StrShuffle(str string) string { _ = "STUB: not implemented"; return "" }

// Trim trim()
func Trim(str string, characterMask ...string) string { _ = "STUB: not implemented"; return "" }

// Ltrim ltrim()
func Ltrim(str string, characterMask ...string) string { _ = "STUB: not implemented"; return "" }

// Rtrim rtrim()
func Rtrim(str string, characterMask ...string) string { _ = "STUB: not implemented"; return "" }

// Explode explode()
func Explode(delimiter, str string) []string { _ = "STUB: not implemented"; return nil }

// Chr chr()
func Chr(ascii int) string { _ = "STUB: not implemented"; return "" }

// Ord ord()
func Ord(char string) int { _ = "STUB: not implemented"; return 0 }

// Nl2br nl2br()
// \n\r, \r\n, \r, \n
func Nl2br(str string, isXhtml bool) string { _ = "STUB: not implemented"; return "" }

// JSONDecode json_decode()
func JSONDecode(data []byte, val interface{}) error { _ = "STUB: not implemented"; return nil }

// JSONEncode json_encode()
func JSONEncode(val interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Addslashes addslashes()
		nil
}

func Addslashes(str string) string { _ = "STUB: not implemented"; return "" }

// Stripslashes stripslashes()
func Stripslashes(str string) string { _ = "STUB: not implemented"; return "" }

// Quotemeta quotemeta()
func Quotemeta(str string) string { _ = "STUB: not implemented"; return "" }

// Htmlentities htmlentities()
func Htmlentities(str string) string { _ = "STUB: not implemented"; return "" }

// HTMLEntityDecode html_entity_decode()
func HTMLEntityDecode(str string) string { _ = "STUB: not implemented"; return "" }

// Md5 md5()
func Md5(str string) string { _ = "STUB: not implemented"; return "" }

// Md5File md5_file()
func Md5File(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// 1M

// Sha1 sha1()
func Sha1(str string) string { _ = "STUB: not implemented"; return "" }

// Sha1File sha1_file()
func Sha1File(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Crc32 crc32()
func Crc32(str string) uint32 { _ = "STUB: not implemented"; return 0 }

// Levenshtein levenshtein()
// costIns: Defines the cost of insertion.
// costRep: Defines the cost of replacement.
// costDel: Defines the cost of deletion.
func Levenshtein(str1, str2 string, costIns, costRep, costDel int) int {
	_ = "STUB: not implemented"
	return 0
}

// SimilarText similar_text()
func SimilarText(first, second string, percent *float64) int { _ = "STUB: not implemented"; return 0 }

// Find the longest segment of the same section in two strings

// Soundex soundex()
// Calculate the soundex key of a string.
func Soundex(str string) string { _ = "STUB: not implemented"; return "" }

// A, B, C, D

// E, F, G

// H

// I, J, K, L, M, N

// O, P, Q, R, S, T

// U, V

// W, X

// Y, Z

// build soundex string

// ToUpper

// pad with "0"

//////////// Multibyte String Functions ////////////

// MbStrlen mb_strlen()
func MbStrlen(str string) int { _ = "STUB: not implemented"; return 0 }

// MbStrtoupper mb_strtoupper()
// Make a string uppercase
func MbStrtoupper(str string) string { _ = "STUB: not implemented"; return "" }

//////////// URL Functions ////////////

// ParseURL parse_url()
// Parse a URL and return its components
// -1: all; 1: scheme; 2: host; 4: port; 8: user; 16: pass; 32: path; 64: query; 128: fragment
func ParseURL(str string, component int) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// URLEncode urlencode()
func URLEncode(str string) string { _ = "STUB: not implemented"; return "" }

// URLDecode urldecode()
func URLDecode(str string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Rawurlencode rawurlencode()
func Rawurlencode(str string) string { _ = "STUB: not implemented"; return "" }

// Rawurldecode rawurldecode()
func Rawurldecode(str string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// HTTPBuildQuery http_build_query()
func HTTPBuildQuery(queryData url.Values) string { _ = "STUB: not implemented"; return "" }

// Base64Encode base64_encode()
func Base64Encode(str string) string { _ = "STUB: not implemented"; return "" }

// Base64Decode base64_decode()
func Base64Decode(str string) (string, error) { _ = "STUB: not implemented"; return "", nil }

//////////// Array(Slice/Map) Functions ////////////

// ArrayFill array_fill()
func ArrayFill(startIndex int, num uint, value interface{}) map[int]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ArrayFlip array_flip()
func ArrayFlip(m map[interface{}]interface{}) map[interface{}]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ArrayKeys array_keys()
func ArrayKeys(elements map[interface{}]interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ArrayValues array_values()
func ArrayValues(elements map[interface{}]interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ArrayMerge array_merge()
func ArrayMerge(ss ...[]interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

// ArrayChunk array_chunk()
func ArrayChunk(s []interface{}, size int) [][]interface{} { _ = "STUB: not implemented"; return nil }

// ArrayPad array_pad()
func ArrayPad(s []interface{}, size int, val interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ArraySlice array_slice()
func ArraySlice(s []interface{}, offset, length uint) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ArrayRand array_rand()
func ArrayRand(elements []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

// ArrayColumn array_column()
func ArrayColumn(input map[string]map[string]interface{}, columnKey string) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ArrayPush array_push()
// Push one or more elements onto the end of slice
func ArrayPush(s *[]interface{}, elements ...interface{}) int { _ = "STUB: not implemented"; return 0 }

// ArrayPop array_pop()
// Pop the element off the end of slice
func ArrayPop(s *[]interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// ArrayUnshift array_unshift()
// Prepend one or more elements to the beginning of a slice
func ArrayUnshift(s *[]interface{}, elements ...interface{}) int {
	_ = "STUB: not implemented"
	return 0
}

// ArrayShift array_shift()
// Shift an element off the beginning of slice
func ArrayShift(s *[]interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// ArrayKeyExists array_key_exists()
func ArrayKeyExists(key interface{}, m map[interface{}]interface{}) bool {
	_ = "STUB: not implemented"
	return false

	// ArrayCombine array_combine()
}

func ArrayCombine(s1, s2 []interface{}) map[interface{}]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ArrayReverse array_reverse()
func ArrayReverse(s []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

// Implode implode()
func Implode(glue string, pieces []string) string { _ = "STUB: not implemented"; return "" }

// InArray in_array()
// haystack supported types: slice, array or map
func InArray(needle interface{}, haystack interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

//////////// Mathematical Functions ////////////

// Abs abs()
func Abs(number float64) float64 { _ = "STUB: not implemented"; return 0 }

// Rand rand()
// Range: [0, 2147483647]
func Rand(min, max int) int { _ = "STUB: not implemented"; return 0 }

// PHP: getrandmax()

// Round round()
func Round(value float64, precision int) float64 { _ = "STUB: not implemented"; return 0 }

// Floor floor()
func Floor(value float64) float64 { _ = "STUB: not implemented"; return 0 }

// Ceil ceil()
func Ceil(value float64) float64 { _ = "STUB: not implemented"; return 0 }

// Pi pi()
func Pi() float64 {
	_ = "STUB: not implemented"

	// Max max()
	return 0
}

func Max(nums ...float64) float64 { _ = "STUB: not implemented"; return 0 }

// Min min()
func Min(nums ...float64) float64 { _ = "STUB: not implemented"; return 0 }

// Decbin decbin()
func Decbin(number int64) string { _ = "STUB: not implemented"; return "" }

// Bindec bindec()
func Bindec(str string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Hex2bin hex2bin()
func Hex2bin(data string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Bin2hex bin2hex()
func Bin2hex(str string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// If input is not binary number

// Dechex dechex()
func Dechex(number int64) string { _ = "STUB: not implemented"; return "" }

// Hexdec hexdec()
func Hexdec(str string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Decoct decoct()
func Decoct(number int64) string { _ = "STUB: not implemented"; return "" }

// Octdec Octdec()
func Octdec(str string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// BaseConvert base_convert()
func BaseConvert(number string, frombase, tobase int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsNan is_nan()
func IsNan(val float64) bool { _ = "STUB: not implemented"; return false }

//////////// CSPRNG Functions ////////////

// RandomBytes random_bytes()
func RandomBytes(length int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// RandomInt random_int()
func RandomInt(min, max int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//////////// Directory/Filesystem Functions ////////////

// Stat stat()
func Stat(filename string) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *

	// Pathinfo pathinfo()
	// -1: all; 1: dirname; 2: basename; 4: extension; 8: filename
	// Usage:
	// Pathinfo("/home/go/path/src/php2go/php2go.go", 1|2|4|8)
	new(os.FileInfo), nil
}

func Pathinfo(path string, options int) map[string]string { _ = "STUB: not implemented"; return nil }

// FileExists file_exists()
func FileExists(filename string) bool { _ = "STUB: not implemented"; return false }

// IsFile is_file()
func IsFile(filename string) bool { _ = "STUB: not implemented"; return false }

// IsDir is_dir()
func IsDir(filename string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// FileSize filesize()
func FileSize(filename string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// FilePutContents file_put_contents()
func FilePutContents(filename string, data string, mode os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

// FileGetContents file_get_contents()
func FileGetContents(filename string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Unlink unlink()
func Unlink(filename string) error { _ = "STUB: not implemented"; return nil }

// Delete delete()
func Delete(filename string) error { _ = "STUB: not implemented"; return nil }

// Copy copy()
func Copy(source, dest string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// IsReadable is_readable()
func IsReadable(filename string) bool { _ = "STUB: not implemented"; return false }

// IsWriteable is_writeable()
func IsWriteable(filename string) bool { _ = "STUB: not implemented"; return false }

// Rename rename()
func Rename(oldname, newname string) error { _ = "STUB: not implemented"; return nil }

// Touch touch()
func Touch(filename string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Mkdir mkdir()
func Mkdir(filename string, mode os.FileMode) error { _ = "STUB: not implemented"; return nil }

// Getcwd getcwd()
func Getcwd() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Realpath realpath()
func Realpath(path string) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// Basename basename()
		nil
}

func Basename(path string) string { _ = "STUB: not implemented"; return "" }

// Chmod chmod()
func Chmod(filename string, mode os.FileMode) bool { _ = "STUB: not implemented"; return false }

// Chown chown()
func Chown(filename string, uid, gid int) bool { _ = "STUB: not implemented"; return false }

// Fclose fclose()
func Fclose(handle *os.File) error { _ = "STUB: not implemented"; return nil }

// Filemtime filemtime()
func Filemtime(filename string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Fgetcsv fgetcsv()
func Fgetcsv(handle *os.File, length int, delimiter rune) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO length limit

// Glob glob()
func Glob(pattern string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

//////////// Variable handling Functions ////////////

// Empty empty()
func Empty(val interface{}) bool { _ = "STUB: not implemented"; return false }

// IsNumeric is_numeric()
// Numeric strings consist of optional sign, any number of digits, optional decimal part and optional exponential part.
// Thus +0123.45e6 is a valid numeric value.
// In PHP hexadecimal (e.g. 0xf4c3b00c) is not supported, but IsNumeric is supported.
func IsNumeric(val interface{}) bool { _ = "STUB: not implemented"; return false }

// Trim any whitespace

// hex

// 0-9, Point, Scientific

// Point

// Scientific

//////////// Program execution Functions ////////////

// Exec exec()
// returnVar, 0: succ; 1: fail
// Return the last line from the result of the command.
// command format eg:
//
//	"ls -a"
//	"/bin/bash -c \"ls -a\""
func Exec(command string, output *[]string, returnVar *int) string {
	_ = "STUB: not implemented"
	return ""
}

// remove the " and ' on both sides

// System system()
// returnVar, 0: succ; 1: fail
// Returns the last line of the command output on success, and "" on failure.
func System(command string, returnVar *int) string { _ = "STUB: not implemented"; return "" }

// split command

// remove the " and ' on both sides

// Passthru passthru()
// returnVar, 0: succ; 1: fail
func Passthru(command string, returnVar *int) { _ = "STUB: not implemented"; return }

// remove the " and ' on both sides

//////////// Network Functions ////////////

// Gethostname gethostname()
func Gethostname() (string, error) {
	_ = "STUB: not implemented"
	return "",

		// Gethostbyname gethostbyname()
		// Get the IPv4 address corresponding to a given Internet host name
		nil
}

func Gethostbyname(hostname string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Gethostbynamel gethostbynamel()
// Get a list of IPv4 addresses corresponding to a given Internet host name
func Gethostbynamel(hostname string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Gethostbyaddr gethostbyaddr()
// Get the Internet host name corresponding to a given IP address
func Gethostbyaddr(ipAddress string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// IP2long ip2long()
// IPv4
func IP2long(ipAddress string) uint32 { _ = "STUB: not implemented"; return 0 }

// Long2ip long2ip()
// IPv4
func Long2ip(properAddress uint32) string { _ = "STUB: not implemented"; return "" }

//////////// Misc. Functions ////////////

// Echo echo
func Echo(args ...interface{}) {
	_ = "STUB: not implemented"

	// Uniqid uniqid()
	return
}

func Uniqid(prefix string) string { _ = "STUB: not implemented"; return "" }

// Exit exit()
func Exit(status int) {
	_ = "STUB: not implemented"

	// Die die()
	return
}

func Die(status int) {
	_ = "STUB: not implemented"

	// Getenv getenv()
	return
}

func Getenv(varname string) string { _ = "STUB: not implemented"; return "" }

// Putenv putenv()
// The setting, like "FOO=BAR"
func Putenv(setting string) error { _ = "STUB: not implemented"; return nil }

// MemoryGetUsage memory_get_usage()
// return in bytes
func MemoryGetUsage(realUsage bool) uint64 { _ = "STUB: not implemented"; return 0 }

// MemoryGetPeakUsage memory_get_peak_usage()
// return in bytes
func MemoryGetPeakUsage(realUsage bool) uint64 { _ = "STUB: not implemented"; return 0 }

// VersionCompare version_compare()
// The possible operators are: <, lt, <=, le, >, gt, >=, ge, ==, =, eq, !=, <>, ne respectively.
// special version strings these are handled in the following order,
// (any string not found) < dev < alpha = a < beta = b < RC = rc < # < pl = p
// Usage:
// VersionCompare("1.2.3-alpha", "1.2.3RC7", '>=')
// VersionCompare("1.2.3-beta", "1.2.3pl", 'lt')
// VersionCompare("1.1_dev", "1.2any", 'eq')
func VersionCompare(version1, version2, operator string) bool {
	_ = "STUB: not implemented"
	return false
}

// version compare

// all is digit

// all digit

// is digit

// canonicalize

// Have the next one

// replace '-', '_', '+' to '.'

// Insert '.' before and after a non-digit

// Non-letters and numbers

// compare special version forms

// (Any string not found) < dev < alpha = a < beta = b < RC = rc < # < pl = p

// ZipOpen zip_open()
func ZipOpen(filename string) (*zip.ReadCloser, error) { _ = "STUB: not implemented"; return nil, nil }

// Pack pack()
func Pack(order binary.ByteOrder, data interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Unpack unpack()
func Unpack(order binary.ByteOrder, data string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ternary Ternary expression
// max := Ternary(a > b, a, b).(int)
func Ternary(condition bool, trueVal, falseVal interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}
