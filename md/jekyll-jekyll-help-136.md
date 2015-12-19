# [type: &quot;posts&quot; in _config.yml default block seems not working](https://github.com/jekyll/jekyll-help/issues/136)

> state: **closed** opened by: **** on: ****

hi, I am trying to create blog site in github page. the site source is in https://github.com/notyy/notyy.github.io

I specified default values in _config.yml
&#x60;&#x60;&#x60;YAML
defaults:
  -
    scope:
      path: &quot;&quot;
      # type: &quot;posts&quot;
    values:
      layout: &quot;post&quot;
      encoding: utf-8
&#x60;&#x60;&#x60;
if I uncomment the line of &#x60;type: &quot;posts&quot;&#x60; then, the layout and encoding is lost.


### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/136#issuecomment-52918711) on: **Thursday, August 21, 2014**

That’s because version 2.3 is not released on GitHub Pages. It’s &#x60;type: post&#x60; with the current version.
---
> from: [**notyy**](https://github.com/jekyll/jekyll-help/issues/136#issuecomment-52919086) on: **Thursday, August 21, 2014**

Yeah! tried and true!  thank you very much.
