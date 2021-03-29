package rhdiff

// https://librsync.github.io/command_8h.html

const (
	KIND_END OpKind = iota + 1000
	KIND_LITERAL
	KIND_SIGNATURE
	KIND_COPY
	KIND_CHECKSUM
	KIND_RESERVED
)

type Op uint8
type OpKind uint16

type Command struct {
	Kind      OpKind
	Immediate uint8
	Len1      uint8
	Len2      uint8
}

const (
	OP_END Op = iota
	OP_LITERAL_1
	OP_LITERAL_2
	OP_LITERAL_3
	OP_LITERAL_4
	OP_LITERAL_5
	OP_LITERAL_6
	OP_LITERAL_7
	OP_LITERAL_8
	OP_LITERAL_9
	OP_LITERAL_10
	OP_LITERAL_11
	OP_LITERAL_12
	OP_LITERAL_13
	OP_LITERAL_14
	OP_LITERAL_15
	OP_LITERAL_16
	OP_LITERAL_17
	OP_LITERAL_18
	OP_LITERAL_19
	OP_LITERAL_20
	OP_LITERAL_21
	OP_LITERAL_22
	OP_LITERAL_23
	OP_LITERAL_24
	OP_LITERAL_25
	OP_LITERAL_26
	OP_LITERAL_27
	OP_LITERAL_28
	OP_LITERAL_29
	OP_LITERAL_30
	OP_LITERAL_31
	OP_LITERAL_32
	OP_LITERAL_33
	OP_LITERAL_34
	OP_LITERAL_35
	OP_LITERAL_36
	OP_LITERAL_37
	OP_LITERAL_38
	OP_LITERAL_39
	OP_LITERAL_40
	OP_LITERAL_41
	OP_LITERAL_42
	OP_LITERAL_43
	OP_LITERAL_44
	OP_LITERAL_45
	OP_LITERAL_46
	OP_LITERAL_47
	OP_LITERAL_48
	OP_LITERAL_49
	OP_LITERAL_50
	OP_LITERAL_51
	OP_LITERAL_52
	OP_LITERAL_53
	OP_LITERAL_54
	OP_LITERAL_55
	OP_LITERAL_56
	OP_LITERAL_57
	OP_LITERAL_58
	OP_LITERAL_59
	OP_LITERAL_60
	OP_LITERAL_61
	OP_LITERAL_62
	OP_LITERAL_63
	OP_LITERAL_64
	OP_LITERAL_N1
	OP_LITERAL_N2
	OP_LITERAL_N4
	OP_LITERAL_N8
	OP_COPY_N1_N1
	OP_COPY_N1_N2
	OP_COPY_N1_N4
	OP_COPY_N1_N8
	OP_COPY_N2_N1
	OP_COPY_N2_N2
	OP_COPY_N2_N4
	OP_COPY_N2_N8
	OP_COPY_N4_N1
	OP_COPY_N4_N2
	OP_COPY_N4_N4
	OP_COPY_N4_N8
	OP_COPY_N8_N1
	OP_COPY_N8_N2
	OP_COPY_N8_N4
	OP_COPY_N8_N8
	OP_RESERVED_85
	OP_RESERVED_86
	OP_RESERVED_87
	OP_RESERVED_88
	OP_RESERVED_89
	OP_RESERVED_90
	OP_RESERVED_91
	OP_RESERVED_92
	OP_RESERVED_93
	OP_RESERVED_94
	OP_RESERVED_95
	OP_RESERVED_96
	OP_RESERVED_97
	OP_RESERVED_98
	OP_RESERVED_99
	OP_RESERVED_100
	OP_RESERVED_101
	OP_RESERVED_102
	OP_RESERVED_103
	OP_RESERVED_104
	OP_RESERVED_105
	OP_RESERVED_106
	OP_RESERVED_107
	OP_RESERVED_108
	OP_RESERVED_109
	OP_RESERVED_110
	OP_RESERVED_111
	OP_RESERVED_112
	OP_RESERVED_113
	OP_RESERVED_114
	OP_RESERVED_115
	OP_RESERVED_116
	OP_RESERVED_117
	OP_RESERVED_118
	OP_RESERVED_119
	OP_RESERVED_120
	OP_RESERVED_121
	OP_RESERVED_122
	OP_RESERVED_123
	OP_RESERVED_124
	OP_RESERVED_125
	OP_RESERVED_126
	OP_RESERVED_127
	OP_RESERVED_128
	OP_RESERVED_129
	OP_RESERVED_130
	OP_RESERVED_131
	OP_RESERVED_132
	OP_RESERVED_133
	OP_RESERVED_134
	OP_RESERVED_135
	OP_RESERVED_136
	OP_RESERVED_137
	OP_RESERVED_138
	OP_RESERVED_139
	OP_RESERVED_140
	OP_RESERVED_141
	OP_RESERVED_142
	OP_RESERVED_143
	OP_RESERVED_144
	OP_RESERVED_145
	OP_RESERVED_146
	OP_RESERVED_147
	OP_RESERVED_148
	OP_RESERVED_149
	OP_RESERVED_150
	OP_RESERVED_151
	OP_RESERVED_152
	OP_RESERVED_153
	OP_RESERVED_154
	OP_RESERVED_155
	OP_RESERVED_156
	OP_RESERVED_157
	OP_RESERVED_158
	OP_RESERVED_159
	OP_RESERVED_160
	OP_RESERVED_161
	OP_RESERVED_162
	OP_RESERVED_163
	OP_RESERVED_164
	OP_RESERVED_165
	OP_RESERVED_166
	OP_RESERVED_167
	OP_RESERVED_168
	OP_RESERVED_169
	OP_RESERVED_170
	OP_RESERVED_171
	OP_RESERVED_172
	OP_RESERVED_173
	OP_RESERVED_174
	OP_RESERVED_175
	OP_RESERVED_176
	OP_RESERVED_177
	OP_RESERVED_178
	OP_RESERVED_179
	OP_RESERVED_180
	OP_RESERVED_181
	OP_RESERVED_182
	OP_RESERVED_183
	OP_RESERVED_184
	OP_RESERVED_185
	OP_RESERVED_186
	OP_RESERVED_187
	OP_RESERVED_188
	OP_RESERVED_189
	OP_RESERVED_190
	OP_RESERVED_191
	OP_RESERVED_192
	OP_RESERVED_193
	OP_RESERVED_194
	OP_RESERVED_195
	OP_RESERVED_196
	OP_RESERVED_197
	OP_RESERVED_198
	OP_RESERVED_199
	OP_RESERVED_200
	OP_RESERVED_201
	OP_RESERVED_202
	OP_RESERVED_203
	OP_RESERVED_204
	OP_RESERVED_205
	OP_RESERVED_206
	OP_RESERVED_207
	OP_RESERVED_208
	OP_RESERVED_209
	OP_RESERVED_210
	OP_RESERVED_211
	OP_RESERVED_212
	OP_RESERVED_213
	OP_RESERVED_214
	OP_RESERVED_215
	OP_RESERVED_216
	OP_RESERVED_217
	OP_RESERVED_218
	OP_RESERVED_219
	OP_RESERVED_220
	OP_RESERVED_221
	OP_RESERVED_222
	OP_RESERVED_223
	OP_RESERVED_224
	OP_RESERVED_225
	OP_RESERVED_226
	OP_RESERVED_227
	OP_RESERVED_228
	OP_RESERVED_229
	OP_RESERVED_230
	OP_RESERVED_231
	OP_RESERVED_232
	OP_RESERVED_233
	OP_RESERVED_234
	OP_RESERVED_235
	OP_RESERVED_236
	OP_RESERVED_237
	OP_RESERVED_238
	OP_RESERVED_239
	OP_RESERVED_240
	OP_RESERVED_241
	OP_RESERVED_242
	OP_RESERVED_243
	OP_RESERVED_244
	OP_RESERVED_245
	OP_RESERVED_246
	OP_RESERVED_247
	OP_RESERVED_248
	OP_RESERVED_249
	OP_RESERVED_250
	OP_RESERVED_251
	OP_RESERVED_252
	OP_RESERVED_253
	OP_RESERVED_254
	OP_RESERVED_255
)

