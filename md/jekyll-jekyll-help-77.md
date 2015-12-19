# [Liquid Exception: Included file not found](https://github.com/jekyll/jekyll-help/issues/77)

> state: **closed** opened by: **** on: ****

Hi all:

With the caveat that I&#x27;m not a Ruby person, and more of a back-end person in general... I&#x27;m trying to broaden my horizons and learn a bit.

I&#x27;ve set up a Jekyll site on Github Pages at http://charliegriefer.github.io. I&#x27;m using the [Minimal Mistakes theme](https://github.com/mmistakes/minimal-mistakes), which is awesome. Everything was going swimmingly until last night, when I wanted to add TravisCI (another new to me technology) to the mix.

Everything still builds and works fine locally. But my Travis build failed with the following:
&gt; Liquid Exception: Included file &#x27;/home/travis/build/charliegriefer/charliegriefer.github.io/_includes/news_item.html&#x27; not found in vendor/bundle/ruby/1.9.1/gems/jekyll-1.5.1/site/news/index.html

I&#x27;ve &quot;fixed&quot; the issue by adding &#x60;news_item.html&#x60; to &#x60;_includes&#x60;, but I think we can all agree that this isn&#x27;t really the proper fix :)

I&#x27;ve verified that Jekyll on my local machine (OS X Mavericks) is 1.5.1, which I believe is in use on Github Pages. My version of ruby shows as &#x60;ruby 2.0.0p451 (2014-02-24 revision 45167) [universal.x86_64-darwin13]&#x60;. I wouldn&#x27;t think this would be an issue, even tho Github Pages uses 2.1.1, but it&#x27;s the only difference that I can see.

In a nutshell, the Minimal Mistakes theme already had a Gemfile and gemfile.lock. I modified those as part of getting Travis CI set up. I&#x27;m probably not intelligent enough (yet) to be modifying Gemfiles. But I&#x27;d imagine that maybe something there went terribly wrong (in spite of the fact that my jekyll build and serve work fine locally).

I don&#x27;t see _anywhere_ in my files where I&#x27;d be calling &#x60;news_item&#x60;, and it&#x27;s driving me nuts.

Any thoughts/suggestions/advice would be greatly appreciated. If there&#x27;s any additional information that I should be providing, please let me know.

Thanks!
Charlie

### Comments

---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/77#issuecomment-47317756) on: **Friday, June 27, 2014**

You must exclude the vendor directory from the build. Add to your &#x60;_config.yml&#x60;:

&#x60;&#x60;&#x60;
exclude: [vendor]
&#x60;&#x60;&#x60;
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/77#issuecomment-47318025) on: **Friday, June 27, 2014**

@penibelst is correct! That should fix your build. Currently working on a CI guide: https://github.com/jekyll/jekyll/pull/2432
