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
\newcommand{\var@bibliography}{參考文獻}
\newcommand{\var@abstract@zh}{摘要}
\newcommand{\var@abstract@en}{Abstract}
\newcommand{\var@acknowledgement}{致謝}

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
\RequirePackage{kvdefinekeys, kvsetkeys, kvoptions}
\RequirePackage{etex}
\RequirePackage{xparse}
\RequirePackage{xeCJK}
\RequirePackage{indentfirst}
\RequirePackage{fontspec}
\RequirePackage[table]{xcolor}
\RequirePackage[compact]{titlesec}
\RequirePackage[titles]{tocloft}
\RequirePackage{titletoc}
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
\RequirePackage{arydshln}
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
\RequirePackage[numbers]{natbib}
\RequirePackage[normalem]{ulem}
\RequirePackage{amsmath, amsthm, amssymb}
\RequirePackage{ulem, CJKulem}
\RequirePackage{lipsum}
\RequirePackage{zhlipsum}
\RequirePackage[T1]{fontenc}
\RequirePackage{algorithm2e}
\RequirePackage{cryptocode}
\RequirePackage[
    scale=0.13,
    angle=0,
]{background}


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

% [引用格式]
%
\bibliographystyle{IEEEtran}

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
\setCJKmonofont{BiauKai}[
    AutoFakeBold=true,
    Path = fonts/chinese/,
    Extension = .ttf
]

% [版面大小]
%
\geometry{top=3cm, bottom=3cm, left=3cm, right=3cm}

% [連結設定]
%
\hypersetup{
    colorlinks=true,
    linkcolor=black,
    urlcolor=black,
    filecolor=gray,
    anchorcolor=blue,
    citecolor=blue
}

% [目錄格式]
%
\addtocontents{toc}{\vskip -3em}
\addtocontents{lof}{\vskip -3em}
\addtocontents{lot}{\vskip -3em}

% \titlecontents{標題名稱}[左間距]{標題格式}{標題內容}{無序號標題}{指引線與頁碼}[下間距]
\titlecontents{chapter}[0em]                                % 修改目錄樣式
{\addvspace{0.5em}\normalsize\bfseries\protect}
{\fboxrule=0pt\fboxsep=0pt\framebox[4.5em][l]{第\CJKnumber{\thecontentslabel}章}}
{}
{\hfill\contentspage}

\setlength{\cftbeforesecskip}{8pt}                          % 修改目錄間距
\setlength{\cftbeforesubsecskip}{8pt}

\setcounter{tocdepth}{2}                                    % 設定目錄深度
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
\titlespacing{\chapter}{0pt}{-3em}{2.4\parskip}
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

% [參考文獻]
% \refmatter
%
\newcommand{\refmatter}{
  \cleardoublepage
  \phantomsection
  \renewcommand{\bibname}{\var@bibliography}
  \addcontentsline{toc}{chapter}{\var@bibliography}
}


% --------------------------------------------------
% 09 自訂環境（Custom Environments）
% --------------------------------------------------

\SetupKeyvalOptions{
  family=var,
  prefix=var@,
  setkeys=\kvsetkeys
}

\newcommand{\varsetup}{
  \kvsetkeys{var}
}

