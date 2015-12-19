# [SCSS compiles but SASS does not (v2.0.2)](https://github.com/jekyll/jekyll-help/issues/34)

> state: **closed** opened by: **** on: ****

In a new project, I removed main.css and replaced it with main.sass.

I included the two lines of &quot;---&quot; and just wrote

    body
      background-color: red

This was copied as main.sass in _site/css when building.

Changing this to main.scss and

    body {
      background-color: red;
    }

It compiles to main.css in _site/css as you would expect.

### Comments

---
> from: [**ajkochanowicz**](https://github.com/jekyll/jekyll-help/issues/34#issuecomment-42464983) on: **Wednesday, May 07, 2014**

Moving to main Jekyll issues.
