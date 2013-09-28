// © 2013 the Ui Authors under the MIT license. See AUTHORS for the list of authors.

package ui

/*
#include "ui.h"
*/
import "C"

import (
	"strconv"
)

type Key C.Uint32

const (
	// "0"
	Key0 Key = C.SDLK_0

	// "1"
	Key1 Key = C.SDLK_1

	// "2"
	Key2 Key = C.SDLK_2

	// "3"
	Key3 Key = C.SDLK_3

	// "4"
	Key4 Key = C.SDLK_4

	// "5"
	Key5 Key = C.SDLK_5

	// "6"
	Key6 Key = C.SDLK_6

	// "7"
	Key7 Key = C.SDLK_7

	// "8"
	Key8 Key = C.SDLK_8

	// "9"
	Key9 Key = C.SDLK_9

	// "A"
	KeyA Key = C.SDLK_a

	// "AC Back" (the Back key (application control keypad))
	KeyACBack Key = C.SDLK_AC_BACK

	// "AC Bookmarks" (the Bookmarks key (application control keypad))
	KeyACBookmarks Key = C.SDLK_AC_BOOKMARKS

	// "AC Forward" (the Forward key (application control keypad))
	KeyACForward Key = C.SDLK_AC_FORWARD

	// "AC Home" (the Home key (application control keypad))
	KeyACHome Key = C.SDLK_AC_HOME

	// "AC Refresh" (the Refresh key (application control keypad))
	KeyACRefresh Key = C.SDLK_AC_REFRESH

	// "AC Search" (the Search key (application control keypad))
	KeyACSearch Key = C.SDLK_AC_SEARCH

	// "AC Stop" (the Stop key (application control keypad))
	KeyACStop Key = C.SDLK_AC_STOP

	// "Again" (the Again key (Redo))
	KeyAgain Key = C.SDLK_AGAIN

	// "AltErase" (Erase-Eaze)
	KeyAltErase Key = C.SDLK_ALTERASE

	// "Application" (the Application / Compose / Context Menu (Windows) key )
	KeyApplication Key = C.SDLK_APPLICATION

	// "AudioMute" (the Mute volume key)
	KeyAudioMute Key = C.SDLK_AUDIOMUTE

	// "AudioNext" (the Next Track media key)
	KeyAudioNext Key = C.SDLK_AUDIONEXT

	// "AudioPlay" (the Play media key)
	KeyAudioPlay Key = C.SDLK_AUDIOPLAY

	// "AudioPrev" (the Previous Track media key)
	KeyAudioPrev Key = C.SDLK_AUDIOPREV

	// "AudioStop" (the Stop media key)
	KeyAudioStop Key = C.SDLK_AUDIOSTOP

	// "B"
	KeyB Key = C.SDLK_b

	// "`"
	KeyBackQuote Key = C.SDLK_BACKQUOTE

	// "\" (Located at the lower left of the return key on ISO keyboards and
	// at the right end of the QWERTY row on ANSI keyboards. Produces REVERSE
	// SOLIDUS (backslash) and VERTICAL LINE in a US layout, REVERSE SOLIDUS
	// and VERTICAL LINE in a UK Mac layout, NUMBER SIGN and TILDE in a UK
	// Windows layout, DOLLAR SIGN and POUND SIGN in a Swiss German layout,
	// NUMBER SIGN and APOSTROPHE in a German layout, GRAVE ACCENT and POUND
	// SIGN in a French Mac layout, and ASTERISK and MICRO SIGN in a French
	// Windows layout.)
	KeyBackslash Key = C.SDLK_BACKSLASH

	// "Backspace"
	KeyBackSpace Key = C.SDLK_BACKSPACE

	// "BrightnessDown" (the Brightness Down key)
	KeyBrightnessDown Key = C.SDLK_BRIGHTNESSDOWN

	// "BrightnessUp" (the Brightness Up key)
	KeyBrightnessUp Key = C.SDLK_BRIGHTNESSUP

	// "C"
	KeyC Key = C.SDLK_c

	// "Calculator" (the Calculator key)
	KeyCalculator Key = C.SDLK_CALCULATOR

	// "Cancel"
	KeyCancel Key = C.SDLK_CANCEL

	// "CapsLock"
	KeyCapsLock Key = C.SDLK_CAPSLOCK

	// "Clear"
	KeyClear Key = C.SDLK_CLEAR

	// "Clear / Again"
	KeyClearAgain Key = C.SDLK_CLEARAGAIN

	// ","
	KeyComma Key = C.SDLK_COMMA

	// "Computer" (the My Computer key)
	KeyComputer Key = C.SDLK_COMPUTER

	// "Copy"
	KeyCopy Key = C.SDLK_COPY

	// "CrSel"
	KeyCrSel Key = C.SDLK_CRSEL

	// "CurrencySubUnit" (the Currency Subunit key)
	KeyCurrencySubunit Key = C.SDLK_CURRENCYSUBUNIT

	// "CurrencyUnit" (the Currency Unit key)
	KeyCurrencyUnit Key = C.SDLK_CURRENCYUNIT

	// "Cut"
	KeyCut Key = C.SDLK_CUT

	// "D"
	KeyD Key = C.SDLK_d

	// "DecimalSeparator" (the Decimal Separator key)
	KeyDecimalSeparator Key = C.SDLK_DECIMALSEPARATOR

	// "Delete"
	KeyDelete Key = C.SDLK_DELETE

	// "DisplaySwitch" (display mirroring/dual display switch, video mode switch)
	KeyDisplaySwitch Key = C.SDLK_DISPLAYSWITCH

	// "Down" (the Down arrow key (navigation keypad))
	KeyDown Key = C.SDLK_DOWN

	// "E"
	KeyE Key = C.SDLK_e

	// "Eject" (the Eject key)
	KeyEject Key = C.SDLK_EJECT

	// "End"
	KeyEnd Key = C.SDLK_END

	// "="
	KeyEquals Key = C.SDLK_EQUALS

	// "Escape" (the Esc key)
	KeyEscape Key = C.SDLK_ESCAPE

	// "Execute"
	KeyExecute Key = C.SDLK_EXECUTE

	// "ExSel"
	KeyExSel Key = C.SDLK_EXSEL

	// "F"
	KeyF Key = C.SDLK_f

	// "F1"
	KeyF1 Key = C.SDLK_F1

	// "F10"
	KeyF10 Key = C.SDLK_F10

	// "F11"
	KeyF11 Key = C.SDLK_F11

	// "F12"
	KeyF12 Key = C.SDLK_F12

	// "F13"
	KeyF13 Key = C.SDLK_F13

	// "F14"
	KeyF14 Key = C.SDLK_F14

	// "F15"
	KeyF15 Key = C.SDLK_F15

	// "F16"
	KeyF16 Key = C.SDLK_F16

	// "F17"
	KeyF17 Key = C.SDLK_F17

	// "F18"
	KeyF18 Key = C.SDLK_F18

	// "F19"
	KeyF19 Key = C.SDLK_F19

	// "F2"
	KeyF2 Key = C.SDLK_F2

	// "F20"
	KeyF20 Key = C.SDLK_F20

	// "F21"
	KeyF21 Key = C.SDLK_F21

	// "F22"
	KeyF22 Key = C.SDLK_F22

	// "F23"
	KeyF23 Key = C.SDLK_F23

	// "F24"
	KeyF24 Key = C.SDLK_F24

	// "F3"
	KeyF3 Key = C.SDLK_F3

	// "F4"
	KeyF4 Key = C.SDLK_F4

	// "F5"
	KeyF5 Key = C.SDLK_F5

	// "F6"
	KeyF6 Key = C.SDLK_F6

	// "F7"
	KeyF7 Key = C.SDLK_F7

	// "F8"
	KeyF8 Key = C.SDLK_F8

	// "F9"
	KeyF9 Key = C.SDLK_F9

	// "Find"
	KeyFind Key = C.SDLK_FIND

	// "G"
	KeyG Key = C.SDLK_g

	// "H"
	KeyH Key = C.SDLK_h

	// "Help"
	KeyHelp Key = C.SDLK_HELP

	// "Home"
	KeyHome Key = C.SDLK_HOME

	// "I"
	KeyI Key = C.SDLK_i

	// "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
	KeyInsert Key = C.SDLK_INSERT

	// "J"
	KeyJ Key = C.SDLK_j

	// "K"
	KeyK Key = C.SDLK_k

	// "KBDIllumDown" (the Keyboard Illumination Down key)
	KeyKBDillumDown Key = C.SDLK_KBDILLUMDOWN

	// "KBDIllumToggle" (the Keyboard Illumination Toggle key)
	KeyKBDillumToggle Key = C.SDLK_KBDILLUMTOGGLE

	// "KBDIllumUp" (the Keyboard Illumination Up key)
	KeyKBDillumUp Key = C.SDLK_KBDILLUMUP

	// "Keypad 0" (the 0 key (numeric keypad))
	KeyKP0 Key = C.SDLK_KP_0

	// "Keypad 00" (the 00 key (numeric keypad))
	KeyKP00 Key = C.SDLK_KP_00

	// "Keypad 000" (the 000 key (numeric keypad))
	KeyKP000 Key = C.SDLK_KP_000

	// "Keypad 1" (the 1 key (numeric keypad))
	KeyKP1 Key = C.SDLK_KP_1

	// "Keypad 2" (the 2 key (numeric keypad))
	KeyKP2 Key = C.SDLK_KP_2

	// "Keypad 3" (the 3 key (numeric keypad))
	KeyKP3 Key = C.SDLK_KP_3

	// "Keypad 4" (the 4 key (numeric keypad))
	KeyKP4 Key = C.SDLK_KP_4

	// "Keypad 5" (the 5 key (numeric keypad))
	KeyKP5 Key = C.SDLK_KP_5

	// "Keypad 6" (the 6 key (numeric keypad))
	KeyKP6 Key = C.SDLK_KP_6

	// "Keypad 7" (the 7 key (numeric keypad))
	KeyKP7 Key = C.SDLK_KP_7

	// "Keypad 8" (the 8 key (numeric keypad))
	KeyKP8 Key = C.SDLK_KP_8

	// "Keypad 9" (the 9 key (numeric keypad))
	KeyKP9 Key = C.SDLK_KP_9

	// "Keypad A" (the A key (numeric keypad))
	KeyKPA Key = C.SDLK_KP_A

	// "Keypad &" (the & key (numeric keypad))
	KeyKPAmpersand Key = C.SDLK_KP_AMPERSAND

	// "Keypad @" (the @ key (numeric keypad))
	KeyKPAt Key = C.SDLK_KP_AT

	// "Keypad B" (the B key (numeric keypad))
	KeyKPB Key = C.SDLK_KP_B

	// "Keypad Backspace" (the Backspace key (numeric keypad))
	KeyKPBackspace Key = C.SDLK_KP_BACKSPACE

	// "Keypad Binary" (the Binary key (numeric keypad))
	KeyKPBinary Key = C.SDLK_KP_BINARY

	// "Keypad C" (the C key (numeric keypad))
	KeyKPC Key = C.SDLK_KP_C

	// "Keypad Clear" (the Clear key (numeric keypad))
	KeyKPClear Key = C.SDLK_KP_CLEAR

	// "Keypad ClearEntry" (the Clear Entry key (numeric keypad))
	KeyKPClearEntry Key = C.SDLK_KP_CLEARENTRY

	// "Keypad :" (the : key (numeric keypad))
	KeyKPColon Key = C.SDLK_KP_COLON

	// "Keypad ," (the Comma key (numeric keypad))
	KeyKPComma Key = C.SDLK_KP_COMMA

	// "Keypad D" (the D key (numeric keypad))
	KeyKPD Key = C.SDLK_KP_D

	// "Keypad &&" (the && key (numeric keypad))
	KeyKPDblampersand Key = C.SDLK_KP_DBLAMPERSAND

	// "Keypad ||" (the || key (numeric keypad))
	KeyKPDblVerticalBar Key = C.SDLK_KP_DBLVERTICALBAR

	// "Keypad Decimal" (the Decimal key (numeric keypad))
	KeyKPDecimal Key = C.SDLK_KP_DECIMAL

	// "Keypad /" (the / key (numeric keypad))
	KeyKPDivide Key = C.SDLK_KP_DIVIDE

	// "Keypad E" (the E key (numeric keypad))
	KeyKPE Key = C.SDLK_KP_E

	// "Keypad Enter" (the Enter key (numeric keypad))
	KeyKPEnter Key = C.SDLK_KP_ENTER

	// "Keypad =" (the = key (numeric keypad))
	KeyKPEquals Key = C.SDLK_KP_EQUALS

	// "Keypad = (AS400)" (the Equals AS400 key (numeric keypad))
	KeyKPEqualsAS400 Key = C.SDLK_KP_EQUALSAS400

	// "Keypad !" (the ! key (numeric keypad))
	KeyKPExclam Key = C.SDLK_KP_EXCLAM

	// "Keypad F" (the F key (numeric keypad))
	KeyKPF Key = C.SDLK_KP_F

	// "Keypad >" (the Greater key (numeric keypad))
	KeyKPGreater Key = C.SDLK_KP_GREATER

	// "Keypad #" (the # key (numeric keypad))
	KeyKPHash Key = C.SDLK_KP_HASH

	// "Keypad Hexadecimal" (the Hexadecimal key (numeric keypad))
	KeyKPHexadecimal Key = C.SDLK_KP_HEXADECIMAL

	// "Keypad {" (the Left Brace key (numeric keypad))
	KeyKPLeftBrace Key = C.SDLK_KP_LEFTBRACE

	// "Keypad (" (the Left Parenthesis key (numeric keypad))
	KeyKPLeftParen Key = C.SDLK_KP_LEFTPAREN

	// "Keypad <" (the Less key (numeric keypad))
	KeyKPLess Key = C.SDLK_KP_LESS

	// "Keypad MemAdd" (the Mem Add key (numeric keypad))
	KeyKPMemAdd Key = C.SDLK_KP_MEMADD

	// "Keypad MemClear" (the Mem Clear key (numeric keypad))
	KeyKPMemClear Key = C.SDLK_KP_MEMCLEAR

	// "Keypad MemDivide" (the Mem Divide key (numeric keypad))
	KeyKPMemDivide Key = C.SDLK_KP_MEMDIVIDE

	// "Keypad MemMultiply" (the Mem Multiply key (numeric keypad))
	KeyKPMemMultiply Key = C.SDLK_KP_MEMMULTIPLY

	// "Keypad MemRecall" (the Mem Recall key (numeric keypad))
	KeyKPMemRecall Key = C.SDLK_KP_MEMRECALL

	// "Keypad MemStore" (the Mem Store key (numeric keypad))
	KeyKPMemStore Key = C.SDLK_KP_MEMSTORE

	// "Keypad MemSubtract" (the Mem Subtract key (numeric keypad))
	KeyKPMemSubtract Key = C.SDLK_KP_MEMSUBTRACT

	// "Keypad -" (the - key (numeric keypad))
	KeyKPMinus Key = C.SDLK_KP_MINUS

	// "Keypad *" (the * key (numeric keypad))
	KeyKPMultiply Key = C.SDLK_KP_MULTIPLY

	// "Keypad Octal" (the Octal key (numeric keypad))
	KeyKPOctal Key = C.SDLK_KP_OCTAL

	// "Keypad %" (the Percent key (numeric keypad))
	KeyKPPercent Key = C.SDLK_KP_PERCENT

	// "Keypad ." (the . key (numeric keypad))
	KeyKPPeriod Key = C.SDLK_KP_PERIOD

	// "Keypad +" (the + key (numeric keypad))
	KeyKPPlus Key = C.SDLK_KP_PLUS

	// "Keypad +/-" (the +/- key (numeric keypad))
	KeyKPPlusMinus Key = C.SDLK_KP_PLUSMINUS

	// "Keypad ^" (the Power key (numeric keypad))
	KeyKPPower Key = C.SDLK_KP_POWER

	// "Keypad }" (the Right Brace key (numeric keypad))
	KeyKPRightBrace Key = C.SDLK_KP_RIGHTBRACE

	// "Keypad )" (the Right Parenthesis key (numeric keypad))
	KeyKPRightParen Key = C.SDLK_KP_RIGHTPAREN

	// "Keypad Space" (the Space key (numeric keypad))
	KeyKPSpace Key = C.SDLK_KP_SPACE

	// "Keypad Tab" (the Tab key (numeric keypad))
	KeyKPTab Key = C.SDLK_KP_TAB

	// "Keypad |" (the | key (numeric keypad))
	KeyKPVerticalBar Key = C.SDLK_KP_VERTICALBAR

	// "Keypad XOR" (the XOR key (numeric keypad))
	KeyKPXOR Key = C.SDLK_KP_XOR

	// "L"
	KeyL Key = C.SDLK_l

	// "Left Alt" (alt, option)
	KeyLAlt Key = C.SDLK_LALT

	// "Left Ctrl"
	KeyLCtrl Key = C.SDLK_LCTRL

	// "Left" (the Left arrow key (navigation keypad))
	KeyLeft Key = C.SDLK_LEFT

	// "["
	KeyLeftBracket Key = C.SDLK_LEFTBRACKET

	// "Left GUI" (windows, command (apple), meta)
	KeyLGUI Key = C.SDLK_LGUI

	// "Left Shift"
	KeyLShift Key = C.SDLK_LSHIFT

	// "M"
	KeyM Key = C.SDLK_m

	// "Mail" (the Mail/eMail key)
	KeyMail Key = C.SDLK_MAIL

	// "MediaSelect" (the Media Select key)
	KeyMediaSelect Key = C.SDLK_MEDIASELECT

	// "Menu"
	KeyMenu Key = C.SDLK_MENU

	// "-"
	KeyMinus Key = C.SDLK_MINUS

	// "ModeSwitch" (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here)
	KeyMode Key = C.SDLK_MODE

	// "Mute"
	KeyMute Key = C.SDLK_MUTE

	// "N"
	KeyN Key = C.SDLK_n

	// "Numlock" (the Num Lock key (PC) / the Clear key (Mac))
	KeyNumLockClear Key = C.SDLK_NUMLOCKCLEAR

	// "O"
	KeyO Key = C.SDLK_o

	// "Oper"
	KeyOper Key = C.SDLK_OPER

	// "Out"
	KeyOut Key = C.SDLK_OUT

	// "P"
	KeyP Key = C.SDLK_p

	// "PageDown"
	KeyPageDown Key = C.SDLK_PAGEDOWN

	// "PageUp"
	KeyPageUp Key = C.SDLK_PAGEUP

	// "Paste"
	KeyPaste Key = C.SDLK_PASTE

	// "Pause" (the Pause / Break key)
	KeyPause Key = C.SDLK_PAUSE

	// "."
	KeyPeriod Key = C.SDLK_PERIOD

	// "Power" (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.)
	KeyPower Key = C.SDLK_POWER

	// "PrintScreen"
	KeyPrintScreen Key = C.SDLK_PRINTSCREEN

	// "Prior"
	KeyPrior Key = C.SDLK_PRIOR

	// "Q"
	KeyQ Key = C.SDLK_q

	// "'"
	KeyQuote Key = C.SDLK_QUOTE

	// "R"
	KeyR Key = C.SDLK_r

	// "Right Alt" (alt gr, option)
	KeyRAlt Key = C.SDLK_RALT

	// "Right Ctrl"
	KeyRCtrl Key = C.SDLK_RCTRL

	// "Return" (the Enter key (main keyboard))
	KeyReturn Key = C.SDLK_RETURN

	// "Return"
	KeyReturn2 Key = C.SDLK_RETURN2

	// "Right GUI" (windows, command (apple), meta)
	KeyRGUI Key = C.SDLK_RGUI

	// "Right" (the Right arrow key (navigation keypad))
	KeyRight Key = C.SDLK_RIGHT

	// "]"
	KeyRightBracket Key = C.SDLK_RIGHTBRACKET

	// "Right Shift"
	KeyRShift Key = C.SDLK_RSHIFT

	// "S"
	KeyS Key = C.SDLK_s

	// "ScrollLock"
	KeyScrollLock Key = C.SDLK_SCROLLLOCK

	// "Select"
	KeySelect Key = C.SDLK_SELECT

	// ";"
	KeySemicolon Key = C.SDLK_SEMICOLON

	// "Separator"
	KeySeparator Key = C.SDLK_SEPARATOR

	// "/"
	KeySlash Key = C.SDLK_SLASH

	// "Sleep" (the Sleep key)
	KeySleep Key = C.SDLK_SLEEP

	// "Space" (the Space Bar key(s))
	KeySpace Key = C.SDLK_SPACE

	// "Stop"
	KeyStop Key = C.SDLK_STOP

	// "SysReq" (the SysReq key)
	KeySysReq Key = C.SDLK_SYSREQ

	// "T"
	KeyT Key = C.SDLK_t

	// "Tab" (the Tab key)
	KeyTab Key = C.SDLK_TAB

	// "ThousandsSeparator" (the Thousands Separator key)
	KeyThousandsSeparator Key = C.SDLK_THOUSANDSSEPARATOR

	// "U"
	KeyU Key = C.SDLK_u

	// "Undo"
	KeyUndo Key = C.SDLK_UNDO

	// "" (no name, empty string)
	KeyUnknown Key = C.SDLK_UNKNOWN

	// "Up" (the Up arrow key (navigation keypad))
	KeyUp Key = C.SDLK_UP

	// "V"
	KeyV Key = C.SDLK_v

	// "VolumeDown"
	KeyVolumeDown Key = C.SDLK_VOLUMEDOWN

	// "VolumeUp"
	KeyVolumeUp Key = C.SDLK_VOLUMEUP

	// "W"
	KeyW Key = C.SDLK_w

	// "WWW" (the WWW/World Wide Web key)
	KeyWWW Key = C.SDLK_WWW

	// "X"
	KeyX Key = C.SDLK_x

	// "Y"
	KeyY Key = C.SDLK_y

	// "Z"
	KeyZ Key = C.SDLK_z
)

