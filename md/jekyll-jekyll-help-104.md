# [Sass compiler not working](https://github.com/jekyll/jekyll-help/issues/104)

> state: **closed** opened by: **** on: ****

Jekyll seems to ignore my &#x60;*.scss&#x60; files and doesn&#x27;t compile any CSS.

This is my project:

&#x60;&#x60;&#x60;&#x60;
.
├── LICENSE
├── README.md
├── _config.yml
├── _scss
│   ├── _borders.scss
│   ├── _button-themes.scss
│   ├── _buttons.scss
│   ├── _colors.scss
│   ├── _form-themes.scss
│   ├── _forms.scss
│   ├── _grid.scss
│   ├── _lists.scss
│   ├── _margins.scss
│   ├── _nav-themes.scss
│   ├── _navs.scss
│   ├── _padding.scss
│   ├── _positions.scss
│   ├── _reset.scss
│   ├── _table-object.scss
│   ├── _table-themes.scss
│   ├── _tables.scss
│   ├── _theme.scss
│   ├── _type.scss
│   ├── _utilities.scss
│   ├── _variables.scss
│   ├── basscss-lite.scss
│   └── basscss.scss
├── css
├── gulpfile.js
├── index.html
├── js
│   ├── app.js
│   └── min
│       └── app.js
└── package.json

4 directories, 31 files
&#x60;&#x60;&#x60;&#x60;

This is my &#x60;_config.yml&#x60;:

&#x60;&#x60;&#x60;&#x60;
# Site settings
title: trim-jekyll-basscss
email: your-email@domain.com
description: &quot;Combinig Jekyll, Trim and BassCSS&quot;
baseurl: &quot;/&quot;
url: &quot;http://https://github.com/mrzool/trim-jekyll-basscss&quot;
twitter_username: __zool
github_username:  mrzool

# Build settings
markdown: kramdown
permalink: pretty

# Sass settings
sass:
    sass_dir: _scss
    style: :compressed
&#x60;&#x60;&#x60;&#x60;

The Sass on my system: &#x60;Sass 3.3.11 (Maptastic Maple)&#x60;
Jekyll: &#x60;jekyll 2.1.1&#x60;

Am I doing something wrong?

Thanks for any help.

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/104#issuecomment-50481229) on: **Tuesday, July 29, 2014**

Have you prepended dashes on your basscss.scss?

&#x60;&#x60;&#x60;
---
---
// content of your basscss.scss
&#x60;&#x60;&#x60;
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/104#issuecomment-50481670) on: **Tuesday, July 29, 2014**

No folders that start with an underscore will be read in. Place your base scss files in the desired output directory, such as &#x60;css&#x60;
---
> from: [**mrzool**](https://github.com/jekyll/jekyll-help/issues/104#issuecomment-50485172) on: **Tuesday, July 29, 2014**

@kleinfreund I did. Matter of fact, every single *.scss file in the scss folder starts with two lines of triple dashes.

@parkr I thought it was the other way around, i.e. that every folder not prefixed with an underscore would get mirrored into the generated site without getting processed. Also,  [Jekyll&#x27;s documentation](http://jekyllrb.com/docs/assets/) uses an underscore-prefixed directory as an example:

&gt; If you are using Sass @import statements, you’ll need to ensure that your sass_dir is set to the base directory that contains your Sass files. You can do that thusly:

        sass:
            sass_dir: _sass

&gt; The Sass converter will default to _sass.

That said, I did as you said and now Jekyll is processing the Sass (actually it isn&#x27;t because I&#x27;m getting errors, but at least it sees the files).

Does that mean that there is an error in the documentation? It is very unclear how to proceed here, I think that page should be expanded quite a bit.
 
---
> from: [**mrzool**](https://github.com/jekyll/jekyll-help/issues/104#issuecomment-50486933) on: **Tuesday, July 29, 2014**

UPDATE: I was getting errors because i prepended dashes also to the partials, while they where required in the main file only. Everything is working now! 

That&#x27;s another thing that should be better dealt with in the docs though, IMHO.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/104#issuecomment-50508996) on: **Tuesday, July 29, 2014**

@mrzool Sorry for the confusion! I&#x27;ll rewrite the docs page.
---
> from: [**mrzool**](https://github.com/jekyll/jekyll-help/issues/104#issuecomment-50509751) on: **Tuesday, July 29, 2014**

@parkr Thanks a lot.
