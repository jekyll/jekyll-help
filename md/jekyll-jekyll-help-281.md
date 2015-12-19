# [Jekyll server running, html showing, CSS won&#x27;t render.](https://github.com/jekyll/jekyll-help/issues/281)

> state: **closed** opened by: **** on: ****

Hello,

I have been trying to get sass working with jekyll for a long time. I have tried about every solution
shown on google. From uninstalling gems, updating/changing yaml files, --- ---, reinstalling related gems, uninstalling reinstalling many different versions of jekyll, but still sass won&#x27;t show. 
Occasionally, changing css properties in _whatever.sass or changing text in index.html will regenerate
css on the localhost. But no sooner had I kept on working than sass freezes in the application not showing styles anymore.
I can provide code examples but rather ask if this is a known issue by now (march 2015) that I can look into and therefore fix the problem.  This leads me into despair and I really don&#x27;t want to give up on Jekyll.  Thanks for any help, and I apologize for bringing this issue up again.  

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/281#issuecomment-77633174) on: **Friday, March 06, 2015**

What do you mean by “sass won&#x27;t show”?
---
> from: [**Elemshalit**](https://github.com/jekyll/jekyll-help/issues/281#issuecomment-77647851) on: **Friday, March 06, 2015**

&quot;sass won&#x27;t show&quot; meaning jekyll won&#x27;t spit out css styles into the index.html. I am using sass and bourbon, and trying to get jekyll to compile  sass to css.  Maybe I should try grunt for the task.  When I revert the project to an earlier version where css clearly is showing,I try to understand what triggers the problem. and sometimes, if I change a value or a text in either _includes, index.html, or 3-modules I get jekyll to show css styles. I don&#x27;t know if its either a build,listen,generate,compile, or a watch issue or something else. but I can&#x27;t get css to work properly. This issue has been documented a lot in this forum. And I have spent like 15 hours trying every avenue in sight. I even tried reinstalling rvm and all of its dependencies as that proved to be effective in other posts. 
---
> from: [**willnorris**](https://github.com/jekyll/jekyll-help/issues/281#issuecomment-77648730) on: **Friday, March 06, 2015**

This certainly isn&#x27;t a known issue, at least in the sense that lots of people generate sass without any issue.  An example of a non-working site would be helpful... it sounds like you&#x27;re looking in the right places, but without knowing exactly what it is you&#x27;re changing, it&#x27;s impossible to know what the problem is.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/281#issuecomment-77679457) on: **Saturday, March 07, 2015**

I still don’t know what your problem is.

&gt; jekyll won&#x27;t spit out css styles into the index.html

You have a Sass file somewhere that starts with the triple dashes (i.e. has front matter) that imports other Sass partials, right? For example, I have a file inside &#x60;/css&#x60; called &#x60;kleinfreund.scss&#x60; that starts like this:

&#x60;&#x60;&#x60;sass
---
---
@charset &quot;utf-8&quot;;
@import
        &quot;base&quot;,
        &quot;typography&quot;,
        // ...
;
&#x60;&#x60;&#x60;

This will spit out a CSS file called &#x60;kleinfreund.css&#x60; in the same directory on building the site.

This file is usually referenced in the &#x60;default.html&#x60; layout with a &#x60;link&#x60; tag. Do you have all this?
---
> from: [**Elemshalit**](https://github.com/jekyll/jekyll-help/issues/281#issuecomment-77741643) on: **Sunday, March 08, 2015**

https://github.com/Elemshalit/JekyllTest001 Here it is..

I am still trying to figure out why css styles are not rendering in the browser.
---
> from: [**willnorris**](https://github.com/jekyll/jekyll-help/issues/281#issuecomment-77742906) on: **Sunday, March 08, 2015**

okay, so you&#x27;ve got a few problems...

First, remove [&#x60;assets/css/main.css&#x60;](https://github.com/Elemshalit/JekyllTest001/blob/master/assets/css/main.css).  Because this static file exists, this is what is always getting loaded instead of your sass file being converted into CSS.  I&#x27;m not sure if main.css is a compiled version of main.sass?  If so, you don&#x27;t need to check it in at all, just let Jekyll build it when needed.  If it&#x27;s something different, then you should rename it so that it doesn&#x27;t conflict.

You also have a trailing space in the first line of [main.sass](https://github.com/Elemshalit/JekyllTest001/blob/master/assets/css/main.sass) which is preventing it from getting processed by jekyll.  Remove that trailing space character.

Once you do that, when you run jekyll you should get the error:

&#x60;&#x60;&#x60;
Conversion error: Jekyll::Converters::Sass encountered an error while converting &#x27;assets/css/main.sass&#x27;:
Invalid CSS after &quot;&quot;: expected selector, was &quot;---&quot; on line 1
&#x60;&#x60;&#x60;

This is because [1-tools/_-tools-index.sass](https://github.com/Elemshalit/JekyllTest001/blob/master/assets/css/1-tools/_-tools-index.sass), [2-base/_-base-index.sass](https://github.com/Elemshalit/JekyllTest001/blob/master/assets/css/2-base/_-base-index.sass), and [3-modules/_-modules-index.sass](https://github.com/Elemshalit/JekyllTest001/blob/master/assets/css/3-modules/_-modules-index.sass) all have front-matter defined.  They shouldn&#x27;t.  Included sass files shouldn&#x27;t have front matter, only the main file that does all of the importing (in your case, &#x60;main.sass&#x60;).  So remove the &#x60;--- ---&#x60; from the top of those files.

Once you do that, you should get the error:

&#x60;&#x60;&#x60;
Conversion error: Jekyll::Converters::Sass encountered an error while converting &#x27;assets/css/main.sass&#x27;:
Invalid CSS after &quot;padding:&quot;: expected pseudoclass or pseudoelement, was &quot;30px&quot; on line 6
&#x60;&#x60;&#x60;

This is because you are missing a space after the &quot;:&quot; in [3-modules/_footer.sass](https://github.com/Elemshalit/JekyllTest001/blob/master/assets/css/3-modules/_footer.sass#L6).  Add a space after the colon.  (I honestly had no idea sass was so sensitive to whitespace in this regard... one of the reasons I much prefer the scss dialect I guess)

After that, your sass file should build fine.  I&#x27;ll send a pull request on your test repo with all of these changes, but I wanted to explain them here as well so you understand what went wrong.  Let me know if you have trouble applying these fixes.

You should probably also add &#x60;.sass-cache&#x60; and &#x60;.jekyll-metadata&#x60; to your &#x60;.gitignore&#x60; file... neither of those should really be checked in to version control.
---
> from: [**willnorris**](https://github.com/jekyll/jekyll-help/issues/281#issuecomment-77743129) on: **Sunday, March 08, 2015**

And just to show that these changes are working:

![screen shot 2015-03-08 at 3 47 16 am](https://cloud.githubusercontent.com/assets/1112/6545460/e62bb3da-c545-11e4-8dee-9b736575cfcb.png)

---
> from: [**Elemshalit**](https://github.com/jekyll/jekyll-help/issues/281#issuecomment-77744487) on: **Sunday, March 08, 2015**

Perfect! Thank you so much for getting me up and running!To both of you. @kleinfreund @willnorris 
Everything is MUCH clearer now. I&#x27;ll update more if I may throughout the development stage if any other issues arise. I will look into your fix more closely tomorrow, but as of now, I can attest to your assumption that sass is indeed sensitive in this instance. Don&#x27;t use spaces sparingly when you indent that&#x27;s for sure. 
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/281#issuecomment-85239719) on: **Monday, March 23, 2015**

:metal: Thanks guys for the assist!
