# [ When I serve my site, I get an error &quot;Liquid Exception: Failed to get header. in _posts&quot;](https://github.com/jekyll/jekyll-help/issues/198)

> state: **closed** opened by: **** on: ****

Here are my steps:

1. I installed jekyll (gem install jekyll)
2. I created a new site (jekyll new sitename)
3. I navigated to that site folder (cd sitename)
4. I served the site (jekyll serve)

I then get this error:  Liquid Exception: Failed to get header. in _posts/2014-11-18-welcome-to-jekyll.markdown

I have not touched the folders in the new site. The post that is causing the problem was not created by me but came with the install. 

I am using a Mac OSX 10.8.5; python 3.6.2 and pygments 0.6.0

Any idea what is causing this issue? 

Many thanks!

Susan



### Comments

---
> from: [**susanholland**](https://github.com/jekyll/jekyll-help/issues/198#issuecomment-63507038) on: **Tuesday, November 18, 2014**

UPDATE: I deleted that post (which was originally added during the install) and my serve worked perfectly. 

I&#x27;m still curious why a file I didn&#x27;t create caused the problem, though. 
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/198#issuecomment-63563151) on: **Tuesday, November 18, 2014**

@susanholland I *believe* you need to running &lt; 3.x Python for Jekyll to work properly.

@parkr Can you confirm, please?
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/198#issuecomment-63563417) on: **Tuesday, November 18, 2014**

The process you listed above works fine for me running 2.7.6:

&#x60;&#x60;&#x60;
Troys-MacBook-Air:~ troy$ ruby --version
ruby 2.0.0p481 (2014-05-08 revision 45883) [x86_64-darwin13.3.0]
Troys-MacBook-Air:~ troy$ python --version
Python 2.7.6
Troys-MacBook-Air:~ troy$ jekyll --version
jekyll 2.5.1
Troys-MacBook-Air:~ troy$ cd Projects/
Troys-MacBook-Air:Projects troy$ jekyll new sitename
New jekyll site installed in /Users/troy/Projects/sitename. 
Troys-MacBook-Air:Projects troy$ cd sitename/
Troys-MacBook-Air:sitename troy$ jekyll serve
Configuration file: /Users/troy/Projects/sitename/_config.yml
            Source: /Users/troy/Projects/sitename
       Destination: /Users/troy/Projects/sitename/_site
      Generating... 
                    done.
 Auto-regeneration: enabled for &#x27;/Users/troy/Projects/sitename&#x27;
Configuration file: /Users/troy/Projects/sitename/_config.yml
    Server address: http://127.0.0.1:4000/
  Server running... press ctrl-c to stop.
&#x60;&#x60;&#x60;
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/198#issuecomment-63563863) on: **Tuesday, November 18, 2014**

Only Python 2.7.x seems to work for Pygments. Can you try &#x60;rouge&#x60;?
---
> from: [**susanholland**](https://github.com/jekyll/jekyll-help/issues/198#issuecomment-63636974) on: **Wednesday, November 19, 2014**

Thanks for the feedback. I got it to work fine with Python 3 by uninstalling the blog post in the post file that was automatically added during the install. That was what was causing the problem, though I am curious what about it was causing the problem. 

When I get some time, I&#x27;ll try going down to Python 2.7.6, adding this file back and seeing if it works. I&quot;m curious. 

The file was called &quot;2014-11-18-welcome-to-jekyll.markdown. Here&#x27;s what was in the file:

---
layout: post
title:  &quot;Welcome to Jekyll!&quot;
date:   2014-11-18 11:08:13
categories: jekyll update
---
You’ll find this post in your &#x60;_posts&#x60; directory. Go ahead and edit it and re-build the site to see your changes. You can rebuild the site in many different ways, but the most common way is to run &#x60;jekyll serve&#x60;, which launches a web server and auto-regenerates your site when a file is updated.

To add new posts, simply add a file in the &#x60;_posts&#x60; directory that follows the convention &#x60;YYYY-MM-DD-name-of-post.ext&#x60; and includes the necessary front matter. Take a look at the source for this post to get an idea about how it works.

Jekyll also offers powerful support for code snippets:

{% highlight ruby %}
def print_hi(name)
  puts &quot;Hi, #{name}&quot;
end
print_hi(&#x27;Tom&#x27;)
#=&gt; prints &#x27;Hi, Tom&#x27; to STDOUT.
{% endhighlight %}

Check out the [Jekyll docs][jekyll] for more info on how to get the most out of Jekyll. File all bugs/feature requests at [Jekyll’s GitHub repo][jekyll-gh]. If you have questions, you can ask them on [Jekyll’s dedicated Help repository][jekyll-help].

[jekyll]:      http://jekyllrb.com
[jekyll-gh]:   https://github.com/jekyll/jekyll
[jekyll-help]: https://github.com/jekyll/jekyll-help
---
> from: [**susanholland**](https://github.com/jekyll/jekyll-help/issues/198#issuecomment-63645923) on: **Wednesday, November 19, 2014**

FYI - I downgraded to Python 2.7.8. I put the document that was causing the problem back into the posts folder and I was able to serve the site! 

Thanks for the help. 
