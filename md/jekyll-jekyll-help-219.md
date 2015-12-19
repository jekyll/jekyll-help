# [Jekyll Removing Changes To Edited Files](https://github.com/jekyll/jekyll-help/issues/219)

> state: **closed** opened by: **** on: ****

Hello,

So far Jekyll has been fun to use and I am starting to get more into it rather than creating basic posts. However, I have made various changes to files such as the main.scss or the index.html in the &quot;_site&quot; folder and find these edited files to display correctly while my server is running, but automatically return to there original state when I restart/regenerate Jekyll. In other words, I completely lose all of my changes. This is very frustrating as I intend to make many UI/functional changes on my site. Has anyone experienced similar problems? I appreciate any advice you can give! 

Thank you!

### Comments

---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/219#issuecomment-68120325) on: **Thursday, December 25, 2014**

When I am developing, I usually run Jekyll in serve &amp; watch mode with &#x60;jekyll serve -w&#x60;. This way, I can edit the source files and preview changes immediately (similar to editing the files in &#x60;_site&#x60;), while still keeping my changes permanent (i.e. not editing the &#x60;_site&#x60; files themselves).

Hope that helps!
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/219#issuecomment-68660633) on: **Sunday, January 04, 2015**

@szier1 You should consider the &#x60;_site&#x60; folder untouchable. The code in Jekyll specifically removes everything in the &#x60;_site&#x60; folder every time it is regenerated. Depending on the platform you are using, there are ways to have your web browser reload once Jekyll is done regenerating the site.

I personally use [CodeKit](https://incident57.com/codekit/) and point it to the &#x60;_site&#x60; folder in my project. It watches for changes on those files and automatically refreshes Chrome. Once incremental regeneration is finished, it will work even better.

&gt; When I am developing, I usually run Jekyll in serve &amp; watch mode with &#x60;jekyll serve -w&#x60;

@alfredxing I don&#x27;t remember which version of Jekyll made the change, but the &#x60;serve&#x60; command enables auto-regeneration by default.
