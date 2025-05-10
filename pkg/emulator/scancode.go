package emulator

const (
	SDL_SCANCODE_UNKNOWN = 0

	/**
	 *  \name Usage page 0x07
	 *
	 *  These values are from usage page 0x07 (USB keyboard page).
	 */
	/* @{ */

	SDL_SCANCODE_APPLICATION = 101 /**< windows contextual menu compose */
	SDL_SCANCODE_POWER       = 102 /**< The USB document says this is a status flag
	 *   not a physical key - but some Mac keyboards
	 *   do have a power key. */

	/* not sure whether there's a reason to enable these */
	/*     SDL_SCANCODE_LOCKINGCAPSLOCK = 130  */
	/*     SDL_SCANCODE_LOCKINGNUMLOCK = 131 */
	/*     SDL_SCANCODE_LOCKINGSCROLLLOCK = 132 */
	SDL_SCANCODE_KP_COMMA       = 133
	SDL_SCANCODE_KP_EQUALSAS400 = 134

	SDL_SCANCODE_INTERNATIONAL1 = 135 /**< used on Asian keyboards see
	  footnotes in USB doc */
	SDL_SCANCODE_INTERNATIONAL2 = 136
	SDL_SCANCODE_INTERNATIONAL3 = 137 /**< Yen */
	SDL_SCANCODE_INTERNATIONAL4 = 138
	SDL_SCANCODE_INTERNATIONAL5 = 139
	SDL_SCANCODE_INTERNATIONAL6 = 140
	SDL_SCANCODE_INTERNATIONAL7 = 141
	SDL_SCANCODE_INTERNATIONAL8 = 142
	SDL_SCANCODE_INTERNATIONAL9 = 143
	SDL_SCANCODE_LANG1          = 144 /**< Hangul/English toggle */
	SDL_SCANCODE_LANG2          = 145 /**< Hanja conversion */
	SDL_SCANCODE_LANG3          = 146 /**< Katakana */
	SDL_SCANCODE_LANG4          = 147 /**< Hiragana */
	SDL_SCANCODE_LANG5          = 148 /**< Zenkaku/Hankaku */
	SDL_SCANCODE_LANG6          = 149 /**< reserved */
	SDL_SCANCODE_LANG7          = 150 /**< reserved */
	SDL_SCANCODE_LANG8          = 151 /**< reserved */
	SDL_SCANCODE_LANG9          = 152 /**< reserved */

	SDL_SCANCODE_ALTERASE   = 153 /**< Erase-Eaze */
	SDL_SCANCODE_SYSREQ     = 154
	SDL_SCANCODE_CANCEL     = 155 /**< AC Cancel */
	SDL_SCANCODE_CLEAR      = 156
	SDL_SCANCODE_PRIOR      = 157
	SDL_SCANCODE_RETURN2    = 158
	SDL_SCANCODE_SEPARATOR  = 159
	SDL_SCANCODE_OUT        = 160
	SDL_SCANCODE_OPER       = 161
	SDL_SCANCODE_CLEARAGAIN = 162
	SDL_SCANCODE_CRSEL      = 163
	SDL_SCANCODE_EXSEL      = 164

	SDL_SCANCODE_KP_00              = 176
	SDL_SCANCODE_KP_000             = 177
	SDL_SCANCODE_THOUSANDSSEPARATOR = 178
	SDL_SCANCODE_DECIMALSEPARATOR   = 179
	SDL_SCANCODE_CURRENCYUNIT       = 180
	SDL_SCANCODE_CURRENCYSUBUNIT    = 181
	SDL_SCANCODE_KP_LEFTPAREN       = 182
	SDL_SCANCODE_KP_RIGHTPAREN      = 183
	SDL_SCANCODE_KP_LEFTBRACE       = 184
	SDL_SCANCODE_KP_RIGHTBRACE      = 185
	SDL_SCANCODE_KP_TAB             = 186
	SDL_SCANCODE_KP_BACKSPACE       = 187
	SDL_SCANCODE_KP_A               = 188
	SDL_SCANCODE_KP_B               = 189
	SDL_SCANCODE_KP_C               = 190
	SDL_SCANCODE_KP_D               = 191
	SDL_SCANCODE_KP_E               = 192
	SDL_SCANCODE_KP_F               = 193
	SDL_SCANCODE_KP_XOR             = 194
	SDL_SCANCODE_KP_POWER           = 195
	SDL_SCANCODE_KP_PERCENT         = 196
	SDL_SCANCODE_KP_LESS            = 197
	SDL_SCANCODE_KP_GREATER         = 198
	SDL_SCANCODE_KP_AMPERSAND       = 199
	SDL_SCANCODE_KP_DBLAMPERSAND    = 200
	SDL_SCANCODE_KP_VERTICALBAR     = 201
	SDL_SCANCODE_KP_DBLVERTICALBAR  = 202
	SDL_SCANCODE_KP_COLON           = 203
	SDL_SCANCODE_KP_HASH            = 204
	SDL_SCANCODE_KP_SPACE           = 205
	SDL_SCANCODE_KP_AT              = 206
	SDL_SCANCODE_KP_EXCLAM          = 207
	SDL_SCANCODE_KP_MEMSTORE        = 208
	SDL_SCANCODE_KP_MEMRECALL       = 209
	SDL_SCANCODE_KP_MEMCLEAR        = 210
	SDL_SCANCODE_KP_MEMADD          = 211
	SDL_SCANCODE_KP_MEMSUBTRACT     = 212
	SDL_SCANCODE_KP_MEMMULTIPLY     = 213
	SDL_SCANCODE_KP_MEMDIVIDE       = 214
	SDL_SCANCODE_KP_PLUSMINUS       = 215
	SDL_SCANCODE_KP_CLEAR           = 216
	SDL_SCANCODE_KP_CLEARENTRY      = 217
	SDL_SCANCODE_KP_BINARY          = 218
	SDL_SCANCODE_KP_OCTAL           = 219
	SDL_SCANCODE_KP_DECIMAL         = 220
	SDL_SCANCODE_KP_HEXADECIMAL     = 221

	SDL_SCANCODE_MODE = 257 /**< I'm not sure if this is really not covered
	 *   by any of the above but since there's a
	 *   special KMOD_MODE for it I'm adding it here
	 */

	/* @} */ /* Usage page 0x07 */

	/**
	 *  \name Usage page 0x0C
	 *
	 *  These values are mapped from usage page 0x0C (USB consumer page).
	 *  See https://usb.org/sites/default/files/hut1_2.pdf
	 *
	 *  There are way more keys in the spec than we can represent in the
	 *  current scancode range so pick the ones that commonly come up in
	 *  real world usage.
	 */
	/* @{ */

	SDL_SCANCODE_AUDIONEXT    = 258
	SDL_SCANCODE_AUDIOPREV    = 259
	SDL_SCANCODE_AUDIOSTOP    = 260
	SDL_SCANCODE_AUDIOPLAY    = 261
	SDL_SCANCODE_AUDIOMUTE    = 262
	SDL_SCANCODE_MEDIASELECT  = 263
	SDL_SCANCODE_WWW          = 264 /**< AL Internet Browser */
	SDL_SCANCODE_MAIL         = 265
	SDL_SCANCODE_CALCULATOR   = 266 /**< AL Calculator */
	SDL_SCANCODE_COMPUTER     = 267
	SDL_SCANCODE_AC_SEARCH    = 268 /**< AC Search */
	SDL_SCANCODE_AC_HOME      = 269 /**< AC Home */
	SDL_SCANCODE_AC_BACK      = 270 /**< AC Back */
	SDL_SCANCODE_AC_FORWARD   = 271 /**< AC Forward */
	SDL_SCANCODE_AC_STOP      = 272 /**< AC Stop */
	SDL_SCANCODE_AC_REFRESH   = 273 /**< AC Refresh */
	SDL_SCANCODE_AC_BOOKMARKS = 274 /**< AC Bookmarks */

	/* @} */ /* Usage page 0x0C */

	/**
	 *  \name Walther keys
	 *
	 *  These are values that Christian Walther added (for mac keyboard?).
	 */
	/* @{ */

	SDL_SCANCODE_BRIGHTNESSDOWN = 275
	SDL_SCANCODE_BRIGHTNESSUP   = 276
	SDL_SCANCODE_DISPLAYSWITCH  = 277 /**< display mirroring/dual display
	  switch video mode switch */
	SDL_SCANCODE_KBDILLUMTOGGLE = 278
	SDL_SCANCODE_KBDILLUMDOWN   = 279
	SDL_SCANCODE_KBDILLUMUP     = 280
	SDL_SCANCODE_EJECT          = 281
	SDL_SCANCODE_SLEEP          = 282 /**< SC System Sleep */

	SDL_SCANCODE_APP1 = 283
	SDL_SCANCODE_APP2 = 284

	/* @} */ /* Walther keys */

	/**
	 *  \name Usage page 0x0C (additional media keys)
	 *
	 *  These values are mapped from usage page 0x0C (USB consumer page).
	 */
	/* @{ */

	SDL_SCANCODE_AUDIOREWIND      = 285
	SDL_SCANCODE_AUDIOFASTFORWARD = 286

	/* @} */ /* Usage page 0x0C (additional media keys) */

	/**
	 *  \name Mobile keys
	 *
	 *  These are values that are often used on mobile phones.
	 */
	/* @{ */

	SDL_SCANCODE_SOFTLEFT = 287 /**< Usually situated below the display on phones and
	  used as a multi-function feature key for selecting
	  a software defined function shown on the bottom left
	  of the display. */
	SDL_SCANCODE_SOFTRIGHT = 288 /**< Usually situated below the display on phones and
	  used as a multi-function feature key for selecting
	  a software defined function shown on the bottom right
	  of the display. */
	SDL_SCANCODE_CALL    = 289 /**< Used for accepting phone calls. */
	SDL_SCANCODE_ENDCALL = 290 /**< Used for rejecting phone calls. */

	/* @} */ /* Mobile keys */

	/* Add any other keys here. */

	SDL_NUM_SCANCODES = 512 /**< not a key just marks the number of scancodes
	  for array bounds */
)