var op2cmd = []Command{
	{KIND_END, 0, 0, 0},        /*            OP_END =    0 */
	{KIND_LITERAL, 1, 0, 0},    /*      OP_LITERAL_1 =  0x1 */
	{KIND_LITERAL, 2, 0, 0},    /*      OP_LITERAL_2 =  0x2 */
	{KIND_LITERAL, 3, 0, 0},    /*      OP_LITERAL_3 =  0x3 */
	{KIND_LITERAL, 4, 0, 0},    /*      OP_LITERAL_4 =  0x4 */
	{KIND_LITERAL, 5, 0, 0},    /*      OP_LITERAL_5 =  0x5 */
	{KIND_LITERAL, 6, 0, 0},    /*      OP_LITERAL_6 =  0x6 */
	{KIND_LITERAL, 7, 0, 0},    /*      OP_LITERAL_7 =  0x7 */
	{KIND_LITERAL, 8, 0, 0},    /*      OP_LITERAL_8 =  0x8 */
	{KIND_LITERAL, 9, 0, 0},    /*      OP_LITERAL_9 =  0x9 */
	{KIND_LITERAL, 10, 0, 0},   /*     OP_LITERAL_10 =  0xa */
	{KIND_LITERAL, 11, 0, 0},   /*     OP_LITERAL_11 =  0xb */
	{KIND_LITERAL, 12, 0, 0},   /*     OP_LITERAL_12 =  0xc */
	{KIND_LITERAL, 13, 0, 0},   /*     OP_LITERAL_13 =  0xd */
	{KIND_LITERAL, 14, 0, 0},   /*     OP_LITERAL_14 =  0xe */
	{KIND_LITERAL, 15, 0, 0},   /*     OP_LITERAL_15 =  0xf */
	{KIND_LITERAL, 16, 0, 0},   /*     OP_LITERAL_16 = 0x10 */
	{KIND_LITERAL, 17, 0, 0},   /*     OP_LITERAL_17 = 0x11 */
	{KIND_LITERAL, 18, 0, 0},   /*     OP_LITERAL_18 = 0x12 */
	{KIND_LITERAL, 19, 0, 0},   /*     OP_LITERAL_19 = 0x13 */
	{KIND_LITERAL, 20, 0, 0},   /*     OP_LITERAL_20 = 0x14 */
	{KIND_LITERAL, 21, 0, 0},   /*     OP_LITERAL_21 = 0x15 */
	{KIND_LITERAL, 22, 0, 0},   /*     OP_LITERAL_22 = 0x16 */
	{KIND_LITERAL, 23, 0, 0},   /*     OP_LITERAL_23 = 0x17 */
	{KIND_LITERAL, 24, 0, 0},   /*     OP_LITERAL_24 = 0x18 */
	{KIND_LITERAL, 25, 0, 0},   /*     OP_LITERAL_25 = 0x19 */
	{KIND_LITERAL, 26, 0, 0},   /*     OP_LITERAL_26 = 0x1a */
	{KIND_LITERAL, 27, 0, 0},   /*     OP_LITERAL_27 = 0x1b */
	{KIND_LITERAL, 28, 0, 0},   /*     OP_LITERAL_28 = 0x1c */
	{KIND_LITERAL, 29, 0, 0},   /*     OP_LITERAL_29 = 0x1d */
	{KIND_LITERAL, 30, 0, 0},   /*     OP_LITERAL_30 = 0x1e */
	{KIND_LITERAL, 31, 0, 0},   /*     OP_LITERAL_31 = 0x1f */
	{KIND_LITERAL, 32, 0, 0},   /*     OP_LITERAL_32 = 0x20 */
	{KIND_LITERAL, 33, 0, 0},   /*     OP_LITERAL_33 = 0x21 */
	{KIND_LITERAL, 34, 0, 0},   /*     OP_LITERAL_34 = 0x22 */
	{KIND_LITERAL, 35, 0, 0},   /*     OP_LITERAL_35 = 0x23 */
	{KIND_LITERAL, 36, 0, 0},   /*     OP_LITERAL_36 = 0x24 */
	{KIND_LITERAL, 37, 0, 0},   /*     OP_LITERAL_37 = 0x25 */
	{KIND_LITERAL, 38, 0, 0},   /*     OP_LITERAL_38 = 0x26 */
	{KIND_LITERAL, 39, 0, 0},   /*     OP_LITERAL_39 = 0x27 */
	{KIND_LITERAL, 40, 0, 0},   /*     OP_LITERAL_40 = 0x28 */
	{KIND_LITERAL, 41, 0, 0},   /*     OP_LITERAL_41 = 0x29 */
	{KIND_LITERAL, 42, 0, 0},   /*     OP_LITERAL_42 = 0x2a */
	{KIND_LITERAL, 43, 0, 0},   /*     OP_LITERAL_43 = 0x2b */
	{KIND_LITERAL, 44, 0, 0},   /*     OP_LITERAL_44 = 0x2c */
	{KIND_LITERAL, 45, 0, 0},   /*     OP_LITERAL_45 = 0x2d */
	{KIND_LITERAL, 46, 0, 0},   /*     OP_LITERAL_46 = 0x2e */
	{KIND_LITERAL, 47, 0, 0},   /*     OP_LITERAL_47 = 0x2f */
	{KIND_LITERAL, 48, 0, 0},   /*     OP_LITERAL_48 = 0x30 */
	{KIND_LITERAL, 49, 0, 0},   /*     OP_LITERAL_49 = 0x31 */
	{KIND_LITERAL, 50, 0, 0},   /*     OP_LITERAL_50 = 0x32 */
	{KIND_LITERAL, 51, 0, 0},   /*     OP_LITERAL_51 = 0x33 */
	{KIND_LITERAL, 52, 0, 0},   /*     OP_LITERAL_52 = 0x34 */
	{KIND_LITERAL, 53, 0, 0},   /*     OP_LITERAL_53 = 0x35 */
	{KIND_LITERAL, 54, 0, 0},   /*     OP_LITERAL_54 = 0x36 */
	{KIND_LITERAL, 55, 0, 0},   /*     OP_LITERAL_55 = 0x37 */
	{KIND_LITERAL, 56, 0, 0},   /*     OP_LITERAL_56 = 0x38 */
	{KIND_LITERAL, 57, 0, 0},   /*     OP_LITERAL_57 = 0x39 */
	{KIND_LITERAL, 58, 0, 0},   /*     OP_LITERAL_58 = 0x3a */
	{KIND_LITERAL, 59, 0, 0},   /*     OP_LITERAL_59 = 0x3b */
	{KIND_LITERAL, 60, 0, 0},   /*     OP_LITERAL_60 = 0x3c */
	{KIND_LITERAL, 61, 0, 0},   /*     OP_LITERAL_61 = 0x3d */
	{KIND_LITERAL, 62, 0, 0},   /*     OP_LITERAL_62 = 0x3e */
	{KIND_LITERAL, 63, 0, 0},   /*     OP_LITERAL_63 = 0x3f */
	{KIND_LITERAL, 64, 0, 0},   /*     OP_LITERAL_64 = 0x40 */
	{KIND_LITERAL, 0, 1, 0},    /*     OP_LITERAL_N1 = 0x41 */
	{KIND_LITERAL, 0, 2, 0},    /*     OP_LITERAL_N2 = 0x42 */
	{KIND_LITERAL, 0, 4, 0},    /*     OP_LITERAL_N4 = 0x43 */
	{KIND_LITERAL, 0, 8, 0},    /*     OP_LITERAL_N8 = 0x44 */
	{KIND_COPY, 0, 1, 1},       /*     OP_COPY_N1_N1 = 0x45 */
	{KIND_COPY, 0, 1, 2},       /*     OP_COPY_N1_N2 = 0x46 */
	{KIND_COPY, 0, 1, 4},       /*     OP_COPY_N1_N4 = 0x47 */
	{KIND_COPY, 0, 1, 8},       /*     OP_COPY_N1_N8 = 0x48 */
	{KIND_COPY, 0, 2, 1},       /*     OP_COPY_N2_N1 = 0x49 */
	{KIND_COPY, 0, 2, 2},       /*     OP_COPY_N2_N2 = 0x4a */
	{KIND_COPY, 0, 2, 4},       /*     OP_COPY_N2_N4 = 0x4b */
	{KIND_COPY, 0, 2, 8},       /*     OP_COPY_N2_N8 = 0x4c */
	{KIND_COPY, 0, 4, 1},       /*     OP_COPY_N4_N1 = 0x4d */
	{KIND_COPY, 0, 4, 2},       /*     OP_COPY_N4_N2 = 0x4e */
	{KIND_COPY, 0, 4, 4},       /*     OP_COPY_N4_N4 = 0x4f */
	{KIND_COPY, 0, 4, 8},       /*     OP_COPY_N4_N8 = 0x50 */
	{KIND_COPY, 0, 8, 1},       /*     OP_COPY_N8_N1 = 0x51 */
	{KIND_COPY, 0, 8, 2},       /*     OP_COPY_N8_N2 = 0x52 */
	{KIND_COPY, 0, 8, 4},       /*     OP_COPY_N8_N4 = 0x53 */
	{KIND_COPY, 0, 8, 8},       /*     OP_COPY_N8_N8 = 0x54 */
	{KIND_RESERVED, 85, 0, 0},  /*    OP_RESERVED_85 = 0x55 */
	{KIND_RESERVED, 86, 0, 0},  /*    OP_RESERVED_86 = 0x56 */
	{KIND_RESERVED, 87, 0, 0},  /*    OP_RESERVED_87 = 0x57 */
	{KIND_RESERVED, 88, 0, 0},  /*    OP_RESERVED_88 = 0x58 */
	{KIND_RESERVED, 89, 0, 0},  /*    OP_RESERVED_89 = 0x59 */
	{KIND_RESERVED, 90, 0, 0},  /*    OP_RESERVED_90 = 0x5a */
	{KIND_RESERVED, 91, 0, 0},  /*    OP_RESERVED_91 = 0x5b */
	{KIND_RESERVED, 92, 0, 0},  /*    OP_RESERVED_92 = 0x5c */
	{KIND_RESERVED, 93, 0, 0},  /*    OP_RESERVED_93 = 0x5d */
	{KIND_RESERVED, 94, 0, 0},  /*    OP_RESERVED_94 = 0x5e */
	{KIND_RESERVED, 95, 0, 0},  /*    OP_RESERVED_95 = 0x5f */
	{KIND_RESERVED, 96, 0, 0},  /*    OP_RESERVED_96 = 0x60 */
	{KIND_RESERVED, 97, 0, 0},  /*    OP_RESERVED_97 = 0x61 */
	{KIND_RESERVED, 98, 0, 0},  /*    OP_RESERVED_98 = 0x62 */
	{KIND_RESERVED, 99, 0, 0},  /*    OP_RESERVED_99 = 0x63 */
	{KIND_RESERVED, 100, 0, 0}, /*   OP_RESERVED_100 = 0x64 */
	{KIND_RESERVED, 101, 0, 0}, /*   OP_RESERVED_101 = 0x65 */
	{KIND_RESERVED, 102, 0, 0}, /*   OP_RESERVED_102 = 0x66 */
	{KIND_RESERVED, 103, 0, 0}, /*   OP_RESERVED_103 = 0x67 */
	{KIND_RESERVED, 104, 0, 0}, /*   OP_RESERVED_104 = 0x68 */
	{KIND_RESERVED, 105, 0, 0}, /*   OP_RESERVED_105 = 0x69 */
	{KIND_RESERVED, 106, 0, 0}, /*   OP_RESERVED_106 = 0x6a */
	{KIND_RESERVED, 107, 0, 0}, /*   OP_RESERVED_107 = 0x6b */
	{KIND_RESERVED, 108, 0, 0}, /*   OP_RESERVED_108 = 0x6c */
	{KIND_RESERVED, 109, 0, 0}, /*   OP_RESERVED_109 = 0x6d */
	{KIND_RESERVED, 110, 0, 0}, /*   OP_RESERVED_110 = 0x6e */
	{KIND_RESERVED, 111, 0, 0}, /*   OP_RESERVED_111 = 0x6f */
	{KIND_RESERVED, 112, 0, 0}, /*   OP_RESERVED_112 = 0x70 */
	{KIND_RESERVED, 113, 0, 0}, /*   OP_RESERVED_113 = 0x71 */
	{KIND_RESERVED, 114, 0, 0}, /*   OP_RESERVED_114 = 0x72 */
	{KIND_RESERVED, 115, 0, 0}, /*   OP_RESERVED_115 = 0x73 */
	{KIND_RESERVED, 116, 0, 0}, /*   OP_RESERVED_116 = 0x74 */
	{KIND_RESERVED, 117, 0, 0}, /*   OP_RESERVED_117 = 0x75 */
	{KIND_RESERVED, 118, 0, 0}, /*   OP_RESERVED_118 = 0x76 */
	{KIND_RESERVED, 119, 0, 0}, /*   OP_RESERVED_119 = 0x77 */
	{KIND_RESERVED, 120, 0, 0}, /*   OP_RESERVED_120 = 0x78 */
	{KIND_RESERVED, 121, 0, 0}, /*   OP_RESERVED_121 = 0x79 */
	{KIND_RESERVED, 122, 0, 0}, /*   OP_RESERVED_122 = 0x7a */
	{KIND_RESERVED, 123, 0, 0}, /*   OP_RESERVED_123 = 0x7b */
	{KIND_RESERVED, 124, 0, 0}, /*   OP_RESERVED_124 = 0x7c */
	{KIND_RESERVED, 125, 0, 0}, /*   OP_RESERVED_125 = 0x7d */
	{KIND_RESERVED, 126, 0, 0}, /*   OP_RESERVED_126 = 0x7e */
	{KIND_RESERVED, 127, 0, 0}, /*   OP_RESERVED_127 = 0x7f */
	{KIND_RESERVED, 128, 0, 0}, /*   OP_RESERVED_128 = 0x80 */
	{KIND_RESERVED, 129, 0, 0}, /*   OP_RESERVED_129 = 0x81 */
	{KIND_RESERVED, 130, 0, 0}, /*   OP_RESERVED_130 = 0x82 */
	{KIND_RESERVED, 131, 0, 0}, /*   OP_RESERVED_131 = 0x83 */
	{KIND_RESERVED, 132, 0, 0}, /*   OP_RESERVED_132 = 0x84 */
	{KIND_RESERVED, 133, 0, 0}, /*   OP_RESERVED_133 = 0x85 */
	{KIND_RESERVED, 134, 0, 0}, /*   OP_RESERVED_134 = 0x86 */
	{KIND_RESERVED, 135, 0, 0}, /*   OP_RESERVED_135 = 0x87 */
	{KIND_RESERVED, 136, 0, 0}, /*   OP_RESERVED_136 = 0x88 */
	{KIND_RESERVED, 137, 0, 0}, /*   OP_RESERVED_137 = 0x89 */
	{KIND_RESERVED, 138, 0, 0}, /*   OP_RESERVED_138 = 0x8a */
	{KIND_RESERVED, 139, 0, 0}, /*   OP_RESERVED_139 = 0x8b */
	{KIND_RESERVED, 140, 0, 0}, /*   OP_RESERVED_140 = 0x8c */
	{KIND_RESERVED, 141, 0, 0}, /*   OP_RESERVED_141 = 0x8d */
	{KIND_RESERVED, 142, 0, 0}, /*   OP_RESERVED_142 = 0x8e */
	{KIND_RESERVED, 143, 0, 0}, /*   OP_RESERVED_143 = 0x8f */
	{KIND_RESERVED, 144, 0, 0}, /*   OP_RESERVED_144 = 0x90 */
	{KIND_RESERVED, 145, 0, 0}, /*   OP_RESERVED_145 = 0x91 */
	{KIND_RESERVED, 146, 0, 0}, /*   OP_RESERVED_146 = 0x92 */
	{KIND_RESERVED, 147, 0, 0}, /*   OP_RESERVED_147 = 0x93 */
	{KIND_RESERVED, 148, 0, 0}, /*   OP_RESERVED_148 = 0x94 */
	{KIND_RESERVED, 149, 0, 0}, /*   OP_RESERVED_149 = 0x95 */
	{KIND_RESERVED, 150, 0, 0}, /*   OP_RESERVED_150 = 0x96 */
	{KIND_RESERVED, 151, 0, 0}, /*   OP_RESERVED_151 = 0x97 */
	{KIND_RESERVED, 152, 0, 0}, /*   OP_RESERVED_152 = 0x98 */
	{KIND_RESERVED, 153, 0, 0}, /*   OP_RESERVED_153 = 0x99 */
	{KIND_RESERVED, 154, 0, 0}, /*   OP_RESERVED_154 = 0x9a */
	{KIND_RESERVED, 155, 0, 0}, /*   OP_RESERVED_155 = 0x9b */
	{KIND_RESERVED, 156, 0, 0}, /*   OP_RESERVED_156 = 0x9c */
	{KIND_RESERVED, 157, 0, 0}, /*   OP_RESERVED_157 = 0x9d */
	{KIND_RESERVED, 158, 0, 0}, /*   OP_RESERVED_158 = 0x9e */
	{KIND_RESERVED, 159, 0, 0}, /*   OP_RESERVED_159 = 0x9f */
	{KIND_RESERVED, 160, 0, 0}, /*   OP_RESERVED_160 = 0xa0 */
	{KIND_RESERVED, 161, 0, 0}, /*   OP_RESERVED_161 = 0xa1 */
	{KIND_RESERVED, 162, 0, 0}, /*   OP_RESERVED_162 = 0xa2 */
	{KIND_RESERVED, 163, 0, 0}, /*   OP_RESERVED_163 = 0xa3 */
	{KIND_RESERVED, 164, 0, 0}, /*   OP_RESERVED_164 = 0xa4 */
	{KIND_RESERVED, 165, 0, 0}, /*   OP_RESERVED_165 = 0xa5 */
	{KIND_RESERVED, 166, 0, 0}, /*   OP_RESERVED_166 = 0xa6 */
	{KIND_RESERVED, 167, 0, 0}, /*   OP_RESERVED_167 = 0xa7 */
	{KIND_RESERVED, 168, 0, 0}, /*   OP_RESERVED_168 = 0xa8 */
	{KIND_RESERVED, 169, 0, 0}, /*   OP_RESERVED_169 = 0xa9 */
	{KIND_RESERVED, 170, 0, 0}, /*   OP_RESERVED_170 = 0xaa */
	{KIND_RESERVED, 171, 0, 0}, /*   OP_RESERVED_171 = 0xab */
	{KIND_RESERVED, 172, 0, 0}, /*   OP_RESERVED_172 = 0xac */
	{KIND_RESERVED, 173, 0, 0}, /*   OP_RESERVED_173 = 0xad */
	{KIND_RESERVED, 174, 0, 0}, /*   OP_RESERVED_174 = 0xae */
	{KIND_RESERVED, 175, 0, 0}, /*   OP_RESERVED_175 = 0xaf */
	{KIND_RESERVED, 176, 0, 0}, /*   OP_RESERVED_176 = 0xb0 */
	{KIND_RESERVED, 177, 0, 0}, /*   OP_RESERVED_177 = 0xb1 */
	{KIND_RESERVED, 178, 0, 0}, /*   OP_RESERVED_178 = 0xb2 */
	{KIND_RESERVED, 179, 0, 0}, /*   OP_RESERVED_179 = 0xb3 */
	{KIND_RESERVED, 180, 0, 0}, /*   OP_RESERVED_180 = 0xb4 */
	{KIND_RESERVED, 181, 0, 0}, /*   OP_RESERVED_181 = 0xb5 */
	{KIND_RESERVED, 182, 0, 0}, /*   OP_RESERVED_182 = 0xb6 */
	{KIND_RESERVED, 183, 0, 0}, /*   OP_RESERVED_183 = 0xb7 */
	{KIND_RESERVED, 184, 0, 0}, /*   OP_RESERVED_184 = 0xb8 */
	{KIND_RESERVED, 185, 0, 0}, /*   OP_RESERVED_185 = 0xb9 */
	{KIND_RESERVED, 186, 0, 0}, /*   OP_RESERVED_186 = 0xba */
	{KIND_RESERVED, 187, 0, 0}, /*   OP_RESERVED_187 = 0xbb */
	{KIND_RESERVED, 188, 0, 0}, /*   OP_RESERVED_188 = 0xbc */
	{KIND_RESERVED, 189, 0, 0}, /*   OP_RESERVED_189 = 0xbd */
	{KIND_RESERVED, 190, 0, 0}, /*   OP_RESERVED_190 = 0xbe */
	{KIND_RESERVED, 191, 0, 0}, /*   OP_RESERVED_191 = 0xbf */
	{KIND_RESERVED, 192, 0, 0}, /*   OP_RESERVED_192 = 0xc0 */
	{KIND_RESERVED, 193, 0, 0}, /*   OP_RESERVED_193 = 0xc1 */
	{KIND_RESERVED, 194, 0, 0}, /*   OP_RESERVED_194 = 0xc2 */
	{KIND_RESERVED, 195, 0, 0}, /*   OP_RESERVED_195 = 0xc3 */
	{KIND_RESERVED, 196, 0, 0}, /*   OP_RESERVED_196 = 0xc4 */
	{KIND_RESERVED, 197, 0, 0}, /*   OP_RESERVED_197 = 0xc5 */
	{KIND_RESERVED, 198, 0, 0}, /*   OP_RESERVED_198 = 0xc6 */
	{KIND_RESERVED, 199, 0, 0}, /*   OP_RESERVED_199 = 0xc7 */
	{KIND_RESERVED, 200, 0, 0}, /*   OP_RESERVED_200 = 0xc8 */
	{KIND_RESERVED, 201, 0, 0}, /*   OP_RESERVED_201 = 0xc9 */
	{KIND_RESERVED, 202, 0, 0}, /*   OP_RESERVED_202 = 0xca */
	{KIND_RESERVED, 203, 0, 0}, /*   OP_RESERVED_203 = 0xcb */
	{KIND_RESERVED, 204, 0, 0}, /*   OP_RESERVED_204 = 0xcc */
	{KIND_RESERVED, 205, 0, 0}, /*   OP_RESERVED_205 = 0xcd */
	{KIND_RESERVED, 206, 0, 0}, /*   OP_RESERVED_206 = 0xce */
	{KIND_RESERVED, 207, 0, 0}, /*   OP_RESERVED_207 = 0xcf */
	{KIND_RESERVED, 208, 0, 0}, /*   OP_RESERVED_208 = 0xd0 */
	{KIND_RESERVED, 209, 0, 0}, /*   OP_RESERVED_209 = 0xd1 */
	{KIND_RESERVED, 210, 0, 0}, /*   OP_RESERVED_210 = 0xd2 */
	{KIND_RESERVED, 211, 0, 0}, /*   OP_RESERVED_211 = 0xd3 */
	{KIND_RESERVED, 212, 0, 0}, /*   OP_RESERVED_212 = 0xd4 */
	{KIND_RESERVED, 213, 0, 0}, /*   OP_RESERVED_213 = 0xd5 */
	{KIND_RESERVED, 214, 0, 0}, /*   OP_RESERVED_214 = 0xd6 */
	{KIND_RESERVED, 215, 0, 0}, /*   OP_RESERVED_215 = 0xd7 */
	{KIND_RESERVED, 216, 0, 0}, /*   OP_RESERVED_216 = 0xd8 */
	{KIND_RESERVED, 217, 0, 0}, /*   OP_RESERVED_217 = 0xd9 */
	{KIND_RESERVED, 218, 0, 0}, /*   OP_RESERVED_218 = 0xda */
	{KIND_RESERVED, 219, 0, 0}, /*   OP_RESERVED_219 = 0xdb */
	{KIND_RESERVED, 220, 0, 0}, /*   OP_RESERVED_220 = 0xdc */
	{KIND_RESERVED, 221, 0, 0}, /*   OP_RESERVED_221 = 0xdd */
	{KIND_RESERVED, 222, 0, 0}, /*   OP_RESERVED_222 = 0xde */
	{KIND_RESERVED, 223, 0, 0}, /*   OP_RESERVED_223 = 0xdf */
	{KIND_RESERVED, 224, 0, 0}, /*   OP_RESERVED_224 = 0xe0 */
	{KIND_RESERVED, 225, 0, 0}, /*   OP_RESERVED_225 = 0xe1 */
	{KIND_RESERVED, 226, 0, 0}, /*   OP_RESERVED_226 = 0xe2 */
	{KIND_RESERVED, 227, 0, 0}, /*   OP_RESERVED_227 = 0xe3 */
	{KIND_RESERVED, 228, 0, 0}, /*   OP_RESERVED_228 = 0xe4 */
	{KIND_RESERVED, 229, 0, 0}, /*   OP_RESERVED_229 = 0xe5 */
	{KIND_RESERVED, 230, 0, 0}, /*   OP_RESERVED_230 = 0xe6 */
	{KIND_RESERVED, 231, 0, 0}, /*   OP_RESERVED_231 = 0xe7 */
	{KIND_RESERVED, 232, 0, 0}, /*   OP_RESERVED_232 = 0xe8 */
	{KIND_RESERVED, 233, 0, 0}, /*   OP_RESERVED_233 = 0xe9 */
	{KIND_RESERVED, 234, 0, 0}, /*   OP_RESERVED_234 = 0xea */
	{KIND_RESERVED, 235, 0, 0}, /*   OP_RESERVED_235 = 0xeb */
	{KIND_RESERVED, 236, 0, 0}, /*   OP_RESERVED_236 = 0xec */
	{KIND_RESERVED, 237, 0, 0}, /*   OP_RESERVED_237 = 0xed */
	{KIND_RESERVED, 238, 0, 0}, /*   OP_RESERVED_238 = 0xee */
	{KIND_RESERVED, 239, 0, 0}, /*   OP_RESERVED_239 = 0xef */
	{KIND_RESERVED, 240, 0, 0}, /*   OP_RESERVED_240 = 0xf0 */
	{KIND_RESERVED, 241, 0, 0}, /*   OP_RESERVED_241 = 0xf1 */
	{KIND_RESERVED, 242, 0, 0}, /*   OP_RESERVED_242 = 0xf2 */
	{KIND_RESERVED, 243, 0, 0}, /*   OP_RESERVED_243 = 0xf3 */
	{KIND_RESERVED, 244, 0, 0}, /*   OP_RESERVED_244 = 0xf4 */
	{KIND_RESERVED, 245, 0, 0}, /*   OP_RESERVED_245 = 0xf5 */
	{KIND_RESERVED, 246, 0, 0}, /*   OP_RESERVED_246 = 0xf6 */
	{KIND_RESERVED, 247, 0, 0}, /*   OP_RESERVED_247 = 0xf7 */
	{KIND_RESERVED, 248, 0, 0}, /*   OP_RESERVED_248 = 0xf8 */
	{KIND_RESERVED, 249, 0, 0}, /*   OP_RESERVED_249 = 0xf9 */
	{KIND_RESERVED, 250, 0, 0}, /*   OP_RESERVED_250 = 0xfa */
	{KIND_RESERVED, 251, 0, 0}, /*   OP_RESERVED_251 = 0xfb */
	{KIND_RESERVED, 252, 0, 0}, /*   OP_RESERVED_252 = 0xfc */
	{KIND_RESERVED, 253, 0, 0}, /*   OP_RESERVED_253 = 0xfd */
	{KIND_RESERVED, 254, 0, 0}, /*   OP_RESERVED_254 = 0xfe */
	{KIND_RESERVED, 255, 0, 0}, /*   OP_RESERVED_255 = 0xff */
}
