\NeedsTeXFormat{LaTeX2e}
\ProvidesClass{style/ntust}[2021/2/29 NTUST Thesis]

\DeclareOption*{\PassOptionsToClass{\CurrentOption}{report}}
\ProcessOptions\relax
\LoadClass[a4paper, 12pt]{report}

\RequirePackage{graphicx}
\RequirePackage{setspace}
\RequirePackage{geometry}
\RequirePackage{xeCJK}
\RequirePackage{fontspec}
\RequirePackage{float}

\setmainfont{Times New Roman}[
    Path = fonts/english/,
    Extension = .ttf ,
    BoldFont = *-Bold ,
    ItalicFont = *-Italic,
    BoldItalicFont = *-BoldItalic
]

\setCJKmainfont{BiauKai}[
    AutoFakeBold=true,
    Path = fonts/chinese/,
    Extension = .ttf
]