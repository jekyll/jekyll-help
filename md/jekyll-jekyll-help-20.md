# [Template layouts](https://github.com/jekyll/jekyll-help/issues/20)

> state: **closed** opened by: **** on: ****

Similar to https://github.com/jekyll/jekyll/wiki/Sites would it be useful to have a list of some skeleton setups with different frameworks? For instance, I&#x27;ve been working on a new site, and there&#x27;s still nothing site-specific in the repo, uploading it to a repo as an example and  skeleton setup for Jekyll + Compass + ZURB Foundation 5 could be useful to others.

### Comments

---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-40652933) on: **Wednesday, April 16, 2014**

good idea. I&#x27;ve been working on one with bourbon. 
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-40656265) on: **Wednesday, April 16, 2014**

Here&#x27;s an initial attempt: https://github.com/jekyll/jekyll/wiki/Template-Websites
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-40719596) on: **Thursday, April 17, 2014**

Check out @imathis&#x27;s octopress-ink: it&#x27;s a way of shipping your themes as gems and allowing seamless inheritance of various layout elements. https://github.com/octopress/ink
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-40726125) on: **Thursday, April 17, 2014**

I was thinking of something a lot less complete than entire themes. More
like starting points with various frameworks pre-installed as a jumping
point for an actual theme. Does octopress-ink require octopress? From the
gemspec it doesn&#x27;t seem to. And for future readers, the correct link is
https://github.com/octopress/ink


On 18 April 2014 00:26, Parker Moore &lt;notifications@github.com&gt; wrote:

&gt; Check out @imathis &lt;https://github.com/imathis&gt;&#x27;s octopress-ink: it&#x27;s a
&gt; way of shipping your themes as gems and allowing seamless inheritance of
&gt; various layout elements. https://github.com/octopress/octopress-ink
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/help/issues/20#issuecomment-40719596&gt;
&gt; .
&gt;
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-40726487) on: **Thursday, April 17, 2014**

It does *not* require Octopress at all! Octopress has turned into tooling around making Jekyll even easier to use, and less about providing default themes and such like it is now.
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-40850408) on: **Friday, April 18, 2014**

@mscharley I am working on [one with Foundation](https://github.com/penibelst/jekyll-noitadnouf).

&gt; I was thinking of something a lot less complete than entire themes.

I’m not sure where to stop. What do you expect from such minimal setup?
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-40850519) on: **Friday, April 18, 2014**

@penibelst I suppose that depends entirely on what you can get away with generically... have a look at my Foundation one: https://github.com/mscharley/template-jekyll-foundation5
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-40851618) on: **Friday, April 18, 2014**

@mscharley Haha, Foundation gets much love. Maybe we can work together? I switch the discussion to your repository.
---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-42647223) on: **Friday, May 09, 2014**

What I love about Jekyll is that it doesn&#x27;t include any junk. I&#x27;d hate to see a change like this.

It&#x27;s also a question of maintenance, why incorporate what is essentially third-party code and create a maintenance burden in the process? Someone would have to ensure that there was parity between the upstream third-party code and the Jekyll version. It&#x27;s creates dependencies that needn&#x27;t exist.

I recently built a Jekyll site using Foundation 5 as a start-point and incorporating it was as easy as opening terminal, cd&#x27;ing to my project folder and typing &#x60;bower install foundation --save&#x60;. Third-party dependencies should, in my opinion, always be ring-fenced and treated as external assets, not internalised.
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-42647855) on: **Friday, May 09, 2014**

@nternetinspired Noone&#x27;s suggesting this would ever touch jekyll core in any form. My original concept was more along the line of &#x27;themes&#x27; in other CMS&#x27;s. And yes, you can install foundation that quickly and easily. But in my case, there&#x27;s other things going on as well. Setting up jekyll-compass to work, adding RequireJS to support jQuery and Foundation, and a few other bits and pieces. Now I can just clone that repository, fork it and do whatever I need without having to worry about setting up all those dependencies for my new project. Maintenance of each theme would be provided (or not!) by the original creator of it. Since it&#x27;s just a git repository of an actual site, there&#x27;s no added burden to jekyll at all, and if you don&#x27;t want to use any of them then that&#x27;s your prerogative.

As an aside, I think that Octopress Ink is probably going to be the way to approach this in the future.
---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-42654213) on: **Friday, May 09, 2014**

Gotcha. Sorry, too early in my day I totally misunderstood. 

There&#x27;s plenty such themes out there, I agree a centralised list of them would be cool. Maybe just start one and then send a pull request?
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-42654274) on: **Friday, May 09, 2014**

I did start one on the wiki, I&#x27;m not sure if people have added to it or not
since then though.


On 9 May 2014 20:53, Seth Warburton &lt;notifications@github.com&gt; wrote:

&gt; Gotcha. Sorry, too early in my day I totally misunderstood.
&gt;
&gt; There&#x27;s plenty such themes out there, I agree a centralised list of them
&gt; would be cool. Maybe just start one and then send a pull request?
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/20#issuecomment-42654213&gt;
&gt; .
&gt;
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-42658826) on: **Friday, May 09, 2014**

I was thinking of adding to it, though I was thinking of doing one that includes a few items that I find helpful (config file stuff to, for instance, put in your google analytics id, or toggle production state), and to add some test content ala Wordpress unit tests. So, something that stands between a full-blown theme and the more skeleton-like starters you were suggesting, @mscharley. 
---
> from: [**klepas**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-44661687) on: **Friday, May 30, 2014**

@budparr Hey, I’m interested in doing the same, providing Bourbon and Neat. I’m actually having some problems trying to get Jekyll to load the SCSS files. I saw your repo at https://github.com/budparr/jekyll-starter — figured this is the one you’re working on? (:
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-44701129) on: **Friday, May 30, 2014**

Hi @klepas I don&#x27;t think I&#x27;ve worked on that since 2.x came out. I think I&#x27;m going to abandon jekyll-starter in favor of jekyll-on-the-rocks, but that one too use the pages gem (i.e. 1.5.1). You can look at any of my client repos under my company name. The top two are things I&#x27;m working on currently so are 2.x (that said, they too are drafts and quite messy): https://github.com/sonnetmedia
---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-48096450) on: **Saturday, July 05, 2014**

I added [JekyllKB](https://github.com/AJAlabs/jekyllkb) to the [Jekyll Template wiki page](https://github.com/jekyll/jekyll/wiki/Template-Websites). Right now JekyllKB is just a starter site for Jekyll using Bootstrap and Font Awesome. The goal is to make it a starting point for your Jekyll powered Knowledge Base or documentation site. Contributions are welcome.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-55623905) on: **Monday, September 15, 2014**

@parkr Is support for themes something that might appear in Jekyll 3.0? I seem to recall a conversation about including a &quot;theme gem&quot; in one&#x27;s config file, but I&#x27;m not sure where that went.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-55637273) on: **Monday, September 15, 2014**

Eventually, I think we&#x27;ll go with @imathis&#x27;s Octopress Ink project which is absolutely baller. We won&#x27;t ever include it in Jekyll core (it would stay a plugin but be a runtime dependency), but I have concerns over it ever making it onto GitHub Pages (security, security, security).

So keep moving on this front – it&#x27;s a sure-fire way to see &quot;themes&quot; like this on GHP. Octopress Ink will come later.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/20#issuecomment-55641664) on: **Monday, September 15, 2014**

:boom: 