\newcommand{\var@define@key}[1]{
  \kvsetkeys{var@key}{#1}
}

\kv@set@family@handler{var@key}{%
  \@namedef{var@#1@@name}{#1}%
  \def\var@@default{}%
  \def\var@@choices{}%
  \kv@define@key{var@value}{name}{%
    \@namedef{var@#1@@name}{##1}%
  }%
  \kv@define@key{var@value}{code}{%
    \@namedef{var@#1@@code}{##1}%
  }%
  \@namedef{var@#1@@check}{}%
  \@namedef{var@#1@@code}{}%
  \@namedef{var@#1@@hook}{%
    \expandafter\ifx\csname\@currname.\@currext-h@@k\endcsname\relax
      \@nameuse{var@#1@@code}%
    \else
      \AtEndOfClass{%
        \@nameuse{var@#1@@code}%
      }%
    \fi
  }%
  \kv@define@key{var@value}{choices}{%
    \def\var@@choices{##1}%
    \@namedef{var@#1@@reset}{}%
    \@namedef{var@#1@@check}{%
      \@ifundefined{%
        ifvar@\@nameuse{var@#1@@name}@\@nameuse{var@\@nameuse{var@#1@@name}}%
      }{%
        \var@error{Invalid value "#1 = \@nameuse{var@\@nameuse{var@#1@@name}}"}%
      }%
      \@nameuse{var@#1@@reset}%
      \@nameuse{var@\@nameuse{var@#1@@name}@\@nameuse{var@\@nameuse{var@#1@@name}}true}%
    }%
  }%
  \kv@define@key{var@value}{default}{%
    \def\var@@default{##1}%
  }%
  \kvsetkeys{var@value}{#2}%
  \@namedef{var@\@nameuse{var@#1@@name}}{}%
  \kv@set@family@handler{var@choice}{%
    \ifx\var@@default\@empty
      \def\var@@default{##1}%
    \fi
    \expandafter\newif\csname ifvar@\@nameuse{var@#1@@name}@##1\endcsname
    \expandafter\g@addto@macro\csname var@#1@@reset\endcsname{%
      \@nameuse{var@\@nameuse{var@#1@@name}@##1false}%
    }%
  }%
  \kvsetkeys@expandafter{var@choice}{\var@@choices}%
  \expandafter\let\csname var@\@nameuse{var@#1@@name}\endcsname\var@@default
  \expandafter\ifx\csname var@\@nameuse{var@#1@@name}\endcsname\@empty\else
    \@nameuse{var@#1@@check}%
    \@nameuse{var@#1@@hook}%
  \fi
  \kv@define@key{var}{#1}{%
    \@namedef{var@\@nameuse{var@#1@@name}}{##1}%
    \@nameuse{var@#1@@check}%
    \@nameuse{var@#1@@hook}%
  }%
}

\newcommand\var@option@hook[2]{
  \expandafter\g@addto@macro\csname var@#1@@code\endcsname{#2}
}

% 命令: 鍵值陣列使用
% Command: Retrieve the array value with key-value pair.
%
\newcommand\var@clist@use[2]{%
  \def\var@@tmp{}%
  \kv@set@family@handler{var@clist}{%
    \ifx\var@@tmp\@empty
      \def\var@@tmp{#2}%
    \else
      #2%
    \fi
    ##1%
  }%
  \kvsetkeys@expandafter{var@clist}{#1}%
}

\var@define@key{
    keywords = {
        default = {關鍵字},
    },
    keywords* = {
        default = {Keywords},
        name    = keywords@en,
    },
}

\newlist{steps}{enumerate}{1}
\setlist[steps, 1]{
    label = {步驟 \arabic*:},
    leftmargin=\parindent+\labelwidth+\labelsep,
    align=left,
}

\newlist{pmsgs}{enumerate}{1}
\setlist[pmsgs, 1]{
    label = {$\pmb M_{\arabic*}:$},
    labelsep = 0em,
    leftmargin=\parindent+\labelwidth,
    align=left,
}

\newlist{pmsgsi}{enumerate}{1}
\setlist[pmsgsi, 1]{
    label = {$\pmb M_{\arabic*}^{i}:$},
    labelsep = 0em,
    leftmargin=\parindent+\labelwidth,
    align=left,
}

\definecolor{whitesmoke}{rgb}{0.9, 0.9, 0.9}

% [摘要環境]
% Abstract Environment
%
\renewenvironment{abstract}{
  \fontsize{12}{18}\selectfont
  \chapter*{\centering \var@abstract@zh}
  \addcontentsline{toc}{chapter}{\var@abstract@zh}
}{
  \bigbreak \bigbreak
  \noindent \textbf{關鍵字：}
  \var@clist@use{\var@keywords}{、}
}

\newenvironment{abstract*}{
  \fontsize{12}{24}\selectfont
  \setstretch{1.1}
  \chapter*{\centering \var@abstract@en}
  \addcontentsline{toc}{chapter}{\var@abstract@en}
}{
  \bigbreak \bigbreak
  \noindent \textbf{Keywords:~}
  \var@clist@use{\var@keywords@en}{,~}
}

% [誌謝環境]
% Acknowledgement Environment
%
\newenvironment{acknowledgement}{
  \chapter*{\centering \var@acknowledgement}
  \addcontentsline{toc}{chapter}{\var@acknowledgement}
  \fontsize{12}{18}\selectfont
}


% --------------------------------------------------
% 10 文檔處理（Document Processing）
% --------------------------------------------------

% 套用翻譯
\renewcommand{\figurename}{圖}
\renewcommand{\tablename}{表}