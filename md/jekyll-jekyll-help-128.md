# [Upgrading Jekyll from 1.4.3 to 2.3.0,  Error:  undefined method &#x60;glob_include?&#x27; for [&quot;.htaccess&quot;]:Array](https://github.com/jekyll/jekyll-help/issues/128)

> state: **open** opened by: **** on: ****

I updated my project from 1.4.2 to 2.3.0 and got this error: 

jekyll 2.3.0 | Error:  undefined method &#x60;glob_include?&#x27; for [&quot;.htaccess&quot;]:Array

I don&#x27;t have an .htaccess file nor do I use one, so this was puzzling. 

Turns out.. after some debugging and help from jaybe on irc, it was because I had the jekyll-sass gem. The later versions of Jekyll have sass support built in. Removing jekyll-sass from the Gemfile fixed it!

Hope this saves someone a few hours :) 

### Comments

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/128#issuecomment-52430912) on: **Sunday, August 17, 2014**

:thumbsup: 