var keyNames = map[Key]string{
	Key0:                  "Key0",
	Key1:                  "Key1",
	Key2:                  "Key2",
	Key3:                  "Key3",
	Key4:                  "Key4",
	Key5:                  "Key5",
	Key6:                  "Key6",
	Key7:                  "Key7",
	Key8:                  "Key8",
	Key9:                  "Key9",
	KeyA:                  "KeyA",
	KeyACBack:             "KeyACBack",
	KeyACBookmarks:        "KeyACBookmarks",
	KeyACForward:          "KeyACForward",
	KeyACHome:             "KeyACHome",
	KeyACRefresh:          "KeyACRefresh",
	KeyACSearch:           "KeyACSearch",
	KeyACStop:             "KeyACStop",
	KeyAgain:              "KeyAgain",
	KeyAltErase:           "KeyAltErase",
	KeyApplication:        "KeyApplication",
	KeyAudioMute:          "KeyAudioMute",
	KeyAudioNext:          "KeyAudioNext",
	KeyAudioPlay:          "KeyAudioPlay",
	KeyAudioPrev:          "KeyAudioPrev",
	KeyAudioStop:          "KeyAudioStop",
	KeyB:                  "KeyB",
	KeyBackQuote:          "KeyBackQuote",
	KeyBackslash:          "KeyBackslash",
	KeyBackSpace:          "KeyBackSpace",
	KeyBrightnessDown:     "KeyBrightnessDown",
	KeyBrightnessUp:       "KeyBrightnessUp",
	KeyC:                  "KeyC",
	KeyCalculator:         "KeyCalculator",
	KeyCancel:             "KeyCancel",
	KeyCapsLock:           "KeyCapsLock",
	KeyClear:              "KeyClear",
	KeyClearAgain:         "KeyClearAgain",
	KeyComma:              "KeyComma",
	KeyComputer:           "KeyComputer",
	KeyCopy:               "KeyCopy",
	KeyCrSel:              "KeyCrSel",
	KeyCurrencySubunit:    "KeyCurrencySubunit",
	KeyCurrencyUnit:       "KeyCurrencyUnit",
	KeyCut:                "KeyCut",
	KeyD:                  "KeyD",
	KeyDecimalSeparator:   "KeyDecimalSeparator",
	KeyDelete:             "KeyDelete",
	KeyDisplaySwitch:      "KeyDisplaySwitch",
	KeyDown:               "KeyDown",
	KeyE:                  "KeyE",
	KeyEject:              "KeyEject",
	KeyEnd:                "KeyEnd",
	KeyEquals:             "KeyEquals",
	KeyEscape:             "KeyEscape",
	KeyExecute:            "KeyExecute",
	KeyExSel:              "KeyExSel",
	KeyF:                  "KeyF",
	KeyF1:                 "KeyF1",
	KeyF10:                "KeyF10",
	KeyF11:                "KeyF11",
	KeyF12:                "KeyF12",
	KeyF13:                "KeyF13",
	KeyF14:                "KeyF14",
	KeyF15:                "KeyF15",
	KeyF16:                "KeyF16",
	KeyF17:                "KeyF17",
	KeyF18:                "KeyF18",
	KeyF19:                "KeyF19",
	KeyF2:                 "KeyF2",
	KeyF20:                "KeyF20",
	KeyF21:                "KeyF21",
	KeyF22:                "KeyF22",
	KeyF23:                "KeyF23",
	KeyF24:                "KeyF24",
	KeyF3:                 "KeyF3",
	KeyF4:                 "KeyF4",
	KeyF5:                 "KeyF5",
	KeyF6:                 "KeyF6",
	KeyF7:                 "KeyF7",
	KeyF8:                 "KeyF8",
	KeyF9:                 "KeyF9",
	KeyFind:               "KeyFind",
	KeyG:                  "KeyG",
	KeyH:                  "KeyH",
	KeyHelp:               "KeyHelp",
	KeyHome:               "KeyHome",
	KeyI:                  "KeyI",
	KeyInsert:             "KeyInsert",
	KeyJ:                  "KeyJ",
	KeyK:                  "KeyK",
	KeyKBDillumDown:       "KeyKBDillumDown",
	KeyKBDillumToggle:     "KeyKBDillumToggle",
	KeyKBDillumUp:         "KeyKBDillumUp",
	KeyKP0:                "KeyKP0",
	KeyKP00:               "KeyKP00",
	KeyKP000:              "KeyKP000",
	KeyKP1:                "KeyKP1",
	KeyKP2:                "KeyKP2",
	KeyKP3:                "KeyKP3",
	KeyKP4:                "KeyKP4",
	KeyKP5:                "KeyKP5",
	KeyKP6:                "KeyKP6",
	KeyKP7:                "KeyKP7",
	KeyKP8:                "KeyKP8",
	KeyKP9:                "KeyKP9",
	KeyKPA:                "KeyKPA",
	KeyKPAmpersand:        "KeyKPAmpersand",
	KeyKPAt:               "KeyKPAt",
	KeyKPB:                "KeyKPB",
	KeyKPBackspace:        "KeyKPBackspace",
	KeyKPBinary:           "KeyKPBinary",
	KeyKPC:                "KeyKPC",
	KeyKPClear:            "KeyKPClear",
	KeyKPClearEntry:       "KeyKPClearEntry",
	KeyKPColon:            "KeyKPColon",
	KeyKPComma:            "KeyKPComma",
	KeyKPD:                "KeyKPD",
	KeyKPDblampersand:     "KeyKPDblampersand",
	KeyKPDblVerticalBar:   "KeyKPDblVerticalBar",
	KeyKPDecimal:          "KeyKPDecimal",
	KeyKPDivide:           "KeyKPDivide",
	KeyKPE:                "KeyKPE",
	KeyKPEnter:            "KeyKPEnter",
	KeyKPEquals:           "KeyKPEquals",
	KeyKPEqualsAS400:      "KeyKPEqualsAS400",
	KeyKPExclam:           "KeyKPExclam",
	KeyKPF:                "KeyKPF",
	KeyKPGreater:          "KeyKPGreater",
	KeyKPHash:             "KeyKPHash",
	KeyKPHexadecimal:      "KeyKPHexadecimal",
	KeyKPLeftBrace:        "KeyKPLeftBrace",
	KeyKPLeftParen:        "KeyKPLeftParen",
	KeyKPLess:             "KeyKPLess",
	KeyKPMemAdd:           "KeyKPMemAdd",
	KeyKPMemClear:         "KeyKPMemClear",
	KeyKPMemDivide:        "KeyKPMemDivide",
	KeyKPMemMultiply:      "KeyKPMemMultiply",
	KeyKPMemRecall:        "KeyKPMemRecall",
	KeyKPMemStore:         "KeyKPMemStore",
	KeyKPMemSubtract:      "KeyKPMemSubtract",
	KeyKPMinus:            "KeyKPMinus",
	KeyKPMultiply:         "KeyKPMultiply",
	KeyKPOctal:            "KeyKPOctal",
	KeyKPPercent:          "KeyKPPercent",
	KeyKPPeriod:           "KeyKPPeriod",
	KeyKPPlus:             "KeyKPPlus",
	KeyKPPlusMinus:        "KeyKPPlusMinus",
	KeyKPPower:            "KeyKPPower",
	KeyKPRightBrace:       "KeyKPRightBrace",
	KeyKPRightParen:       "KeyKPRightParen",
	KeyKPSpace:            "KeyKPSpace",
	KeyKPTab:              "KeyKPTab",
	KeyKPVerticalBar:      "KeyKPVerticalBar",
	KeyKPXOR:              "KeyKPXOR",
	KeyL:                  "KeyL",
	KeyLAlt:               "KeyLAlt",
	KeyLCtrl:              "KeyLCtrl",
	KeyLeft:               "KeyLeft",
	KeyLeftBracket:        "KeyLeftBracket",
	KeyLGUI:               "KeyLGUI",
	KeyLShift:             "KeyLShift",
	KeyM:                  "KeyM",
	KeyMail:               "KeyMail",
	KeyMediaSelect:        "KeyMediaSelect",
	KeyMenu:               "KeyMenu",
	KeyMinus:              "KeyMinus",
	KeyMode:               "KeyMode",
	KeyMute:               "KeyMute",
	KeyN:                  "KeyN",
	KeyNumLockClear:       "KeyNumLockClear",
	KeyO:                  "KeyO",
	KeyOper:               "KeyOper",
	KeyOut:                "KeyOut",
	KeyP:                  "KeyP",
	KeyPageDown:           "KeyPageDown",
	KeyPageUp:             "KeyPageUp",
	KeyPaste:              "KeyPaste",
	KeyPause:              "KeyPause",
	KeyPeriod:             "KeyPeriod",
	KeyPower:              "KeyPower",
	KeyPrintScreen:        "KeyPrintScreen",
	KeyPrior:              "KeyPrior",
	KeyQ:                  "KeyQ",
	KeyQuote:              "KeyQuote",
	KeyR:                  "KeyR",
	KeyRAlt:               "KeyRAlt",
	KeyRCtrl:              "KeyRCtrl",
	KeyReturn:             "KeyReturn",
	KeyReturn2:            "KeyReturn2",
	KeyRGUI:               "KeyRGUI",
	KeyRight:              "KeyRight",
	KeyRightBracket:       "KeyRightBracket",
	KeyRShift:             "KeyRShift",
	KeyS:                  "KeyS",
	KeyScrollLock:         "KeyScrollLock",
	KeySelect:             "KeySelect",
	KeySemicolon:          "KeySemicolon",
	KeySeparator:          "KeySeparator",
	KeySlash:              "KeySlash",
	KeySleep:              "KeySleep",
	KeySpace:              "KeySpace",
	KeyStop:               "KeyStop",
	KeySysReq:             "KeySysReq",
	KeyT:                  "KeyT",
	KeyTab:                "KeyTab",
	KeyThousandsSeparator: "KeyThousandsSeparator",
	KeyU:          "KeyU",
	KeyUndo:       "KeyUndo",
	KeyUnknown:    "KeyUnknown",
	KeyUp:         "KeyUp",
	KeyV:          "KeyV",
	KeyVolumeDown: "KeyVolumeDown",
	KeyVolumeUp:   "KeyVolumeUp",
	KeyW:          "KeyW",
	KeyWWW:        "KeyWWW",
	KeyX:          "KeyX",
	KeyY:          "KeyY",
	KeyZ:          "KeyZ",
}

