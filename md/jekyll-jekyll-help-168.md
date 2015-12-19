# [Collections site object is empty](https://github.com/jekyll/jekyll-help/issues/168)

> state: **closed** opened by: **** on: ****

I&#x27;m building a new site using Jekyll and I&#x27;m having trouble getting collections to work. I&#x27;m trying to use collections to display a portfolio of projects.

I&#x27;ve added this to my &#x60;_config.yml&#x60; file:
&#x60;&#x60;&#x60;yaml
collections:
  projects:
    output: true
&#x60;&#x60;&#x60;&#x60;
I then created a folder called &#x60;_projects&#x60; and within it I have one file named &#x60;2014-10-08-test-project.md&#x60;

At this point, I&#x27;m just hoping to have an html file created, or a way to reference it from another page (for instance, the home page).

However, when I build this, no project html files are created. And if I output &#x60;{{ site }}&#x60; on a page, this is the only reference to my collection that I get: &#x60;&quot;collections&quot;=&gt;{&quot;projects&quot;=&gt;{&quot;output&quot;=&gt;true}}&#x60;

I installed Jekyll this week, so I shouldn&#x27;t be on an old version. I&#x27;m on OSX 10.9.5. When I run &#x60;ruby -v&#x60; on the terminal, I get this &#x60;ruby 1.9.3p362 (2012-12-25 revision 38607) [x86_64-darwin12.2.0]&#x60;. I&#x27;m also using this in conjunction with yeoman.

I&#x27;ve even tried using some sample code from blog posts that I&#x27;ve found (ex, https://github.com/tay1orjones/jekyllTest), and I get the same result - the collections aren&#x27;t getting processed.

I&#x27;ve reread the documentation, and I&#x27;ve been searching through stackoverflow and the github issues with no luck. I&#x27;m worried that this is an issue with my machine/setup more than Jekyll but not sure where to look.

Any ideas or suggestions would be greatly appreciated.

thanks!

### Comments

---
> from: [**tylercraft**](https://github.com/jekyll/jekyll-help/issues/168#issuecomment-60037347) on: **Tuesday, October 21, 2014**

Looks like this was something unique to my system. A fresh install of Yosemite and all was fixed.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/168#issuecomment-61438671) on: **Sunday, November 02, 2014**

Glad you got it sorted out!
