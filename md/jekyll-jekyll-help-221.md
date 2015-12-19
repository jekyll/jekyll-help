# [css/main.scss contains syntax errors when I push to gh-pages](https://github.com/jekyll/jekyll-help/issues/221)

> state: **closed** opened by: **** on: ****

I keep getting this error but i works fine with i run it with -&gt; bundle exec jekyll serve

This is the file -&gt; https://github.com/lucachaco/sourcequarter-lander/blob/gh-pages/project/website/css/main.scss

It is the same that came with the default theme.

I already tried this workaround but it didnt work -&gt; https://github.com/jekyll/jekyll-help/issues/185

Any help would be appreciate it. 

Thanks.

### Comments

---
> from: [**lucachaco**](https://github.com/jekyll/jekyll-help/issues/221#issuecomment-68371082) on: **Tuesday, December 30, 2014**

I realised I had to move the project to the root dir to make it work.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/221#issuecomment-68660160) on: **Sunday, January 04, 2015**

I think you could have worked around this by adding a &#x60;_config.yml&#x60; file in the root of the repo and adding a &#x60;source&#x60; property that points to the relative position of the root of the Jekyll source... Not entirely sure if that would work though.

Glad you got it working, either way.
---
> from: [**lucachaco**](https://github.com/jekyll/jekyll-help/issues/221#issuecomment-68669618) on: **Sunday, January 04, 2015**

Thank you troyswanson for your comment.
I was going to try that but I saw github overwrites that settings: https://help.github.com/articles/using-jekyll-with-pages/#configuring-jekyll

So, moving the files to the root I think was my only option a the moment.

Thanks.
