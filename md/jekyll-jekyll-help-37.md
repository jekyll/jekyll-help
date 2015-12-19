# [Template no longer included in site after upgrade from 1.4.3 to 2.0.2](https://github.com/jekyll/jekyll-help/issues/37)

> state: **closed** opened by: **** on: ****

Greetings,

I was excited to upgrade to 2.0.2, but found that my generated index.html file no longer included the template code. I reverted back to 1.4.3 and everything works again.

Here&#x27;s the repo:

https://github.com/morea-framework/basic-template

It just generates the index.html without any content in the &lt;head&gt; element.

What am I doing wrong, and/or what should I change for 2.0 to make this site work correctly? 

Or if it works fine for you, that would be useful info as well. 

Thanks,
Philip



### Comments

---
> from: [**philipmjohnson**](https://github.com/jekyll/jekyll-help/issues/37#issuecomment-42581186) on: **Thursday, May 08, 2014**

My basic-template repo includes a ruby plugin (MoreaGenerator.rb).  I disabled the plugin (by renaming it to MoreaGenerator.txt) and found that the problem still exists under Jekyll 2.0.2  but not under 1.4.3.  So there must be something simple that I&#x27;m doing wrong that causes my [morea.html](https://github.com/morea-framework/basic-template/blob/master/src/_layouts/morea.html) layout to not be included when [index.md](https://github.com/morea-framework/basic-template/blob/master/src/index.md) is compiled to index.html.  
---
> from: [**philipmjohnson**](https://github.com/jekyll/jekyll-help/issues/37#issuecomment-42614370) on: **Thursday, May 08, 2014**

I just created a more detailed description of my problem with Jekyll 2.0.2 here:

http://stackoverflow.com/questions/23553629/why-does-my-site-work-using-jekyll-1-4-3-but-not-jekyll-2-0-2

Please respond in StackExchange.  Thanks!
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/37#issuecomment-42629170) on: **Thursday, May 08, 2014**

I have run your site in Jekyll 2.0.2 and I&#x27;m seeing this:

![screen shot 2014-05-08 at 10 26 44 pm](https://cloud.githubusercontent.com/assets/451539/2923983/9c9cc9d4-d725-11e3-9789-7ee88b5b9048.png)

It seems it&#x27;s working, no?
---
> from: [**philipmjohnson**](https://github.com/jekyll/jekyll-help/issues/37#issuecomment-42638417) on: **Thursday, May 08, 2014**

No, that image indicates it&#x27;s broken (and that you&#x27;re using the master branch, not the jekyll-2.0 branch).  When correctly generated (as happens in 1.4.3) you should see a page styled using bootstrap similar to this:

![basic-template-1 4 3](https://cloud.githubusercontent.com/assets/290288/2925195/abee0f8a-d746-11e3-96ad-12961ca2a6d2.png)

---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/37#issuecomment-42645060) on: **Friday, May 09, 2014**

Your template has no default layout, though it is referenced from your post layout:

https://github.com/morea-framework/basic-template/blob/jekyll-2.0/src/_layouts/post.html


---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/37#issuecomment-42645618) on: **Friday, May 09, 2014**

I can&#x27;t see how this could have worked in 1.4.3 either. You need to have a default layout.
---
> from: [**philipmjohnson**](https://github.com/jekyll/jekyll-help/issues/37#issuecomment-42685174) on: **Friday, May 09, 2014**

Thanks, renaming the layout morea.html to default.html and updating
references did fix my problem.

Since index.md was a page, not a post, it was not obvious to me that the
reference to a layout called default in post.html would matter to
index.html generation (and, obviously, it did not matter under Jekyll
1.4.3).

Maybe update the docs at some point to specify that a layout called
default.html is required?

I appreciate your quick response!

Philip


On Thu, May 8, 2014 at 10:52 PM, Seth Warburton &lt;notifications@github.com&gt;wrote:

&gt; I can&#x27;t see how this could have worked in 1.4.3 either. You need to have a
&gt; default layout.
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/37#issuecomment-42645618&gt;
&gt; .
&gt;
---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/37#issuecomment-42688041) on: **Friday, May 09, 2014**

No worries, I&#x27;m glad I was able to resolve it for you.

AFAIK, nothing has changed in this behaviour though, there always has to be a default layout file whether it is called default or not. You were asking Jekyll to render content to a file that didn&#x27;t exist, that will always cause a break.

The only thing that defines if the output of the index file is a page, post or anything else is the layout specified in the files front-matter.