func (k Key) String() string {
	if n, ok := keyNames[k]; ok {
		return n
	}
	return "Unknown(" + strconv.Itoa(int(k)) + ")"
}

type KeyMod C.SDL_Keymod

const (
	// ModNone means no modifier is applicable.
	ModNone = C.KMOD_NONE

	// ModLShift means that the left Shift key is down.
	ModLShift KeyMod = C.KMOD_LSHIFT

	// ModRShift means that the right Shift key is down.
	ModRShift KeyMod = C.KMOD_RSHIFT

	// ModLCtrl means that the left Ctrl (Control) key is down.
	ModLCtrl KeyMod = C.KMOD_LCTRL

	// ModRCtrl means that the right Ctrl (Control) key is down.
	ModRCtrl KeyMod = C.KMOD_RCTRL

	// ModLAlt means that the left Alt key is down.
	ModLAlt KeyMod = C.KMOD_LALT

	// ModRAlt means that the right Alt key is down.
	ModRAlt KeyMod = C.KMOD_RALT

	// ModLGUI means that the left GUI key (often the Windows key) is down.
	ModLGUI KeyMod = C.KMOD_LGUI

	// ModRGUI means that the right GUI key (often the Windows key) is down.
	ModRGUI KeyMod = C.KMOD_RGUI

	// ModNum means that the Num Lock key (may be located on an extended keypad) is down.
	ModNum KeyMod = C.KMOD_NUM

	// ModCaps means that the Caps Lock key is down.
	ModCaps KeyMod = C.KMOD_CAPS

	// ModMode means that the AltGr key is down.
	ModMode KeyMod = C.KMOD_MODE

	// ModCtrl is an alias for (KMOD_LCTRL|KMOD_RCTRL).
	ModCtrl KeyMod = C.KMOD_CTRL

	// ModShift is an alias for (KMOD_LSHIFT|KMOD_RSHIFT).
	ModShift KeyMod = C.KMOD_SHIFT

	// ModAlt is an alias for (KMOD_LALT|KMOD_RALT).
	ModAlt KeyMod = C.KMOD_ALT

	// ModGUI is an alias for (KMOD_LGUI|KMOD_RGUI).
	ModGUI KeyMod = C.KMOD_GUI
)

var keyModNames = map[KeyMod]string{
	ModNone:   "ModNone",
	ModLShift: "ModLShift",
	ModRShift: "ModRShift",
	ModLCtrl:  "ModLCtrl",
	ModRCtrl:  "ModRCtrl",
	ModLAlt:   "ModLAlt",
	ModRAlt:   "ModRAlt",
	ModLGUI:   "ModLGUI",
	ModRGUI:   "ModRGUI",
	ModNum:    "ModNum",
	ModCaps:   "ModCaps",
	ModMode:   "ModMode",
	ModCtrl:   "ModCtrl",
	ModShift:  "ModShift",
	ModAlt:    "ModAlt",
	ModGUI:    "ModGUI",
}

func (mod KeyMod) String() string {
	s := ""
	for m, n := range keyModNames {
		if mod&m == 0 {
			continue
		}
		if s != "" {
			s += " | "
		}
		s += n
	}
	if s == "" {
		return keyModNames[ModNone]
	}
	return s
}
