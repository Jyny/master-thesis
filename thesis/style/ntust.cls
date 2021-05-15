% --------------------------------------------------
% 01 類別宣告（Class Declaration）
% --------------------------------------------------
\NeedsTeXFormat{LaTeX2e}
\ProvidesClass{style/ntust}[2021/2/29 NTUST Thesis]

% --------------------------------------------------
% 02 基礎設定（Basic Settings）
% --------------------------------------------------
\newcommand{\var@tableofcontents}{目錄}
\newcommand{\var@listoftables}{表目錄}
\newcommand{\var@listoffigures}{圖目錄}

% --------------------------------------------------
% 03 編譯檢查（Compile Checking）
% --------------------------------------------------

% --------------------------------------------------
% 04 論文設定（Document Configs）
% --------------------------------------------------

% --------------------------------------------------
% 05 選項設定（Options Settings）
% --------------------------------------------------
\DeclareOption*{\PassOptionsToClass{\CurrentOption}{report}}
\ProcessOptions*\relax
\LoadClass[a4paper, 12pt]{report}



% --------------------------------------------------
% 06 載入套件 (Include Packages)
% --------------------------------------------------
\RequirePackage{xparse}
\RequirePackage{xeCJK}
\RequirePackage{indentfirst}
\RequirePackage{fontspec}
\RequirePackage{xcolor}
\RequirePackage{titlesec}
\RequirePackage{titletoc, tocloft}
\RequirePackage{fancyhdr}
\RequirePackage{setspace}
\RequirePackage{graphicx}
\RequirePackage{tikz}
\RequirePackage{eso-pic}
\RequirePackage{pdfpages}
\RequirePackage{tabularx}
\RequirePackage{booktabs}
\RequirePackage{multirow}
\RequirePackage{diagbox}
\RequirePackage{array}
\RequirePackage{longtable}
\RequirePackage{datetime}
\RequirePackage{CJKnumb}
\RequirePackage{fp, calc}
\RequirePackage{hyperref}
\RequirePackage{paralist}
\RequirePackage{footmisc}
\RequirePackage{caption}
\RequirePackage[labelformat=simple]{subcaption}
\RequirePackage{geometry}
\RequirePackage{float}
\RequirePackage{enumitem}

% --------------------------------------------------
% 07 套件設定 (Package Settings)
% --------------------------------------------------
% [圖表目錄]
\graphicspath{{./figures/}}

% [文章排版]
%
\XeTeXlinebreaklocale "zh"
\XeTeXlinebreakskip = 0pt plus 1pt
\setlength{\parskip}{10pt}
\setlength{\parindent}{2em}

% 設定行距
% 1.2：1.5 倍行距
% 1.6：2.0 倍行距
\setstretch{1.2}

% [字體設定]
%
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

% [版面大小]
%
\RequirePackage{geometry}
\geometry{top=3cm, bottom=3cm, left=3cm, right=3cm}

% [連結設定]
%
\hypersetup{
    colorlinks=true,
    linkcolor=black,
    filecolor=gray,
    anchorcolor=blue,
    citecolor=blue
}

% [目錄格式]
%
\addtocontents{toc}                                         % 修改目錄樣式
{~\hfill\textbf{Page}\par}
{}
{}
{}

% \titlecontents{標題名稱}[左間距]{標題格式}{標題內容}{無序號標題}{指引線與頁碼}[下間距]
\titlecontents{chapter}[0em]                                % 修改目錄樣式
{\addvspace{0.5em}\normalsize\bfseries\protect}
{\fboxrule=0pt\fboxsep=0pt\framebox[4.5em][l]{第\CJKnumber{\thecontentslabel}章}}
{}
{\hfill\contentspage}

\setlength{\cftbeforesecskip}{8pt}                          % 修改目錄間距
\setlength{\cftbeforesubsecskip}{8pt}

\setcounter{tocdepth}{3}                                    % 設定目錄深度
\setcounter{secnumdepth}{3}                                 % 設定編號深度

% \cftsetindents{entry}{indent}{numwidth}
\cftsetindents{section}{1.5em}{3.0em}
\cftsetindents{subsection}{3.0em}{3.0em}

% [頁首頁尾]
%

% [標題格式]
% 載入套件包 titlesec 設定各層標題格式
% \titleformat{command}[shape]{format}{label}{sep}{before}[after]
% - command: 代表欲重新定義之標題階層，包含 \part, \chapter, \section, \paragraph
% - shape: 設定段落形狀，可選參數有 hang, block, display
% - format: 定義標題外觀，如：位置、字體大小粗細…
% - label: 定義標籤前綴標號
% - sep: 定義標號與標題內容間隔距離
% - before: 標題內容前增加文字
% - after: 標題內容後增加文字
%
\titleformat{\chapter}                      % 設置 Chapter 格式
{\centering\Huge\bfseries}                  % 定義 format
{第\CJKnumber{\thechapter}章}          	     % 定義 label
{1em}                                       % 定義 sep
{}                                          % before

% [圖表標題]
%
\setlength{\intextsep}{12bp \@plus4pt \@minus2pt}
\captionsetup{}
\captionsetup[table]{position=top,belowskip={12bp-\intextsep},aboveskip=6bp}
\captionsetup[figure]{position=bottom,belowskip={12bp-\intextsep},aboveskip=6bp}
\captionsetup[sub]{skip=6bp}

% [列表]
%
\setlist*[]{
    topsep=0pt,
    partopsep=0pt,
    listparindent=\parindent,
}

% --------------------------------------------------
% 08 自訂命令（Custom Commands）
% --------------------------------------------------

% [生成目錄]
% Generate table of contents page.
%
\renewcommand{\contentsname}{\centerline{\var@tableofcontents}}
\newcommand{\maketableofcontents}{
    \cleardoublepage
    \phantomsection
    \addcontentsline{toc}{chapter}{\var@tableofcontents}
    \tableofcontents
}

\renewcommand{\listtablename}{\centerline{\var@listoftables}}
\newcommand{\makelistoftables}{
    \cleardoublepage
    \phantomsection
    \addcontentsline{toc}{chapter}{\var@listoftables}
    \renewcommand{\numberline}{\tablename~\oldnumberline}
    \listoftables
}

\let\oldnumberline\numberline
\renewcommand{\listfigurename}{\centerline{\var@listoffigures}}
\newcommand{\makelistoffigures}{
    \cleardoublepage
    \phantomsection
    \addcontentsline{toc}{chapter}{\var@listoffigures}
    \renewcommand{\numberline}{\figurename~\oldnumberline}
    \listoffigures
}

% [正文開始]
% \mainmatter
%
\newcommand{\mainmatter}{
    \cleardoublepage
    \expandafter\newgeometry\expandafter{top=2.5cm, bottom=2.5cm, left=3cm, right=2.5cm}
    \pagenumbering{arabic}
    \fontsize{12}{18}\selectfont
}

% --------------------------------------------------
% 09 自訂環境（Custom Environments）
% --------------------------------------------------

\newlist{steps}{enumerate}{1}
\setlist[steps, 1]{
    label = {\bf Step \arabic*},
    leftmargin=\parindent+\labelwidth+\labelsep,
    align=left,
}

% --------------------------------------------------
% 10 文檔處理（Document Processing）
% --------------------------------------------------

% 套用翻譯
\renewcommand{\figurename}{圖}
\renewcommand{\tablename}{表}